#!/bin/bash

#set -x

ENDPOINT=127.0.0.1:8090/api

source ./scripts/test-util.sh

# Check Requirements
requirement

# Check Healthcheck
healthcheck


## API TEST
### Login Error
echo '[Scenario] Login Error due to no password: 400 must be returned'
STATUS=`getStatus "${ENDPOINT}/auth/login email=hiroki@goa.com"`
handleResult 400 $STATUS

echo '[Scenario] Login Error due to wrong password: 400 must be returned'
STATUS=`getStatus "${ENDPOINT}/auth/login email=hiroki@goa.com password=foobar"`
handleResult 400 $STATUS

echo '[Scenario] Success Login: 200 must be returned'
STATUS=`getStatus "${ENDPOINT}/auth/login email=hiroki@goa.com password=password"`
handleResult 200 $STATUS

### Check Token
TOKEN=`getToken`
if [ -z "$TOKEN" ]; then
  echo 'fail to get TOKEN. script needs to be fixed'
  exit 1
fi

### User Test
#### User Status Test
echo '[Scenario] Fail to get User List without Token: 401 must be returned'
STATUS=`getStatus ${ENDPOINT}/user`
handleResult 401 $STATUS

echo '[Scenario] Get User List: 200 must be returned'
STATUS=`getStatusWithToken ${ENDPOINT}/user ${TOKEN}`
handleResult 200 $STATUS

echo '[Scenario] Get User: 200 must be returned'
STATUS=`getStatusWithToken ${ENDPOINT}/user/1 ${TOKEN}`
handleResult 200 $STATUS

echo '[Scenario] Get User: 404 must be returned'
STATUS=`getStatusWithToken ${ENDPOINT}/user/99999 ${TOKEN}`
handleResult 404 $STATUS

#### User Body Test for user
json=`getBodyWithToken ${ENDPOINT}/user/1 ${TOKEN}`
echo '[Scenario] Get User Body: user_id'
userID=$(echo $json | jq '.user_id')
isEqual 1 $userID

echo '[Scenario] Get User Body: user_name'
userName=$(echo $json | jq -r '.user_name')
echo $userName
isEqual "Hiroki Yasui" "$userName"

echo '[Scenario] Get User Body: email'
email=$(echo $json | jq -r '.email')
echo $email
isEqual hiroki@goa.com $email

#### User Body Test for userlist
json=`getBodyWithToken ${ENDPOINT}/user ${TOKEN}`
echo '[Scenario] Get User List Body: Length'
len=$(echo $json | jq '. | length')
isEqual 2 $len

echo '[Scenario] Get User List Body: index:0 user_id'
echo $json
userID=$(echo $json | jq '.[0].user_id')
isEqual 1 $userID



#http $(ENDPOINT)/user 'Authorization: Bearer $(TOKEN)'
#http $(ENDPOINT)/user/1 'Authorization: Bearer $(TOKEN)'
#http POST $(ENDPOINT)/user user_name=harry email=newuser01@foo.com password=secret123 'Authorization: Bearer $(TOKEN)'
#http PUT $(ENDPOINT)/user/5 email=updateduser01@bar.com password=secret456 'Authorization: Bearer $(TOKEN)'
#http DELETE $(ENDPOINT)/user/5 'Authorization: Bearer $(TOKEN)'
#http $(ENDPOINT)/user/5 'Authorization: Bearer $(TOKEN)'

# Result
echo "Result status: $isTestFailed"