package main

import (
	"context"
	"encoding/json"
	"fmt"
	"log"
	"os"

	"github.com/aws/aws-lambda-go/events"
	"github.com/aws/aws-lambda-go/lambda"
	"golang.org/x/oauth2/google"
	"google.golang.org/api/option"
	"google.golang.org/api/sheets/v4"
)

func main() {
	lambda.Start(HandleRequest)
}

type RequestBody struct {
	FirstName string `json:"First_Name"`
	LastName  string `json:"Last_Name"`
	Email     string `json:"Email"`
}

func HandleRequest(ctx context.Context, request events.APIGatewayProxyRequest) (events.APIGatewayProxyResponse, error) {
	var input RequestBody

	requestBodyString := string(request.Body)

	log.Printf("Received request body: %s", requestBodyString)
	// Unmarshal the JSON request body into a map
	if err := json.Unmarshal([]byte(request.Body), &input); err != nil {
		return events.APIGatewayProxyResponse{
			StatusCode: 400,
			Body:       "Invalid JSON input",
		}, err
	}

	// Access JSON data
	first_name := input.FirstName
	last_name := input.LastName
	email := input.Email

	// Your Lambda function logic goes here
	appender(first_name, last_name, email)

	// Create a response struct
	response := map[string]string{
		"message": fmt.Sprintf("Received JSON"),
	}
	// Marshal the response struct to JSON
	responseBody, err := json.Marshal(response)
	if err != nil {
		return events.APIGatewayProxyResponse{
			StatusCode: 500,
			Body:       "Internal Server Error",
		}, err
	}

	return events.APIGatewayProxyResponse{
		StatusCode: 200,
		Headers:    map[string]string{"Content-Type": "application/json"},
		Body:       string(responseBody),
	}, nil
}

func appender(value1, value2, value3 string) {
	// Load your Google Sheets API credentials JSON file
	credsFile := os.Getenv("GOOGLE_CRED")
	creds, err := google.CredentialsFromJSON(context.Background(), []byte(credsFile), sheets.SpreadsheetsScope)
	if err != nil {
		log.Fatalf("Unable to read credentials: %v", err)
	}

	// Create a Google Sheets Service client using option.WithHTTPClient
	sheetsService, err := sheets.NewService(context.Background(), option.WithCredentials(creds))
	if err != nil {
		log.Fatalf("Unable to create Sheets service: %v", err)
	}

	// The ID of the spreadsheet you want to append data to
	spreadsheetID := os.Getenv("GOOGLE_SPREADSHEET_ID")

	// The name of the sheet you want to append data to
	sheetName := os.Getenv("GOOGLE_SHEET_NAME")

	// Data to append to the next row
	data := []interface{}{value1, value2, value3}

	// Get the current values in the sheet
	resp, err := sheetsService.Spreadsheets.Values.Get(spreadsheetID, sheetName).Context(context.Background()).Do()
	if err != nil {
		log.Fatalf("Unable to retrieve data from sheet: %v", err)
	}

	// Calculate the next empty row
	nextRow := len(resp.Values) + 1

	// Create a ValueRange object to append the data to the next row
	vr := &sheets.ValueRange{
		Values: [][]interface{}{data},
	}

	// Append data to the Google Sheet
	_, err = sheetsService.Spreadsheets.Values.Update(spreadsheetID, sheetName+"!A"+fmt.Sprint(nextRow), vr).ValueInputOption("RAW").Context(context.Background()).Do()
	if err != nil {
		log.Fatalf("Unable to append data to sheet: %v", err)
	}

	fmt.Println("Data appended successfully to Google Sheet.")
}
