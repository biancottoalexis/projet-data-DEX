package models

type Fighter struct {
	Name         string `json:"name"`
	Nickname     string `json:"nickname"`
	Category     string `json:"category"`
	Wins         string `json:"wins"`
	Losses       string `json:"losses"`
	Draws        string `json:"draws"`
	ImgUrl       string `json:"imgUrl"`
	Status       string `json:"status"`
	PlaceOfBirth string `json:"placeOfBirth"`
	TrainsAt     string `json:"trainsAt"`
	FightingStyle string `json:"fightingStyle"`
	Age          string `json:"age"`
	Height       string `json:"height"`
	Weight       string `json:"weight"`
	OctagonDebut string `json:"octagonDebut"`
	Reach        string `json:"reach"`
	LegReach     string `json:"legReach"`
}