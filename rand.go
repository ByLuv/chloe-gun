package main

/*
#include <stdlib.h>
*/
import "C"

func Seed(i int) {
	C.srandom(C.uint(i))
}

func Random() int {
	return int(C.random())
}
