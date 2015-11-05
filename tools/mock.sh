#!/bin/bash

echo "Mocking mediaos objects ..."

# make sure mockery is installed
go get github.com/vektra/mockery
go install github.com/vektra/mockery

OLDDIR=`pwd`

cd `dirname "${BASH_SOURCE[0]}"`

cd ..

mockery -all -inpkg .

cd $OLDDIR

echo "done."
