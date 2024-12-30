package pokeapi

type Response struct {
	Count   int        `json:"count"`
	Next    string     `json:"next"`
	Prev    string     `json:"previous"`
	Results []Location `json:"results"`
}

type Location struct {
	Name string `json:"name"`
	Url  string `json:"url"`
}

type Links struct {
	NextUrl string
	PrevUrl string
}

type PokemonResult struct {
	Encounter_method_rates any                `json:"encounter_method_rates"`
	Game_index             int                `json:"game_index"`
	Id                     int                `json:"id"`
	Location               any                `json:"location"`
	Name                   string             `json:"name"`
	Names                  any                `json:"names"`
	Pokemon_encounters     []PokemonEncounter `json:"pokemon_encounters"`
}

type PokemonEncounter struct {
	Pokemon         PokemonInEncounter `json:"pokemon"`
	Version_details any                `json:"version_details"`
}

type PokemonInEncounter struct {
	Name string `json:"name"`
	Url  string `json:"url"`
}

type Pokemon struct {
	Abilities                any    `json:"abilities"`
	Base_experience          int    `json:"base_experience"`
	Cries                    any    `json:"cries"`
	Forms                    any    `json:"forms"`
	Game_indices             any    `json:"game_indices"`
	Height                   int    `json:"height"`
	Held_items               any    `json:"held_items"`
	Id                       int    `json:"id"`
	Is_default               bool   `json:"is_default"`
	Location_area_encounters string `json:"location_area_encounters"`
	Moves                    any    `json:"moves"`
	Name                     string `json:"name"`
	Order                    int    `json:"order"`
	Past_abilities           any    `json:"past_abilities"`
	Past_types               any    `json:"past_types"`
	Species                  any    `json:"species"`
	Sprites                  any    `json:"sprites"`
	Stats                    any    `json:"stats"`
	Types                    any    `json:"types"`
	Weight                   int    `json:"weight"`
}
