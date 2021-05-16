package evaluate

import (
	"math"
	"strconv"

	"github.com/cloudfstrife/algs/fundamentals/dst/stack"
)

// DijkstraEvluate Dijkstra’s Two-Stack Algorithm for Expression Evaluation
// Page number - Chinese Edition : 80
// Page number : 129
func DijkstraEvluate(slist []string) (float64, error) {
	ops := stack.NewStack()
	vals := stack.NewStack()
	for _, v := range slist {
		switch v {
		case "(":
			continue
		case "+":
			ops.Push(v)
		case "-":
			ops.Push(v)
		case "*":
			ops.Push(v)
		case "/":
			ops.Push(v)
		case "sqrt":
			ops.Push(v)
		case ")":
			v1, err := vals.Pop()
			if err != nil {
				return 0, err
			}
			opt, err := ops.Pop()
			if err != nil {
				return 0, err
			}
			switch opt {
			case "+":
				v2, err := vals.Pop()
				if err != nil {
					return 0, err
				}
				vals.Push(v2.(float64) + v1.(float64))
			case "-":
				v2, err := vals.Pop()
				if err != nil {
					return 0, err
				}
				vals.Push(v2.(float64) - v1.(float64))
			case "*":
				v2, err := vals.Pop()
				if err != nil {
					return 0, err
				}
				vals.Push(v2.(float64) * v1.(float64))
			case "/":
				v2, err := vals.Pop()
				if err != nil {
					return 0, err
				}
				vals.Push(v2.(float64) / v1.(float64))
			case "sqrt":
				vals.Push(math.Sqrt(v1.(float64)))
			}
		default:
			val, err := strconv.ParseFloat(v, 64)
			if err != nil {
				return 0, err
			}
			vals.Push(val)
		}
	}
	result, err := vals.Pop()
	if err != nil {
		return 0, err
	}
	return result.(float64), nil
}
