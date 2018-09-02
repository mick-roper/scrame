package main

import (
	"encoding/json"
	"strings"

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
	resStatusCode := 200
	resBody := ""

	payload, err := decode(req.Body)

	if err != nil {
		resStatusCode = 400
		resBody = "bad input data"
	} else {
		resBody, resStatusCode = process(payload)
	}

	res := Response{
		Body:       resBody,
		StatusCode: resStatusCode,
	}

	return res, nil
}

func decode(s string) (*RequestPayload, error) {
	reader := strings.NewReader(s)

	var p RequestPayload

	err := json.NewDecoder(reader).Decode(&p)

	return &p, err
}

func process(p *RequestPayload) (string, int) {
	switch p.Kind {
	case "sci-fi":
		return "sith sloths", 200
	default:
		return "unknown kind", 400
	}
}
