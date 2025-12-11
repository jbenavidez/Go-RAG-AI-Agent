package models

type Document struct {
	Text          string
	EmbeddingText []float64
	ProjectName   string
	Description   string
	Distance      float64
}
