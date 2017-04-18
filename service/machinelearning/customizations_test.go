package machinelearning_test

import (
	"bytes"
	"io/ioutil"
	"net/http"
	"testing"

	"github.com/stretchr/testify/assert"

	"github.com/yangb8/aws-sdk-go/aws"
	"github.com/yangb8/aws-sdk-go/aws/request"
	"github.com/yangb8/aws-sdk-go/awstesting/unit"
	"github.com/yangb8/aws-sdk-go/service/machinelearning"
)

func TestPredictEndpoint(t *testing.T) {
	ml := machinelearning.New(unit.Session)
	ml.Handlers.Send.Clear()
	ml.Handlers.Send.PushBack(func(r *request.Request) {
		r.HTTPResponse = &http.Response{
			StatusCode: 200,
			Header:     http.Header{},
			Body:       ioutil.NopCloser(bytes.NewReader([]byte("{}"))),
		}
	})

	req, _ := ml.PredictRequest(&machinelearning.PredictInput{
		PredictEndpoint: aws.String("https://localhost/endpoint"),
		MLModelId:       aws.String("id"),
		Record:          map[string]*string{},
	})
	err := req.Send()

	assert.Nil(t, err)
	assert.Equal(t, "https://localhost/endpoint", req.HTTPRequest.URL.String())
}
