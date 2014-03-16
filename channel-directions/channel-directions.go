// 当使用同通道最为函数的参数时，你可以指定一个通道是不是
// 只用来发送或者接收值的。这个特性提升了程序的类型安全性。

package main

import "fmt"

// `ping` 函数只允许一个用来发送数据的通道。尝试使用这个通
// 道来接受数据将会得到一个编译时错误。
func ping(pings chan<- string, msg string) {
    pings <- msg
}

// `pong` 函数允许通道（`pings`）来接收数据，另一通道
// （`pongs`）来发送数据。
func pong(pings <-chan string, pongs chan<- string) {
    msg := <-pings
    pongs <- msg
}

func main() {
    pings := make(chan string, 1)
    pongs := make(chan string, 1)
    ping(pings, "passed message")
    pong(pings, pongs)
    fmt.Println(<-pongs)
}
