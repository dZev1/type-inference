package main

import (
	"fmt"
	inference "github.com/dZev1/type-inference/typeinference"
)

func main() {
	myExpression := inference.If{
		Cond: inference.App{
			Func: inference.Abs{
				Param:     "x",
				ParamType: inference.FreshTypeVar(),
				Body:      inference.Var{Name: "x"},
			},
			Arg: inference.True{},
		},
		Then: inference.Zero{},
		Else: inference.Succ{N: inference.Zero{}},
	}

	fmt.Println(myExpression)

	a, b := inference.Infer(inference.Context{}, myExpression)
	fmt.Printf("Type: %v\n", a)
	fmt.Println("Constraints:")

	PrintConstraints(b)
}

func PrintConstraints(cs []inference.Constraint) {
	for i, c := range cs {
		fmt.Printf("%2d: %v =? %v\n", i+1, c.Left, c.Right)
	}
}
