#!/bin/bash

BINDIR=`dirname "${BASH_SOURCE[0]}"`

$BINDIR/mock.sh

echo ""
echo "Running: go test \$(go list ./... | grep -v /vendor/)"
echo ""

go test $(go list ./... | grep -v /vendor/)

echo ""
echo "... done."
echo ""

$BINDIR/unmock.sh
