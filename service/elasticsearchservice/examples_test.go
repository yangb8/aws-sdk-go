// Code generated by private/model/cli/gen-api/main.go. DO NOT EDIT.

package elasticsearchservice_test

import (
	"bytes"
	"fmt"
	"time"

	"github.com/yangb8/aws-sdk-go/aws"
	"github.com/yangb8/aws-sdk-go/aws/session"
	"github.com/yangb8/aws-sdk-go/service/elasticsearchservice"
)

var _ time.Duration
var _ bytes.Buffer

func ExampleElasticsearchService_AddTags() {
	sess := session.Must(session.NewSession())

	svc := elasticsearchservice.New(sess)

	params := &elasticsearchservice.AddTagsInput{
		ARN: aws.String("ARN"), // Required
		TagList: []*elasticsearchservice.Tag{ // Required
			{ // Required
				Key:   aws.String("TagKey"),   // Required
				Value: aws.String("TagValue"), // Required
			},
			// More values...
		},
	}
	resp, err := svc.AddTags(params)

	if err != nil {
		// Print the error, cast err to awserr.Error to get the Code and
		// Message from an error.
		fmt.Println(err.Error())
		return
	}

	// Pretty-print the response data.
	fmt.Println(resp)
}

func ExampleElasticsearchService_CreateElasticsearchDomain() {
	sess := session.Must(session.NewSession())

	svc := elasticsearchservice.New(sess)

	params := &elasticsearchservice.CreateElasticsearchDomainInput{
		DomainName:     aws.String("DomainName"), // Required
		AccessPolicies: aws.String("PolicyDocument"),
		AdvancedOptions: map[string]*string{
			"Key": aws.String("String"), // Required
			// More values...
		},
		EBSOptions: &elasticsearchservice.EBSOptions{
			EBSEnabled: aws.Bool(true),
			Iops:       aws.Int64(1),
			VolumeSize: aws.Int64(1),
			VolumeType: aws.String("VolumeType"),
		},
		ElasticsearchClusterConfig: &elasticsearchservice.ElasticsearchClusterConfig{
			DedicatedMasterCount:   aws.Int64(1),
			DedicatedMasterEnabled: aws.Bool(true),
			DedicatedMasterType:    aws.String("ESPartitionInstanceType"),
			InstanceCount:          aws.Int64(1),
			InstanceType:           aws.String("ESPartitionInstanceType"),
			ZoneAwarenessEnabled:   aws.Bool(true),
		},
		ElasticsearchVersion: aws.String("ElasticsearchVersionString"),
		SnapshotOptions: &elasticsearchservice.SnapshotOptions{
			AutomatedSnapshotStartHour: aws.Int64(1),
		},
	}
	resp, err := svc.CreateElasticsearchDomain(params)

	if err != nil {
		// Print the error, cast err to awserr.Error to get the Code and
		// Message from an error.
		fmt.Println(err.Error())
		return
	}

	// Pretty-print the response data.
	fmt.Println(resp)
}

func ExampleElasticsearchService_DeleteElasticsearchDomain() {
	sess := session.Must(session.NewSession())

	svc := elasticsearchservice.New(sess)

	params := &elasticsearchservice.DeleteElasticsearchDomainInput{
		DomainName: aws.String("DomainName"), // Required
	}
	resp, err := svc.DeleteElasticsearchDomain(params)

	if err != nil {
		// Print the error, cast err to awserr.Error to get the Code and
		// Message from an error.
		fmt.Println(err.Error())
		return
	}

	// Pretty-print the response data.
	fmt.Println(resp)
}

func ExampleElasticsearchService_DescribeElasticsearchDomain() {
	sess := session.Must(session.NewSession())

	svc := elasticsearchservice.New(sess)

	params := &elasticsearchservice.DescribeElasticsearchDomainInput{
		DomainName: aws.String("DomainName"), // Required
	}
	resp, err := svc.DescribeElasticsearchDomain(params)

	if err != nil {
		// Print the error, cast err to awserr.Error to get the Code and
		// Message from an error.
		fmt.Println(err.Error())
		return
	}

	// Pretty-print the response data.
	fmt.Println(resp)
}

func ExampleElasticsearchService_DescribeElasticsearchDomainConfig() {
	sess := session.Must(session.NewSession())

	svc := elasticsearchservice.New(sess)

	params := &elasticsearchservice.DescribeElasticsearchDomainConfigInput{
		DomainName: aws.String("DomainName"), // Required
	}
	resp, err := svc.DescribeElasticsearchDomainConfig(params)

	if err != nil {
		// Print the error, cast err to awserr.Error to get the Code and
		// Message from an error.
		fmt.Println(err.Error())
		return
	}

	// Pretty-print the response data.
	fmt.Println(resp)
}

func ExampleElasticsearchService_DescribeElasticsearchDomains() {
	sess := session.Must(session.NewSession())

	svc := elasticsearchservice.New(sess)

	params := &elasticsearchservice.DescribeElasticsearchDomainsInput{
		DomainNames: []*string{ // Required
			aws.String("DomainName"), // Required
			// More values...
		},
	}
	resp, err := svc.DescribeElasticsearchDomains(params)

	if err != nil {
		// Print the error, cast err to awserr.Error to get the Code and
		// Message from an error.
		fmt.Println(err.Error())
		return
	}

	// Pretty-print the response data.
	fmt.Println(resp)
}

