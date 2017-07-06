package server

import (
	"container/ring"
)

type rrPicker struct {
	r *ring.Ring
}


func NewRRpicker(rs []*Remote) RemotePiker{
	p:=&rrPicker{ ring.New(len(rs))}
		for _, s := range rs {
			p.r.Value = s
			p.r = p.r.Next()
		}

	return p
}

func (p *rrPicker) Pick() *Remote{
	for i:=0;i<p.r.Len();i++{
		n := p.r.Value.(*Remote)
		p.r = p.r.Next()
		if n.isActive(){
			return n
		}
	}
	return nil
}

