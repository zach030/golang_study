package main

import (
	"reflect"
	"testing"
)

func TestSum(t *testing.T){
	nums1 := []int{1, 2, 3, 4}
	nums2 := []int{1, 3, 2, 6}
	nums3 := []int{4, 6, 2, 3}
	expected := []int{10, 12, 15}
	var res []int
	res = append(res, sum(nums1...))
	res = append(res, sum(nums2...))
	res = append(res, sum(nums3...))
	if !reflect.DeepEqual(res, expected){
		t.Errorf("want:%v,actual:%v",expected,res)
	}
}