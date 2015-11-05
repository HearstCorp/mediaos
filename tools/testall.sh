#!/bin/bash

BINDIR=`dirname "${BASH_SOURCE[0]}"`

$BINDIR/mock.sh

echo ""
echo "Running go test github.com/Hearst-DD/mediaos/..."
echo ""
go test github.com/Hearst-DD/mediaos/...
echo ""
echo "... done."
echo ""

$BINDIR/unmock.sh
