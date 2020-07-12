package main

import (
	"reflect"
	"testing"
)

func TestZeroPtr(t *testing.T){
	expected := []int{2, 3, 0}
	var res []int
	i := 2
	zeroval(i)
	res = append(res, i)
	j := 3
	zeroval(j)
	res = append(res, j)
	temp := 2
	res = append(res, zeroptr(&temp))
	if !reflect.DeepEqual(res, expected){
		t.Errorf("want:%v,actual:%v", expected, res)
	}
}