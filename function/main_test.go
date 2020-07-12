package main

import (
	"reflect"
	"testing"
)

func TestPlus(t *testing.T){
	expected := []int{4, 6, 8, 10, 12, 16}
	var res []int
	res = append(res, plus(1,3))
	res = append(res, plus(3,3))
	res = append(res, plus(5,3))
	res = append(res, plus(4,6))
	res = append(res, plusPlus(3,5,4))
	res = append(res, plusPlus(4,7,5))
	if !reflect.DeepEqual(expected, res){
		t.Errorf("want: %v, actual: %v", expected,res)
	}
}