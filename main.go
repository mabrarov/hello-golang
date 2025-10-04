package main

import (
	"fmt"
	"runtime"
	"weak"
)

func main() {
	var t weak.Pointer[Obj]
	var p *int
	{
		obj := &Obj{
			Name: "hello",
			Age:  20,
		}
		p = &obj.Age
		t = weak.Make(obj)
		obj = nil
		runtime.GC()
	}

	fmt.Printf("Pointer to struct member: %p -> %v\n", p, *p)

	obj := t.Value()
	if obj == nil {
		fmt.Printf("Pointer to struct: %p\n", obj)
	} else {
		fmt.Printf("Pointer to struct: %p -> %+v\n", obj, *obj)
	}
}

type Obj struct {
	Name string
	Age  int
}
