package day01

type Elf struct {
	Name     string
	Calories int
}

type ElfList []Elf

func (pl ElfList) Less(i, j int) bool {
	return pl[i].Calories > pl[j].Calories
}

// Swap tauscht die Elemente mit den Indizes i und j aus
func (pl ElfList) Swap(i, j int) {
	pl[i], pl[j] = pl[j], pl[i]
}

func (pl ElfList) Len() int {
	return len(pl)
}

func CreateElf(name string) Elf {
	return Elf{
		Name:     name,
		Calories: 0,
	}
}
