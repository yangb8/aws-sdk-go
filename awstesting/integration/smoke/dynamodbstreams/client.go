// +build integration

//Package dynamodbstreams provides gucumber integration tests support.
package dynamodbstreams

import (
	"github.com/yangb8/aws-sdk-go/awstesting/integration/smoke"
	"github.com/yangb8/aws-sdk-go/service/dynamodbstreams"
	"github.com/gucumber/gucumber"
)

func init() {
	gucumber.Before("@dynamodbstreams", func() {
		gucumber.World["client"] = dynamodbstreams.New(smoke.Session)
	})
}
