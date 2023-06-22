package domain

type Product struct {
	Name     string `bson:"name" json:"name"`
	Quantity string `bson:"quantity" json:"quantity"`
}
