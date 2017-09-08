package main

import (
	"fmt"
	"unsafe"
)

func main() {
	var ptr *int
	var x int

	ptr = &x
	*ptr = 0

	fmt.Println(" x = ", x)
	fmt.Println(" *ptr = ", *ptr)

	*ptr += 5
	fmt.Println(" x = ", x)
	fmt.Println(" *ptr = ", *ptr)

	(*ptr)++
	fmt.Println(" x = ", x)
	fmt.Println(" *ptr = ", *ptr)
	x++
	fmt.Println(" x = ", x)
	fmt.Println(" *ptr = ", *ptr)

	var fun string = "funQuiz"
	var funPtr *string = &fun
	fmt.Println("funPtr =", *funPtr)

	yo := string([]rune(fun)[2])
	var xptr *string = &yo
	fmt.Println("yo =", *xptr)

	primes := [6]int{2, 3, 5, 7, 11, 13}
	var p *int = &primes[0];
	fmt.Println("p =", *p)

	*p = 67
	var p2 *[6]int = &primes;
	fmt.Println("p2 =", *p2) //

	//p = p + 1 No pointer arithmetic is supported
	intArray := [...]int{1, 2}

	fmt.Printf("intArray: %v", intArray)

	intPtr := &intArray[0]
	fmt.Printf("intPtr=%p, *intPtr=%d.", intPtr, *intPtr) //add, 1

	addressHolder := uintptr(unsafe.Pointer(intPtr)) + unsafe.Sizeof(intArray[0])

	intPtr = (*int)(unsafe.Pointer(addressHolder))

	fmt.Printf("intPtr=%p, *intPtr=%d.", intPtr, *intPtr)

	h := make([]int, 4)
	h = append(h, 5)
	fmt.Println("h= ", h)
	var hptr *int = &h[3]
	fmt.Println("hptr =", *hptr)
	var dptr *[]int = &h;
	*hptr = 40
	fmt.Println("dptr =", *dptr)

	type Vertex struct {
		X int
		Y int
	}

	v := Vertex{1, 2}
	var pVertex *Vertex
	pVertex = &v
	pVertex.X = 200
	fmt.Println("Vertex =", v)
	fmt.Println("Vertex X = ", pVertex.X)
	var vptr *int = &v.X
	*vptr = 10
	fmt.Println("Vertex again = ", v)
	*pVertex = Vertex{20, 30}
	fmt.Println("Vertex reloaded = ", v)

	type Person struct {
		name   string
		age    int
		gender bool
	}

	who := Person{"John", 30, true}
	pp := unsafe.Pointer(&who)
	pname := (*string)(unsafe.Pointer(uintptr(pp) + unsafe.Offsetof(who.name)))
	page := (*int)(unsafe.Pointer(uintptr(pp) + unsafe.Offsetof(who.age)))
	pgender := (*bool)(unsafe.Pointer(uintptr(pp) + unsafe.Offsetof(who.gender)))
	*pname = "Alice"
	*page = 28
	*pgender = false
	fmt.Println("who =", who)
}
