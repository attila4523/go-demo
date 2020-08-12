package main

// import "fmt"
import (
    "fmt"
    hello "github.com/attila4523/go-demo-module"
)

func main() {
    var test int = 123
    f := "hello"
    fmt.Println("hello world", test, "asd", f)
    demo()

    hello.HelloWorld()
}
