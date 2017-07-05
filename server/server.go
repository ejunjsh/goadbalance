package server

import (
	"net"
	"time"
	"sync"
	"io"
)

type Server struct {
	listener   net.Listener
	HealthChkInterval time.Duration

	done chan struct{}

	locker sync.Mutex

	remotes []*remote

	piker remotePiker
}

type remotePiker interface {
	 pick() *remote
}



func NewServer(l net.Listener,adresses ...string) *Server {
	s:=&Server{}
	s.listener=l
	rs:=make([]*remote,0)
	for _,adr:=range adresses{
		r:=remote{}
		r.address=adr
		rs=append(rs, &r)
	}
	s.done=make(chan struct{})
	if s.HealthChkInterval==0{
		s.HealthChkInterval=1*time.Minute
	}

	return s
}

func (s *Server) Run() {
	go s.healthCheck()

	for{
		in,err:=s.listener.Accept()
		if err==nil{
			go s.serve(in)
		}
	}
}

func (s *Server) serve(in net.Conn){
    var(
		err error
		out net.Conn
	)

	for{
		s.locker.Lock()
        r:=s.piker.pick()
		s.locker.Unlock()
		if r==nil{
			break
		}

		out,err=net.Dial("tcp",r.address)

		if err==nil{
			break
		}

		r.inactivate()
	}

	if out ==nil{
		in.Close()
        return
	}

	go func() {
		io.Copy(in,out)
		in.Close()
		out.Close()
	}()

	io.Copy(out,in)
	out.Close()
	in.Close()
}

func (s *Server) healthCheck(){
    for{
		select {
		    case <-time.After(s.HealthChkInterval):
			         s.locker.Lock()
				for _, rem := range s.remotes {
					if rem.isActive() {
						continue
					}
					go func(r *remote) {
						r.tryReactivate()
					}(rem)
				}
				s.locker.Unlock()
		case <-s.done:
			return
		}
	}
}

func (s *Server) Stop() {
	s.listener.Close()
	close(s.done)
}