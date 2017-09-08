package multiple

func GetMyValueBack(value int)  int {
	var p *int
	var pp **int
	p = &value
	pp = &p
	return **pp
}
