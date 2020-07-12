package main

import (
	"reflect"
	"testing"
)

func TestAdd(t *testing.T) {
	expected := 9
	if res:= Add([]int{2, 3, 4}); !reflect.DeepEqual(res, expected){
		t.Errorf("want: %v, actual: %v", expected, res)
	}
}