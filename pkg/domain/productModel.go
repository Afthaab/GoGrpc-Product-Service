package domain

type Size struct {
	ID    string `bson:"_id,omitempty" json:"id"`
	Name  string `bson:"name" json:"email"`
	Price string `bson:"price" json:"size"`
}
