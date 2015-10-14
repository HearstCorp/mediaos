#!/bin/bash

echo "Deleting mediaos mock packages ..."

OLDDIR=`pwd`

cd `dirname "${BASH_SOURCE[0]}"`/..

echo "removing mocks ..."

rm mock_*.go

cd $OLDDIR

echo "... done."
