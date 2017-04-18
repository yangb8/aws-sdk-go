// Code generated by private/model/cli/gen-api/main.go. DO NOT EDIT.

package cloudformation

import (
	"time"

	"github.com/yangb8/aws-sdk-go/aws"
	"github.com/yangb8/aws-sdk-go/aws/request"
)

// WaitUntilChangeSetCreateComplete uses the AWS CloudFormation API operation
// DescribeChangeSet to wait for a condition to be met before returning.
// If the condition is not meet within the max attempt window an error will
// be returned.
func (c *CloudFormation) WaitUntilChangeSetCreateComplete(input *DescribeChangeSetInput) error {
	return c.WaitUntilChangeSetCreateCompleteWithContext(aws.BackgroundContext(), input)
}

// WaitUntilChangeSetCreateCompleteWithContext is an extended version of WaitUntilChangeSetCreateComplete.
// With the support for passing in a context and options to configure the
// Waiter and the underlying request options.
//
// The context must be non-nil and will be used for request cancellation. If
// the context is nil a panic will occur. In the future the SDK may create
// sub-contexts for http.Requests. See https://golang.org/pkg/context/
// for more information on using Contexts.
func (c *CloudFormation) WaitUntilChangeSetCreateCompleteWithContext(ctx aws.Context, input *DescribeChangeSetInput, opts ...request.WaiterOption) error {
	w := request.Waiter{
		Name:        "WaitUntilChangeSetCreateComplete",
		MaxAttempts: 120,
		Delay:       request.ConstantWaiterDelay(30 * time.Second),
		Acceptors: []request.WaiterAcceptor{
			{
				State:   request.SuccessWaiterState,
				Matcher: request.PathWaiterMatch, Argument: "Status",
				Expected: "CREATE_COMPLETE",
			},
			{
				State:   request.FailureWaiterState,
				Matcher: request.PathWaiterMatch, Argument: "Status",
				Expected: "FAILED",
			},
			{
				State:    request.FailureWaiterState,
				Matcher:  request.ErrorWaiterMatch,
				Expected: "ValidationError",
			},
		},
		Logger: c.Config.Logger,
		NewRequest: func(opts []request.Option) (*request.Request, error) {
			var inCpy *DescribeChangeSetInput
			if input != nil {
				tmp := *input
				inCpy = &tmp
			}
			req, _ := c.DescribeChangeSetRequest(inCpy)
			req.SetContext(ctx)
			req.ApplyOptions(opts...)
			return req, nil
		},
	}
	w.ApplyOptions(opts...)

	return w.WaitWithContext(ctx)
}

// WaitUntilStackCreateComplete uses the AWS CloudFormation API operation
// DescribeStacks to wait for a condition to be met before returning.
// If the condition is not meet within the max attempt window an error will
// be returned.
func (c *CloudFormation) WaitUntilStackCreateComplete(input *DescribeStacksInput) error {
	return c.WaitUntilStackCreateCompleteWithContext(aws.BackgroundContext(), input)
}

// WaitUntilStackCreateCompleteWithContext is an extended version of WaitUntilStackCreateComplete.
// With the support for passing in a context and options to configure the
// Waiter and the underlying request options.
//
// The context must be non-nil and will be used for request cancellation. If
// the context is nil a panic will occur. In the future the SDK may create
// sub-contexts for http.Requests. See https://golang.org/pkg/context/
// for more information on using Contexts.
func (c *CloudFormation) WaitUntilStackCreateCompleteWithContext(ctx aws.Context, input *DescribeStacksInput, opts ...request.WaiterOption) error {
	w := request.Waiter{
		Name:        "WaitUntilStackCreateComplete",
		MaxAttempts: 120,
		Delay:       request.ConstantWaiterDelay(30 * time.Second),
		Acceptors: []request.WaiterAcceptor{
			{
				State:   request.SuccessWaiterState,
				Matcher: request.PathAllWaiterMatch, Argument: "Stacks[].StackStatus",
				Expected: "CREATE_COMPLETE",
			},
			{
				State:   request.FailureWaiterState,
				Matcher: request.PathAnyWaiterMatch, Argument: "Stacks[].StackStatus",
				Expected: "CREATE_FAILED",
			},
			{
				State:   request.FailureWaiterState,
				Matcher: request.PathAnyWaiterMatch, Argument: "Stacks[].StackStatus",
				Expected: "DELETE_COMPLETE",
			},
			{
				State:   request.FailureWaiterState,
				Matcher: request.PathAnyWaiterMatch, Argument: "Stacks[].StackStatus",
				Expected: "DELETE_FAILED",
			},
			{
				State:   request.FailureWaiterState,
				Matcher: request.PathAnyWaiterMatch, Argument: "Stacks[].StackStatus",
				Expected: "ROLLBACK_FAILED",
			},
			{
				State:   request.FailureWaiterState,
				Matcher: request.PathAnyWaiterMatch, Argument: "Stacks[].StackStatus",
				Expected: "ROLLBACK_COMPLETE",
			},
			{
				State:    request.FailureWaiterState,
				Matcher:  request.ErrorWaiterMatch,
				Expected: "ValidationError",
			},
		},
		Logger: c.Config.Logger,
		NewRequest: func(opts []request.Option) (*request.Request, error) {
			var inCpy *DescribeStacksInput
			if input != nil {
				tmp := *input
				inCpy = &tmp
			}
			req, _ := c.DescribeStacksRequest(inCpy)
			req.SetContext(ctx)
			req.ApplyOptions(opts...)
			return req, nil
		},
	}
	w.ApplyOptions(opts...)

	return w.WaitWithContext(ctx)
}

