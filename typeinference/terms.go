package typeinference

import "fmt"

// Expr represents an expression in the language
type Expr interface {
	isExpr()
	String() string // Ensure all expressions implement String()
}

// Variable reference
type Var struct {
	Name string
}

// Lambda abstraction
type Abs struct {
	Param     string
	ParamType Type
	Body      Expr
}

// Function application
type App struct {
	Func Expr
	Arg  Expr
}

// Conditional expression
type If struct {
	Cond Expr
	Then Expr
	Else Expr
}

// Boolean literals
type (
	True  struct{}
	False struct{}
)

// Natural number expressions
type (
	Zero struct{}
	Succ struct {
		N Expr
	}
)

// Expr interface implementations
func (Var) isExpr()   {}
func (Abs) isExpr()   {}
func (App) isExpr()   {}
func (If) isExpr()    {}
func (True) isExpr()  {}
func (False) isExpr() {}
func (Zero) isExpr()  {}
func (Succ) isExpr()  {}

// String representations for expressions
func (v Var) String() string   { return v.Name }
func (t True) String() string  { return "True" }
func (f False) String() string { return "False" }
func (z Zero) String() string  { return "Zero" }

func (abs Abs) String() string {
	return fmt.Sprintf("(λ%s : %v . %v)", abs.Param, abs.ParamType, abs.Body)
}

func (app App) String() string {
	return fmt.Sprintf("(%v %v)", app.Func, app.Arg)
}

func (ite If) String() string {
	return fmt.Sprintf("if %v then %v else %v", ite.Cond, ite.Then, ite.Else)
}

func (s Succ) String() string {
	return fmt.Sprintf("Succ(%v)", s.N)
}

