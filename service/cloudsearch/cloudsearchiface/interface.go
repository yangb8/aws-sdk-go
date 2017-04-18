// Code generated by private/model/cli/gen-api/main.go. DO NOT EDIT.

// Package cloudsearchiface provides an interface to enable mocking the Amazon CloudSearch service client
// for testing your code.
//
// It is important to note that this interface will have breaking changes
// when the service model is updated and adds new API operations, paginators,
// and waiters.
package cloudsearchiface

import (
	"github.com/yangb8/aws-sdk-go/aws"
	"github.com/yangb8/aws-sdk-go/aws/request"
	"github.com/yangb8/aws-sdk-go/service/cloudsearch"
)

// CloudSearchAPI provides an interface to enable mocking the
// cloudsearch.CloudSearch service client's API operation,
// paginators, and waiters. This make unit testing your code that calls out
// to the SDK's service client's calls easier.
//
// The best way to use this interface is so the SDK's service client's calls
// can be stubbed out for unit testing your code with the SDK without needing
// to inject custom request handlers into the the SDK's request pipeline.
//
//    // myFunc uses an SDK service client to make a request to
//    // Amazon CloudSearch.
//    func myFunc(svc cloudsearchiface.CloudSearchAPI) bool {
//        // Make svc.BuildSuggesters request
//    }
//
//    func main() {
//        sess := session.New()
//        svc := cloudsearch.New(sess)
//
//        myFunc(svc)
//    }
//
// In your _test.go file:
//
//    // Define a mock struct to be used in your unit tests of myFunc.
//    type mockCloudSearchClient struct {
//        cloudsearchiface.CloudSearchAPI
//    }
//    func (m *mockCloudSearchClient) BuildSuggesters(input *cloudsearch.BuildSuggestersInput) (*cloudsearch.BuildSuggestersOutput, error) {
//        // mock response/functionality
//    }
//
//    func TestMyFunc(t *testing.T) {
//        // Setup Test
//        mockSvc := &mockCloudSearchClient{}
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
type CloudSearchAPI interface {
	BuildSuggesters(*cloudsearch.BuildSuggestersInput) (*cloudsearch.BuildSuggestersOutput, error)
	BuildSuggestersWithContext(aws.Context, *cloudsearch.BuildSuggestersInput, ...request.Option) (*cloudsearch.BuildSuggestersOutput, error)
	BuildSuggestersRequest(*cloudsearch.BuildSuggestersInput) (*request.Request, *cloudsearch.BuildSuggestersOutput)

	CreateDomain(*cloudsearch.CreateDomainInput) (*cloudsearch.CreateDomainOutput, error)
	CreateDomainWithContext(aws.Context, *cloudsearch.CreateDomainInput, ...request.Option) (*cloudsearch.CreateDomainOutput, error)
	CreateDomainRequest(*cloudsearch.CreateDomainInput) (*request.Request, *cloudsearch.CreateDomainOutput)

	DefineAnalysisScheme(*cloudsearch.DefineAnalysisSchemeInput) (*cloudsearch.DefineAnalysisSchemeOutput, error)
	DefineAnalysisSchemeWithContext(aws.Context, *cloudsearch.DefineAnalysisSchemeInput, ...request.Option) (*cloudsearch.DefineAnalysisSchemeOutput, error)
	DefineAnalysisSchemeRequest(*cloudsearch.DefineAnalysisSchemeInput) (*request.Request, *cloudsearch.DefineAnalysisSchemeOutput)

	DefineExpression(*cloudsearch.DefineExpressionInput) (*cloudsearch.DefineExpressionOutput, error)
	DefineExpressionWithContext(aws.Context, *cloudsearch.DefineExpressionInput, ...request.Option) (*cloudsearch.DefineExpressionOutput, error)
	DefineExpressionRequest(*cloudsearch.DefineExpressionInput) (*request.Request, *cloudsearch.DefineExpressionOutput)

	DefineIndexField(*cloudsearch.DefineIndexFieldInput) (*cloudsearch.DefineIndexFieldOutput, error)
	DefineIndexFieldWithContext(aws.Context, *cloudsearch.DefineIndexFieldInput, ...request.Option) (*cloudsearch.DefineIndexFieldOutput, error)
	DefineIndexFieldRequest(*cloudsearch.DefineIndexFieldInput) (*request.Request, *cloudsearch.DefineIndexFieldOutput)

	DefineSuggester(*cloudsearch.DefineSuggesterInput) (*cloudsearch.DefineSuggesterOutput, error)
	DefineSuggesterWithContext(aws.Context, *cloudsearch.DefineSuggesterInput, ...request.Option) (*cloudsearch.DefineSuggesterOutput, error)
	DefineSuggesterRequest(*cloudsearch.DefineSuggesterInput) (*request.Request, *cloudsearch.DefineSuggesterOutput)

	DeleteAnalysisScheme(*cloudsearch.DeleteAnalysisSchemeInput) (*cloudsearch.DeleteAnalysisSchemeOutput, error)
	DeleteAnalysisSchemeWithContext(aws.Context, *cloudsearch.DeleteAnalysisSchemeInput, ...request.Option) (*cloudsearch.DeleteAnalysisSchemeOutput, error)
	DeleteAnalysisSchemeRequest(*cloudsearch.DeleteAnalysisSchemeInput) (*request.Request, *cloudsearch.DeleteAnalysisSchemeOutput)

	DeleteDomain(*cloudsearch.DeleteDomainInput) (*cloudsearch.DeleteDomainOutput, error)
	DeleteDomainWithContext(aws.Context, *cloudsearch.DeleteDomainInput, ...request.Option) (*cloudsearch.DeleteDomainOutput, error)
	DeleteDomainRequest(*cloudsearch.DeleteDomainInput) (*request.Request, *cloudsearch.DeleteDomainOutput)

	DeleteExpression(*cloudsearch.DeleteExpressionInput) (*cloudsearch.DeleteExpressionOutput, error)
	DeleteExpressionWithContext(aws.Context, *cloudsearch.DeleteExpressionInput, ...request.Option) (*cloudsearch.DeleteExpressionOutput, error)
	DeleteExpressionRequest(*cloudsearch.DeleteExpressionInput) (*request.Request, *cloudsearch.DeleteExpressionOutput)

	DeleteIndexField(*cloudsearch.DeleteIndexFieldInput) (*cloudsearch.DeleteIndexFieldOutput, error)
	DeleteIndexFieldWithContext(aws.Context, *cloudsearch.DeleteIndexFieldInput, ...request.Option) (*cloudsearch.DeleteIndexFieldOutput, error)
	DeleteIndexFieldRequest(*cloudsearch.DeleteIndexFieldInput) (*request.Request, *cloudsearch.DeleteIndexFieldOutput)

	DeleteSuggester(*cloudsearch.DeleteSuggesterInput) (*cloudsearch.DeleteSuggesterOutput, error)
	DeleteSuggesterWithContext(aws.Context, *cloudsearch.DeleteSuggesterInput, ...request.Option) (*cloudsearch.DeleteSuggesterOutput, error)
	DeleteSuggesterRequest(*cloudsearch.DeleteSuggesterInput) (*request.Request, *cloudsearch.DeleteSuggesterOutput)

	DescribeAnalysisSchemes(*cloudsearch.DescribeAnalysisSchemesInput) (*cloudsearch.DescribeAnalysisSchemesOutput, error)
	DescribeAnalysisSchemesWithContext(aws.Context, *cloudsearch.DescribeAnalysisSchemesInput, ...request.Option) (*cloudsearch.DescribeAnalysisSchemesOutput, error)
	DescribeAnalysisSchemesRequest(*cloudsearch.DescribeAnalysisSchemesInput) (*request.Request, *cloudsearch.DescribeAnalysisSchemesOutput)

	DescribeAvailabilityOptions(*cloudsearch.DescribeAvailabilityOptionsInput) (*cloudsearch.DescribeAvailabilityOptionsOutput, error)
	DescribeAvailabilityOptionsWithContext(aws.Context, *cloudsearch.DescribeAvailabilityOptionsInput, ...request.Option) (*cloudsearch.DescribeAvailabilityOptionsOutput, error)
	DescribeAvailabilityOptionsRequest(*cloudsearch.DescribeAvailabilityOptionsInput) (*request.Request, *cloudsearch.DescribeAvailabilityOptionsOutput)

	DescribeDomains(*cloudsearch.DescribeDomainsInput) (*cloudsearch.DescribeDomainsOutput, error)
	DescribeDomainsWithContext(aws.Context, *cloudsearch.DescribeDomainsInput, ...request.Option) (*cloudsearch.DescribeDomainsOutput, error)
	DescribeDomainsRequest(*cloudsearch.DescribeDomainsInput) (*request.Request, *cloudsearch.DescribeDomainsOutput)

	DescribeExpressions(*cloudsearch.DescribeExpressionsInput) (*cloudsearch.DescribeExpressionsOutput, error)
	DescribeExpressionsWithContext(aws.Context, *cloudsearch.DescribeExpressionsInput, ...request.Option) (*cloudsearch.DescribeExpressionsOutput, error)
	DescribeExpressionsRequest(*cloudsearch.DescribeExpressionsInput) (*request.Request, *cloudsearch.DescribeExpressionsOutput)

	DescribeIndexFields(*cloudsearch.DescribeIndexFieldsInput) (*cloudsearch.DescribeIndexFieldsOutput, error)
	DescribeIndexFieldsWithContext(aws.Context, *cloudsearch.DescribeIndexFieldsInput, ...request.Option) (*cloudsearch.DescribeIndexFieldsOutput, error)
	DescribeIndexFieldsRequest(*cloudsearch.DescribeIndexFieldsInput) (*request.Request, *cloudsearch.DescribeIndexFieldsOutput)

	DescribeScalingParameters(*cloudsearch.DescribeScalingParametersInput) (*cloudsearch.DescribeScalingParametersOutput, error)
	DescribeScalingParametersWithContext(aws.Context, *cloudsearch.DescribeScalingParametersInput, ...request.Option) (*cloudsearch.DescribeScalingParametersOutput, error)
	DescribeScalingParametersRequest(*cloudsearch.DescribeScalingParametersInput) (*request.Request, *cloudsearch.DescribeScalingParametersOutput)

	DescribeServiceAccessPolicies(*cloudsearch.DescribeServiceAccessPoliciesInput) (*cloudsearch.DescribeServiceAccessPoliciesOutput, error)
	DescribeServiceAccessPoliciesWithContext(aws.Context, *cloudsearch.DescribeServiceAccessPoliciesInput, ...request.Option) (*cloudsearch.DescribeServiceAccessPoliciesOutput, error)
	DescribeServiceAccessPoliciesRequest(*cloudsearch.DescribeServiceAccessPoliciesInput) (*request.Request, *cloudsearch.DescribeServiceAccessPoliciesOutput)

	DescribeSuggesters(*cloudsearch.DescribeSuggestersInput) (*cloudsearch.DescribeSuggestersOutput, error)
	DescribeSuggestersWithContext(aws.Context, *cloudsearch.DescribeSuggestersInput, ...request.Option) (*cloudsearch.DescribeSuggestersOutput, error)
	DescribeSuggestersRequest(*cloudsearch.DescribeSuggestersInput) (*request.Request, *cloudsearch.DescribeSuggestersOutput)

	IndexDocuments(*cloudsearch.IndexDocumentsInput) (*cloudsearch.IndexDocumentsOutput, error)
	IndexDocumentsWithContext(aws.Context, *cloudsearch.IndexDocumentsInput, ...request.Option) (*cloudsearch.IndexDocumentsOutput, error)
	IndexDocumentsRequest(*cloudsearch.IndexDocumentsInput) (*request.Request, *cloudsearch.IndexDocumentsOutput)

	ListDomainNames(*cloudsearch.ListDomainNamesInput) (*cloudsearch.ListDomainNamesOutput, error)
	ListDomainNamesWithContext(aws.Context, *cloudsearch.ListDomainNamesInput, ...request.Option) (*cloudsearch.ListDomainNamesOutput, error)
	ListDomainNamesRequest(*cloudsearch.ListDomainNamesInput) (*request.Request, *cloudsearch.ListDomainNamesOutput)

	UpdateAvailabilityOptions(*cloudsearch.UpdateAvailabilityOptionsInput) (*cloudsearch.UpdateAvailabilityOptionsOutput, error)
	UpdateAvailabilityOptionsWithContext(aws.Context, *cloudsearch.UpdateAvailabilityOptionsInput, ...request.Option) (*cloudsearch.UpdateAvailabilityOptionsOutput, error)
	UpdateAvailabilityOptionsRequest(*cloudsearch.UpdateAvailabilityOptionsInput) (*request.Request, *cloudsearch.UpdateAvailabilityOptionsOutput)

	UpdateScalingParameters(*cloudsearch.UpdateScalingParametersInput) (*cloudsearch.UpdateScalingParametersOutput, error)
	UpdateScalingParametersWithContext(aws.Context, *cloudsearch.UpdateScalingParametersInput, ...request.Option) (*cloudsearch.UpdateScalingParametersOutput, error)
	UpdateScalingParametersRequest(*cloudsearch.UpdateScalingParametersInput) (*request.Request, *cloudsearch.UpdateScalingParametersOutput)

	UpdateServiceAccessPolicies(*cloudsearch.UpdateServiceAccessPoliciesInput) (*cloudsearch.UpdateServiceAccessPoliciesOutput, error)
	UpdateServiceAccessPoliciesWithContext(aws.Context, *cloudsearch.UpdateServiceAccessPoliciesInput, ...request.Option) (*cloudsearch.UpdateServiceAccessPoliciesOutput, error)
	UpdateServiceAccessPoliciesRequest(*cloudsearch.UpdateServiceAccessPoliciesInput) (*request.Request, *cloudsearch.UpdateServiceAccessPoliciesOutput)
}

var _ CloudSearchAPI = (*cloudsearch.CloudSearch)(nil)
