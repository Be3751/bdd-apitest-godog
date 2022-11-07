package main

import (
	"bdd-apitest-godog/util"
	"fmt"
	"net/http"
	"net/url"

	"github.com/cucumber/godog"
	"github.com/xeipuuv/gojsonschema"
)

type address struct {
	respBody   string
	statusCode int
}

func (adrs *address) iSendRequestWithTheQueryParam(zipcode string) (err error) {
	reqParams := url.Values{}
	reqParams.Add("zipcode", zipcode)
	actualResp, err := http.PostForm("https://zipcloud.ibsnet.co.jp/api/search", reqParams)
	if err != nil {
		return
	}

	adrs.statusCode = actualResp.StatusCode
	adrs.respBody, err = util.GetStrRespBody(actualResp.Body)
	if err != nil {
		return
	}

	return
}

func (adrs *address) theStatusCodeShouldBe(statusCode int) (err error) {
	return util.AssertEqual(statusCode, adrs.statusCode)
}

func (adrs *address) theFieldsOfTheResponseJSONShouldMeetTheRestriction() error {
	return godog.ErrPending
}

func (adrs *address) theResponseJSONShouldMatchTheJSONFile(arg1 string) error {
	return godog.ErrPending
}

func (adrs *address) theResponseJSONShouldMatchTheSchema() error {
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

func InitializeScenario(ctx *godog.ScenarioContext) {
	adrs := address{}

	ctx.Step(`^I send request with the query param "([^"]*)"$`, adrs.iSendRequestWithTheQueryParam)
	ctx.Step(`^the status code should be (\d+)$`, adrs.theStatusCodeShouldBe)
	ctx.Step(`^the fields of the response JSON should meet the restriction:$`, adrs.theFieldsOfTheResponseJSONShouldMeetTheRestriction)
	ctx.Step(`^the response JSON should match the JSON file "([^"]*)"$`, adrs.theResponseJSONShouldMatchTheJSONFile)
	ctx.Step(`^the response JSON should match the schema:$`, adrs.theResponseJSONShouldMatchTheSchema)
}
