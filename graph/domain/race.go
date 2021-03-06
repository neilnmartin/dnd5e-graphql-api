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
	ID             string
	Name           string
	AbilityBonuses []AbilityBonus
}

//AbilityBonus describes a subrace ability bonus
type AbilityBonus struct {
	Name  string
	Bonus int
}

//Trait describes a racial trait
type Trait struct {
	ID          string
	Name        string
	Description string
	Races       []Race
}
