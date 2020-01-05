package db

type Student struct {
	// MongoID   primitive.ObjectID `json:"_id,omitempty" bson:"_id,omitempty"`
	ID        int    `json:"id,omitempty" bson:"id,omitempty"`
	FirstName string `json:"first_name,omitempty" bson:"first_name"`
	LastName  string `json:"last_name,omitempty" bson:"last_name"`
	Age       int    `json:"age,omitempty" bson:"age"`
	Email     string `json:"email,omitempty" bson:"email"`
}

type StudentSearchRequest struct {
	// ID        primitive.ObjectID `json:"id,omitempty" bson:"_id,omitempty"`
	FirstName string `json:"first_name,omitempty"`
	LastName  string `json:"last_name,omitempty"`
	Age       int    `json:"age,omitempty"`
	Email     string `json:"email,omitempty"`
}

type Counters struct {
	NameID string `json:"_id,omitempty" bson:"_id,omitempty"`
	ID     int    `json:"id,omitempty" bson:"id,omitempty"`
}