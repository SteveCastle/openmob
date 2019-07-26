#!/bin/bash
echo "hello"
sed -i -e '/CUSTOMMESSAGES/r shrike_messages.proto' $1