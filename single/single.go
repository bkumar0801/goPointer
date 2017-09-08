package single

func GetValueBack(value int)  int {
	var ptr *int
	ptr = &value
	return *ptr
}

func GetOriginalValue(p *int)  int {
	*p = *p * 10
	var ptr *int
	ptr = p
	*ptr = *ptr / 10
	return *p
}

func GetValueRef(val int) *int {
	var i int = val;
	return &i
}

func FunQuiz(p *int)  {
	var value int = 20;
	p = &value;
}
