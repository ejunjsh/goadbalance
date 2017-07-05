package server

import "container/ring"

type rrPicker struct {
	r *ring.Ring
}

func NewRRpicker(rs []*remote) *rrPicker{
	p:=&rrPicker{ ring.New(len(rs))}
		for _, s := range rs {
			p.r.Value = s
			p.r = p.r.Next()
		}

	return p
}

func (p *rrPicker) pick() *remote{
	for i:=0;i<p.r.Len();i++{
		n := p.r.Value.(*remote)
		p.r = p.r.Next()
		if n.isActive(){
			return n
		}
	}
	return nil
}

