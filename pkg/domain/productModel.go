package domain

type Size struct {
	Id    string `bson:"_id,omitempty" json:"id"`
	Name  string `bson:"name" json:"name"`
	Price string `bson:"price" json:"price"`
}

type Category struct {
	Id           string `bson:"_id,omitempty" json:"id"`
	Categoryname string `bson:"categoryname" json:"categoryname"`
	Imageurl     string `bson:"imageurl" json:"imageurl"`
}

type Products struct {
	ID           string   `bson:"_id,omitempty" json:"_id"`
	Productname  string   `bson:"productname" json:"productname"`
	Calories     string   `bson:"calories" json:"calories"`
	Availibility bool     `bson:"availibilty" json:"availibility"`
	Categoryid   string   `bson:"categoryid" json:"categoryid"`
	Typeid       string   `bson:"typeid" json:"categoryname"`
	Baseprice    float64  `bson:"baseprice" json:"baseprice"`
	Sizeid       []string `bson:"sizesid" json:"sizeid"`
	Imageurls    []string `bson:"imageurls" json:"imageurls"`
}

type Foodtype struct {
	Id       string `bson:"_id,omitempty" json:"id"`
	Foodtype string `bson:"foodtype" json:"foodtype"`
	Imageurl string `bson:"imageurl" json:"imageurl"`
}
