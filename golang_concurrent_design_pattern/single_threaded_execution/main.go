package main

import (
	"fmt"
	"sync"
	"time"
)

type Gate struct {
	count   int
	name    string
	address string
	m       *sync.Mutex
}

func (self *Gate) Pass(name string, address string) {
	// 加锁即可保障不会有竞争的情况出现
	self.m.Lock()
	self.count++
	self.name = name
	self.address = address
	self.check()
	self.m.Unlock()
}

func (self *Gate) toString() string {
	return fmt.Sprintf("No %d %s %s", self.count, self.address, self.name)
}

func (self *Gate) check() {
	// 延迟执行 捕获异常
	defer func() {
		if err := recover(); err != nil {
			fmt.Println(err)
		}
	}()
	if self.name[0] != self.address[0] {
		// 用出错的形式报出
		panic(fmt.Sprintf("*** BROKEN *** %s", self.toString()))
	}
}

func run(gate *Gate, name string, address string) {
	fmt.Printf("start ==> %s %s\n", name, address)
	for {
		gate.Pass(name, address)
	}
}

func main() {
	// 最好用工厂模式 New***** 获取，这里偷懒直接获取使用
	gate := &Gate{
		m: new(sync.Mutex),
	}

	go run(gate, "Alice", "Alaska")
	go run(gate, "Bobby", "Brazil")
	go run(gate, "Chris", "Canada")

	time.Sleep(1000 * time.Second)
}