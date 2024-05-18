package domain

type Pokemon struct {
	ID             int
	Identifier     string
	SpeciesID      int
	Height         int
	Weight         int
	BaseExperience int
	Order          int
	IsDefault      bool
}
