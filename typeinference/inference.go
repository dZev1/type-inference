package typeinference

import (
	"fmt"
	"maps"
)

var typeVarCounter = 0

func Infer(context Context, expr Expr) (Type, []Constraint) {
	switch e := expr.(type) {
	case True:
		return BoolType{}, nil
	case False:
		return BoolType{}, nil
	case Zero:
		return NatType{}, nil
	case Succ:
		tN, cN := Infer(context, e.N)
		succConstraint := Constraint{Left: tN, Right: NatType{}}

		allConstraints := append([]Constraint{}, cN...)
		allConstraints = append(allConstraints, succConstraint)

		return NatType{}, allConstraints
	case Var:
		t, ok := context[e.Name]
		if !ok {
			panic("Invalid expression")
		}
		return t, nil
	case Abs:
		newCtx := copyContext(context)
		newCtx[e.Param] = e.ParamType
		tBody, cBody := Infer(newCtx, e.Body)

		return ArrowType{From: e.ParamType, To: tBody}, cBody
	case If:
		tCond, cCond := Infer(context, e.Cond)
		tThen, cThen := Infer(context, e.Then)
		tElse, cElse := Infer(context, e.Else)

		boolConstraint := Constraint{Left: tCond, Right: BoolType{}}
		branchConstraint := Constraint{Left: tThen, Right: tElse}

		allConstraints := append([]Constraint{}, cCond...)
		allConstraints = append(allConstraints, cThen...)
		allConstraints = append(allConstraints, cElse...)
		allConstraints = append(allConstraints, boolConstraint, branchConstraint)

		return tThen, allConstraints
	case App:
		retType := FreshTypeVar()

		tFunc, cFunc := Infer(context, e.Func)
		tArg, cArg := Infer(context, e.Arg)

		appConstraint := Constraint{Left: tFunc, Right: ArrowType{From: tArg, To: retType}}

		allConstraints := append([]Constraint{}, cFunc...)
		allConstraints = append(allConstraints, cArg...)
		allConstraints = append(allConstraints, appConstraint)

		return retType, allConstraints
	default:
		panic("Unknown expression")
	}
}

func copyContext(context Context) Context {
	copiedContext := make(Context)
	maps.Copy(copiedContext, context)
	return copiedContext
}

func FreshTypeVar() VarType {
	typeVarCounter++
	return VarType{Name: fmt.Sprintf("X%d", typeVarCounter)}
}
