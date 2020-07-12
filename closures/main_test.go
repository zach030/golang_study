package main

import (
	"reflect"
	"testing"
)

func TestIntSeq(t *testing.T){
	expected := []int{1, 2, 3, 1}
	var res []int
	nextInt := intSeq()
	res = append(res, nextInt())
	res = append(res, nextInt())
	res = append(res, nextInt())
	nextInts := intSeq()
	res = append(res, nextInts())
	if !reflect.DeepEqual(res, expected){
		t.Errorf("want:%v,actual:%v",expected,res)
	}
}