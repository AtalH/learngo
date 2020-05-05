package main

import (
	"github.com/AtalH/learngo/errorAndResource/deferdemo"
	"github.com/AtalH/learngo/errorAndResource/webserver"
)

func main() {
	deferdemo.Test()
	webserver.Start()
}
