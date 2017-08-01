package main_test

import (
	"bytes"
	"fmt"
	"io"
	"io/ioutil"
	"net/http"
	"net/http/httputil"
	//"net/http/httptest"
	"encoding/json"
	"errors"
	"net/url"
	"os"
	"strings"
	"testing"
	"time"

	lg "github.com/hiromaily/golibs/log"
	r "github.com/hiromaily/golibs/runtimes"
)

const SERVER_HOST = "http://localhost:8080"

//These struct were copied from ./goa/app/contexts.go
type loginAuthPayload struct {
	// E-mail of user
	Email *string `form:"email,omitempty" json:"email,omitempty" xml:"email,omitempty"`
	// Password
	Password *string `form:"password,omitempty" json:"password,omitempty" xml:"password,omitempty"`
}

//TODO:change
type createUserHyUserPayload struct {
	// E-mail of user
	Email *string `form:"email,omitempty" json:"email,omitempty" xml:"email,omitempty"`
	// Password
	Password *string `form:"password,omitempty" json:"password,omitempty" xml:"password,omitempty"`
	// First name
	UserName *string `form:"user_name,omitempty" json:"user_name,omitempty" xml:"user_name,omitempty"`
}

var (
	jwtToken          string
	errRedirect       = errors.New("redirect")
	contentTypeForm   = map[string]string{"Content-Type": "application/x-www-form-urlencoded"}
	contentTypeJson   = map[string]string{"Content-Type": "application/json"}
	jwtAuth           = map[string]string{"Authorization": "Bearer %s"}
	loginHeaders      = []map[string]string{contentTypeJson}
	loginWrongHeaders = []map[string]string{contentTypeForm}
	userHeaders       = []map[string]string{jwtAuth}
	userJsonHeaders   = []map[string]string{jwtAuth, contentTypeJson}
)

type TableTest struct {
	url      string
	code     int
	method   string
	headers  []map[string]string
	nextPage string
	err      error
}

var getTests = []TableTest{
	{"/", http.StatusOK, "GET", nil, "", nil},
	{"/swagger.json", http.StatusOK, "GET", nil, "", nil},
	{"/swagger-ui/", http.StatusOK, "GET", nil, "", nil},
	{"/api/_ah/health", http.StatusOK, "GET", nil, "", nil},
}

type LoginAPITest struct {
	TableTest
	email    string
	password string
}

var loginAPITests = []LoginAPITest{
	{TableTest{"/api/auth/login", http.StatusBadRequest, "POST", loginHeaders, "", nil}, "", ""},
	{TableTest{"/api/auth/login", http.StatusBadRequest, "POST", loginHeaders, "", nil}, "hiro", ""},
	{TableTest{"/api/auth/login", http.StatusBadRequest, "POST", loginHeaders, "", nil}, "wrong", "something"},
	{TableTest{"/api/auth/login", http.StatusNotFound, "GET", loginHeaders, "", nil}, "aaaa1@test.jp", "password"},
	{TableTest{"/api/auth/login", http.StatusBadRequest, "POST", nil, "", nil}, "aaaa1@test.jp", "password"},
	{TableTest{"/api/auth/login", http.StatusBadRequest, "POST", loginWrongHeaders, "", nil}, "aaaa1@test.jp", "password"},
	{TableTest{"/api/auth/login", http.StatusOK, "POST", loginHeaders, "", nil}, "aaaa1@test.jp", "password"},
}

type UserAPITest struct {
	TableTest
	email    string
	password string
	userName string
}

var userAPITests = []UserAPITest{
	{TableTest{"/api/user", http.StatusOK, "GET", userHeaders, "", nil}, "", "", ""},
	{TableTest{"/api/user/999999", http.StatusNotFound, "GET", userHeaders, "", nil}, "", "", ""},
	{TableTest{"/api/user/1", http.StatusOK, "GET", userHeaders, "", nil}, "", "", ""},
	{TableTest{"/api/user", http.StatusBadRequest, "POST", userJsonHeaders, "", nil}, "", "", ""},
	//TODO:this test return 200 though http header doesn't include [Content-Type: application/json]
	{TableTest{"/api/user", http.StatusOK, "POST", userHeaders, "", nil},
		"fromtest01@test.com",
		"testtest01",
		"testuser01",
	},
	{TableTest{"/api/user", http.StatusOK, "POST", userJsonHeaders, "", nil},
		"fromtest02@test.com",
		"testtest02",
		"testuser02",
	},
	//{"/api/user/%s", http.StatusOK, "GET", userHeaders, "", nil},
	//{"/api/user/%s", http.StatusOK, "PUT", userHeaders, "", nil},
	//{"/api/user/%s", http.StatusOK, "GET", userHeaders, "", nil},
	//{"/api/user/%s", http.StatusOK, "DELETE", userHeaders, "", nil},
	//{"/api/user/%s", http.StatusOK, "GET", userHeaders, "", nil},
}

