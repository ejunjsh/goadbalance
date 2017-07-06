# goadbalance
[![Build Status](https://travis-ci.org/ejunjsh/goadbalance.svg?branch=master)](https://travis-ci.org/ejunjsh/goadbalance)
[![baby-gopher](https://raw.githubusercontent.com/drnic/babygopher-site/gh-pages/images/babygopher-badge.png)](http://www.babygopher.org)
a simple load balance with go, supports reconnecting to healthy backends while some backends are down,and reactivating some failed backends while they are up.

## install
````
go get github.com/ejunjsh/goadbalance
````

## run
````
$GOPATH/bin/goadbalance -a :8090 -b [backend_ip:port,backend_ip:port1,...]
````
