#!/usr/bin/env sh

set -x

curl -X POST 127.0.0.1:8080 -d '{"metadata":"some metadata", "value": "Hello world!"}'
