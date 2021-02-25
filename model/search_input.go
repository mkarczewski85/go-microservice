package model

type WordsSearchInput struct {
	Letters         string `json:"letters" binding:"required"`
	AuxiliaryLetter string `json:"auxiliaryLetter" binding:"required"`
	Option          uint8  `json:"option" binding:"required"`
}
