package mods

import (
	"os"
	"path"
	"slices"
	"strings"
)

var BUILT_IN_ESPS = []string{
	"Oblivion.esm",
	"DLCBattlehornCastle.esp",
	"DLCFrostcrag.esp",
	"DLCHorseArmor.esp",
	"DLCMehrunesRazor.esp",
	"DLCOrrery.esp",
	"DLCShiveringIsles.esp",
	"DLCSpellTomes.esp",
	"DLCThievesDen.esp",
	"DLCVileLair.esp",
	"Knights.esp",
	"AltarESPMain.esp",
	"AltarESPLocal.esp",
	"AltarGymNavigation.esp",
	"AltarDeluxe.esp",
}

func OBVDataLoadOrderSuggestion(mods []Mod) string {
	vanilla := strings.Join(BUILT_IN_ESPS, "\n")

	mods = slices.DeleteFunc(mods, func(m Mod) bool {
		return m.LoadPriority < 0
	})

	slices.SortFunc(mods, func(a Mod, b Mod) int {
		return b.LoadPriority - a.LoadPriority
	})

	modOrdering := make([]string, 0)
	for _, m := range mods {
		modOrdering = append(modOrdering, m.Name)
	}

	modLoadOrder := strings.Join(modOrdering, "\n")
	return vanilla + "\n" + modLoadOrder
}

// Scans for mods in OBVDATA dir
func OBVDataScan(root string) ([]Mod, error) {
	dir, err := os.ReadDir(path.Join(root, string(OBVDATA)))
	if err != nil {
		return nil, err
	}
	m := make([]Mod, 0)
	for _, d := range dir {
		if d.IsDir() {
			continue
		}
		if slices.Contains(BUILT_IN_ESPS, d.Name()) {
			continue
		} else {
			if !strings.HasSuffix(d.Name(), ".esp") {
				continue
			}
		}
		foundMod := Mod{
			Name:     d.Name(),
			Category: OBVDATA,
		}
		foundMod.Serialize(root)
		m = append(m, foundMod)
	}
	return m, err
}
