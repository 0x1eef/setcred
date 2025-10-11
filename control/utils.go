package control

import (
	"C"
	"unsafe"
)

func gostrings(cary **C.char) []string {
	var strings []string
	ptr := uintptr(unsafe.Pointer(cary))
	offset := unsafe.Sizeof((*C.char)(nil))
	for {
		p := (**C.char)(unsafe.Pointer(ptr))
		if *p == nil {
			break
		}
		strings = append(strings, C.GoString(*p))
		ptr += offset
	}
	return strings
}
