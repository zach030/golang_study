package main

import "testing"

func TestIf(t *testing.T){
	res := ifelse(6)
	if res == true{
		t.Log("The test is correct!")
	}else {
		t.Fatal("The result is wrong!")
	}
}
