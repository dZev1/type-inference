package typeinference

import "fmt"

type Expr interface{ isExpr() }

type Var struct{ Name string }

func (Var) isExpr() {}

type Abs struct {
	Param     string
	ParamType Type
	Body      Expr
}

func (Abs) isExpr() {}

type App struct {
	Func Expr
	Arg  Expr
}

func (App) isExpr() {}

type If struct {
	Cond Expr
	Then Expr
	Else Expr
}

func (If) isExpr() {}

type True struct{}

func (True) isExpr() {}

type False struct{}

func (False) isExpr() {}

type Zero struct{}

func (Zero) isExpr() {}

type Succ struct{ N Expr }

func (Succ) isExpr() {}

// String functions

func (ite If) String() string {
	return fmt.Sprintf("if %v then %v else %v", ite.Cond, ite.Then, ite.Else)
}

func (v Var) String() string {
	return v.Name 
}

func (abs Abs) String() string {
	return fmt.Sprintf("(λ%v : %v . %v)", abs.Param, abs.ParamType, abs.Body)
}

func (app App) String() string {
	return fmt.Sprintf("(%v %v)", app.Func, app.Arg)
}

func (True) String() string {
	return "True"
}

func (False) String() string {
	return "False"
}

func (Zero) String() string {
	return "Zero"
}

func (s Succ) String() string {
	return fmt.Sprintf("Succ(%v)", s.N)
}