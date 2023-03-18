package models

type Paraphrase struct {
	Text string `json:"text" binding:"required"`
}
