package operate

import (
	"strconv"
	"sync"
)

const (
	_              int = iota
	LogicPriority      // 逻辑运算 and or
	LogicPriority2     // 逻辑运算2 > < >= <= !
	Arithmetic1        // 四则运算1 + -
	Arithmetic2        // 四则运算2 * / %
)

var globalOperateMap = map[string]IOperate{}

type IOperate interface {
	GetOperateSymbol() []string                         // operate symbol
	Execute(v1 interface{}, v2 interface{}) interface{} // execute func
	GetPriority() int
	GetRegexMatch() string //string parse by regex
}

var lock sync.RWMutex

//Register operate
func Register(name string, operate IOperate) {
	lock.Lock()
	defer lock.Unlock()
	globalOperateMap[name] = operate
}

//GetOperate
//todo reflect
func GetOperate(symbol string) IOperate {
	for _, iOperate := range globalOperateMap {
		for _, v := range iOperate.GetOperateSymbol() {
			if v == symbol {
				return iOperate
			}
		}
	}
	return nil
}

func GetAllOperate() (operateList []IOperate) {
	for _, operate := range globalOperateMap {
		operateList = append(operateList, operate)
	}
	return
}

var _ IOperate = baseOperate{}

type baseOperate struct {
	name       string
	execute    func(v1, v2 interface{}) interface{}
	symbol     []string
	regexMatch string
	priority   int
}

func (b baseOperate) GetOperateSymbol() []string {
	return b.symbol
}

func (b baseOperate) GetRegexMatch() string {
	return b.regexMatch
}

func (b baseOperate) Execute(v1 interface{}, v2 interface{}) interface{} {
	return b.execute(v1, v2)
}

func (b baseOperate) GetOperate() []string {
	return b.symbol
}

func (b baseOperate) GetPriority() int {
	return b.priority
}

func ensureFloat64(s interface{}) float64 {
	switch s.(type) {
	case string:
		floatValue, err := strconv.ParseFloat(s.(string), 64)
		if err != nil {
			panic(err)
		}
		return floatValue

	case float64:
		return s.(float64)
	}
	panic("value format err")
}

func ensureBool(s interface{}) bool {
	switch s.(type) {
	case string:
		boolValue, err := strconv.ParseBool(s.(string))
		if err != nil {
			panic(err)
		}
		return boolValue

	case bool:
		return s.(bool)
	}
	panic("value format err")
}
