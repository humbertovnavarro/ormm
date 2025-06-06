package mods

import (
	"encoding/json"
	"os"
	"path"
)

type ModType string

const BINARIES_PATH = "Binaries/Win64/"
const CONTENT_PATH = "Content/"

const (
	OBSE    ModType = BINARIES_PATH + "OBSE/Plugins"
	UE4SS   ModType = BINARIES_PATH + "ue4ss/Mods"
	PAK     ModType = CONTENT_PATH + "Paks/~mods"
	OBVDATA ModType = CONTENT_PATH + "Dev/ObvData/Data"
)

const PLUGINS_TXT = string(OBVDATA + "/Plugins.txt")

type Mod struct {
	Name         string `json:"-"`
	Category     ModType
	LoadPriority int
}

func (m *Mod) Serialize(root string) (err error) {
	metaPath := path.Join(root, string(m.Category), m.Name+".meta.json")
	j, _ := os.ReadFile(metaPath)
	json.Unmarshal(j, m)
	encoding, err := json.MarshalIndent(m, "", "	")
	if err != nil {
		panic(err)
	}
	err = os.WriteFile(metaPath, encoding, 0755)
	return err
}
