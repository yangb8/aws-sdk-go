// Code generated by private/model/cli/gen-api/main.go. DO NOT EDIT.

package opsworkscm

import (
	"github.com/yangb8/aws-sdk-go/aws"
	"github.com/yangb8/aws-sdk-go/aws/client"
	"github.com/yangb8/aws-sdk-go/aws/client/metadata"
	"github.com/yangb8/aws-sdk-go/aws/request"
	"github.com/yangb8/aws-sdk-go/aws/signer/v4"
	"github.com/yangb8/aws-sdk-go/private/protocol/jsonrpc"
)

// AWS OpsWorks for Chef Automate is a service that runs and manages configuration
// management servers.
//
// Glossary of terms
//
//    * Server: A configuration management server that can be highly-available.
//    The configuration manager runs on your instances by using various AWS
//    services, such as Amazon Elastic Compute Cloud (EC2), and potentially
//    Amazon Relational Database Service (RDS). A server is a generic abstraction
//    over the configuration manager that you want to use, much like Amazon
//    RDS. In AWS OpsWorks for Chef Automate, you do not start or stop servers.
//    After you create servers, they continue to run until they are deleted.
//
//    * Engine: The specific configuration manager that you want to use (such
//    as Chef) is the engine.
//
//    * Backup: This is an application-level backup of the data that the configuration
//    manager stores. A backup creates a .tar.gz file that is stored in an Amazon
//    Simple Storage Service (S3) bucket in your account. AWS OpsWorks for Chef
//    Automate creates the S3 bucket when you launch the first instance. A backup
//    maintains a snapshot of all of a server's important attributes at the
//    time of the backup.
//
//    * Events: Events are always related to a server. Events are written during
//    server creation, when health checks run, when backups are created, etc.
//    When you delete a server, the server's events are also deleted.
//
//    * AccountAttributes: Every account has attributes that are assigned in
//    the AWS OpsWorks for Chef Automate database. These attributes store information
//    about configuration limits (servers, backups, etc.) and your customer
//    account.
//
// Endpoints
//
// AWS OpsWorks for Chef Automate supports the following endpoints, all HTTPS.
// You must connect to one of the following endpoints. Chef servers can only
// be accessed or managed within the endpoint in which they are created.
//
//    * opsworks-cm.us-east-1.amazonaws.com
//
//    * opsworks-cm.us-west-2.amazonaws.com
//
//    * opsworks-cm.eu-west-1.amazonaws.com
//
// Throttling limits
//
// All API operations allow for five requests per second with a burst of 10
// requests per second.
// The service client's operations are safe to be used concurrently.
// It is not safe to mutate any of the client's properties though.
// Please also see https://docs.aws.amazon.com/goto/WebAPI/opsworkscm-2016-11-01
type OpsWorksCM struct {
	*client.Client
}

// Used for custom client initialization logic
var initClient func(*client.Client)

// Used for custom request initialization logic
var initRequest func(*request.Request)

// Service information constants
const (
	ServiceName = "opsworks-cm" // Service endpoint prefix API calls made to.
	EndpointsID = ServiceName   // Service ID for Regions and Endpoints metadata.
)

// New creates a new instance of the OpsWorksCM client with a session.
// If additional configuration is needed for the client instance use the optional
// aws.Config parameter to add your extra config.
//
// Example:
//     // Create a OpsWorksCM client from just a session.
//     svc := opsworkscm.New(mySession)
//
//     // Create a OpsWorksCM client with additional configuration
//     svc := opsworkscm.New(mySession, aws.NewConfig().WithRegion("us-west-2"))
func New(p client.ConfigProvider, cfgs ...*aws.Config) *OpsWorksCM {
	c := p.ClientConfig(EndpointsID, cfgs...)
	return newClient(*c.Config, c.Handlers, c.Endpoint, c.SigningRegion, c.SigningName)
}

// newClient creates, initializes and returns a new service client instance.
func newClient(cfg aws.Config, handlers request.Handlers, endpoint, signingRegion, signingName string) *OpsWorksCM {
	if len(signingName) == 0 {
		signingName = "opsworks-cm"
	}
	svc := &OpsWorksCM{
		Client: client.New(
			cfg,
			metadata.ClientInfo{
				ServiceName:   ServiceName,
				SigningName:   signingName,
				SigningRegion: signingRegion,
				Endpoint:      endpoint,
				APIVersion:    "2016-11-01",
				JSONVersion:   "1.1",
				TargetPrefix:  "OpsWorksCM_V2016_11_01",
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

// newRequest creates a new request for a OpsWorksCM operation and runs any
// custom request initialization.
func (c *OpsWorksCM) newRequest(op *request.Operation, params, data interface{}) *request.Request {
	req := c.NewRequest(op, params, data)

	// Run custom request initialization if present
	if initRequest != nil {
		initRequest(req)
	}

	return req
}
