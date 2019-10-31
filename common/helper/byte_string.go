package helper

import (
	"reflect"
	"unsafe"
)

//StringHeader String Header
type StringHeader struct {
	Data uintptr
	Len  int
}

//SliceHeader Slice Header
type SliceHeader struct {
	Data uintptr
	Len  int
	Cap  int
}

//B2S []byte -> string
func B2S(b []byte) string {
	return *(*string)(unsafe.Pointer(&b))
}

//S2B string -> []byte 转换出来的[]byte不能有修改操作
func S2B(s string) []byte {
	sh := (*reflect.StringHeader)(unsafe.Pointer(&s))
	bh := reflect.SliceHeader{
		Data: sh.Data,
		Len:  sh.Len,
		Cap:  sh.Len,
	}
	return *(*[]byte)(unsafe.Pointer(&bh))
}