func ExampleElasticsearchService_DescribeElasticsearchInstanceTypeLimits() {
	sess := session.Must(session.NewSession())

	svc := elasticsearchservice.New(sess)

	params := &elasticsearchservice.DescribeElasticsearchInstanceTypeLimitsInput{
		ElasticsearchVersion: aws.String("ElasticsearchVersionString"), // Required
		InstanceType:         aws.String("ESPartitionInstanceType"),    // Required
		DomainName:           aws.String("DomainName"),
	}
	resp, err := svc.DescribeElasticsearchInstanceTypeLimits(params)

	if err != nil {
		// Print the error, cast err to awserr.Error to get the Code and
		// Message from an error.
		fmt.Println(err.Error())
		return
	}

	// Pretty-print the response data.
	fmt.Println(resp)
}

func ExampleElasticsearchService_ListDomainNames() {
	sess := session.Must(session.NewSession())

	svc := elasticsearchservice.New(sess)

	var params *elasticsearchservice.ListDomainNamesInput
	resp, err := svc.ListDomainNames(params)

	if err != nil {
		// Print the error, cast err to awserr.Error to get the Code and
		// Message from an error.
		fmt.Println(err.Error())
		return
	}

	// Pretty-print the response data.
	fmt.Println(resp)
}

func ExampleElasticsearchService_ListElasticsearchInstanceTypes() {
	sess := session.Must(session.NewSession())

	svc := elasticsearchservice.New(sess)

	params := &elasticsearchservice.ListElasticsearchInstanceTypesInput{
		ElasticsearchVersion: aws.String("ElasticsearchVersionString"), // Required
		DomainName:           aws.String("DomainName"),
		MaxResults:           aws.Int64(1),
		NextToken:            aws.String("NextToken"),
	}
	resp, err := svc.ListElasticsearchInstanceTypes(params)

	if err != nil {
		// Print the error, cast err to awserr.Error to get the Code and
		// Message from an error.
		fmt.Println(err.Error())
		return
	}

	// Pretty-print the response data.
	fmt.Println(resp)
}

func ExampleElasticsearchService_ListElasticsearchVersions() {
	sess := session.Must(session.NewSession())

	svc := elasticsearchservice.New(sess)

	params := &elasticsearchservice.ListElasticsearchVersionsInput{
		MaxResults: aws.Int64(1),
		NextToken:  aws.String("NextToken"),
	}
	resp, err := svc.ListElasticsearchVersions(params)

	if err != nil {
		// Print the error, cast err to awserr.Error to get the Code and
		// Message from an error.
		fmt.Println(err.Error())
		return
	}

	// Pretty-print the response data.
	fmt.Println(resp)
}

func ExampleElasticsearchService_ListTags() {
	sess := session.Must(session.NewSession())

	svc := elasticsearchservice.New(sess)

	params := &elasticsearchservice.ListTagsInput{
		ARN: aws.String("ARN"), // Required
	}
	resp, err := svc.ListTags(params)

	if err != nil {
		// Print the error, cast err to awserr.Error to get the Code and
		// Message from an error.
		fmt.Println(err.Error())
		return
	}

	// Pretty-print the response data.
	fmt.Println(resp)
}

func ExampleElasticsearchService_RemoveTags() {
	sess := session.Must(session.NewSession())

	svc := elasticsearchservice.New(sess)

	params := &elasticsearchservice.RemoveTagsInput{
		ARN: aws.String("ARN"), // Required
		TagKeys: []*string{ // Required
			aws.String("String"), // Required
			// More values...
		},
	}
	resp, err := svc.RemoveTags(params)

	if err != nil {
		// Print the error, cast err to awserr.Error to get the Code and
		// Message from an error.
		fmt.Println(err.Error())
		return
	}

	// Pretty-print the response data.
	fmt.Println(resp)
}

func ExampleElasticsearchService_UpdateElasticsearchDomainConfig() {
	sess := session.Must(session.NewSession())

	svc := elasticsearchservice.New(sess)

	params := &elasticsearchservice.UpdateElasticsearchDomainConfigInput{
		DomainName:     aws.String("DomainName"), // Required
		AccessPolicies: aws.String("PolicyDocument"),
		AdvancedOptions: map[string]*string{
			"Key": aws.String("String"), // Required
			// More values...
		},
		EBSOptions: &elasticsearchservice.EBSOptions{
			EBSEnabled: aws.Bool(true),
			Iops:       aws.Int64(1),
			VolumeSize: aws.Int64(1),
			VolumeType: aws.String("VolumeType"),
		},
		ElasticsearchClusterConfig: &elasticsearchservice.ElasticsearchClusterConfig{
			DedicatedMasterCount:   aws.Int64(1),
			DedicatedMasterEnabled: aws.Bool(true),
			DedicatedMasterType:    aws.String("ESPartitionInstanceType"),
			InstanceCount:          aws.Int64(1),
			InstanceType:           aws.String("ESPartitionInstanceType"),
			ZoneAwarenessEnabled:   aws.Bool(true),
		},
		SnapshotOptions: &elasticsearchservice.SnapshotOptions{
			AutomatedSnapshotStartHour: aws.Int64(1),
		},
	}
	resp, err := svc.UpdateElasticsearchDomainConfig(params)

	if err != nil {
		// Print the error, cast err to awserr.Error to get the Code and
		// Message from an error.
		fmt.Println(err.Error())
		return
	}

	// Pretty-print the response data.
	fmt.Println(resp)
}
