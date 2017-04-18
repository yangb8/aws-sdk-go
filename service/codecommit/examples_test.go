// Code generated by private/model/cli/gen-api/main.go. DO NOT EDIT.

package codecommit_test

import (
	"bytes"
	"fmt"
	"time"

	"github.com/yangb8/aws-sdk-go/aws"
	"github.com/yangb8/aws-sdk-go/aws/session"
	"github.com/yangb8/aws-sdk-go/service/codecommit"
)

var _ time.Duration
var _ bytes.Buffer

func ExampleCodeCommit_BatchGetRepositories() {
	sess := session.Must(session.NewSession())

	svc := codecommit.New(sess)

	params := &codecommit.BatchGetRepositoriesInput{
		RepositoryNames: []*string{ // Required
			aws.String("RepositoryName"), // Required
			// More values...
		},
	}
	resp, err := svc.BatchGetRepositories(params)

	if err != nil {
		// Print the error, cast err to awserr.Error to get the Code and
		// Message from an error.
		fmt.Println(err.Error())
		return
	}

	// Pretty-print the response data.
	fmt.Println(resp)
}

func ExampleCodeCommit_CreateBranch() {
	sess := session.Must(session.NewSession())

	svc := codecommit.New(sess)

	params := &codecommit.CreateBranchInput{
		BranchName:     aws.String("BranchName"),     // Required
		CommitId:       aws.String("CommitId"),       // Required
		RepositoryName: aws.String("RepositoryName"), // Required
	}
	resp, err := svc.CreateBranch(params)

	if err != nil {
		// Print the error, cast err to awserr.Error to get the Code and
		// Message from an error.
		fmt.Println(err.Error())
		return
	}

	// Pretty-print the response data.
	fmt.Println(resp)
}

func ExampleCodeCommit_CreateRepository() {
	sess := session.Must(session.NewSession())

	svc := codecommit.New(sess)

	params := &codecommit.CreateRepositoryInput{
		RepositoryName:        aws.String("RepositoryName"), // Required
		RepositoryDescription: aws.String("RepositoryDescription"),
	}
	resp, err := svc.CreateRepository(params)

	if err != nil {
		// Print the error, cast err to awserr.Error to get the Code and
		// Message from an error.
		fmt.Println(err.Error())
		return
	}

	// Pretty-print the response data.
	fmt.Println(resp)
}

func ExampleCodeCommit_DeleteRepository() {
	sess := session.Must(session.NewSession())

	svc := codecommit.New(sess)

	params := &codecommit.DeleteRepositoryInput{
		RepositoryName: aws.String("RepositoryName"), // Required
	}
	resp, err := svc.DeleteRepository(params)

	if err != nil {
		// Print the error, cast err to awserr.Error to get the Code and
		// Message from an error.
		fmt.Println(err.Error())
		return
	}

	// Pretty-print the response data.
	fmt.Println(resp)
}

func ExampleCodeCommit_GetBlob() {
	sess := session.Must(session.NewSession())

	svc := codecommit.New(sess)

	params := &codecommit.GetBlobInput{
		BlobId:         aws.String("ObjectId"),       // Required
		RepositoryName: aws.String("RepositoryName"), // Required
	}
	resp, err := svc.GetBlob(params)

	if err != nil {
		// Print the error, cast err to awserr.Error to get the Code and
		// Message from an error.
		fmt.Println(err.Error())
		return
	}

	// Pretty-print the response data.
	fmt.Println(resp)
}

func ExampleCodeCommit_GetBranch() {
	sess := session.Must(session.NewSession())

	svc := codecommit.New(sess)

	params := &codecommit.GetBranchInput{
		BranchName:     aws.String("BranchName"),
		RepositoryName: aws.String("RepositoryName"),
	}
	resp, err := svc.GetBranch(params)

	if err != nil {
		// Print the error, cast err to awserr.Error to get the Code and
		// Message from an error.
		fmt.Println(err.Error())
		return
	}

	// Pretty-print the response data.
	fmt.Println(resp)
}

