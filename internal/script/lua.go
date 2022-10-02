package script

import (
	"errors"
	"fmt"
	"reflect"
	"strconv"
	"time"

	golua "github.com/yuin/gopher-lua"
)

// @see https://morioh.com/p/f9f6ad4de5b8
// @see http://dnaeon.github.io/extending-lua-with-go-types/

// Represents LUA script
type LuaScript struct {
	state *golua.LState
}

// NewLuaScript
func NewLuaScript() *LuaScript {
	luaScript := LuaScript{
		state: golua.NewState(),
	}
	return &luaScript
}

// Load LUA script from file
func (lua *LuaScript) LoadFile(filename string) error {
	return lua.state.DoFile(filename)
}

// Load LUA script from string
func (lua *LuaScript) LoadString(script string) error {
	return lua.state.DoString(script)
}

// Release LUA state
func (lua *LuaScript) Release() {
	lua.state.Close()
}

// RegisterFunction
// Sets a Go function in LUA script
func (lua *LuaScript) RegisterFunction(name string, f interface{}) error {
	if reflect.TypeOf(f).Kind() != reflect.Func {
		return errors.New("is not a function")
	}
	lua.state.SetGlobal(name, lua.state.NewFunction(lua.toLuaFunc(f)))
	return nil
}

// RegisterInt
// Sets a global named value in LUA script
func (lua *LuaScript) RegisterInt(name string, v int) {
	lua.state.SetGlobal(name, golua.LNumber(v))
}

// RegisterFloat
// Sets a global named value in LUA script
func (lua *LuaScript) RegisterFloat(name string, v float64) {
	lua.state.SetGlobal(name, golua.LNumber(v))
}

// RegisterString
// Sets a global named value in LUA script
func (lua *LuaScript) RegisterString(name string, v string) {
	lua.state.SetGlobal(name, golua.LString(v))
}

// RegisterType
// http://dnaeon.github.io/extending-lua-with-go-types/
// https://morioh.com/p/f9f6ad4de5b8
func (lua *LuaScript) RegisterType(name string, factory interface{}, methods map[string]interface{}) {
	mt := lua.state.NewTypeMetatable(name)

	lua.state.SetGlobal(name, mt)

	// Static factory method
	lua.state.SetField(mt, "new", lua.state.NewFunction(func(L *golua.LState) int {
		rf := reflect.ValueOf(factory)
		ft := reflect.TypeOf(factory)
		numargs := ft.NumIn()
		var args []reflect.Value

		// Convert LUA func args to a GO func
		for i := 0; i < numargs; i++ {
			var gv interface{}
			lt := L.Get(i + 1).Type()

			switch lt {
			case golua.LTString:
				gv = L.ToString(i + 1)
			case golua.LTNumber:
				gv = L.ToInt(i + 1)
			case golua.LTBool:
				gv = L.ToBool(i + 1)
			case golua.LTUserData:
				gv = L.ToUserData(i + 1).Value
			default:
				panic("Could not convert LUA type to GO type")
			}
			args = append(args, reflect.ValueOf(gv))
		}

		// Call a Factory method
		newUd := rf.Call(args)

		if len(newUd) != 1 {
			panic("A factory method must return one value")
		}

		ud := L.NewUserData()
		ud.Value = newUd[0].Interface()
		L.SetMetatable(ud, L.GetTypeMetatable(name))

		// Return value to LUA
		L.Push(ud)

		return 1
	}))

	// Methods
	luaMethods := make(map[string]golua.LGFunction)
	for k, f := range methods {
		luaMethods[k] = lua.toLuaFunc(f)
	}

	lua.state.SetField(mt, "__index", lua.state.SetFuncs(lua.state.NewTable(), luaMethods))
}

// CallFunction
// Calls a LUA function
func (lua *LuaScript) CallFunction(fn string, args ...interface{}) (interface{}, error) {
	// Lookup for a LUA function
	lVal := lua.state.GetGlobal(fn)
	if lVal.Type() != golua.LTFunction {
		return nil, fmt.Errorf("[%s] is not a LUA function", fn)
	}

	// Convert args to LUA values
	fnArgs := make([]golua.LValue, len(args))
	for i, arg := range args {
		fnArgs[i] = lua.toLuaValue(arg)
	}

	// Call a LUA function
	if err := lua.state.CallByParam(golua.P{
		Fn:      lVal, // name of LUA function
		NRet:    1,    // number of returned values
		Protect: true, // return error or panic
	}, fnArgs...); err != nil {
		return nil, err
	}

	ret := lua.state.Get(-1) // returned value
	lua.state.Pop(1)         // remove received value

	return lua.toValue(ret)
}

