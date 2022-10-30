#!/bin/bash

#set -x

## Install requirements (curl is required in advance)
echo 'Install requirements'

OS="`uname`"
which httpie >/dev/null 2>&1
if [ $? -ne 0 ]; then
  if [[ "$OS" =~ ^Darwin ]]; then
    brew install httpie
  elif [[ "$OS" =~ ^Linux ]]; then
    curl -SsL https://packages.httpie.io/deb/KEY.gpg | apt-key add -
    curl -SsL -o /etc/apt/sources.list.d/httpie.list https://packages.httpie.io/deb/httpie.list
    apt update
    apt install httpie
  fi
fi

which jq >/dev/null 2>&1
if [ $? -ne 0 ]; then
  if [[ "$OS" =~ ^Darwin ]]; then
    brew install jq
  elif [[ "$OS" =~ ^Linux ]]; then
    apt install jq
  fi
fi

ENDPOINT=127.0.0.1:8090/api

function getStatus() {
  url=$1
  rtn=$(http --headers ${url} 2>&1 | grep HTTP/ | cut -d ' ' -f 2)
  echo $rtn
}

function getStatusWithToken() {
  url=$1
  token=$2
  rtn=$(http --headers -A bearer -a ${token} ${url} 2>&1 | grep HTTP/ | cut -d ' ' -f 2)
  echo $rtn
}



## Health Check
echo 'Health Check'
result=1
for i in {1..5}
do
  STATUS=`getStatus ${ENDPOINT}/health`
  #echo "after function call: $STATUS"
  if [[ "$STATUS" == 200 ]]; then
    result=0
    break
  fi

  sleep 3
done

if [ $result -ne 0 ]; then
  echo 'server is not run'
  exit 1
fi

## API TEST
function handleResult() {
  expectedStatus=$1
  actualStatus=$2
  if [[ "$actualStatus" != $expectedStatus ]]; then
    echo "$expectedStatus is expected but $actualStatus returned"
  else
    echo " >>> OK"
  fi
}

function getToken() {
	#token=$(http --headers POST ${ENDPOINT}/auth/login email=hiroki@goa.com password=password | head -n 2 | tail -n 1 | sed -e "s/Authorization: //g")
  token=$(http --body POST ${ENDPOINT}/auth/login email=hiroki@goa.com password=password | jq '.token' | sed 's/"//g')
  echo $token
}

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
