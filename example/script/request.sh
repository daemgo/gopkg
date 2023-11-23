#!/bin/zsh

curl -v 127.0.0.1:3000/api/v1/user

curl -v -X PUT 127.0.0.1:3000/api/v1/user/1024