// toLuaFunc
// GO -> LUA
// Converts a Go function to a LUA function
func (lua *LuaScript) toLuaFunc(f interface{}) golua.LGFunction {
	return func(L *golua.LState) int {
		rf := reflect.ValueOf(f)
		ft := reflect.TypeOf(f)
		numargs := ft.NumIn()
		var args []reflect.Value

		// Convert LUA func args to a GO func
		for i := 0; i < numargs; i++ {
			var gv interface{}
			lt := L.Get(i + 1).Type()

			switch lt {
			case golua.LTString:
				gv = L.ToString(i + 1)
			case golua.LTNumber:
				gv = L.ToInt(i + 1)
			case golua.LTBool:
				gv = L.ToBool(i + 1)
			case golua.LTUserData:
				gv = L.ToUserData(i + 1).Value
			default:
				panic("Could not convert LUA type to GO type")
			}
			args = append(args, reflect.ValueOf(gv))
		}

		// Call a GO func
		rvs := rf.Call(args)

		// Convert GO result to LUA
		for i := 0; i < len(rvs); i++ {
			rv := rvs[i]
			switch rv.Kind() {
			case reflect.Int:
				L.Push(golua.LNumber(rv.Int()))
			case reflect.String:
				L.Push(golua.LString(rv.String()))
			default:
				ud := L.NewUserData()
				ud.Value = rv.Interface()
				L.Push(ud)
			}
		}

		return len(rvs)
	}
}

// toLuaValue
// GO -> LUA
// Converts a Go value to a LUA value
func (lua *LuaScript) toLuaValue(v interface{}) golua.LValue {
	if v == nil {
		return golua.LNil
	}

	switch v.(type) {
	case float64:
		return golua.LNumber(v.(float64))
	case int:
		return golua.LNumber(v.(int))
	case int64:
		return golua.LNumber(v.(int64))
	case string:
		return golua.LString(v.(string))
	case bool:
		return golua.LBool(v.(bool))
	case []byte:
		return golua.LString(string(v.([]byte)))
	case map[string]interface{}:
		return lua.toLuaTable(v.(map[string]interface{}))
	case time.Time:
		return golua.LNumber(v.(time.Time).Unix())
	case []map[string]interface{}:
		sliceTable := &golua.LTable{}
		for _, s := range v.([]map[string]interface{}) {
			tble := lua.toLuaTable(s)
			sliceTable.Append(tble)
		}
		return sliceTable
	case []interface{}:
		sliceTable := &golua.LTable{}
		for _, s := range v.([]interface{}) {
			sliceTable.Append(lua.toLuaValue(s))
		}
		return sliceTable
	}

	panic(fmt.Sprintf("could not map [%s] => lua type", reflect.ValueOf(v).Kind().String()))
}

// toLuaTable
// GO -> LUA
// Converts a Go map to a LUA table
func (lua *LuaScript) toLuaTable(m map[string]interface{}) *golua.LTable {
	resultTable := &golua.LTable{}

	for key, value := range m {
		resultTable.RawSetString(key, lua.toLuaValue(value))
	}

	return resultTable
}

// toGoValue
// LUA -> GO
// Converts a LUA value to a GO value
func (lua *LuaScript) toValue(lval golua.LValue) (interface{}, error) {

	switch lval.Type() {
	case golua.LTNil:
		return nil, nil
	case golua.LTBool:
		return golua.LVAsBool(lval), nil
	case golua.LTNumber:
		return str2number(lval.String())
	case golua.LTString:
		return golua.LVAsString(lval), nil
	}

	return nil, errors.New("could not convert a LUA value to a GO interface")
}

// str2number
func str2number(s string) (interface{}, error) {
	if iv, err := strconv.Atoi(s); err == nil {
		return iv, nil
	}
	return strconv.ParseFloat(s, 64)
}
