# goadbalance
[![Build Status](https://travis-ci.org/ejunjsh/goadbalance.svg?branch=master)](https://travis-ci.org/ejunjsh/goadbalance)

[![baby-gopher](https://raw.githubusercontent.com/drnic/babygopher-site/gh-pages/images/babygopher-badge.png)](http://www.babygopher.org)

a simple load balance with go, supports below feature:

* use random and round robin to select backend servers
* health check
* reconnect the failed backend servers.

## install
````
go get github.com/ejunjsh/goadbalance
````

## run
````
$GOPATH/bin/goadbalance -a :8090 -b [backend_ip:port,backend_ip:port1,...]
````
