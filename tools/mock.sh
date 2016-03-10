#!/bin/bash

echo "Mocking mediaos objects ..."

# make sure mockery is installed
go get github.com/vektra/mockery
go install github.com/vektra/mockery

OLDDIR=`pwd`

cd `dirname "${BASH_SOURCE[0]}"`

cd ..

mockery -name '.*' -inpkg .

for dir in $(ls); do
  if [ -d "${dir}" ] && [ "${dir}" != "vendor" ]; then
    cd $dir
    mockery -all -inpkg .
    cd ..
  fi
done

cd $OLDDIR

echo "done."
