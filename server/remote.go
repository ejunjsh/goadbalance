package server

import (
	"sync"
	"net"
)

type Remote struct {
	locker   sync.Mutex
	Address  string
	inactive bool
}

func (r *Remote) inactivate()  {
	r.locker.Lock()
	defer r.locker.Unlock()
	r.inactive=true
}

func (r *Remote) tryReactivate() error {
	conn,err:=net.Dial("tcp",r.Address)
	if err!=nil{
		return err
	}
	conn.Close()
	r.locker.Lock()
	defer r.locker.Unlock()
	r.inactive=false
	return  nil
}

func (r *Remote) isActive() bool{
	r.locker.Lock()
	defer r.locker.Unlock()
	return !r.inactive
}
