package model

type WordsSearchInput struct {
	Letters         string `json:"letters" binding:"required"`
	AuxiliaryLetter string `json:"auxiliaryLetter"`
	Mode            int    `json:"mode" binding:"required"`
}