//-----------------------------------------------------------------------------
// Test Framework
//-----------------------------------------------------------------------------
// Initialize
func init() {
	lg.InitializeLog(lg.DebugStatus, lg.LogOff, 99, "[GoGOA]", "/var/log/go/test.log")
}

func setup() {
}

func teardown() {
}

func TestMain(m *testing.M) {
	setup()

	code := m.Run()

	teardown()

	os.Exit(code)
}

//-----------------------------------------------------------------------------
// function utility
//-----------------------------------------------------------------------------
// SkipLog is to skip test with func name
func skipLog(t *testing.T) {
	//t.Skip(fmt.Sprintf("skipping %s", r.CurrentFunc(1)))
	t.Skip(fmt.Sprintf("skipping %s", r.CurrentFunc(2)))
}

func convertJson(model interface{}) ([]byte, error) {
	data, err := json.Marshal(model)
	if err != nil {
		return nil, fmt.Errorf("[ERROR] When calling `json.Marshal`: %v\n", err)
	}
	//fmt.Println("[Debug] Json Data:", string(data))

	return data, nil
}

//-----------------------------------------------------------------------------
// function
//-----------------------------------------------------------------------------
// Set HTTP Header
func setHTTPHeaders(req *http.Request, headers []map[string]string) {
	//req.Header.Set("Authorization", "Bearer access-token")
	for _, header := range headers {
		for k, v := range header {
			req.Header.Set(k, v)
		}
	}
}

func searchHTTPHeaders(key string, headers []map[string]string) string {
	//req.Header.Set("Authorization", "Bearer access-token")
	for _, header := range headers {
		for k, v := range header {
			if k == key {
				return v
			}
		}
	}
	return ""
}

func dumpHTTP(req *http.Request) {
	b, err := httputil.DumpRequestOut(req, true)
	if err != nil {
		fmt.Printf("[dumpHTTP] error: %s\n", err)
	} else {
		fmt.Printf("[dumpHTTP] headers:\n%s\n", b)
	}
}

func checkError(t *testing.T, err error, code int, header http.Header, tt TableTest, idx int) {
	//t.Logf("%#v", err)
	//&url.Error{Op:"Get", URL:"/", Err:(*errors.errorString)(0xc8202101b0)}
	if err == nil && tt.err == nil {
		//check expected status code
		if code != tt.code {
			t.Errorf("%d [%s] status code is not correct. \n return code is %d / expected %d", idx, tt.url, code, tt.code)
		}
		return
	} else if err != nil && tt.err == nil || err == nil && tt.err != nil {
		t.Errorf("%d [%s] unexpected error. \n error:%s", idx, tt.url, err)
		return
	}

	//Condtion is both err and tt.err has error object
	urlError, isURLErr := err.(*url.Error)
	if isURLErr && urlError.Err.Error() != tt.err.Error() {
		t.Errorf("%d [%s] this error is not expected. \n error:%s", idx, tt.url, urlError.Err)
		return
	} else {
		//check expected status code
		if code != tt.code {
			t.Errorf("%d [%s] status code is not correct. \n return code is %d / expected %d", idx, tt.url, code, tt.code)
		}
	}

	//check next page
	if isURLErr && urlError.Err.Error() == errRedirect.Error() {
		//t.Log(res.Header["Location"])
		if header == nil {
			t.Errorf("[%s] response header should be set.", tt.url)
		} else if tt.nextPage != header["Location"][0] {
			t.Errorf("[%s] redirect url is not correct. \n url is %s / expected %s", tt.url, header["Location"][0], tt.nextPage)
		}
	}
}

func getJWT(body []byte) (string, error) {
	type ResJWT struct {
		Code  uint8  `json:"code"`
		Token string `json:"token"`
	}
	var jwt ResJWT

	err := json.Unmarshal(body, &jwt)
	if err != nil {
		return "", err
	}
	return jwt.Token, nil
}

