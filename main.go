package main

import (
	"fmt"
	"net/http"

	lua "github.com/yuin/gopher-lua"
)

func handler(w http.ResponseWriter, r *http.Request) {
	L := lua.NewState()
	defer L.Close()

	// Beispiel-Daten aus dem Backend
	data := map[string]string{
		"title":   "Hello, World! Perter123",
		"message": "Welcome to my website! Oder auch nicht",
	}

	// Daten manuell in eine Lua-Tabelle umwandeln
	luaData := L.NewTable()
	for key, value := range data {
		L.SetTable(luaData, lua.LString(key), lua.LString(value))
	}
	L.SetGlobal("data", luaData)

	if err := L.DoFile("template.lua"); err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	luaValue := L.GetGlobal("Render")
	if luaValue.Type() == lua.LTFunction {
		if err := L.CallByParam(lua.P{
			Fn:      luaValue,
			NRet:    1,
			Protect: true,
		}); err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}
		rendered := L.Get(-1).String()
		L.Pop(1)
		fmt.Fprintf(w, rendered)
	}
}

func main() {
	http.HandleFunc("/", handler)

	http.ListenAndServe(":3000", nil)
}
