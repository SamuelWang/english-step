package models

import (
	"time"
)

// SynonymExplanation represents an explanation of synonyms for a set of vocabulary words.
type SynonymExplanation struct {
	ID           uint           `gorm:"primaryKey"`
	Vocabularies []string       `gorm:"type:json"` // Array or JSON of vocabulary words (unique constraint on sorted values)
	Explanation  string         `gorm:"type:text"`
	CreatedAt    time.Time      `gorm:"autoCreateTime"`
}

// TableName overrides the table name used by GORM
func (SynonymExplanation) TableName() string {
	return "synonym_explanations"
}
