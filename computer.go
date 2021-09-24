package go_utils

import "strconv"

// Computer 计算后缀表达式
func Computer(expression string) float64 {
	mpnExpression := parse2mpn(expression)
	rpnExpression := parse2rpn(mpnExpression)
	resultStack := NewStack()
	for _, v := range rpnExpression {
		switch v {
		case "+", "-", "*", "/":
			v1, _ := resultStack.Pop().(float64)
			v2, _ := resultStack.Pop().(float64)
			switch v {
			case "+":
				resultStack.Push(v1 + v2)
			case "-":
				resultStack.Push(v2 - v1)
			case "*":
				resultStack.Push(v2 * v1)
			case "/":
				resultStack.Push(v2 / v1)
			}
		default:
			float, err := strconv.ParseFloat(v, 64)
			if err != nil {
				panic(err)
			}
			resultStack.Push(float)
		}
	}
	return resultStack.Pop().(float64)
}

// 解析中缀表达式数组
func parse2mpn(express string) []string {
	//compile, _ := regexp.Compile("[\\d\\+\\-\\*\\/\\(\\)]")
	//return compile.FindAllString(express, -1)
	result := make([]string, 0)
	s := ""
	for i, v := range express {
		if v == 32 {
			continue
		}
		if v > 47 && v < 59 || v == 46 {
			s += string(v)
			if i == len(express)-1 {
				result = append(result, s)
			}
		} else {
			// +-*/ ( )
			if s != "" {
				result = append(result, s)
				s = ""
			}
			result = append(result, string(v))
		}
	}
	return result
}

// parse2rpn 解析中缀表达式-> 后缀表达式
func parse2rpn(express []string) []string {
	stact := NewStack()
	rpnList := make([]string, 0, len(express))
	var s string
	for i := 0; i < len(express); i++ {
		s = express[i]
		switch s {
		case "+", "-":
			for {
				if stact.IsEmpty() || stact.Peek().(string) == "(" {
					stact.Push(s)
					break
				} else {
					rpnList = append(rpnList, stact.Pop().(string))
				}
			}
		case "*", "/":
			for {
				if stact.IsEmpty() || (stact.Peek().(string) != "*" && stact.Peek().(string) != "/") {
					stact.Push(s)
					break
				} else {
					rpnList = append(rpnList, stact.Pop().(string))
				}
			}

		case "(":
			stact.Push(s)

		case ")":

			for stact.Peek().(string) != "(" {
				rpnList = append(rpnList, stact.Pop().(string))
			}
			stact.Pop()
		default:
			rpnList = append(rpnList, s)
		}
	}

	for stact.Size() > 0 {
		rpnList = append(rpnList, stact.Pop().(string))
	}
	return rpnList
}
