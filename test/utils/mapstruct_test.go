package utils

import (
	"rd-assist/utils"
	"testing"
)

func TestFill(t *testing.T) {
	type src struct {
		Id int
		Cname string
		Age int
	}
	type dst struct {
		Id int
		Name string
		Age int
	}
	src1 := src{Id:1,Cname:"ss",Age:3}
	var dst1 dst
	err := utils.Fill(src1, &dst1)
	if  err!=nil || dst1.Age != 3 {
		t.Errorf("Fill() error = %v, wantErr %v", err, dst1)
	}
}
