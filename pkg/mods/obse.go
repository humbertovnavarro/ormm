package mods

import (
	"os"
	"path"
)

func OBSEDataScan(root string) ([]Mod, error) {
	dir, err := os.ReadDir(path.Join(root, string(OBSE)))
	if err != nil {
		return nil, err
	}
	m := make([]Mod, 0)
	for _, d := range dir {
		if d.IsDir() {
			m = append(m, Mod{
				Name:     d.Name(),
				Category: OBSE,
			})
		}
	}
	return m, err
}
