package server

import (
	"math/rand"
	"time"
)

type ramdomPicker struct {
	random *rand.Rand
	remotes []*Remote
}

func NewRandomPicker(rs []*Remote) RemotePiker {
	p := &ramdomPicker{rand.New(rand.NewSource(time.Now().UnixNano())),rs}
	return p
}

func (p *ramdomPicker) Pick() *Remote {
    rs:=getActive(p.remotes)
	i:=p.random.Intn(len(rs))
	return rs[i]
}

func getActive(rs []*Remote) []*Remote {
	result:=[]*Remote{}
    for _,r:=range rs{
        if r.isActive(){
			result=append(result, r)
		}
     }

	return result
}
