package main

import (
	"fmt"

	"github.com/cucumber/godog"
	"github.com/xeipuuv/gojsonschema"
)

func iSendRequest() error {
	return godog.ErrPending
}

func theFieldsOfTheResponseJSONShouldMeetTheRestriction() error {
	return godog.ErrPending
}

func theResponseJSONShouldMatchTheJSONFile(arg1 string) error {
	return godog.ErrPending
}

func theResponseJSONShouldMatchTheSchema() error {
	schemaLoader := gojsonschema.NewStringLoader(`{"type":"string"}`)

	fmt.Println("---- Valid JSON")
	jsonLoader := gojsonschema.NewStringLoader(`"string"`)
	result, _ := gojsonschema.Validate(schemaLoader, jsonLoader)
	if result.Valid() {
		fmt.Println("Valid!")
	} else {
		fmt.Println("Invalid!")
	}

	fmt.Println("---- Invalid JSON")
	jsonLoader = gojsonschema.NewStringLoader(`{}`)
	result, _ = gojsonschema.Validate(schemaLoader, jsonLoader)
	if result.Valid() {
		fmt.Println("Valid!")
	} else {
		fmt.Println("Invalid!")
	}

	return godog.ErrPending
}

func InitializeScenarioAddress(ctx *godog.ScenarioContext) {
	ctx.Step(`^I send request$`, iSendRequest)
	ctx.Step(`^the fields of the response JSON should meet the restriction:$`, theFieldsOfTheResponseJSONShouldMeetTheRestriction)
	ctx.Step(`^the response JSON should match the JSON file "([^"]*)"$`, theResponseJSONShouldMatchTheJSONFile)
	ctx.Step(`^the response JSON should match the schema:$`, theResponseJSONShouldMatchTheSchema)
}
