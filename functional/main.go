package main

import (
	"github.com/AtalH/learngo/functional/closure"
	"github.com/AtalH/learngo/functional/implinterface"
	"github.com/AtalH/learngo/functional/traverse"
)

func main() {
	closure.TestAdder()
	implinterface.TestFib()
	implinterface.TestGenerator()
	traverse.TestOnTraverse()
}
