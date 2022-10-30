package main

import (
	"fmt"
	"net/http"
	"net/url"

	"github.com/cucumber/godog"
	"github.com/xeipuuv/gojsonschema"

	"bdd-apitest-godog/util"
)

var (
	adrsRespBody   string
	adrsStatusCode int
)

func iSendRequestWithTheQueryParam(zipcode string) (err error) {
	reqParams := url.Values{}
	reqParams.Add("zipcode", zipcode)
	actualResp, err := http.PostForm("https://zipcloud.ibsnet.co.jp/api/search", reqParams)
	if err != nil {
		return
	}

	adrsStatusCode = actualResp.StatusCode
	adrsRespBody, err = util.GetStrRespBody(actualResp.Body)
	if err != nil {
		return
	}

	return
}

func theResponseCodeShouldBe(statusCode int) (err error) {
	if statusCode != adrsStatusCode {
		return fmt.Errorf("expected status code was %v, but actually %v", statusCode, adrsStatusCode)
	}
	return
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
	ctx.Step(`^I send request with the query param "([^"]*)"$`, iSendRequestWithTheQueryParam)
	ctx.Step(`^the response code should be (\d^)`, theResponseCodeShouldBe)
	ctx.Step(`^the fields of the response JSON should meet the restriction:$`, theFieldsOfTheResponseJSONShouldMeetTheRestriction)
	ctx.Step(`^the response JSON should match the JSON file "([^"]*)"$`, theResponseJSONShouldMatchTheJSONFile)
	ctx.Step(`^the response JSON should match the schema:$`, theResponseJSONShouldMatchTheSchema)
}
