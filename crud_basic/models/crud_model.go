package models

type Customer struct {
	Id       string `json:"id" bson:"id"`
	Name     string `json:"name" bson:"name"`
	Password string `json:"password" bson:"password"`
	Contact  string `json:"contact" bson:"contact"`
}
