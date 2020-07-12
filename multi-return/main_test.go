package main

import (
	"reflect"
	"testing"
)

func TestVals(t *testing.T){
	expected := []int{3,7}
	re1, re2 := vals()
	var res []int
	res = append(res, re1)
	res = append(res, re2)
	if !reflect.DeepEqual(res, expected){
		t.Errorf("want:%v, actual:%v", expected, res)
	}
}