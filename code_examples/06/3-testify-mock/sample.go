package testify

type MyService interface {
	DoSomething(int) (bool, error)
}

func targetFuncThatDoesSomethingWithObj(srv MyService) (bool, error) {
	return srv.DoSomething(123)
}
