package structs

type Stringify interface {
	GetString() string
}

type A struct {
	Public  int
	Private string
}

func Example() A {
	a := A{}
	a.Private = "test"
	a.Public = 1

	return a
}

func (a *A) GetString() string {
	return a.Private
}
