package server

import (
	"testing"
	"fmt"
	"strconv"
)

func TestNewRRpicker(t *testing.T) {
	rs:=[]*Remote{}
	for i:=0;i<5;i++{
		rs=append(rs,&Remote{Address:strconv.Itoa(i)})
	}
	p:=NewRRpicker(rs)
	for a:=0;a<100;a++{
		fmt.Println(p.Pick().Address)
	}
}
