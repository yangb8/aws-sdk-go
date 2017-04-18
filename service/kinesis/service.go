// Code generated by private/model/cli/gen-api/main.go. DO NOT EDIT.

package kinesis

import (
	"github.com/yangb8/aws-sdk-go/aws"
	"github.com/yangb8/aws-sdk-go/aws/client"
	"github.com/yangb8/aws-sdk-go/aws/client/metadata"
	"github.com/yangb8/aws-sdk-go/aws/request"
	"github.com/yangb8/aws-sdk-go/aws/signer/v4"
	"github.com/yangb8/aws-sdk-go/private/protocol/jsonrpc"
)

// Amazon Kinesis Streams is a managed service that scales elastically for real
// time processing of streaming big data.
// The service client's operations are safe to be used concurrently.
// It is not safe to mutate any of the client's properties though.
// Please also see https://docs.aws.amazon.com/goto/WebAPI/kinesis-2013-12-02
type Kinesis struct {
	*client.Client
}

// Used for custom client initialization logic
var initClient func(*client.Client)

// Used for custom request initialization logic
var initRequest func(*request.Request)

// Service information constants
const (
	ServiceName = "kinesis"   // Service endpoint prefix API calls made to.
	EndpointsID = ServiceName // Service ID for Regions and Endpoints metadata.
)

// New creates a new instance of the Kinesis client with a session.
// If additional configuration is needed for the client instance use the optional
// aws.Config parameter to add your extra config.
//
// Example:
//     // Create a Kinesis client from just a session.
//     svc := kinesis.New(mySession)
//
//     // Create a Kinesis client with additional configuration
//     svc := kinesis.New(mySession, aws.NewConfig().WithRegion("us-west-2"))
func New(p client.ConfigProvider, cfgs ...*aws.Config) *Kinesis {
	c := p.ClientConfig(EndpointsID, cfgs...)
	return newClient(*c.Config, c.Handlers, c.Endpoint, c.SigningRegion, c.SigningName)
}

// newClient creates, initializes and returns a new service client instance.
func newClient(cfg aws.Config, handlers request.Handlers, endpoint, signingRegion, signingName string) *Kinesis {
	svc := &Kinesis{
		Client: client.New(
			cfg,
			metadata.ClientInfo{
				ServiceName:   ServiceName,
				SigningName:   signingName,
				SigningRegion: signingRegion,
				Endpoint:      endpoint,
				APIVersion:    "2013-12-02",
				JSONVersion:   "1.1",
				TargetPrefix:  "Kinesis_20131202",
			},
			handlers,
		),
	}

	// Handlers
	svc.Handlers.Sign.PushBackNamed(v4.SignRequestHandler)
	svc.Handlers.Build.PushBackNamed(jsonrpc.BuildHandler)
	svc.Handlers.Unmarshal.PushBackNamed(jsonrpc.UnmarshalHandler)
	svc.Handlers.UnmarshalMeta.PushBackNamed(jsonrpc.UnmarshalMetaHandler)
	svc.Handlers.UnmarshalError.PushBackNamed(jsonrpc.UnmarshalErrorHandler)

	// Run custom client initialization if present
	if initClient != nil {
		initClient(svc.Client)
	}

	return svc
}

// newRequest creates a new request for a Kinesis operation and runs any
// custom request initialization.
func (c *Kinesis) newRequest(op *request.Operation, params, data interface{}) *request.Request {
	req := c.NewRequest(op, params, data)

	// Run custom request initialization if present
	if initRequest != nil {
		initRequest(req)
	}

	return req
}