func sendRequest(url, method string, bd io.Reader, headers []map[string]string) ([]byte, int, http.Header, error) {

	//1. prepare NewRequest data
	req, err := http.NewRequest(
		method,
		url,
		bd,
	)
	if err != nil {
		return nil, 0, nil, err
	}

	//2. set http header
	//req.Header.Set("Content-Type", "application/json; charset=utf-8")
	if headers != nil {
		setHTTPHeaders(req, headers)
	}

	//dump
	//dumpHTTP(req)

	//3. send
	client := &http.Client{
		Timeout: time.Duration(3) * time.Second,
		CheckRedirect: func(req *http.Request, via []*http.Request) error {
			return errors.New("redirect")
		},
	}

	resp, err := client.Do(req)
	if err != nil {
		return nil, 0, nil, err
	}
	defer resp.Body.Close()

	//4. read response
	body, _ := ioutil.ReadAll(resp.Body)

	return body, resp.StatusCode, resp.Header, nil
}

//-----------------------------------------------------------------------------
// Test
//-----------------------------------------------------------------------------
func TestGetRequestOnTable(t *testing.T) {
	//TODO: this test is run on docker with webserver, so httptest is not used.
	for i, tt := range getTests {
		fmt.Printf("%d [%s] %s\n", i+1, tt.method, SERVER_HOST+tt.url)

		//send request
		_, code, header, err := sendRequest(SERVER_HOST+tt.url, tt.method, nil, tt.headers)
		checkError(t, err, code, header, tt, i+1)
	}
}

func TestLoginOnTable(t *testing.T) {
	//loginAPITests
	var ioReader io.Reader
	for i, tt := range loginAPITests {
		fmt.Printf("%d [%s] %s\n", i+1, tt.method, SERVER_HOST+tt.url)

		//body
		//ioReader = nil
		if searchHTTPHeaders("Content-Type", tt.headers) == "application/json" {
			//json
			//this pattern can't work somehow
			//var jsonData string = `{"email": "%s","password": "%s"}`
			//jsonData = fmt.Sprintf(jsonData, tt.email, tt.password)
			//jsonByte, err := json.Marshal(jsonData)

			loginData := loginAuthPayload{}
			loginData.Email = &tt.email
			loginData.Password = &tt.password
			jsonByte, err := convertJson(&loginData)
			if err != nil {
				t.Fatalf("[%s] json data for request is invalid.", tt.url)
			}
			ioReader = bytes.NewReader(jsonByte)
		} else {
			values := url.Values{}
			values.Set("email", tt.email)
			values.Set("password", tt.password)
			ioReader = strings.NewReader(values.Encode())
		}

		//send request
		body, code, header, err := sendRequest(SERVER_HOST+tt.url, tt.method, ioReader, tt.headers)
		checkError(t, err, code, header, tt.TableTest, i+1)

		//jwt
		if jwtToken == "" && code == 200 {
			jwtToken, err = getJWT(body)
			if err != nil {
				t.Fatalf("[%s] jwt token could not be retrieved from body.", tt.url)
				return
			}
			jwtAuth["Authorization"] = fmt.Sprintf(jwtAuth["Authorization"], jwtToken)
		}
	}
}

func TestUserAPIOnTable(t *testing.T) {
	//skipLog(t)

	//http localhost:8080/api/user 'Authorization: Bearer $(TOKEN)'
	//http localhost:8080/api/user/1 'Authorization: Bearer $(TOKEN)'
	//http POST http://localhost:8080/api/user name=Harry email=test@oo.bb 'Authorization: Bearer $(TOKEN)'
	//http PUT http://localhost:8080/api/user/1 name=Harry email=test@oo.bb 'Authorization: Bearer $(TOKEN)'
	//http DELETE http://localhost:8080/api/user/1 'Authorization: Bearer $(TOKEN)'

	//userAPITests
	var ioReader io.Reader
	for i, tt := range userAPITests {
		fmt.Printf("%d [%s] %s\n", i+1, tt.method, SERVER_HOST+tt.url)
		//send request
		ioReader = nil
		if tt.method == "POST" {
			//json
			userData := createUserHyUserPayload{}
			userData.Email = &tt.email
			userData.Password = &tt.password
			userData.UserName = &tt.userName
			jsonByte, err := convertJson(&userData)
			if err != nil {
				t.Fatalf("[%s] json data for request is invalid.", tt.url)
			}
			ioReader = bytes.NewReader(jsonByte)
		}
		body, code, header, err := sendRequest(SERVER_HOST+tt.url, tt.method, ioReader, tt.headers)
		checkError(t, err, code, header, tt.TableTest, i+1)

		fmt.Println("[Debug] body:", string(body))
	}
}

//-----------------------------------------------------------------------------
// Benchmark
//-----------------------------------------------------------------------------
func BenchmarSomething(b *testing.B) {
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		//
		//_ = CallSomething()
		//
	}
	b.StopTimer()
}
