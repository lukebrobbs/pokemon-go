package pokemon

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"

	errorhandler "github.com/lukebrobbs/pokemonServer/errorHandler"
)

type NameUrl struct {
	Name string `json: string`
	Url  string `json: url`
}

type AbilityItem struct {
	Ability   NameUrl `json: ability`
	Is_hidden bool    `json: is_hidden`
	Slot      int64   `json: slot`
}

type GameIndeces struct {
	Game_index int64   `json: game_index`
	Version    NameUrl `json: version`
}

type HeldItem struct {
	Item NameUrl `json: item`
}

type VersionGroupDetail struct {
	Level_learned_at  int64   `json: level_learned_at`
	Move_learn_method NameUrl `json: move_learn_method`
	Version_group     NameUrl `json: version_group`
}

type Move struct {
	Move                  NameUrl              `json: move`
	Version_group_details []VersionGroupDetail `json: version_group_details`
}

type Sprites struct {
	Back_default       string `json: back_default`
	Back_female        string `json: back_female`
	Back_shiny         string `json: back_shiny`
	Back_shiny_female  string `json: back_shiny_female`
	Front_default      string `json: front_default`
	Front_female       string `json: front_female`
	Front_shiny        string `json: front_shiny`
	Front_shiny_female string `json: front_shiny_female`
}

type Stat struct {
	Base_stat int64 `json: base_stat`
	Effort    int64 `json: effort`
}

type Type struct {
	Slot int64   `json: slot`
	Type NameUrl `json: type`
}

type PokemonReponse struct {
	Abilities                []AbilityItem `json: abilities`
	Base_experience          int64         `json: base_experience`
	Height                   int64         `json: height`
	Forms                    []NameUrl     `json: forms`
	Game_indices             []GameIndeces `json: game_indices`
	Held_items               []HeldItem    `json: held_items`
	Id                       int64         `json: id`
	Is_default               bool          `json: is_default`
	Location_area_encounters string        `json: location_area_encounters`
	Moves                    []Move        `json: moves`
	Name                     string        `json: name`
	Order                    int64         `json: order`
	Species                  NameUrl       `json: species`
	Sprites                  Sprites       `json: sprites`
	Stats                    []Stat        `json: stats`
	Types                    []Type        `json: types`
	Weight                   int64         `json: weight`
}

// GetPokemon formats the raw response from the PokeAPI
// and returns a Pokemons response struct
func GetPokemon(body []byte) (*PokemonReponse, error) {
	var s = new(PokemonReponse)
	err := json.Unmarshal(body, &s)
	if err != nil {
		fmt.Println("whoops:", err)
	}
	return s, err
}

// Finder Takes a pokemon name and returns data about that
// given pokemon.
func Finder(w http.ResponseWriter, r *http.Request) {
	if r.Method != "POST" {
		errorhandler.BadRequest(w)
		return
	}

	resp, err := http.Get("https://pokeapi.co/api/v2/pokemon/ditto")

	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	defer resp.Body.Close()
	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	p, err := GetPokemon([]byte(body))
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	js, err := json.Marshal(p)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	w.Write(js)

}
