package main_test

import (
	"fmt"
	"io/ioutil"
	"net/http"
	//"net/http/httptest"
	"bytes"
	"errors"
	"net/url"
	"os"
	"testing"
	"time"

	lg "github.com/hiromaily/golibs/log"
)

const SERVER_HOST = "http://localhost:8080"

var (
	errRedirect = errors.New("redirect")
)

type GetTest struct {
	url      string
	code     int
	method   string
	headers  []map[string]string
	nextPage string
	err      error
}

var getTests = []GetTest{
	{"/", http.StatusOK, "GET", nil, "", nil},
	{"/swagger.json", http.StatusOK, "GET", nil, "", nil},
	{"/swagger-ui/", http.StatusOK, "GET", nil, "", nil},
	{"/api/_ah/health", http.StatusOK, "GET", nil, "", nil},
}

//http --body POST http://localhost:8080/api/auth/login username=hiro password=xxxxxxxx

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

func checkError(t *testing.T, err error, code int, header http.Header, tt GetTest) {
	//t.Logf("%#v", err)
	//&url.Error{Op:"Get", URL:"/", Err:(*errors.errorString)(0xc8202101b0)}
	if err == nil && tt.err == nil {
		//check expected status code
		if code != tt.code {
			t.Errorf("[%s] status code is not correct. \n return code is %d / expected %d", tt.url, code, tt.code)
		}
		return
	} else if err != nil && tt.err == nil || err == nil && tt.err != nil {
		t.Errorf("[%s] unexpected error. \n error:%s", tt.url, err)
		return
	}

	//Condtion is both err and tt.err has error object
	urlError, isURLErr := err.(*url.Error)
	if isURLErr && urlError.Err.Error() != tt.err.Error() {
		t.Errorf("[%s] this error is not expected. \n error:%s", tt.url, urlError.Err)
		return
	} else {
		//check expected status code
		if code != tt.code {
			t.Errorf("[%s] status code is not correct. \n return code is %d / expected %d", tt.url, code, tt.code)
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

func sendRequest(url, method string, data []byte, headers []map[string]string) ([]byte, int, http.Header, error) {

	//1. prepare NewRequest data
	req, err := http.NewRequest(
		//"POST",
		method,
		url,
		//bytes.NewBuffer(jsonStr),
		bytes.NewReader(data),
	)
	if err != nil {
		return nil, 0, nil, err
	}

	//2. set http header
	//req.Header.Set("Content-Type", "application/json; charset=utf-8")
	//req.Header.Set("Cookie", fmt.Sprintf("%s=%s", cookieKey, cookie))
	if headers != nil {
		setHTTPHeaders(req, headers)
	}

	//3. send
	//client := &http.Client{}
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
func TestSomething(t *testing.T) {
	//if err != nil {
	//	t.Errorf("TestMain error: %s", err)
	//}
}

func TestGetRequestOnTable(t *testing.T) {
	//TODO: this test is run on docker with webserver, so httptest is not used.
	//request

	for i, tt := range getTests {
		fmt.Printf("%d [%s] %s\n", i+1, tt.method, SERVER_HOST+tt.url)

		//send request
		_, code, header, err := sendRequest(SERVER_HOST+tt.url, tt.method, nil, tt.headers)
		checkError(t, err, code, header, tt)
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
