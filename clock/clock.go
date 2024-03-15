package clock

import (
	"fmt"
	"time"
)

type Clock struct {
	h int
	m int
}

func New(h, m int) Clock {
	now := time.Date(0, 0, 0, h, m, 0, 0, time.Now().Location())

	return Clock{now.Hour(), now.Minute()}
}

func (c Clock) Add(m int) Clock {
	now := time.Date(0, 0, 0, c.h, c.m, 0, 0, time.Now().Location())
	added := now.Add(time.Minute * time.Duration(m))
	return New(added.Hour(), added.Minute())
}

func (c Clock) Subtract(m int) Clock {
	return c.Add(-m)
}

func (c Clock) String() string {
	return fmt.Sprintf("%.2d:%.2d", c.h, c.m)
}
