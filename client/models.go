package client

type Pokemon struct {
	ID             int    `json:"id"`
	Name           string `json:"name"`
	BaseExperience int    `json:"base_experience"`
	Height         int    `json:"height"`
	Weight         int    `json:"weight"`
}

/*
id:35
name:"clefairy"
base_experience:113
height:6
is_default:true
order:56
weight:75
*/
