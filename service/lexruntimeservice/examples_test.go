// Code generated by private/model/cli/gen-api/main.go. DO NOT EDIT.

package lexruntimeservice_test

import (
	"bytes"
	"fmt"
	"time"

	"github.com/yangb8/aws-sdk-go/aws"
	"github.com/yangb8/aws-sdk-go/aws/session"
	"github.com/yangb8/aws-sdk-go/service/lexruntimeservice"
)

var _ time.Duration
var _ bytes.Buffer

func ExampleLexRuntimeService_PostContent() {
	sess := session.Must(session.NewSession())

	svc := lexruntimeservice.New(sess)

	params := &lexruntimeservice.PostContentInput{
		BotAlias:          aws.String("BotAlias"),             // Required
		BotName:           aws.String("BotName"),              // Required
		ContentType:       aws.String("HttpContentType"),      // Required
		InputStream:       bytes.NewReader([]byte("PAYLOAD")), // Required
		UserId:            aws.String("UserId"),               // Required
		Accept:            aws.String("Accept"),
		SessionAttributes: aws.JSONValue{"key": "value"},
	}
	resp, err := svc.PostContent(params)

	if err != nil {
		// Print the error, cast err to awserr.Error to get the Code and
		// Message from an error.
		fmt.Println(err.Error())
		return
	}

	// Pretty-print the response data.
	fmt.Println(resp)
}

func ExampleLexRuntimeService_PostText() {
	sess := session.Must(session.NewSession())

	svc := lexruntimeservice.New(sess)

	params := &lexruntimeservice.PostTextInput{
		BotAlias:  aws.String("BotAlias"), // Required
		BotName:   aws.String("BotName"),  // Required
		InputText: aws.String("Text"),     // Required
		UserId:    aws.String("UserId"),   // Required
		SessionAttributes: map[string]*string{
			"Key": aws.String("String"), // Required
			// More values...
		},
	}
	resp, err := svc.PostText(params)

	if err != nil {
		// Print the error, cast err to awserr.Error to get the Code and
		// Message from an error.
		fmt.Println(err.Error())
		return
	}

	// Pretty-print the response data.
	fmt.Println(resp)
}
