package express

import (
	"fmt"
	"strconv"

	go_utils "github.com/huangxingx/go-utils"
	"github.com/huangxingx/go-utils/express/operate"
)

type Express struct {
	inputExpress string
}

func NewExpress(express string) *Express {
	return &Express{
		inputExpress: express,
	}
}

func (e *Express) Execute(param map[string]interface{}) interface{} {
	mpn := parse2mpn(e.inputExpress)
	suffixExpress := parseSuffixExpress(mpn)
	stack := go_utils.NewStack()
	for _, v := range suffixExpress {
		// number or keywork
		if IsNum(v) || isKeyWork(v) {
			stack.Push(v)
			continue
		}
		// op
		iOperate := operate.GetOperate(v)
		if iOperate == nil {
			// var
			value, ok := param[v]
			if !ok {
				panic(fmt.Sprintf("var %s not value", v))
			}
			stack.Push(value)
			continue
		}
		v2 := stack.Pop()
		v1 := stack.Pop()

		result := iOperate.Execute(v1, v2)
		stack.Push(result)
	}
	result := stack.Pop()
	if !stack.IsEmpty() {
		stack.Print()
		panic(fmt.Sprintf("execute err: stack.len > 0"))
	}
	// todo string -> float64
	switch result.(type) {
	case string:
		floatValue, err := strconv.ParseFloat(result.(string), 64)
		if err != nil {
			panic(err)
		}
		return floatValue
	case float64:
		return result.(float64)
	case bool:
		return result.(bool)
	}
	panic("value format err")
}

func IsNum(s string) bool {
	_, err := strconv.ParseFloat(s, 64)
	return err == nil
}
