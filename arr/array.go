package arr

import "fmt"

func DoubleFib(fib *[6]int)  {
	for i:=0; i < 6; i++ {
		var p *int = &fib[i]
		*p = *p*2
	}
}

func StringManip(fib *[]string)  {
	for i:=0; i < 6; i++ {
		fmt.Println(fib)
	}
}
