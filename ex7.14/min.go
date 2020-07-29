package eval

import (
	"math"
)

// Min() retrieves the minimum value of the set of this expressions operands
func (opExpr OpExpr) Min() literal {
	currentMin := opExpr.min(literal(math.MaxFloat64))

	return currentMin
}

func (opExpr OpExpr) min(currentMin literal) literal {
	// check for each type

	switch v := opExpr.expr.(type) {
	case literal:
		if v <= currentMin {
			currentMin = v
		}
		return currentMin
	case unary:
		opExpr = OpExpr{v.x}
		return opExpr.min(currentMin)
	case binary:
		opExprX := OpExpr{v.x}
		opExprY := OpExpr{v.y}
		minX := opExprX.min(currentMin)
		minY := opExprY.min(currentMin)
		if minX <= minY {
			return minX
		}
		return minY
	case call:
		// loop through args creating objects and return the min of this set
		minOfSet := currentMin
		for _, val := range v.args {
			current := OpExpr{val}
			min := current.min(currentMin)
			if min <= minOfSet {
				minOfSet = min
			}
		}
		return minOfSet
	default:
		return currentMin

	}

}
