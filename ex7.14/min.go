package eval

import (
	"math"
)

// retrieves the minimum value of the set of this expressions operands
func (opExpr OpExpr) Min() literal {
	currentMin := opExpr.min(literal(math.MaxFloat64))

	return currentMin
}

func (opExpr OpExpr) min(currentMin literal) literal {
	// check for each type

	if v, ok := opExpr.expr.(literal); ok {
		if v <= currentMin {
			currentMin = v
		}
		return currentMin
	} else if v, ok := opExpr.expr.(unary); ok {
		opExpr = OpExpr{v.x}
		return opExpr.min(currentMin)
	} else if v, ok := opExpr.expr.(binary); ok {
		opExprX := OpExpr{v.x}
		opExprY := OpExpr{v.y}
		minX := opExprX.min(currentMin)
		minY := opExprY.min(currentMin)
		if minX <= minY {
			return minX
		}
		return minY
	} else { // Var type
		return currentMin
	}
	// else if v, ok := opExpr.expr.(call); ok {
	// 	buffer := make(int[], len(v.args))
	// 	for i, v := range v.args {
	// 		opExpr = OpExpr{v.args[i]}
	// 		opExpr.min()
	// 	}
	// 	return opExpr.min(currentMin)
	// }

}
