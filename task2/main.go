package main

import (
	"fmt"
	"math"
	"sync"
	"sync/atomic"
	"time"
)

//指针 题目1

func addTen(n *int) {
	*n += 10
}

// 指针 题目2
func elementMultiplyTwo(s []int) {
	for i, _ := range s {
		s[i] *= 2
	}
}

//goroutine 题目1

func printNumParallel(n int) {
	go func() {
		for i := 1; i <= 10; i += 2 {
			fmt.Println(i)
		}
	}()

	go func() {
		for i := 0; i <= 10; i += 2 {
			fmt.Println(i)
		}
	}()
}

// goroutine 题目2
func goExecutor(n int) {
	//统计执行时间
	t := time.Now().UnixMilli()
	time.Sleep(time.Millisecond * time.Duration(n*10))
	fmt.Println("task", n, " running in", (time.Now().UnixMilli() - t), "ms")

}

func dispatch() {
	for i := 0; i < 10; i++ {
		go goExecutor(i)
	}
}

// 面向对象 题目1
type Shape interface {
	Area() float64
	Perimeter() float64
}

type Rectangle struct {
	width  float64
	height float64
}

func (r *Rectangle) Area() float64 {
	return r.width * r.height
}

func (r *Rectangle) Perimeter() float64 {
	return 2 * (r.width + r.height)
}

type Circle struct {
	radius float64
}

func (c *Circle) Area() float64 {
	return math.Pi * math.Pow(c.radius, 2)
}

func (c *Circle) Perimeter() float64 {
	return 2 * math.Pi * c.radius
}

// 面向对象 题目2
type Person struct {
	Name string
	Age  int
}

type Employee struct {
	Person
	EmployeeID string
}

func (e *Employee) PrintInfo() {
	fmt.Printf("Name:%s, Age:%d, EmployeeID:%s\n", e.Name, e.Age, e.EmployeeID)
}

// channel 題目1
func chTest() {
	ch := make(chan int)
	go func() {
		for i := range ch {
			fmt.Println("receive:", i)
		}
		fmt.Println("channel closed")
	}()

	for i := 1; i <= 10; i++ {
		ch <- i
	}
	close(ch)
	time.Sleep(1 * time.Second)
}

//channel 題目2

func receive(ch <-chan int) {
	t := time.After(5 * time.Second)
	for {
		select {
		case v := <-ch:
			fmt.Println("receive:", v)
		case <-t:
			fmt.Println("timeout")
			break
			//default:
			//	fmt.Println("receive no data")
		}
	}
}

func send() {
	//創建channel，創建緩冲区
	ch := make(chan int, 10)
	go receive(ch)
	for i := 1; i <= 100; i++ {
		fmt.Println("send:", i, "before")
		ch <- i
		fmt.Println("send:", i, "after")
	}
	time.Sleep(10 * time.Second)
}

//锁，题目1
func lockTest1() {
	counter := 0
	m := sync.Mutex{}

	for i := 0; i < 10; i++ {
		go func() {
			for j := 0; j < 1000; j++ {
				m.Lock()
				counter++
				m.Unlock()
			}
		}()
	}
	time.Sleep(2 * time.Second)
	fmt.Println(counter)
}

//锁，题目2
func lockTest2() {
	counter := atomic.Int32{}
	fmt.Println("init counter:", counter.Load())
	for i := 0; i < 10; i++ {
		go func() {
			for j := 0; j < 1000; j++ {
				counter.Add(1)
			}
		}()
	}
	time.Sleep(2 * time.Second)
	fmt.Println("final counter:", counter.Load())
}

func main() {
	lockTest2()
}
