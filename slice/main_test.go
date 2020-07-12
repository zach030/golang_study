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

func TestInterceptslice(t *testing.T) {
	expected := []string{"hi", "go"}
	str := []string{"hi", "go", "hello", "!"}
	if res := Interceptslice(str, 0, 2); !reflect.DeepEqual(res,expected){
		t.Errorf("Want: %v, actual: %v", expected, res)
	}
}