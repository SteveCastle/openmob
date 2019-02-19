#!/bin/bash
if [[ $1 == *".go"* ]]; then
  gofmt -w $1
fi