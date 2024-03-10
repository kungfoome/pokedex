package pokeapi

type EncounterMethodRate struct {
	EncounterMethod NamedAPIResource          `json:"encounter_method"`
	VersionDetails  []EncounterVersionDetails `json:"version_details"`
}

type EncounterVersionDetails struct {
	Rate    int              `json:"rate"`
	Version NamedAPIResource `json:"version"`
}

type Name struct {
	Name     string           `json:"name"`
	Language NamedAPIResource `json:"language"`
}

type PokemonEncounter struct {
	Pokemon        NamedAPIResource         `json:"pokemon"`
	VersionDetails []VersionEncounterDetail `json:"version_details"`
}

type VersionEncounterDetail struct {
	Version          NamedAPIResource `json:"version"`
	MaxChannce       int              `json:"max_chance"`
	EncounterDetails []Encounter      `json:"encounter_details"`
}

type Encounter struct {
	MinLevel       int                `json:"min_level"`
	MaxLevel       int                `json:"max_level"`
	ConditionValue []NamedAPIResource `json:"condition_values"`
	Chance         int                `json:"chance"`
	Methon         NamedAPIResource   `json:"method"`
}

type NamedAPIResource struct {
	Name string `json:"name"`
	Url  string `json:"url"`
}
