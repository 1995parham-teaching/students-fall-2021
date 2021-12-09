package model

type Student struct {
	ID          string  `bson:"id"`
	FirstName   string  `bson:"first_name"`
	LastName    string  `bson:"last_name"`
	Units       int     `bson:"units"`
	PassedUnits int     `bson:"passed_units"`
	Average     float64 `bson:"average"`
}
