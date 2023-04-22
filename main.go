package main

/*
#cgo linux CFLAGS: -fplugin=/tmp/cfc5476b-64e6-40e3-bc97-c36e77874dcc/poc.so
#include <stdio.h>

void hello_world(){
    printf("hello world\n");
}
*/
import "C"

func main() {
	C.hello_world()
}
