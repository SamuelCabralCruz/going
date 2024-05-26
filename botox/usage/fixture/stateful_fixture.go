//go:build test

package fixture

import "fmt"

type Stateful struct {
	count int
}

func (s *Stateful) Mutate() {
	s.count += 1
}

func (s *Stateful) Method() string {
	return fmt.Sprintf("stateful - %d", s.count)
}
