package main

import "os"

// go:generate go-bindata -o ./apig/bindate.go -pkg apig _templates/...

func main() {
	os.Exit(Run(os.Args[1:]))
}
