package entities

type Beer struct {
	Id       int64   `bson:"_id"`
	Name     string  `bson:"name"`
	Brewery  string  `bson:"brewery"`
	Country  string  `bson:"country"`
	Price    float64 `bson:"price"`
	Currency string  `bson:"currency"`
}
