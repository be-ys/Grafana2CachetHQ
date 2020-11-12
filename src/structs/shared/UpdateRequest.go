package shared

// Grafana Input
type UpdateRequest struct {
	State string `json:"state"`
	Tags map[string]string `json:"tags"`
	RuleId int `json:"ruleId"`
}


// Cachet output
type CachetUpdate struct {
	Status int `json:"status"`
}

type CachetPageable struct {
	Data []CachetContent `json:"data"`
}

type CachetContent struct {
	Id int `json:"id"`
	Name string `json:"name"`
	Status int `json:"status"`
	Tags map[string]string `json:"tags"`
}
