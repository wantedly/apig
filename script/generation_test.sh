#!/bin/bash

set -eu

echo "===> Generating API server..."
cd _example
../bin/apig gen --all

if [[ ! $(git status | grep 'nothing to commit') ]]; then
  echo " x Generator artifact and example application are different."
  git --no-pager diff
  exit 1
fi

echo "===> Building API server..."
go get ./...
go build

if [[ $? -gt 0 ]]; then
  echo " x Failed to build generated API server."
  exit 1
fi

echo " o Generation test PASSED!"
q
