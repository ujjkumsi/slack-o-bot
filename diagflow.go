package main

import (
	"context"
	"errors"
	"fmt"
	"log"

	dialogflow "cloud.google.com/go/dialogflow/apiv2"
	dialogflowpb "google.golang.org/genproto/googleapis/cloud/dialogflow/v2"
)

func GetIntent(message string) string {
	projectID, sessionID, languageCode := "vision-196212", "BABA1234", "en"
	response, err := DetectIntentText(projectID, sessionID, message, languageCode)
	if err != nil {
		log.Fatal(err)
		return "I am unable to understand what you want to say!"
	}
	return response
}

// [START dialogflow_detect_intent_text]
func DetectIntentText(projectID, sessionID, text, languageCode string) (string, error) {
	ctx := context.Background()

	sessionClient, err := dialogflow.NewSessionsClient(ctx)
	if err != nil {
		return "", err
	}
	defer sessionClient.Close()

	if projectID == "" || sessionID == "" {
		return "", errors.New(fmt.Sprintf("Received empty project (%s) or session (%s)", projectID, sessionID))
	}

	sessionPath := fmt.Sprintf("projects/%s/agent/sessions/%s", projectID, sessionID)
	textInput := dialogflowpb.TextInput{Text: text, LanguageCode: languageCode}
	queryTextInput := dialogflowpb.QueryInput_Text{Text: &textInput}
	queryInput := dialogflowpb.QueryInput{Input: &queryTextInput}
	request := dialogflowpb.DetectIntentRequest{Session: sessionPath, QueryInput: &queryInput}

	response, err := sessionClient.DetectIntent(ctx, &request)
	if err != nil {
		log.Printf("Session Path (%s)", sessionPath)
		return "", err
	}

	queryResult := response.GetQueryResult()
	fulfillmentText := queryResult.GetFulfillmentText()
	return fulfillmentText, nil
}
