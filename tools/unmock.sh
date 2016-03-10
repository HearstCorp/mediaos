#!/bin/bash

echo "Deleting mediaos mock packages ..."

OLDDIR=`pwd`

cd `dirname "${BASH_SOURCE[0]}"`/..

echo "removing mocks ..."

find . -type f -name 'mock_*.go' -delete

cd $OLDDIR

echo "... done."
