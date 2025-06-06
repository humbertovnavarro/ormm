package main

import (
	"os"

	"github.com/humbertovnavarro/obrmm/pkg/mods"
)

func main() {
	m, err := mods.OBVDataScan(".")
	if err != nil {
		panic(err)
	}
	order := mods.OBVDataLoadOrderSuggestion(m)
	backup, err := os.Create(mods.PLUGINS_TXT + ".bak")
	if err != nil {
		panic(err)
	}
	existing, err := os.ReadFile(mods.PLUGINS_TXT)
	if err != nil {
		panic(err)
	}
	_, err = backup.WriteString(string(existing))
	if err != nil {
		panic(err)
	}
	err = os.WriteFile(mods.PLUGINS_TXT, []byte(order), 0755)
	if err != nil {
		panic(err)
	}
}
