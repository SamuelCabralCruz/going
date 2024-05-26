package pkg

type Inter interface {
	Hello(string) (string, error)
	Hello2(string, int) (string, error, int)
	Prepare() Another
}

type Another interface {
	Bye(string)
}
