// Code generated by private/model/cli/gen-api/main.go. DO NOT EDIT.

// Package pinpointiface provides an interface to enable mocking the Amazon Pinpoint service client
// for testing your code.
//
// It is important to note that this interface will have breaking changes
// when the service model is updated and adds new API operations, paginators,
// and waiters.
package pinpointiface

import (
	"github.com/yangb8/aws-sdk-go/aws"
	"github.com/yangb8/aws-sdk-go/aws/request"
	"github.com/yangb8/aws-sdk-go/service/pinpoint"
)

// PinpointAPI provides an interface to enable mocking the
// pinpoint.Pinpoint service client's API operation,
// paginators, and waiters. This make unit testing your code that calls out
// to the SDK's service client's calls easier.
//
// The best way to use this interface is so the SDK's service client's calls
// can be stubbed out for unit testing your code with the SDK without needing
// to inject custom request handlers into the the SDK's request pipeline.
//
//    // myFunc uses an SDK service client to make a request to
//    // Amazon Pinpoint.
//    func myFunc(svc pinpointiface.PinpointAPI) bool {
//        // Make svc.CreateCampaign request
//    }
//
//    func main() {
//        sess := session.New()
//        svc := pinpoint.New(sess)
//
//        myFunc(svc)
//    }
//
// In your _test.go file:
//
//    // Define a mock struct to be used in your unit tests of myFunc.
//    type mockPinpointClient struct {
//        pinpointiface.PinpointAPI
//    }
//    func (m *mockPinpointClient) CreateCampaign(input *pinpoint.CreateCampaignInput) (*pinpoint.CreateCampaignOutput, error) {
//        // mock response/functionality
//    }
//
//    func TestMyFunc(t *testing.T) {
//        // Setup Test
//        mockSvc := &mockPinpointClient{}
//
//        myfunc(mockSvc)
//
//        // Verify myFunc's functionality
//    }
//
// It is important to note that this interface will have breaking changes
// when the service model is updated and adds new API operations, paginators,
// and waiters. Its suggested to use the pattern above for testing, or using
// tooling to generate mocks to satisfy the interfaces.
type PinpointAPI interface {
	CreateCampaign(*pinpoint.CreateCampaignInput) (*pinpoint.CreateCampaignOutput, error)
	CreateCampaignWithContext(aws.Context, *pinpoint.CreateCampaignInput, ...request.Option) (*pinpoint.CreateCampaignOutput, error)
	CreateCampaignRequest(*pinpoint.CreateCampaignInput) (*request.Request, *pinpoint.CreateCampaignOutput)

	CreateImportJob(*pinpoint.CreateImportJobInput) (*pinpoint.CreateImportJobOutput, error)
	CreateImportJobWithContext(aws.Context, *pinpoint.CreateImportJobInput, ...request.Option) (*pinpoint.CreateImportJobOutput, error)
	CreateImportJobRequest(*pinpoint.CreateImportJobInput) (*request.Request, *pinpoint.CreateImportJobOutput)

	CreateSegment(*pinpoint.CreateSegmentInput) (*pinpoint.CreateSegmentOutput, error)
	CreateSegmentWithContext(aws.Context, *pinpoint.CreateSegmentInput, ...request.Option) (*pinpoint.CreateSegmentOutput, error)
	CreateSegmentRequest(*pinpoint.CreateSegmentInput) (*request.Request, *pinpoint.CreateSegmentOutput)

	DeleteApnsChannel(*pinpoint.DeleteApnsChannelInput) (*pinpoint.DeleteApnsChannelOutput, error)
	DeleteApnsChannelWithContext(aws.Context, *pinpoint.DeleteApnsChannelInput, ...request.Option) (*pinpoint.DeleteApnsChannelOutput, error)
	DeleteApnsChannelRequest(*pinpoint.DeleteApnsChannelInput) (*request.Request, *pinpoint.DeleteApnsChannelOutput)

	DeleteCampaign(*pinpoint.DeleteCampaignInput) (*pinpoint.DeleteCampaignOutput, error)
	DeleteCampaignWithContext(aws.Context, *pinpoint.DeleteCampaignInput, ...request.Option) (*pinpoint.DeleteCampaignOutput, error)
	DeleteCampaignRequest(*pinpoint.DeleteCampaignInput) (*request.Request, *pinpoint.DeleteCampaignOutput)

	DeleteEventStream(*pinpoint.DeleteEventStreamInput) (*pinpoint.DeleteEventStreamOutput, error)
	DeleteEventStreamWithContext(aws.Context, *pinpoint.DeleteEventStreamInput, ...request.Option) (*pinpoint.DeleteEventStreamOutput, error)
	DeleteEventStreamRequest(*pinpoint.DeleteEventStreamInput) (*request.Request, *pinpoint.DeleteEventStreamOutput)

	DeleteGcmChannel(*pinpoint.DeleteGcmChannelInput) (*pinpoint.DeleteGcmChannelOutput, error)
	DeleteGcmChannelWithContext(aws.Context, *pinpoint.DeleteGcmChannelInput, ...request.Option) (*pinpoint.DeleteGcmChannelOutput, error)
	DeleteGcmChannelRequest(*pinpoint.DeleteGcmChannelInput) (*request.Request, *pinpoint.DeleteGcmChannelOutput)

	DeleteSegment(*pinpoint.DeleteSegmentInput) (*pinpoint.DeleteSegmentOutput, error)
	DeleteSegmentWithContext(aws.Context, *pinpoint.DeleteSegmentInput, ...request.Option) (*pinpoint.DeleteSegmentOutput, error)
	DeleteSegmentRequest(*pinpoint.DeleteSegmentInput) (*request.Request, *pinpoint.DeleteSegmentOutput)

	GetApnsChannel(*pinpoint.GetApnsChannelInput) (*pinpoint.GetApnsChannelOutput, error)
	GetApnsChannelWithContext(aws.Context, *pinpoint.GetApnsChannelInput, ...request.Option) (*pinpoint.GetApnsChannelOutput, error)
	GetApnsChannelRequest(*pinpoint.GetApnsChannelInput) (*request.Request, *pinpoint.GetApnsChannelOutput)

	GetApplicationSettings(*pinpoint.GetApplicationSettingsInput) (*pinpoint.GetApplicationSettingsOutput, error)
	GetApplicationSettingsWithContext(aws.Context, *pinpoint.GetApplicationSettingsInput, ...request.Option) (*pinpoint.GetApplicationSettingsOutput, error)
	GetApplicationSettingsRequest(*pinpoint.GetApplicationSettingsInput) (*request.Request, *pinpoint.GetApplicationSettingsOutput)

	GetCampaign(*pinpoint.GetCampaignInput) (*pinpoint.GetCampaignOutput, error)
	GetCampaignWithContext(aws.Context, *pinpoint.GetCampaignInput, ...request.Option) (*pinpoint.GetCampaignOutput, error)
	GetCampaignRequest(*pinpoint.GetCampaignInput) (*request.Request, *pinpoint.GetCampaignOutput)

	GetCampaignActivities(*pinpoint.GetCampaignActivitiesInput) (*pinpoint.GetCampaignActivitiesOutput, error)
	GetCampaignActivitiesWithContext(aws.Context, *pinpoint.GetCampaignActivitiesInput, ...request.Option) (*pinpoint.GetCampaignActivitiesOutput, error)
	GetCampaignActivitiesRequest(*pinpoint.GetCampaignActivitiesInput) (*request.Request, *pinpoint.GetCampaignActivitiesOutput)

	GetCampaignVersion(*pinpoint.GetCampaignVersionInput) (*pinpoint.GetCampaignVersionOutput, error)
	GetCampaignVersionWithContext(aws.Context, *pinpoint.GetCampaignVersionInput, ...request.Option) (*pinpoint.GetCampaignVersionOutput, error)
	GetCampaignVersionRequest(*pinpoint.GetCampaignVersionInput) (*request.Request, *pinpoint.GetCampaignVersionOutput)

	GetCampaignVersions(*pinpoint.GetCampaignVersionsInput) (*pinpoint.GetCampaignVersionsOutput, error)
	GetCampaignVersionsWithContext(aws.Context, *pinpoint.GetCampaignVersionsInput, ...request.Option) (*pinpoint.GetCampaignVersionsOutput, error)
	GetCampaignVersionsRequest(*pinpoint.GetCampaignVersionsInput) (*request.Request, *pinpoint.GetCampaignVersionsOutput)

	GetCampaigns(*pinpoint.GetCampaignsInput) (*pinpoint.GetCampaignsOutput, error)
	GetCampaignsWithContext(aws.Context, *pinpoint.GetCampaignsInput, ...request.Option) (*pinpoint.GetCampaignsOutput, error)
	GetCampaignsRequest(*pinpoint.GetCampaignsInput) (*request.Request, *pinpoint.GetCampaignsOutput)

	GetEndpoint(*pinpoint.GetEndpointInput) (*pinpoint.GetEndpointOutput, error)
	GetEndpointWithContext(aws.Context, *pinpoint.GetEndpointInput, ...request.Option) (*pinpoint.GetEndpointOutput, error)
	GetEndpointRequest(*pinpoint.GetEndpointInput) (*request.Request, *pinpoint.GetEndpointOutput)

	GetEventStream(*pinpoint.GetEventStreamInput) (*pinpoint.GetEventStreamOutput, error)
	GetEventStreamWithContext(aws.Context, *pinpoint.GetEventStreamInput, ...request.Option) (*pinpoint.GetEventStreamOutput, error)
	GetEventStreamRequest(*pinpoint.GetEventStreamInput) (*request.Request, *pinpoint.GetEventStreamOutput)

	GetGcmChannel(*pinpoint.GetGcmChannelInput) (*pinpoint.GetGcmChannelOutput, error)
	GetGcmChannelWithContext(aws.Context, *pinpoint.GetGcmChannelInput, ...request.Option) (*pinpoint.GetGcmChannelOutput, error)
	GetGcmChannelRequest(*pinpoint.GetGcmChannelInput) (*request.Request, *pinpoint.GetGcmChannelOutput)

	GetImportJob(*pinpoint.GetImportJobInput) (*pinpoint.GetImportJobOutput, error)
	GetImportJobWithContext(aws.Context, *pinpoint.GetImportJobInput, ...request.Option) (*pinpoint.GetImportJobOutput, error)
	GetImportJobRequest(*pinpoint.GetImportJobInput) (*request.Request, *pinpoint.GetImportJobOutput)

	GetImportJobs(*pinpoint.GetImportJobsInput) (*pinpoint.GetImportJobsOutput, error)
	GetImportJobsWithContext(aws.Context, *pinpoint.GetImportJobsInput, ...request.Option) (*pinpoint.GetImportJobsOutput, error)
	GetImportJobsRequest(*pinpoint.GetImportJobsInput) (*request.Request, *pinpoint.GetImportJobsOutput)

	GetSegment(*pinpoint.GetSegmentInput) (*pinpoint.GetSegmentOutput, error)
	GetSegmentWithContext(aws.Context, *pinpoint.GetSegmentInput, ...request.Option) (*pinpoint.GetSegmentOutput, error)
	GetSegmentRequest(*pinpoint.GetSegmentInput) (*request.Request, *pinpoint.GetSegmentOutput)

	GetSegmentImportJobs(*pinpoint.GetSegmentImportJobsInput) (*pinpoint.GetSegmentImportJobsOutput, error)
	GetSegmentImportJobsWithContext(aws.Context, *pinpoint.GetSegmentImportJobsInput, ...request.Option) (*pinpoint.GetSegmentImportJobsOutput, error)
	GetSegmentImportJobsRequest(*pinpoint.GetSegmentImportJobsInput) (*request.Request, *pinpoint.GetSegmentImportJobsOutput)

	GetSegmentVersion(*pinpoint.GetSegmentVersionInput) (*pinpoint.GetSegmentVersionOutput, error)
	GetSegmentVersionWithContext(aws.Context, *pinpoint.GetSegmentVersionInput, ...request.Option) (*pinpoint.GetSegmentVersionOutput, error)
	GetSegmentVersionRequest(*pinpoint.GetSegmentVersionInput) (*request.Request, *pinpoint.GetSegmentVersionOutput)

	GetSegmentVersions(*pinpoint.GetSegmentVersionsInput) (*pinpoint.GetSegmentVersionsOutput, error)
	GetSegmentVersionsWithContext(aws.Context, *pinpoint.GetSegmentVersionsInput, ...request.Option) (*pinpoint.GetSegmentVersionsOutput, error)
	GetSegmentVersionsRequest(*pinpoint.GetSegmentVersionsInput) (*request.Request, *pinpoint.GetSegmentVersionsOutput)

	GetSegments(*pinpoint.GetSegmentsInput) (*pinpoint.GetSegmentsOutput, error)
	GetSegmentsWithContext(aws.Context, *pinpoint.GetSegmentsInput, ...request.Option) (*pinpoint.GetSegmentsOutput, error)
	GetSegmentsRequest(*pinpoint.GetSegmentsInput) (*request.Request, *pinpoint.GetSegmentsOutput)

	PutEventStream(*pinpoint.PutEventStreamInput) (*pinpoint.PutEventStreamOutput, error)
	PutEventStreamWithContext(aws.Context, *pinpoint.PutEventStreamInput, ...request.Option) (*pinpoint.PutEventStreamOutput, error)
	PutEventStreamRequest(*pinpoint.PutEventStreamInput) (*request.Request, *pinpoint.PutEventStreamOutput)

	UpdateApnsChannel(*pinpoint.UpdateApnsChannelInput) (*pinpoint.UpdateApnsChannelOutput, error)
	UpdateApnsChannelWithContext(aws.Context, *pinpoint.UpdateApnsChannelInput, ...request.Option) (*pinpoint.UpdateApnsChannelOutput, error)
	UpdateApnsChannelRequest(*pinpoint.UpdateApnsChannelInput) (*request.Request, *pinpoint.UpdateApnsChannelOutput)

	UpdateApplicationSettings(*pinpoint.UpdateApplicationSettingsInput) (*pinpoint.UpdateApplicationSettingsOutput, error)
	UpdateApplicationSettingsWithContext(aws.Context, *pinpoint.UpdateApplicationSettingsInput, ...request.Option) (*pinpoint.UpdateApplicationSettingsOutput, error)
	UpdateApplicationSettingsRequest(*pinpoint.UpdateApplicationSettingsInput) (*request.Request, *pinpoint.UpdateApplicationSettingsOutput)

	UpdateCampaign(*pinpoint.UpdateCampaignInput) (*pinpoint.UpdateCampaignOutput, error)
	UpdateCampaignWithContext(aws.Context, *pinpoint.UpdateCampaignInput, ...request.Option) (*pinpoint.UpdateCampaignOutput, error)
	UpdateCampaignRequest(*pinpoint.UpdateCampaignInput) (*request.Request, *pinpoint.UpdateCampaignOutput)

	UpdateEndpoint(*pinpoint.UpdateEndpointInput) (*pinpoint.UpdateEndpointOutput, error)
	UpdateEndpointWithContext(aws.Context, *pinpoint.UpdateEndpointInput, ...request.Option) (*pinpoint.UpdateEndpointOutput, error)
	UpdateEndpointRequest(*pinpoint.UpdateEndpointInput) (*request.Request, *pinpoint.UpdateEndpointOutput)

	UpdateEndpointsBatch(*pinpoint.UpdateEndpointsBatchInput) (*pinpoint.UpdateEndpointsBatchOutput, error)
	UpdateEndpointsBatchWithContext(aws.Context, *pinpoint.UpdateEndpointsBatchInput, ...request.Option) (*pinpoint.UpdateEndpointsBatchOutput, error)
	UpdateEndpointsBatchRequest(*pinpoint.UpdateEndpointsBatchInput) (*request.Request, *pinpoint.UpdateEndpointsBatchOutput)

	UpdateGcmChannel(*pinpoint.UpdateGcmChannelInput) (*pinpoint.UpdateGcmChannelOutput, error)
	UpdateGcmChannelWithContext(aws.Context, *pinpoint.UpdateGcmChannelInput, ...request.Option) (*pinpoint.UpdateGcmChannelOutput, error)
	UpdateGcmChannelRequest(*pinpoint.UpdateGcmChannelInput) (*request.Request, *pinpoint.UpdateGcmChannelOutput)

	UpdateSegment(*pinpoint.UpdateSegmentInput) (*pinpoint.UpdateSegmentOutput, error)
	UpdateSegmentWithContext(aws.Context, *pinpoint.UpdateSegmentInput, ...request.Option) (*pinpoint.UpdateSegmentOutput, error)
	UpdateSegmentRequest(*pinpoint.UpdateSegmentInput) (*request.Request, *pinpoint.UpdateSegmentOutput)
}

var _ PinpointAPI = (*pinpoint.Pinpoint)(nil)
