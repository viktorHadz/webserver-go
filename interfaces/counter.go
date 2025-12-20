package interfaces

type Incrementor interface {
	Inc() // Implement using a pointer receiver 
}

type Counter struct {
	Value int
}
type Points struct { // example ... 
	score int
}

func (c *Counter) Inc() int {
	inc := c.Value
	inc++
	return inc
}

func IncGeneric(i Incrementor) {
	
}