package main

import (
	"encoding/json"
	"errors"
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
	GeneratedName string `json:"generatedName"`
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
	}

	x, err := process(payload)

	if err != nil {
		resStatusCode = 400
		resBody = err.Error()
	}

	jsonData, err := json.Marshal(&x)

	if err != nil {
		resStatusCode = 500
		resBody = err.Error()
	}

	resBody = string(jsonData)

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

func process(p *RequestPayload) (*ResponsePayload, error) {
	switch p.Kind {
	default:
		return nil, errors.New("unknown kind")
	}
}
