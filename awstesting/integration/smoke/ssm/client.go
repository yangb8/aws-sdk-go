// +build integration

//Package ssm provides gucumber integration tests support.
package ssm

import (
	"github.com/yangb8/aws-sdk-go/awstesting/integration/smoke"
	"github.com/yangb8/aws-sdk-go/service/ssm"
	"github.com/gucumber/gucumber"
)

func init() {
	gucumber.Before("@ssm", func() {
		gucumber.World["client"] = ssm.New(smoke.Session)
	})
}
