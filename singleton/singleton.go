package singleton

import (
	"fmt"
	"sync"
)

type config struct {
}

var instance *config
var once sync.Once

func GetInstance() *config {
	once.Do(func() {
		instance = new(config)
	})

	return instance
}

func (this *config) HelloWorld(msg string) {
	fmt.Println(msg)
}
