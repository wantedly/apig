package msg

import (
	"fmt"
)

var Mute = false

func Printf(format string, a ...interface{}) {
	if !Mute {
		fmt.Printf(format, a)
	}
}

func Println(a ...interface{}) {
	if !Mute {
		fmt.Println(a)
	}
}
