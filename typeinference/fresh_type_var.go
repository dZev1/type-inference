package typeinference

import "fmt"

var typeVarCounter = 0

func FreshTypeVar() VarType {
	typeVarCounter++
	return VarType{Name: fmt.Sprintf("X%d", typeVarCounter)}
}