func ExampleCodeCommit_GetCommit() {
	sess := session.Must(session.NewSession())

	svc := codecommit.New(sess)

	params := &codecommit.GetCommitInput{
		CommitId:       aws.String("ObjectId"),       // Required
		RepositoryName: aws.String("RepositoryName"), // Required
	}
	resp, err := svc.GetCommit(params)

	if err != nil {
		// Print the error, cast err to awserr.Error to get the Code and
		// Message from an error.
		fmt.Println(err.Error())
		return
	}

	// Pretty-print the response data.
	fmt.Println(resp)
}

func ExampleCodeCommit_GetDifferences() {
	sess := session.Must(session.NewSession())

	svc := codecommit.New(sess)

	params := &codecommit.GetDifferencesInput{
		AfterCommitSpecifier:  aws.String("CommitName"),     // Required
		RepositoryName:        aws.String("RepositoryName"), // Required
		AfterPath:             aws.String("Path"),
		BeforeCommitSpecifier: aws.String("CommitName"),
		BeforePath:            aws.String("Path"),
		MaxResults:            aws.Int64(1),
		NextToken:             aws.String("NextToken"),
	}
	resp, err := svc.GetDifferences(params)

	if err != nil {
		// Print the error, cast err to awserr.Error to get the Code and
		// Message from an error.
		fmt.Println(err.Error())
		return
	}

	// Pretty-print the response data.
	fmt.Println(resp)
}

func ExampleCodeCommit_GetRepository() {
	sess := session.Must(session.NewSession())

	svc := codecommit.New(sess)

	params := &codecommit.GetRepositoryInput{
		RepositoryName: aws.String("RepositoryName"), // Required
	}
	resp, err := svc.GetRepository(params)

	if err != nil {
		// Print the error, cast err to awserr.Error to get the Code and
		// Message from an error.
		fmt.Println(err.Error())
		return
	}

	// Pretty-print the response data.
	fmt.Println(resp)
}

func ExampleCodeCommit_GetRepositoryTriggers() {
	sess := session.Must(session.NewSession())

	svc := codecommit.New(sess)

	params := &codecommit.GetRepositoryTriggersInput{
		RepositoryName: aws.String("RepositoryName"), // Required
	}
	resp, err := svc.GetRepositoryTriggers(params)

	if err != nil {
		// Print the error, cast err to awserr.Error to get the Code and
		// Message from an error.
		fmt.Println(err.Error())
		return
	}

	// Pretty-print the response data.
	fmt.Println(resp)
}

func ExampleCodeCommit_ListBranches() {
	sess := session.Must(session.NewSession())

	svc := codecommit.New(sess)

	params := &codecommit.ListBranchesInput{
		RepositoryName: aws.String("RepositoryName"), // Required
		NextToken:      aws.String("NextToken"),
	}
	resp, err := svc.ListBranches(params)

	if err != nil {
		// Print the error, cast err to awserr.Error to get the Code and
		// Message from an error.
		fmt.Println(err.Error())
		return
	}

	// Pretty-print the response data.
	fmt.Println(resp)
}

func ExampleCodeCommit_ListRepositories() {
	sess := session.Must(session.NewSession())

	svc := codecommit.New(sess)

	params := &codecommit.ListRepositoriesInput{
		NextToken: aws.String("NextToken"),
		Order:     aws.String("OrderEnum"),
		SortBy:    aws.String("SortByEnum"),
	}
	resp, err := svc.ListRepositories(params)

	if err != nil {
		// Print the error, cast err to awserr.Error to get the Code and
		// Message from an error.
		fmt.Println(err.Error())
		return
	}

	// Pretty-print the response data.
	fmt.Println(resp)
}

