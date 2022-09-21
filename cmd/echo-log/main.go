package main

import (
	"github.com/paveldanilin/echo-log/internal/script"
	lua "github.com/yuin/gopher-lua"
)

func Double() func(L *lua.LState) int {
	return func(L *lua.LState) int {
		lv := L.ToInt(1)            /* get argument */
		L.Push(lua.LNumber(lv * 2)) /* push result */
		return 1
	}
}

func main() {
	script := script.NewLuaScript()

	script.Load("func.lua")

	script.RegisterFunction("double", Double)

	ret := script.CallFunction("test_one")
	if ret.String() == "123" {
		println(">> OK")
	}

	ret2 := script.CallFunction("test")
	println(ret2.String())

}
