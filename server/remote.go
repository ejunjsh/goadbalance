package server

import (
	"sync"
	"net"
)

type remote struct {
	locker   sync.Mutex
	address  string
	inactive bool
}

func (r *remote) inactivate()  {
	r.locker.Lock()
	defer r.locker.Unlock()
	r.inactive=true
}

func (r *remote) tryReactivate() error {
	conn,err:=net.Dial("tcp",r.address)
	if err!=nil{
		return err
	}
	conn.Close()
	r.locker.Lock()
	defer r.locker.Unlock()
	r.inactive=false
	return  nil
}

func (r *remote) isActive() bool{
	r.locker.Lock()
	defer r.locker.Unlock()
	return !r.inactive
}
