package pokemon

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"

	"github.com/gorilla/mux"

	errorhandler "github.com/lukebrobbs/pokemon-go/errorHandler"
)

type nameURL struct {
	Name string `json:"name"`
	URL  string `json:"url"`
}

type abilityItem struct {
	Ability  nameURL `json:"ability"`
	IsHidden bool    `json:"is_hidden"`
	Slot     int64   `json:"slot"`
}

type gameIndeces struct {
	GameIndex int64   `json:"game_index"`
	Version   nameURL `json:"version"`
}

type heldItem struct {
	Item nameURL `json:"item"`
}

type versionGroupDetail struct {
	LevelLearnedAt  int64   `json:"level_learned_at"`
	MoveLearnMethod nameURL `json:"move_learn_method"`
	VersionGroup    nameURL `json:"version_group"`
}

type move struct {
	Move                nameURL              `json:"move"`
	VersionGroupDetails []versionGroupDetail `json:"version_group_details"`
}

type sprites struct {
	BackDefault      string `json:"back_default"`
	BackFemale       string `json:"back_female"`
	BackShiny        string `json:"back_shiny"`
	BackShinyFemale  string `json:"back_shiny_female"`
	FrontDefault     string `json:"front_default"`
	FrontFemale      string `json:"front_female"`
	FrontShiny       string `json:"front_shiny"`
	FrontShinyFemale string `json:"front_shiny_female"`
}

type stat struct {
	BaseStat int64 `json:"base_stat"`
	Effort   int64 `json:"effort"`
}

type types struct {
	Slot int64   `json:"slot"`
	Type nameURL `json:"type"`
}

// Response is the shape of response from the poke api on the `/pokemon/{pokemon}` route
type Response struct {
	Abilities              []abilityItem `json:"abilities"`
	BaseExperience         int64         `json:"base_experience"`
	Height                 int64         `json:"height"`
	Forms                  []nameURL     `json:"forms"`
	GameIndices            []gameIndeces `json:"game_indices"`
	HeldItems              []heldItem    `json:"held_items"`
	ID                     int64         `json:"id"`
	IsDefault              bool          `json:"is_default"`
	LocationAreaEncounters string        `json:"location_area_encounters"`
	Moves                  []move        `json:"moves"`
	Name                   string        `json:"name"`
	Order                  int64         `json:"order"`
	Species                nameURL       `json:"species"`
	Sprites                sprites       `json:"sprites"`
	Stats                  []stat        `json:"stats"`
	Types                  []types       `json:"types"`
	Weight                 int64         `json:"weight"`
}

// GetPokemon formats the raw response from the PokeAPI
// and returns a Pokemons response struct
func GetPokemon(body []byte) (*Response, error) {
	var s = new(Response)
	err := json.Unmarshal(body, &s)
	if err != nil {
		fmt.Println("whoops:", err)
	}
	return s, err
}

// Finder Takes a pokemon name and returns data about that
// given pokemon.
func Finder(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodGet {
		errorhandler.BadRequest(w)
		return
	}
	vars := mux.Vars(r)
	var url = "https://pokeapi.co/api/v2/pokemon/" + vars["pokemon"]
	resp, err := http.Get(url)

	if resp.StatusCode != http.StatusOK {
		http.Error(w, "Unable to fetch Pokemon - Server responded with status "+resp.Status, http.StatusInternalServerError)
		return
	}
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
