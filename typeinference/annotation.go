package typeinference

import "maps"

func TypeAnnotations(expr Expr) (Context, Expr) {
	retContext := make(Context)
	switch e := expr.(type) {
	case Var:
		retContext[e.Name] = FreshTypeVar()
		return retContext, e

	case Abs:
		paramType := FreshTypeVar()
		bodyContext, annotatedBody := TypeAnnotations(e.Body)

		mergedContext := make(Context)
		for name, typ := range bodyContext {
			if name != e.Param {
				mergedContext[name] = typ
			}
		}

		return mergedContext, Abs{
			Param:     e.Param,
			ParamType: paramType,
			Body:      annotatedBody,
		}

	case App:
		funContext, funcAnnotated := TypeAnnotations(e.Func)
		argContext, argAnnotated := TypeAnnotations(e.Arg)

		mergedContext := make(Context)
		maps.Copy(mergedContext, funContext)
		maps.Copy(mergedContext, argContext)

		return mergedContext, App{
			Func: funcAnnotated,
			Arg:  argAnnotated,
		}

	case If:
		condContext, condAnnotated := TypeAnnotations(e.Cond)
		thenContext, thenAnnotated := TypeAnnotations(e.Then)
		elseContext, elseAnnotated := TypeAnnotations(e.Else)

		mergedContext := make(Context)
		maps.Copy(mergedContext, condContext)
		maps.Copy(mergedContext, thenContext)
		maps.Copy(mergedContext, elseContext)

		return mergedContext, If{
			Cond: condAnnotated,
			Then: thenAnnotated,
			Else: elseAnnotated,
		}
	}
	return retContext, expr
}
