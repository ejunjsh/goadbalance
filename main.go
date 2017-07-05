package main

import (
	"flag"
	"strings"
	"github.com/ejunjsh/goadbalance/server"
	"os"
	"os/signal"
	"syscall"
	"net"
)

func main()  {
	var(
		serverAdr  string
		backends string
	)

	flag.StringVar(&serverAdr, "addr", "", "Network host to listen on.")
	flag.StringVar(&serverAdr, "a", "", "Network host to listen on.")
	flag.StringVar(&backends, "b", "", "List of backends")
	flag.StringVar(&backends, "backends", "", "List of backends")

	flag.Parse()

	a:=strings.Split(backends,",")
	l,err:=net.Listen("tcp",serverAdr)
	if err!=nil{
		return
	}
	s:= server.NewServer(l,a...)

	go func() {
		for{
			c := make(chan os.Signal)

			signal.Notify(c)
			sig := <-c

			if sig==syscall.SIGKILL||sig==syscall.SIGTERM{
				s.Stop()
			}
		}
	}()
	s.Run()
}
