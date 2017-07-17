package server

import (
	"testing"
	"fmt"
	"strconv"
)

func TestNewRandomPicker(t *testing.T) {
	fmt.Println("random start")
	rs:=[]*Remote{}
	for i:=0;i<5;i++{
		rs=append(rs,&Remote{Address:strconv.Itoa(i)})
	}
	p:=NewRandomPicker(rs)
	for a:=0;a<100;a++{
		fmt.Println(p.Pick().Address)
	}
	fmt.Println("random end")
}
