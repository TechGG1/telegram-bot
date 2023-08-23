package models

type Beer struct {
	ID               int           `json:"id"`
	Name             string        `json:"name"`
	Tagline          string        `json:"tagline"`
	Description      string        `json:"description"`
	FirstBrewed      string        `json:"first_brewed"`
	ImageURL         string        `json:"image_url"`
	ABV              float32       `json:"abv"`
	IBU              float32       `json:"ibu"`
	AttenuationLevel float32       `json:"attenuation_level"`
	Ingredients      Ingredients   `json:"ingredients"`
	FoodPairing      []FoodPairing `json:"food_pairing"`
	BrewersTips      string        `json:"brewers_tips"`
	ContributedBy    string        `json:"contributed_by"`
}

type Ingredients struct {
	Products []Product `json:"hops"`
	Yeast    string    `json:"yeast"`
}

type Product struct {
	Name      string `json:"name"`
	Add       string `json:"add"`
	Attribute string `json:"attribute"`
}

type FoodPairing string

type Taste string

var (
	Bitter  Taste = "bitter"
	Sweet   Taste = "sweet"
	Neutral Taste = "neutral"
)

var PollQuestionsMood = []string{"fine", "sad", "party"}
var PollQuestionsTaste = []string{"bitter", "sweet", "neutral", "spicy"}

type RespBeer struct {
	Name          string        `json:"name"`
	Tagline       string        `json:"tagline"`
	Description   string        `json:"description"`
	FirstBrewed   string        `json:"first_brewed"`
	ImageURL      string        `json:"image_url"`
	ABV           float32       `json:"abv"`
	IBU           float32       `json:"ibu"`
	FoodPairing   []FoodPairing `json:"food_pairing"`
	ContributedBy string        `json:"contributed_by"`
}
