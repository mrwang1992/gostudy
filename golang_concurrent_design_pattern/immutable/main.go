package main

import (
	"fmt"
	"time"
)

type Person struct {
	name    string
	address string
}

func (self *Person) Name() string {
	return self.name
}

func (self *Person) Address() string {
	return self.address
}

func (self *Person) ToString() string {
	return fmt.Sprintf("[Person: name = %s, address = %s ] \n", self.name, self.address)
}

func NewPerson(name string, address string) *Person {
	return &Person{
		name:    name,
		address: address,
	}
}

func PrintPersonThread(person *Person, id int) {
	for {
		// 官方可以无法暴露 goroutine id 使用外部id代入
		fmt.Printf("Goroutine id: %02d => %s", id, person.ToString())
	}
}

func main() {
	person := NewPerson("Alice", "Alaska")

	for i := 0; i < 3; i++ {
		go PrintPersonThread(person, i)
	}

	time.Sleep(1 * time.Second)
}
