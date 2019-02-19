#!/bin/bash
if [ ! -e /usr/local/bin/protoc ]; then
curl -OL https://github.com/google/protobuf/releases/download/v3.3.0/$(PROTOC_ZIP)
sudo unzip -o $(PROTOC_ZIP) -d /usr/local bin/protoc
rm -f $(PROTOC_ZIP)
fi