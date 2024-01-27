// package main

// import (
// 	"encoding/json"

// 	"github.com/aws/aws-lambda-go/events"
// 	"github.com/aws/aws-lambda-go/lambda"
// )

// type user struct {
// 	Name string `json:"name"`
// 	ID   int    `json:"id"`
// }

// func handler(request events.APIGatewayProxyRequest) (events.APIGatewayProxyResponse, error) {
// 	deven := user{Name: "Deven", ID: 1}
// 	alex := user{Name: "Alex", ID: 2}

// 	users := []user{deven, alex}

// 	body, err := json.Marshal(users)
// 	if err != nil {
// 		return events.APIGatewayProxyResponse{}, err
// 	}

// 	return events.APIGatewayProxyResponse{
// 		Body:       string(body),
// 		StatusCode: 200,
// 	}, nil
// }

// func main() {
// 	lambda.Start(handler)
// }

package main

import (
	"fmt"

	"github.com/aws/aws-lambda-go/lambda"
)

type User struct {
	Name string `json:"name"`
	ID   int    `json:"id"`
}

type Request struct {
	ID int `json:"id"`
}

func handler(request Request) ([]User, error) {

	fmt.Println("request", request)
	fmt.Println("requestID", request.ID)

	deven := User{Name: "Deven", ID: 1}
	alex := User{Name: "Alex", ID: 2}

	users := []User{deven, alex}
	// log.Fatal("lllllllllllllllllllll", request)

	fmt.Println("users", users)

	// Convertendo a lista de usuários para JSON

	fmt.Println("Deu bom")

	return users, nil
}

func main() {
	lambda.Start(handler)
}
