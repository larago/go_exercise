package main 

import (
	"fmt"
)

type Log struct {
	msg string
}

type Customer struct {
	Name string
	log *Log
}

func main() {
	c := new(Customer)
	c.Name = "Vara"
	c.log = new(Log)
	c.log.msg = "1 - Yes we can!"
	// shorter
	c = &Customer{"Vara", &Log{"1 - Yes we can!"}}
	c.Log().Add("2 - after me the world will be a better place!")
	fmt.Println(c.Log())
}

func (l *Log) Add(s string) {
	l.msg += "\n" + s
}

func (l *Log) String() string {
	return l.msg
}

func (c * Customer) Log() *Log {
	return c.log
}