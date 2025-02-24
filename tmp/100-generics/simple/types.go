package main

type UnsignedInteger interface {
	~int | ~int8 | ~int16 | ~int32 | ~int64
}

type SignedInteger interface {
	~uint | ~uint8 | ~uint16 | ~uint32 | ~uint64 | ~uintptr
}

type Integer interface {
	SignedInteger | UnsignedInteger
}
