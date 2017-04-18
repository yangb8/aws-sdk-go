package s3_test

import (
	"bytes"
	"io/ioutil"
	"net/http"
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/yangb8/aws-sdk-go/aws"
	"github.com/yangb8/aws-sdk-go/aws/awsutil"
	"github.com/yangb8/aws-sdk-go/aws/request"
	"github.com/yangb8/aws-sdk-go/awstesting/unit"
	"github.com/yangb8/aws-sdk-go/service/s3"
)

var s3LocationTests = []struct {
	body string
	loc  string
}{
	{`<?xml version="1.0" encoding="UTF-8"?><LocationConstraint xmlns="http://s3.amazonaws.com/doc/2006-03-01/"/>`, ``},
	{`<?xml version="1.0" encoding="UTF-8"?><LocationConstraint xmlns="http://s3.amazonaws.com/doc/2006-03-01/">EU</LocationConstraint>`, `EU`},
}

func TestGetBucketLocation(t *testing.T) {
	for _, test := range s3LocationTests {
		s := s3.New(unit.Session)
		s.Handlers.Send.Clear()
		s.Handlers.Send.PushBack(func(r *request.Request) {
			reader := ioutil.NopCloser(bytes.NewReader([]byte(test.body)))
			r.HTTPResponse = &http.Response{StatusCode: 200, Body: reader}
		})

		resp, err := s.GetBucketLocation(&s3.GetBucketLocationInput{Bucket: aws.String("bucket")})
		assert.NoError(t, err)
		if test.loc == "" {
			assert.Nil(t, resp.LocationConstraint)
		} else {
			assert.Equal(t, test.loc, *resp.LocationConstraint)
		}
	}
}

func TestPopulateLocationConstraint(t *testing.T) {
	s := s3.New(unit.Session)
	in := &s3.CreateBucketInput{
		Bucket: aws.String("bucket"),
	}
	req, _ := s.CreateBucketRequest(in)
	err := req.Build()
	assert.NoError(t, err)
	v, _ := awsutil.ValuesAtPath(req.Params, "CreateBucketConfiguration.LocationConstraint")
	assert.Equal(t, "mock-region", *(v[0].(*string)))
	assert.Nil(t, in.CreateBucketConfiguration) // don't modify original params
}

func TestNoPopulateLocationConstraintIfProvided(t *testing.T) {
	s := s3.New(unit.Session)
	req, _ := s.CreateBucketRequest(&s3.CreateBucketInput{
		Bucket: aws.String("bucket"),
		CreateBucketConfiguration: &s3.CreateBucketConfiguration{},
	})
	err := req.Build()
	assert.NoError(t, err)
	v, _ := awsutil.ValuesAtPath(req.Params, "CreateBucketConfiguration.LocationConstraint")
	assert.Equal(t, 0, len(v))
}

func TestNoPopulateLocationConstraintIfClassic(t *testing.T) {
	s := s3.New(unit.Session, &aws.Config{Region: aws.String("us-east-1")})
	req, _ := s.CreateBucketRequest(&s3.CreateBucketInput{
		Bucket: aws.String("bucket"),
	})
	err := req.Build()
	assert.NoError(t, err)
	v, _ := awsutil.ValuesAtPath(req.Params, "CreateBucketConfiguration.LocationConstraint")
	assert.Equal(t, 0, len(v))
}
