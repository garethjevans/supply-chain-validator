#!/bin/bash

#set -euo pipefail

make build
echo "================================================="

echo "Validate"
ytt -f ./config | ./build/scv validate -
echo "================================================="

echo "Diff"
ytt -f ./config | kubectl diff -f -
echo "================================================="

echo "Apply"
ytt -f ./config | kubectl apply -f -
echo "================================================="
