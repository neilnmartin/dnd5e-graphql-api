// Code generated by github.com/99designs/gqlgen, DO NOT EDIT.

package model

type Ability struct {
	ID          *string   `json:"id"`
	Name        *string   `json:"name"`
	Description []*string `json:"description"`
	Skills      []*Skill  `json:"skills"`
}

type AbilityBonus struct {
	Ability *Ability `json:"ability"`
	Bonus   *string  `json:"bonus"`
}

type Character struct {
	ID        *string      `json:"id"`
	User      *User        `json:"user"`
	Name      string       `json:"name"`
	Age       int          `json:"age"`
	Level     int          `json:"level"`
	Race      *Race        `json:"race"`
	SubRace   *SubRace     `json:"subRace"`
	Class     *Class       `json:"class"`
	SubClass  *SubClass    `json:"subClass"`
	Traits    []*Trait     `json:"traits"`
	Equipment []*Equipment `json:"equipment"`
}

type Class struct {
	ID                 *string             `json:"id"`
	Name               *string             `json:"name"`
	HitDie             *int                `json:"hitDie"`
	ProficiencyChoices *ProficiencyChoices `json:"proficiencyChoices"`
	Proficiencies      []*Proficiency      `json:"proficiencies"`
	SavingThrows       []*Ability          `json:"savingThrows"`
	StartingEquipment  *StartingEquipment  `json:"startingEquipment"`
	ClassLevels        []*ClassLevel       `json:"classLevels"`
	SubClasses         []*SubClass         `json:"subClasses"`
	Spellcasting       *Ability            `json:"spellcasting"`
	URL                *string             `json:"url"`
}

type ClassLevel struct {
	Level    *int       `json:"level"`
	Class    *Class     `json:"class"`
	Features []*Feature `json:"features"`
}

type CreateCharacterInput struct {
	Name  *string `json:"name"`
	Race  *string `json:"race"`
	Class *string `json:"class"`
}

type Equipment struct {
	Name     *string `json:"name"`
	Category *string `json:"category"`
}

type Feature struct {
	ID          *string   `json:"id"`
	Name        *string   `json:"name"`
	Description []*string `json:"description"`
}

type Game struct {
	ID    string `json:"id"`
	Owner *User  `json:"owner"`
}

type Language struct {
	ID   string `json:"id"`
	Name string `json:"name"`
}

type LoginInput struct {
	Email    string `json:"email"`
	Password string `json:"password"`
}

type LoginResponse struct {
	User  *User  `json:"user"`
	Token string `json:"token"`
}

type Name struct {
	GivenName  string  `json:"givenName"`
	FamilyName string  `json:"familyName"`
	Formatted  *string `json:"formatted"`
}

type NameInput struct {
	GivenName  string `json:"givenName"`
	FamilyName string `json:"familyName"`
}

type Party struct {
	ID         string       `json:"id"`
	GameID     *string      `json:"gameId"`
	Characters []*Character `json:"characters"`
}

type Proficiency struct {
	Name *string `json:"name"`
}

type ProficiencyChoices struct {
	Choose *int           `json:"choose"`
	Type   *string        `json:"type"`
	From   []*Proficiency `json:"from"`
}

type Race struct {
	ID                         *string       `json:"id"`
	Name                       *string       `json:"name"`
	Description                []*string     `json:"description"`
	AbilityBonuses             *AbilityBonus `json:"abilityBonuses"`
	Perks                      *string       `json:"perks"`
	SubRaces                   []*SubRace    `json:"subRaces"`
	Age                        *string       `json:"age"`
	Speed                      *int          `json:"speed"`
	Size                       *string       `json:"size"`
	SizeDescription            *string       `json:"sizeDescription"`
	StartingProficiencies      []*string     `json:"startingProficiencies"`
	StartingProficiencyOptions []*string     `json:"startingProficiencyOptions"`
	Languages                  []*Language   `json:"languages"`
	LanguageOptions            []*Language   `json:"languageOptions"`
	LanguageDescription        *string       `json:"languageDescription"`
	Traits                     []*Trait      `json:"traits"`
	TraitOptions               []*Trait      `json:"traitOptions"`
}

type Skill struct {
	ID   *string `json:"id"`
	Name *string `json:"name"`
}

type Spell struct {
	ID            string    `json:"id"`
	Name          string    `json:"name"`
	Description   []string  `json:"description"`
	HigherLevel   []*string `json:"higherLevel"`
	Range         string    `json:"range"`
	Components    []*string `json:"components"`
	Material      *string   `json:"material"`
	Ritual        *bool     `json:"ritual"`
	Duration      *string   `json:"duration"`
	Concentration *string   `json:"concentration"`
}

type StartingEquipment struct {
	Class                    *Class       `json:"class"`
	StartingEquipment        []*Equipment `json:"startingEquipment"`
	StartingEquipmentOptions []*Equipment `json:"startingEquipmentOptions"`
}

type SubClass struct {
	ID          *string `json:"id"`
	Name        *string `json:"name"`
	Description *string `json:"description"`
	Flavor      *string `json:"flavor"`
	Class       *Class  `json:"class"`
}

type SubRace struct {
	ID   string `json:"id"`
	Race *Race  `json:"race"`
	Name string `json:"name"`
}

type Trait struct {
	ID          *string    `json:"id"`
	Races       []*Race    `json:"races"`
	SubRaces    []*SubRace `json:"subRaces"`
	Name        *string    `json:"name"`
	Description *string    `json:"description"`
}

type User struct {
	ID    string `json:"id"`
	Email string `json:"email"`
	Name  *Name  `json:"name"`
}

type UserInput struct {
	Name     *NameInput `json:"name"`
	Email    string     `json:"email"`
	Password string     `json:"password"`
}

type CreatePartyInput struct {
	GameID       *string   `json:"gameId"`
	Name         *string   `json:"name"`
	CharacterIds []*string `json:"characterIds"`
}
