#!/bin/bash

COMMIT=$(git rev-list HEAD | wc -l)
BRANCH=$(git rev-parse --abbrev-ref HEAD | sed 's/\//-/g')
VERSION="0.1"

if [[ "$BRANCH" == "master" ]]; then
    echo $VERSION.$COMMIT
    exit
fi

echo $VERSION.$COMMIT-$BRANCH