package domain

// Race describes the DnD 5E character Race
type Race struct {
	ID                  string
	Name                string
	Age                 string
	Speed               int
	Size                string
	SizeDescription     string
	Alignment           string
	LanguageDescription string
	SubRaces            []SubRace
	Traits              []Trait
}

//SubRace describes the subraces of the Race
type SubRace struct {
}

//Trait describes a racial trait
type Trait struct {
}
