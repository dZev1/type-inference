package main

import (
	"fmt"
	inference "github.com/dZev1/type-inference/typeinference"
)

func main() {
	myExpression := inference.If{
		Cond: inference.App{
			Func: inference.Abs{
				Param: "z",
				Body: inference.Var{Name: "z"},
			},
			Arg: inference.Var{Name: "y"},
		},
		Then: inference.Zero{},
		Else: inference.Succ{N: inference.Zero{}},
	}

	gamma0, m0 := inference.TypeAnnotations(myExpression)
	typeOfTerm, constraints := inference.Infer(gamma0, m0)
	fmt.Printf("%v ⊢ %v : %v\n", gamma0, m0, typeOfTerm)
	fmt.Println("Constraints:")

	PrintConstraints(constraints)
}

func PrintConstraints(cs []inference.Constraint) {
	for i, c := range cs {
		fmt.Printf("%2d: %v =? %v\n", i+1, c.Left, c.Right)
	}
}