// WaitUntilStackDeleteComplete uses the AWS CloudFormation API operation
// DescribeStacks to wait for a condition to be met before returning.
// If the condition is not meet within the max attempt window an error will
// be returned.
func (c *CloudFormation) WaitUntilStackDeleteComplete(input *DescribeStacksInput) error {
	return c.WaitUntilStackDeleteCompleteWithContext(aws.BackgroundContext(), input)
}

// WaitUntilStackDeleteCompleteWithContext is an extended version of WaitUntilStackDeleteComplete.
// With the support for passing in a context and options to configure the
// Waiter and the underlying request options.
//
// The context must be non-nil and will be used for request cancellation. If
// the context is nil a panic will occur. In the future the SDK may create
// sub-contexts for http.Requests. See https://golang.org/pkg/context/
// for more information on using Contexts.
func (c *CloudFormation) WaitUntilStackDeleteCompleteWithContext(ctx aws.Context, input *DescribeStacksInput, opts ...request.WaiterOption) error {
	w := request.Waiter{
		Name:        "WaitUntilStackDeleteComplete",
		MaxAttempts: 120,
		Delay:       request.ConstantWaiterDelay(30 * time.Second),
		Acceptors: []request.WaiterAcceptor{
			{
				State:   request.SuccessWaiterState,
				Matcher: request.PathAllWaiterMatch, Argument: "Stacks[].StackStatus",
				Expected: "DELETE_COMPLETE",
			},
			{
				State:    request.SuccessWaiterState,
				Matcher:  request.ErrorWaiterMatch,
				Expected: "ValidationError",
			},
			{
				State:   request.FailureWaiterState,
				Matcher: request.PathAnyWaiterMatch, Argument: "Stacks[].StackStatus",
				Expected: "DELETE_FAILED",
			},
			{
				State:   request.FailureWaiterState,
				Matcher: request.PathAnyWaiterMatch, Argument: "Stacks[].StackStatus",
				Expected: "CREATE_FAILED",
			},
			{
				State:   request.FailureWaiterState,
				Matcher: request.PathAnyWaiterMatch, Argument: "Stacks[].StackStatus",
				Expected: "ROLLBACK_FAILED",
			},
			{
				State:   request.FailureWaiterState,
				Matcher: request.PathAnyWaiterMatch, Argument: "Stacks[].StackStatus",
				Expected: "UPDATE_ROLLBACK_FAILED",
			},
			{
				State:   request.FailureWaiterState,
				Matcher: request.PathAnyWaiterMatch, Argument: "Stacks[].StackStatus",
				Expected: "UPDATE_ROLLBACK_IN_PROGRESS",
			},
		},
		Logger: c.Config.Logger,
		NewRequest: func(opts []request.Option) (*request.Request, error) {
			var inCpy *DescribeStacksInput
			if input != nil {
				tmp := *input
				inCpy = &tmp
			}
			req, _ := c.DescribeStacksRequest(inCpy)
			req.SetContext(ctx)
			req.ApplyOptions(opts...)
			return req, nil
		},
	}
	w.ApplyOptions(opts...)

	return w.WaitWithContext(ctx)
}

