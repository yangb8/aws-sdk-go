// +build integration

//Package support provides gucumber integration tests support.
package support

import (
	"github.com/yangb8/aws-sdk-go/awstesting/integration/smoke"
	"github.com/yangb8/aws-sdk-go/service/support"
	"github.com/gucumber/gucumber"
)

func init() {
	gucumber.Before("@support", func() {
		gucumber.World["client"] = support.New(smoke.Session)
	})
}
