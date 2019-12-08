package handler

// func UpdateStudent(c echo.Context) error {
// 	type Student struct {
// 		FirstName string `json:"first_name"`
// 		LastName  string `json:"last_name"`
// 		Age       int    `json:"age"`
// 		Email     string `json:"email"`
// 	}

// 	collection := db.Client.Database("nc-student").Collection("student")
// 	ctx, _ := context.WithTimeout(context.Background(), 5*time.Second)
// 	var student Student
// 	res, err := collection.InsertOne(ctx, student)
// 	if err != nil {
// 		log.Faltalf(err)
// 	}

// 	return c.JSON()
// }
