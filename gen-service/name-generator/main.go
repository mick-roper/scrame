package main

import (
	"github.com/aws/aws-lambda-go/events"
	"github.com/aws/aws-lambda-go/lambda"
)

// Request from a client
type Request events.APIGatewayProxyRequest

// Response to a request
type Response events.APIGatewayProxyResponse

// RequestPayload sent from a client
type RequestPayload struct {
	Kind string `json:"kind"`
}

// ResponsePayload returned to the client
type ResponsePayload struct {
	Name string `json:"name"`
}

func main() {
	lambda.Start(Handler)
}

// Handler that handles the request
func Handler(req Request) (Response, error) {
	res := Response{
		Body:       req.Body,
		StatusCode: 200,
	}

	return res, nil
}
