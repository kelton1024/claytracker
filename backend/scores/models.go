package scores

type scoreAddRequest struct {
	Station int `json:"course"`
	Scores  int `json:"score"`
}

type ScoreGetRequest struct {
	Key string `json:"key"`
}
