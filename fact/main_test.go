package main

import (
	"reflect"
	"testing"
)

func TestFact(t *testing.T){
	var expected []int
	expected = append(expected, 6)
	expected = append(expected, 24)
	expected = append(expected, 120)
	var res []int
	res = append(res, fact(3))
	res = append(res, fact(4))
	res = append(res, fact(5))
	if !reflect.DeepEqual(res,expected){
		t.Errorf("want:%v,actual:%v", expected, res)
	}
}