// WaitUntilStackExists uses the AWS CloudFormation API operation
// DescribeStacks to wait for a condition to be met before returning.
// If the condition is not meet within the max attempt window an error will
// be returned.
func (c *CloudFormation) WaitUntilStackExists(input *DescribeStacksInput) error {
	return c.WaitUntilStackExistsWithContext(aws.BackgroundContext(), input)
}

// WaitUntilStackExistsWithContext is an extended version of WaitUntilStackExists.
// With the support for passing in a context and options to configure the
// Waiter and the underlying request options.
//
// The context must be non-nil and will be used for request cancellation. If
// the context is nil a panic will occur. In the future the SDK may create
// sub-contexts for http.Requests. See https://golang.org/pkg/context/
// for more information on using Contexts.
func (c *CloudFormation) WaitUntilStackExistsWithContext(ctx aws.Context, input *DescribeStacksInput, opts ...request.WaiterOption) error {
	w := request.Waiter{
		Name:        "WaitUntilStackExists",
		MaxAttempts: 20,
		Delay:       request.ConstantWaiterDelay(5 * time.Second),
		Acceptors: []request.WaiterAcceptor{
			{
				State:    request.SuccessWaiterState,
				Matcher:  request.StatusWaiterMatch,
				Expected: 200,
			},
			{
				State:    request.RetryWaiterState,
				Matcher:  request.ErrorWaiterMatch,
				Expected: "ValidationError",
			},
		},
		Logger: c.Config.Logger,
		NewRequest: func(opts []request.Option) (*request.Request, error) {
			var inCpy *DescribeStacksInput
			if input != nil {
				tmp := *input
				inCpy = &tmp
			}
			req, _ := c.DescribeStacksRequest(inCpy)
			req.SetContext(ctx)
			req.ApplyOptions(opts...)
			return req, nil
		},
	}
	w.ApplyOptions(opts...)

	return w.WaitWithContext(ctx)
}

// WaitUntilStackUpdateComplete uses the AWS CloudFormation API operation
// DescribeStacks to wait for a condition to be met before returning.
// If the condition is not meet within the max attempt window an error will
// be returned.
func (c *CloudFormation) WaitUntilStackUpdateComplete(input *DescribeStacksInput) error {
	return c.WaitUntilStackUpdateCompleteWithContext(aws.BackgroundContext(), input)
}

// WaitUntilStackUpdateCompleteWithContext is an extended version of WaitUntilStackUpdateComplete.
// With the support for passing in a context and options to configure the
// Waiter and the underlying request options.
//
// The context must be non-nil and will be used for request cancellation. If
// the context is nil a panic will occur. In the future the SDK may create
// sub-contexts for http.Requests. See https://golang.org/pkg/context/
// for more information on using Contexts.
func (c *CloudFormation) WaitUntilStackUpdateCompleteWithContext(ctx aws.Context, input *DescribeStacksInput, opts ...request.WaiterOption) error {
	w := request.Waiter{
		Name:        "WaitUntilStackUpdateComplete",
		MaxAttempts: 120,
		Delay:       request.ConstantWaiterDelay(30 * time.Second),
		Acceptors: []request.WaiterAcceptor{
			{
				State:   request.SuccessWaiterState,
				Matcher: request.PathAllWaiterMatch, Argument: "Stacks[].StackStatus",
				Expected: "UPDATE_COMPLETE",
			},
			{
				State:   request.FailureWaiterState,
				Matcher: request.PathAnyWaiterMatch, Argument: "Stacks[].StackStatus",
				Expected: "UPDATE_FAILED",
			},
			{
				State:   request.FailureWaiterState,
				Matcher: request.PathAnyWaiterMatch, Argument: "Stacks[].StackStatus",
				Expected: "UPDATE_ROLLBACK_FAILED",
			},
			{
				State:   request.FailureWaiterState,
				Matcher: request.PathAnyWaiterMatch, Argument: "Stacks[].StackStatus",
				Expected: "UPDATE_ROLLBACK_COMPLETE",
			},
			{
				State:    request.FailureWaiterState,
				Matcher:  request.ErrorWaiterMatch,
				Expected: "ValidationError",
			},
		},
		Logger: c.Config.Logger,
		NewRequest: func(opts []request.Option) (*request.Request, error) {
			var inCpy *DescribeStacksInput
			if input != nil {
				tmp := *input
				inCpy = &tmp
			}
			req, _ := c.DescribeStacksRequest(inCpy)
			req.SetContext(ctx)
			req.ApplyOptions(opts...)
			return req, nil
		},
	}
	w.ApplyOptions(opts...)

	return w.WaitWithContext(ctx)
}
