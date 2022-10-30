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
echo '[Scenario] Get User List: 200 must be returned'
STATUS=`getStatusWithToken ${ENDPOINT}/user ${TOKEN}`
handleResult 200 $STATUS

#http $(ENDPOINT)/user 'Authorization: Bearer $(TOKEN)'
#http $(ENDPOINT)/user/1 'Authorization: Bearer $(TOKEN)'
#http POST $(ENDPOINT)/user user_name=harry email=newuser01@foo.com password=secret123 'Authorization: Bearer $(TOKEN)'
#http PUT $(ENDPOINT)/user/5 email=updateduser01@bar.com password=secret456 'Authorization: Bearer $(TOKEN)'
#http DELETE $(ENDPOINT)/user/5 'Authorization: Bearer $(TOKEN)'
#http $(ENDPOINT)/user/5 'Authorization: Bearer $(TOKEN)'
