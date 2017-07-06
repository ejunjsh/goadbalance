# goadbalance
[![Build Status](https://travis-ci.org/ejunjsh/goadbalance.svg?branch=master)](https://travis-ci.org/ejunjsh/goadbalance)

a simple load balance with go

## install
````
go get github.com/ejunjsh/goadbalance
````

## run
````
$GOPATH/bin/goadbalance -a :8090 -b [backend_ip:port,backend_ip:port1,...]
````
