package main

import (
    "fmt"
    DEPLOY_imports
)

func someFunction1(a, b int) int {
    return a + b
}

func someOtherFunction(a, b int, f func(int, int) int) int {
    return f(a, b)
}

func mapExample(){
    fmt.Println("mapExample:")
    m := map[string]func(int, int) int {
        "someFunction1": someFunction1,
        "someFunction2": someFunction1,
    }

    fmt.Println(someOtherFunction(111, 12, m["someFunction1"]))
}

func natsExample(f func(string)){
    fmt.Println("Hello from Nats module")
    fmt.Println("Alert!!! from Nats module")
    f("Nats")
}

func alertModule(moduleName string){
    fmt.Println("readJSON")
    fmt.Println("init")
    fmt.Println("call "+moduleName+" endpoints")
}
func alertFunction(moduleName string){
    fmt.Println("callbackFunction: "+moduleName)
    fmt.Println("call alert module")
    alertModule(moduleName)
}

func main() {
    DEPLOY_functions
    // a := demo2()
    // fmt.Println(demo3()+":",a)
    
    //https://medium.com/@gauravsingharoy/asynchronous-programming-with-go-546b96cd50c1
    natsExample(alertFunction)

    // fmt.Println(someOtherFunction(111, 12, someFunction1))
    // mapExample()


}
