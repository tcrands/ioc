package main

import "fmt"

type Random struct {
	val int
	s   string
}

func (r *Random) getVal() int {
	return r.val
}

func main() {
	ioc := NewIOC()

	thing := Random{val: 12}

	ioc.Register("randomStruct1", thing, 0)

	obj, _ := ioc.Resolve("randomStruct1")
	randomObj := obj.(Random)
	fmt.Println(randomObj.getVal())
	randomObj.val = 14
	fmt.Println(randomObj.getVal())

	ioc.Register("randomStruct2", thing, 1)

	obj2, _ := ioc.Resolve("randomStruct2")
	fmt.Println(obj2)
}
