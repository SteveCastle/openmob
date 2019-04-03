#!/bin/bash
if [[ $1 == *".go"* ]]; then
  gofmt -w $1
  goimports -w $1
fi