# goadbalance
[![Build Status](https://travis-ci.org/ejunjsh/goadbalance.svg?branch=master)](https://travis-ci.org/ejunjsh/goadbalance)

## install
````
go get github.com/ejunjsh/goadbalance
````

## run
````
$GOPATH/bin/goadbalance -a :8090 -b "[backend_ip:port,backend_ip:port1,...]"
````
for example : $GOPATH/bin/goadbalance -a :8090 -b "127.0.0.1:8100,127.0.0.1:8101,127.0.0.1:8102"