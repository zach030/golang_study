package main

import (
	"reflect"
	"testing"
)

func TestInitslice(t *testing.T){

	expected := []string{"hi"}
	if res := initslice(1,"hi"); !reflect.DeepEqual(res, expected){
		t.Errorf("slice: %v, actual slice : %v", expected, res)
	}
}
