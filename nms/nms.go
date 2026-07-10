package nms

import "cmp"

type U interface {
	~uint | ~uint8 | ~uint16 | ~uint32 | ~uint64 | ~uintptr
}

type S interface {
	~int | ~int8 | ~int16 | ~int32 | ~int64
}

type F interface {
	~float32 | ~float64
}

type C interface {
	~complex64 | ~complex128
}

type I interface {
	S | U
}

type O = cmp.Ordered

type N interface {
	I | F | C
}
