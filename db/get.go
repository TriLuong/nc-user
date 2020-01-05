package db

import (
	"context"
	"log"
	"time"

	"github.com/Triluong/nc-student/config"
	jwt "github.com/dgrijalva/jwt-go"
	"golang.org/x/crypto/bcrypt"
	"gopkg.in/mgo.v2/bson"
)

// func GetStudents() (*[]Student, error) {
// 	collection := Client.Database("nc-user").Collection("users")
// 	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
// 	defer cancel()
// 	filter := bson.M{}

// 	cur, err := collection.Find(ctx, filter)
// 	if err != nil {
// 		log.Printf("Find error: %v", err)
// 		return nil, err
// 	}
// 	ctx, cancel = context.WithTimeout(context.Background(), 5*time.Second)

// 	var students []Student
// 	err = cur.All(ctx, &students)
// 	if err != nil {
// 		log.Printf("Find error: %v", err)
// 		return nil, err
// 	}
// 	return &students, nil
// }

// func GetStudentByID(id string) (Student, error) {
// 	collection := Client.Database("nc-user").Collection("users")
// 	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
// 	defer cancel()
// 	idInt, error := strconv.Atoi(id)
// 	if error != nil {
// 		return Student{}, error
// 	}
// 	filter := bson.M{"id": idInt}
// 	var result Student
// 	err := collection.FindOne(ctx, filter).Decode(&result)
// 	if err != nil {
// 		return Student{}, err
// 	}
// 	return result, nil
// }

// func DeleteStudentById(id string) (Student, error) {
// 	collection := Client.Database("nc-user").Collection("users")
// 	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
// 	defer cancel()
// 	idInt, error := strconv.Atoi(id)
// 	if error != nil {
// 		return Student{}, error
// 	}
// 	filter := bson.M{"id": idInt}
// 	var result Student
// 	err := collection.FindOneAndDelete(ctx, filter).Decode(&result)
// 	if err != nil {
// 		return Student{}, err
// 	}
// 	return result, nil
// }

// func SerchStudentSimple(req StudentSearchRequest) (*[]Student, error) {
// 	collection := Client.Database("nc-user").Collection("users")
// 	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
// 	defer cancel()
// 	var filter bson.M
// 	bs, err := json.Marshal(req)

// 	if err != nil {
// 		log.Printf("Find error: %v", err)
// 		return nil, err
// 	}
// 	err = json.Unmarshal(bs, &filter)
// 	if err != nil {
// 		log.Printf("Find error: %v", err)
// 		return nil, err
// 	}
// 	fmt.Println(filter)
// 	cur, err := collection.Find(ctx, filter)
// 	if err != nil {
// 		log.Printf("Find error: %v", err)
// 		return nil, err
// 	}
// 	ctx, cancel = context.WithTimeout(context.Background(), 5*time.Second)

// 	var students []Student
// 	err = cur.All(ctx, &students)
// 	if err != nil {
// 		log.Printf("Find error: %v", err)
// 		return nil, err
// 	}
// 	return &students, nil
// }

func Login(req LoginForm) (LoginResponse, error) {
	collection := Client.Database("nc-user").Collection("users")
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()
	filter := bson.M{"email": req.Email}
	var result User
	err := collection.FindOne(ctx, filter).Decode(&result)
	if err != nil {
		log.Printf("Not FOUND Error")
		return LoginResponse{}, err
	}

	isVerify := comparePasswords(result.Password, req.Password)

	if !isVerify {
		log.Println("isVerify Error")
		return LoginResponse{}, err
	}

	token, error := generateToken(req)
	if error != nil {
		log.Println("token Error", error)
		return LoginResponse{}, error
	}

	log.Println(token)
	return LoginResponse{User: result, Token: token}, nil
}

func generateToken(creds LoginForm) (string, error) {
	expirationTime := time.Now().Add(5 * time.Minute)
	// Create the JWT claims, which includes the username and expiry time
	claims := &Claims{
		Email: creds.Email,
		StandardClaims: jwt.StandardClaims{
			// In JWT, the expiry time is expressed as unix milliseconds
			ExpiresAt: expirationTime.Unix(),
		},
	}

	// Declare the token with the algorithm used for signing, and the claims
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	// Create the JWT string
	tokenString, err := token.SignedString([]byte(config.Config.JWTSecret.JWTKey))
	if err != nil {
		return "", err
	}
	return tokenString, nil
}

func comparePasswords(hashedPwd string, plainPwd string) bool {
	// Since we'll be getting the hashed password from the DB it
	// will be a string so we'll need to convert it to a byte slice
	// byteHash, error := HashPassword(plainPwd)
	// if error != nil {
	// 	log.Print(error)
	// 	return false
	// }
	err := bcrypt.CompareHashAndPassword([]byte(hashedPwd), []byte(plainPwd))
	if err != nil {
		log.Println(err)
		return false
	}

	return true
}
