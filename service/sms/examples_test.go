// Code generated by private/model/cli/gen-api/main.go. DO NOT EDIT.

package sms_test

import (
	"bytes"
	"fmt"
	"time"

	"github.com/yangb8/aws-sdk-go/aws"
	"github.com/yangb8/aws-sdk-go/aws/session"
	"github.com/yangb8/aws-sdk-go/service/sms"
)

var _ time.Duration
var _ bytes.Buffer

func ExampleSMS_CreateReplicationJob() {
	sess := session.Must(session.NewSession())

	svc := sms.New(sess)

	params := &sms.CreateReplicationJobInput{
		Frequency:           aws.Int64(1),           // Required
		SeedReplicationTime: aws.Time(time.Now()),   // Required
		ServerId:            aws.String("ServerId"), // Required
		Description:         aws.String("Description"),
		LicenseType:         aws.String("LicenseType"),
		RoleName:            aws.String("RoleName"),
	}
	resp, err := svc.CreateReplicationJob(params)

	if err != nil {
		// Print the error, cast err to awserr.Error to get the Code and
		// Message from an error.
		fmt.Println(err.Error())
		return
	}

	// Pretty-print the response data.
	fmt.Println(resp)
}

func ExampleSMS_DeleteReplicationJob() {
	sess := session.Must(session.NewSession())

	svc := sms.New(sess)

	params := &sms.DeleteReplicationJobInput{
		ReplicationJobId: aws.String("ReplicationJobId"), // Required
	}
	resp, err := svc.DeleteReplicationJob(params)

	if err != nil {
		// Print the error, cast err to awserr.Error to get the Code and
		// Message from an error.
		fmt.Println(err.Error())
		return
	}

	// Pretty-print the response data.
	fmt.Println(resp)
}

func ExampleSMS_DeleteServerCatalog() {
	sess := session.Must(session.NewSession())

	svc := sms.New(sess)

	var params *sms.DeleteServerCatalogInput
	resp, err := svc.DeleteServerCatalog(params)

	if err != nil {
		// Print the error, cast err to awserr.Error to get the Code and
		// Message from an error.
		fmt.Println(err.Error())
		return
	}

	// Pretty-print the response data.
	fmt.Println(resp)
}

func ExampleSMS_DisassociateConnector() {
	sess := session.Must(session.NewSession())

	svc := sms.New(sess)

	params := &sms.DisassociateConnectorInput{
		ConnectorId: aws.String("ConnectorId"), // Required
	}
	resp, err := svc.DisassociateConnector(params)

	if err != nil {
		// Print the error, cast err to awserr.Error to get the Code and
		// Message from an error.
		fmt.Println(err.Error())
		return
	}

	// Pretty-print the response data.
	fmt.Println(resp)
}

func ExampleSMS_GetConnectors() {
	sess := session.Must(session.NewSession())

	svc := sms.New(sess)

	params := &sms.GetConnectorsInput{
		MaxResults: aws.Int64(1),
		NextToken:  aws.String("NextToken"),
	}
	resp, err := svc.GetConnectors(params)

	if err != nil {
		// Print the error, cast err to awserr.Error to get the Code and
		// Message from an error.
		fmt.Println(err.Error())
		return
	}

	// Pretty-print the response data.
	fmt.Println(resp)
}

func ExampleSMS_GetReplicationJobs() {
	sess := session.Must(session.NewSession())

	svc := sms.New(sess)

	params := &sms.GetReplicationJobsInput{
		MaxResults:       aws.Int64(1),
		NextToken:        aws.String("NextToken"),
		ReplicationJobId: aws.String("ReplicationJobId"),
	}
	resp, err := svc.GetReplicationJobs(params)

	if err != nil {
		// Print the error, cast err to awserr.Error to get the Code and
		// Message from an error.
		fmt.Println(err.Error())
		return
	}

	// Pretty-print the response data.
	fmt.Println(resp)
}

func ExampleSMS_GetReplicationRuns() {
	sess := session.Must(session.NewSession())

	svc := sms.New(sess)

	params := &sms.GetReplicationRunsInput{
		ReplicationJobId: aws.String("ReplicationJobId"), // Required
		MaxResults:       aws.Int64(1),
		NextToken:        aws.String("NextToken"),
	}
	resp, err := svc.GetReplicationRuns(params)

	if err != nil {
		// Print the error, cast err to awserr.Error to get the Code and
		// Message from an error.
		fmt.Println(err.Error())
		return
	}

	// Pretty-print the response data.
	fmt.Println(resp)
}

func ExampleSMS_GetServers() {
	sess := session.Must(session.NewSession())

	svc := sms.New(sess)

	params := &sms.GetServersInput{
		MaxResults: aws.Int64(1),
		NextToken:  aws.String("NextToken"),
	}
	resp, err := svc.GetServers(params)

	if err != nil {
		// Print the error, cast err to awserr.Error to get the Code and
		// Message from an error.
		fmt.Println(err.Error())
		return
	}

	// Pretty-print the response data.
	fmt.Println(resp)
}

func ExampleSMS_ImportServerCatalog() {
	sess := session.Must(session.NewSession())

	svc := sms.New(sess)

	var params *sms.ImportServerCatalogInput
	resp, err := svc.ImportServerCatalog(params)

	if err != nil {
		// Print the error, cast err to awserr.Error to get the Code and
		// Message from an error.
		fmt.Println(err.Error())
		return
	}

	// Pretty-print the response data.
	fmt.Println(resp)
}

func ExampleSMS_StartOnDemandReplicationRun() {
	sess := session.Must(session.NewSession())

	svc := sms.New(sess)

	params := &sms.StartOnDemandReplicationRunInput{
		ReplicationJobId: aws.String("ReplicationJobId"), // Required
		Description:      aws.String("Description"),
	}
	resp, err := svc.StartOnDemandReplicationRun(params)

	if err != nil {
		// Print the error, cast err to awserr.Error to get the Code and
		// Message from an error.
		fmt.Println(err.Error())
		return
	}

	// Pretty-print the response data.
	fmt.Println(resp)
}

func ExampleSMS_UpdateReplicationJob() {
	sess := session.Must(session.NewSession())

	svc := sms.New(sess)

	params := &sms.UpdateReplicationJobInput{
		ReplicationJobId:            aws.String("ReplicationJobId"), // Required
		Description:                 aws.String("Description"),
		Frequency:                   aws.Int64(1),
		LicenseType:                 aws.String("LicenseType"),
		NextReplicationRunStartTime: aws.Time(time.Now()),
		RoleName:                    aws.String("RoleName"),
	}
	resp, err := svc.UpdateReplicationJob(params)

	if err != nil {
		// Print the error, cast err to awserr.Error to get the Code and
		// Message from an error.
		fmt.Println(err.Error())
		return
	}

	// Pretty-print the response data.
	fmt.Println(resp)
}
