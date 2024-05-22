package pkg

import "fmt"

type Impl struct{}

var _ Inter = Impl{}

func (i Impl) Hello(a string) (string, error) {
	return "Hello " + a, nil
}

func (i Impl) Prepare() Another {
	return AnotherImpl{}
}

type AnotherImpl struct{}

var _ Another = AnotherImpl{}

func (a AnotherImpl) Bye(v string) {
	fmt.Println("Bye " + v)
}
