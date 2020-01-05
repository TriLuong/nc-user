package db

import jwt "github.com/dgrijalva/jwt-go"

type LoginForm struct {
	Email    string `json:"email" bson:"email"`
	Password string `json:"password" bson:"password"`
}
type User struct {
	// MongoID   primitive.ObjectID `json:"_id,omitempty" bson:"_id,omitempty"`
	ID        int    `json:"id,omitempty" bson:"id,omitempty"`
	FirstName string `json:"first_name,omitempty" bson:"first_name"`
	LastName  string `json:"last_name,omitempty" bson:"last_name"`
	ClassName string `json:"class_name,omitempty" bson:"class_name"`
	Age       int    `json:"age,omitempty" bson:"age"`
	Email     string `json:"email,omitempty" bson:"email"`
	Password  string `json:"password" bson:"password"`
}

type Claims struct {
	Email string `json:"email"`
	jwt.StandardClaims
}
type LoginResponse struct {
	User  User
	Token string `json:"token" bson:"tone"`
}

type Counters struct {
	NameID string `json:"_id,omitempty" bson:"_id,omitempty"`
	ID     int    `json:"id,omitempty" bson:"id,omitempty"`
}
