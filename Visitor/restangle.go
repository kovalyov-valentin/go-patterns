package main 

type restangle struct {
	l int 
	b int 
}

func (r *restangle) accept(v visitor) {
	v.visitForRestangle(r)
}

func (r *restangle) getType() string {
	return "Restangle"
}