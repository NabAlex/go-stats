#!/bin/bash

# start ab tests

server="127.0.0.1"
port="9090"

requests="10000"

ab -n $requests -c 2 http://$server:$port/token
