package typeinference

import "fmt"

type Type interface{ isType() }

type BoolType struct{}

func (BoolType) isType() {}

type NatType struct{}

func (NatType) isType() {}

type VarType struct{ Name string }

func (VarType) isType() {}

type ArrowType struct {
	From Type
	To   Type
}

func (ArrowType) isType() {}

type Context map[string]Type

type Constraint struct {
	Left  Type
	Right Type
}

func (BoolType) String() string {
	return "Bool"
}

func (a ArrowType) String() string {
	return fmt.Sprintf("(%v -> %v)", a.From, a.To)
}

func (v VarType) String() string {
	return v.Name
}

func (NatType) String() string {
	return "Nat"
}

func (c Constraint) String() string {
	return fmt.Sprintf("%v =? %v", c.Left, c.Right)
}