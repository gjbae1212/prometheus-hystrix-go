#!/bin/bash
set -e -o pipefail
trap '[ "$?" -eq 0 ] || echo "Error Line:<$LINENO> Error Function:<${FUNCNAME}>"' EXIT
cd `dirname $0`
CURRENT=`pwd`

function test
{
   go test -v $(go list ./... | grep -v vendor | grep -v test) --count 1
}

function check_heap
{
  # analysis `escape to heap`
  go build -gcflags '-m -m' $CURRENT/prometheus_collector.go
}


function test_server
{
  go build $CURRENT/cmd/test/main.go && $CURRENT/main
}

CMD=$1
shift
$CMD $*
