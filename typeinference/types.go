package typeinference

import (
	"fmt"
	"strings"
)

// Type represents a type in the type system
type Type interface {
	isType()
	String() string // Ensure all types implement String()
}

// Context is a mapping from variable names to their types
type Context map[string]Type

// Constraint represents a type constraint between two types
type Constraint struct {
	Left  Type
	Right Type
}

// Concrete types
type (
	// BoolType represents a boolean type
	BoolType struct{}

	// NatType represents a natural number type
	NatType struct{}

	// VarType represents a type variable
	VarType struct {
		Name string
	}

	// ArrowType represents a function type
	ArrowType struct {
		From Type
		To   Type
	}
)

// Type interface implementations
func (BoolType) isType()  {}
func (NatType) isType()   {}
func (VarType) isType()   {}
func (ArrowType) isType() {}

// String representations
func (BoolType) String() string  { return "Bool" }
func (NatType) String() string   { return "Nat" }
func (v VarType) String() string { return v.Name }

func (a ArrowType) String() string {
	return fmt.Sprintf("(%v -> %v)", a.From, a.To)
}

func (c Constraint) String() string {
	return fmt.Sprintf("%v =? %v", c.Left, c.Right)
}

func (c Context) String() string {
	var sb strings.Builder
    sb.WriteString("{")
    
    first := true
    for k, v := range c {
        if !first {
            sb.WriteString(",")
        }
        sb.WriteString(fmt.Sprintf(" %v : %v ", k, v))
        first = false
    }
    
    sb.WriteString("}")
    return sb.String()
}
