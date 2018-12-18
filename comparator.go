package testProj

type Compare func(interface{}, interface{}) int

func AgreaterThanB(f Compare, A interface{}, B interface{}) int {
	return f(A, B)
}


