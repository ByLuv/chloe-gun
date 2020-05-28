package main

// #cgo CFLAGS: -I./srnd
// #cgo LDFLAGS: -L./srnd -lrand
// #include "rand.h"
import "C"

func Random() int {
	return int(C.get_random_number())
}
