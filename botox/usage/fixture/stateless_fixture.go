//go:build test

package fixture

import "time"

func NewStateless() Stateless {
	return Stateless{
		Timestamp: time.Now().Nanosecond(),
	}
}

type Stateless struct {
	Timestamp int
}

func (s Stateless) Method() string {
	return "stateless"
}
