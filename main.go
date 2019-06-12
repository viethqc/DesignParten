package main

import (
	"./singleton"
)

func main() {
	//singleton demo
	singleton := singleton.GetInstance()
	singleton.HelloWorld("hello, i'm singleton")
}
