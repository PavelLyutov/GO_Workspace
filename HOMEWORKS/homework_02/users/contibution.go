package users

type Contribution struct{
	Name string `json:"username"`
	Count []Count `json:"contributions"`
}

type Count struct {
	Week int `json:"week"`
	Days []Day `json:"days"`
}


type Day struct {
	Count int `json:"count"`

}
