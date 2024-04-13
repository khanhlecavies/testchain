#!/usr/bin/env bash

set -eo pipefail

cd proto

proto_dirs=$(find ./testchain -path . -prune -o -name '*.proto' -print0 | xargs -0 -n1 dirname | sort | uniq)
for dir in $proto_dirs; do
  for file in $(find "${dir}" -maxdepth 1 -name '*.proto'); do
    if grep "option go_package" $file &> /dev/null ; then
      buf generate $file --template buf.gen.gogo.yaml 
    fi
  done
done

cd ..
# move proto files to the right places
# cp -r github.com/Cosmos-SDK-Learning-101/testchain/* ./
# rm -rf github.com