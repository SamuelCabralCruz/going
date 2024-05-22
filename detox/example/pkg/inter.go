package pkg

type Inter interface {
	Hello(string) (string, error)
	Prepare() Another
}

type Another interface {
	Bye(string)
}