func ExampleCodeCommit_PutRepositoryTriggers() {
	sess := session.Must(session.NewSession())

	svc := codecommit.New(sess)

	params := &codecommit.PutRepositoryTriggersInput{
		RepositoryName: aws.String("RepositoryName"), // Required
		Triggers: []*codecommit.RepositoryTrigger{ // Required
			{ // Required
				DestinationArn: aws.String("Arn"), // Required
				Events: []*string{ // Required
					aws.String("RepositoryTriggerEventEnum"), // Required
					// More values...
				},
				Name: aws.String("RepositoryTriggerName"), // Required
				Branches: []*string{
					aws.String("BranchName"), // Required
					// More values...
				},
				CustomData: aws.String("RepositoryTriggerCustomData"),
			},
			// More values...
		},
	}
	resp, err := svc.PutRepositoryTriggers(params)

	if err != nil {
		// Print the error, cast err to awserr.Error to get the Code and
		// Message from an error.
		fmt.Println(err.Error())
		return
	}

	// Pretty-print the response data.
	fmt.Println(resp)
}

func ExampleCodeCommit_TestRepositoryTriggers() {
	sess := session.Must(session.NewSession())

	svc := codecommit.New(sess)

	params := &codecommit.TestRepositoryTriggersInput{
		RepositoryName: aws.String("RepositoryName"), // Required
		Triggers: []*codecommit.RepositoryTrigger{ // Required
			{ // Required
				DestinationArn: aws.String("Arn"), // Required
				Events: []*string{ // Required
					aws.String("RepositoryTriggerEventEnum"), // Required
					// More values...
				},
				Name: aws.String("RepositoryTriggerName"), // Required
				Branches: []*string{
					aws.String("BranchName"), // Required
					// More values...
				},
				CustomData: aws.String("RepositoryTriggerCustomData"),
			},
			// More values...
		},
	}
	resp, err := svc.TestRepositoryTriggers(params)

	if err != nil {
		// Print the error, cast err to awserr.Error to get the Code and
		// Message from an error.
		fmt.Println(err.Error())
		return
	}

	// Pretty-print the response data.
	fmt.Println(resp)
}

func ExampleCodeCommit_UpdateDefaultBranch() {
	sess := session.Must(session.NewSession())

	svc := codecommit.New(sess)

	params := &codecommit.UpdateDefaultBranchInput{
		DefaultBranchName: aws.String("BranchName"),     // Required
		RepositoryName:    aws.String("RepositoryName"), // Required
	}
	resp, err := svc.UpdateDefaultBranch(params)

	if err != nil {
		// Print the error, cast err to awserr.Error to get the Code and
		// Message from an error.
		fmt.Println(err.Error())
		return
	}

	// Pretty-print the response data.
	fmt.Println(resp)
}

func ExampleCodeCommit_UpdateRepositoryDescription() {
	sess := session.Must(session.NewSession())

	svc := codecommit.New(sess)

	params := &codecommit.UpdateRepositoryDescriptionInput{
		RepositoryName:        aws.String("RepositoryName"), // Required
		RepositoryDescription: aws.String("RepositoryDescription"),
	}
	resp, err := svc.UpdateRepositoryDescription(params)

	if err != nil {
		// Print the error, cast err to awserr.Error to get the Code and
		// Message from an error.
		fmt.Println(err.Error())
		return
	}

	// Pretty-print the response data.
	fmt.Println(resp)
}

func ExampleCodeCommit_UpdateRepositoryName() {
	sess := session.Must(session.NewSession())

	svc := codecommit.New(sess)

	params := &codecommit.UpdateRepositoryNameInput{
		NewName: aws.String("RepositoryName"), // Required
		OldName: aws.String("RepositoryName"), // Required
	}
	resp, err := svc.UpdateRepositoryName(params)

	if err != nil {
		// Print the error, cast err to awserr.Error to get the Code and
		// Message from an error.
		fmt.Println(err.Error())
		return
	}

	// Pretty-print the response data.
	fmt.Println(resp)
}
