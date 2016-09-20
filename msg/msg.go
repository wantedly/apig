package msg

import (
	"fmt"
	"sync"
)

var (
	Mute = false
	m    sync.Mutex
)

func Printf(format string, a ...interface{}) {
	if !Mute {
		m.Lock()
		fmt.Printf(format, a...)
		m.Unlock()
	}
}

func Println(a ...interface{}) {
	if !Mute {
		m.Lock()
		fmt.Println(a...)
		m.Unlock()
	}
}
