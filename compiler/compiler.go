package compiler

import (
	"panda/code"
)

type Compiler struct {
	instructions code.Instructions
	constants    []object.Object
}
