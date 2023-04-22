package main

/*
#cgo linux CFLAGS: -fplugin=./poc.so
#include <stdio.h>

void hello_world(){
    printf("hello world\n");
}
*/
import "C"

func main() {
	C.hello_world()
}
