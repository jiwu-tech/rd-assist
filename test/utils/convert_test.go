package utils

import (
	"rd-assist/utils"
	"testing"
)

func TestStr2Int(t *testing.T) {
	if got := utils.Str2Int("11121"); got != 11121 {
		t.Errorf("Str2Int() = %v, want %v", got, "11121" )
	}
	t.Log("Str2Int pass")
}
func TestStr2Int64(t *testing.T) {
	if got := utils.Str2Int64("11121"); got != 11121 {
		t.Errorf("Str2Int64() = %v, want %v", got, "11121" )
	}
	t.Log("Str2Int64 pass")
}
func TestStr2Bool(t *testing.T) {
	if got := utils.Str2Bool("true"); got != true {
		t.Errorf("Str2Bool() = %v, want %v", got, "1" )
	}
	if got := utils.Str2Bool("false"); got != false {
		t.Errorf("Str2Bool() = %v, want %v", got, "0" )
	}
	t.Log("Str2Bool pass")
}
func TestStr2Float32(t *testing.T) {
	if got := utils.Str2Float32("1.22"); got != 1.22 {
		t.Errorf("Str2Float32() = %v, want %v", got, "1.22" )
	}
	t.Log("Str2Float32 pass")
}