package script

import (
	"reflect"
	"time"

	golua "github.com/yuin/gopher-lua"
)

type UserFunction func() func(L *golua.LState) int

type LuaScript struct {
	state *golua.LState
}

func NewLuaScript() *LuaScript {
	luaScript := LuaScript{
		state: golua.NewState(),
	}
	return &luaScript
}

// Load LUA script from file
func (lua *LuaScript) Load(filename string) {
	if err := lua.state.DoFile(filename); err != nil {
		panic(err)
	}
}

// Load LUA script from string
func (lua *LuaScript) LoadString(script string) {
	if err := lua.state.DoString(script); err != nil {
		panic(err)
	}
}

// Close LUA state
func (lua *LuaScript) Close() {
	lua.state.Close()
}

// Register go fucntion
func (lua *LuaScript) RegisterFunction(fn string, userFunction UserFunction) {
	lua.state.SetGlobal(fn, lua.state.NewFunction(userFunction()))
}

// Call function with arguments
func (lua *LuaScript) CallFunction(fn string, args ...interface{}) golua.LValue {

	luaFunc := lua.state.GetGlobal(fn)
	if luaFunc.Type() != golua.LTFunction {
		panic("[" + fn + "] is not a function!")
	}

	fnArgs := make([]golua.LValue, len(args))

	for i, arg := range args {

		if arg == nil {
			fnArgs[i] = golua.LNil
			continue
		}

		switch reflect.ValueOf(arg).Kind() {
		case reflect.Int:
			fnArgs[i] = golua.LNumber(arg.(int))
		case reflect.Float64:
			fnArgs[i] = golua.LNumber(arg.(float64))
		case reflect.String:
			fnArgs[i] = golua.LString(arg.(string))
		case reflect.Bool:
			fnArgs[i] = golua.LBool(arg.(bool))
		case reflect.Map:
			fnArgs[i] = lua.mapToTable(arg.(map[string]interface{}))
		}
	}

	if err := lua.state.CallByParam(golua.P{
		Fn:      luaFunc,
		NRet:    1,
		Protect: true,
	}, fnArgs...); err != nil {
		panic(err)
	}

	ret := lua.state.Get(-1) // returned value
	lua.state.Pop(1)         // remove received value

	return ret
}

// mapToTable converts a Go map to a lua table
func (lua *LuaScript) mapToTable(m map[string]interface{}) *golua.LTable {
	// Main table pointer
	resultTable := &golua.LTable{}

	// Loop map
	for key, element := range m {

		switch element.(type) {
		case float64:
			resultTable.RawSetString(key, golua.LNumber(element.(float64)))
		case int:
			resultTable.RawSetString(key, golua.LNumber(element.(int)))
		case int64:
			resultTable.RawSetString(key, golua.LNumber(element.(int64)))
		case string:
			resultTable.RawSetString(key, golua.LString(element.(string)))
		case bool:
			resultTable.RawSetString(key, golua.LBool(element.(bool)))
		case []byte:
			resultTable.RawSetString(key, golua.LString(string(element.([]byte))))
		case map[string]interface{}:

			// Get table from map
			tble := lua.mapToTable(element.(map[string]interface{}))

			resultTable.RawSetString(key, tble)

		case time.Time:
			resultTable.RawSetString(key, golua.LNumber(element.(time.Time).Unix()))

		case []map[string]interface{}:

			// Create slice table
			sliceTable := &golua.LTable{}

			// Loop element
			for _, s := range element.([]map[string]interface{}) {

				// Get table from map
				tble := lua.mapToTable(s)

				sliceTable.Append(tble)
			}

			// Set slice table
			resultTable.RawSetString(key, sliceTable)

		case []interface{}:
			// Create slice table
			sliceTable := &golua.LTable{}

			// Loop interface slice
			for _, s := range element.([]interface{}) {

				// Switch interface type
				switch s.(type) {
				case map[string]interface{}:
					// Convert map to table
					t := lua.mapToTable(s.(map[string]interface{}))
					// Append result
					sliceTable.Append(t)

				case int:
					sliceTable.Append(golua.LNumber(s.(int)))

				case float64:
					// Append result as number
					sliceTable.Append(golua.LNumber(s.(float64)))

				case string:
					// Append result as string
					sliceTable.Append(golua.LString(s.(string)))

				case bool:
					// Append result as bool
					sliceTable.Append(golua.LBool(s.(bool)))
				}
			}

			// Append to main table
			resultTable.RawSetString(key, sliceTable)
		}
	}

	return resultTable
}
