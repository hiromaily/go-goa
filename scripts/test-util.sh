#!/bin/bash

isTestFailed=0

## Install requirements (curl is required in advance)
function requirement() {
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
}

## Health Check
function healthcheck() {
  echo 'Health Check'
  result=1
  for i in {1..5}
  do
    STATUS=`getStatus GET ${ENDPOINT}/health`
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
}

function getStatus() {
  method=$1
  url=$2
  rtn=$(http --headers --ignore-stdin ${method} ${url} 2>&1 | grep HTTP/ | cut -d ' ' -f 2)
  echo $rtn
}

function getStatusWithToken() {
  method=$1
  url=$2
  token=$3
  rtn=$(http --headers --ignore-stdin -A bearer -a ${token} ${method} ${url} 2>&1 | grep HTTP/ | cut -d ' ' -f 2)
  echo $rtn
}

function getBodyWithToken() {
  method=$1
  url=$2
  token=$3
  json=$(http --ignore-stdin -A bearer -a ${token} ${method} ${url} 2>&1)
  echo $json
}

function handleResult() {
  expectedStatus=$1
  actualStatus=$2
  if [[ "$actualStatus" != $expectedStatus ]]; then
    echo " >>> Fail:  $expectedStatus is expected but $actualStatus returned"
    isTestFailed=1
  else
    echo " >>> OK"
  fi
}

function getToken() {
	#token=$(http --headers POST ${ENDPOINT}/auth/login email=hiroki@goa.com password=password | head -n 2 | tail -n 1 | sed -e "s/Authorization: //g")
  token=$(http --ignore-stdin --body POST ${ENDPOINT}/auth/login email=hiroki@goa.com password=password | jq '.token' | sed 's/"//g')
  echo $token
}

function isEqual() {
  expected=$1
  actual=$2
  #if [ $expected == $actual ]; then
  if [ "$expected" = "$actual" ]; then
    echo " >>> OK"
  else
    echo " >>> Fail   $expected is expected but $actual returned"
    isTestFailed=1
  fi
}