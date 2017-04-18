// Code generated by private/model/cli/gen-api/main.go. DO NOT EDIT.

package ses_test

import (
	"bytes"
	"fmt"
	"time"

	"github.com/yangb8/aws-sdk-go/aws"
	"github.com/yangb8/aws-sdk-go/aws/session"
	"github.com/yangb8/aws-sdk-go/service/ses"
)

var _ time.Duration
var _ bytes.Buffer

func ExampleSES_CloneReceiptRuleSet() {
	sess := session.Must(session.NewSession())

	svc := ses.New(sess)

	params := &ses.CloneReceiptRuleSetInput{
		OriginalRuleSetName: aws.String("ReceiptRuleSetName"), // Required
		RuleSetName:         aws.String("ReceiptRuleSetName"), // Required
	}
	resp, err := svc.CloneReceiptRuleSet(params)

	if err != nil {
		// Print the error, cast err to awserr.Error to get the Code and
		// Message from an error.
		fmt.Println(err.Error())
		return
	}

	// Pretty-print the response data.
	fmt.Println(resp)
}

func ExampleSES_CreateConfigurationSet() {
	sess := session.Must(session.NewSession())

	svc := ses.New(sess)

	params := &ses.CreateConfigurationSetInput{
		ConfigurationSet: &ses.ConfigurationSet{ // Required
			Name: aws.String("ConfigurationSetName"), // Required
		},
	}
	resp, err := svc.CreateConfigurationSet(params)

	if err != nil {
		// Print the error, cast err to awserr.Error to get the Code and
		// Message from an error.
		fmt.Println(err.Error())
		return
	}

	// Pretty-print the response data.
	fmt.Println(resp)
}

func ExampleSES_CreateConfigurationSetEventDestination() {
	sess := session.Must(session.NewSession())

	svc := ses.New(sess)

	params := &ses.CreateConfigurationSetEventDestinationInput{
		ConfigurationSetName: aws.String("ConfigurationSetName"), // Required
		EventDestination: &ses.EventDestination{ // Required
			MatchingEventTypes: []*string{ // Required
				aws.String("EventType"), // Required
				// More values...
			},
			Name: aws.String("EventDestinationName"), // Required
			CloudWatchDestination: &ses.CloudWatchDestination{
				DimensionConfigurations: []*ses.CloudWatchDimensionConfiguration{ // Required
					{ // Required
						DefaultDimensionValue: aws.String("DefaultDimensionValue"), // Required
						DimensionName:         aws.String("DimensionName"),         // Required
						DimensionValueSource:  aws.String("DimensionValueSource"),  // Required
					},
					// More values...
				},
			},
			Enabled: aws.Bool(true),
			KinesisFirehoseDestination: &ses.KinesisFirehoseDestination{
				DeliveryStreamARN: aws.String("AmazonResourceName"), // Required
				IAMRoleARN:        aws.String("AmazonResourceName"), // Required
			},
		},
	}
	resp, err := svc.CreateConfigurationSetEventDestination(params)

	if err != nil {
		// Print the error, cast err to awserr.Error to get the Code and
		// Message from an error.
		fmt.Println(err.Error())
		return
	}

	// Pretty-print the response data.
	fmt.Println(resp)
}

func ExampleSES_CreateReceiptFilter() {
	sess := session.Must(session.NewSession())

	svc := ses.New(sess)

	params := &ses.CreateReceiptFilterInput{
		Filter: &ses.ReceiptFilter{ // Required
			IpFilter: &ses.ReceiptIpFilter{ // Required
				Cidr:   aws.String("Cidr"),                // Required
				Policy: aws.String("ReceiptFilterPolicy"), // Required
			},
			Name: aws.String("ReceiptFilterName"), // Required
		},
	}
	resp, err := svc.CreateReceiptFilter(params)

	if err != nil {
		// Print the error, cast err to awserr.Error to get the Code and
		// Message from an error.
		fmt.Println(err.Error())
		return
	}

	// Pretty-print the response data.
	fmt.Println(resp)
}

func ExampleSES_CreateReceiptRule() {
	sess := session.Must(session.NewSession())

	svc := ses.New(sess)

	params := &ses.CreateReceiptRuleInput{
		Rule: &ses.ReceiptRule{ // Required
			Name: aws.String("ReceiptRuleName"), // Required
			Actions: []*ses.ReceiptAction{
				{ // Required
					AddHeaderAction: &ses.AddHeaderAction{
						HeaderName:  aws.String("HeaderName"),  // Required
						HeaderValue: aws.String("HeaderValue"), // Required
					},
					BounceAction: &ses.BounceAction{
						Message:       aws.String("BounceMessage"),       // Required
						Sender:        aws.String("Address"),             // Required
						SmtpReplyCode: aws.String("BounceSmtpReplyCode"), // Required
						StatusCode:    aws.String("BounceStatusCode"),
						TopicArn:      aws.String("AmazonResourceName"),
					},
					LambdaAction: &ses.LambdaAction{
						FunctionArn:    aws.String("AmazonResourceName"), // Required
						InvocationType: aws.String("InvocationType"),
						TopicArn:       aws.String("AmazonResourceName"),
					},
					S3Action: &ses.S3Action{
						BucketName:      aws.String("S3BucketName"), // Required
						KmsKeyArn:       aws.String("AmazonResourceName"),
						ObjectKeyPrefix: aws.String("S3KeyPrefix"),
						TopicArn:        aws.String("AmazonResourceName"),
					},
					SNSAction: &ses.SNSAction{
						TopicArn: aws.String("AmazonResourceName"), // Required
						Encoding: aws.String("SNSActionEncoding"),
					},
					StopAction: &ses.StopAction{
						Scope:    aws.String("StopScope"), // Required
						TopicArn: aws.String("AmazonResourceName"),
					},
					WorkmailAction: &ses.WorkmailAction{
						OrganizationArn: aws.String("AmazonResourceName"), // Required
						TopicArn:        aws.String("AmazonResourceName"),
					},
				},
				// More values...
			},
			Enabled: aws.Bool(true),
			Recipients: []*string{
				aws.String("Recipient"), // Required
				// More values...
			},
			ScanEnabled: aws.Bool(true),
			TlsPolicy:   aws.String("TlsPolicy"),
		},
		RuleSetName: aws.String("ReceiptRuleSetName"), // Required
		After:       aws.String("ReceiptRuleName"),
	}
	resp, err := svc.CreateReceiptRule(params)

	if err != nil {
		// Print the error, cast err to awserr.Error to get the Code and
		// Message from an error.
		fmt.Println(err.Error())
		return
	}

	// Pretty-print the response data.
	fmt.Println(resp)
}

func ExampleSES_CreateReceiptRuleSet() {
	sess := session.Must(session.NewSession())

	svc := ses.New(sess)

	params := &ses.CreateReceiptRuleSetInput{
		RuleSetName: aws.String("ReceiptRuleSetName"), // Required
	}
	resp, err := svc.CreateReceiptRuleSet(params)

	if err != nil {
		// Print the error, cast err to awserr.Error to get the Code and
		// Message from an error.
		fmt.Println(err.Error())
		return
	}

	// Pretty-print the response data.
	fmt.Println(resp)
}

func ExampleSES_DeleteConfigurationSet() {
	sess := session.Must(session.NewSession())

	svc := ses.New(sess)

	params := &ses.DeleteConfigurationSetInput{
		ConfigurationSetName: aws.String("ConfigurationSetName"), // Required
	}
	resp, err := svc.DeleteConfigurationSet(params)

	if err != nil {
		// Print the error, cast err to awserr.Error to get the Code and
		// Message from an error.
		fmt.Println(err.Error())
		return
	}

	// Pretty-print the response data.
	fmt.Println(resp)
}

func ExampleSES_DeleteConfigurationSetEventDestination() {
	sess := session.Must(session.NewSession())

	svc := ses.New(sess)

	params := &ses.DeleteConfigurationSetEventDestinationInput{
		ConfigurationSetName: aws.String("ConfigurationSetName"), // Required
		EventDestinationName: aws.String("EventDestinationName"), // Required
	}
	resp, err := svc.DeleteConfigurationSetEventDestination(params)

	if err != nil {
		// Print the error, cast err to awserr.Error to get the Code and
		// Message from an error.
		fmt.Println(err.Error())
		return
	}

	// Pretty-print the response data.
	fmt.Println(resp)
}

func ExampleSES_DeleteIdentity() {
	sess := session.Must(session.NewSession())

	svc := ses.New(sess)

	params := &ses.DeleteIdentityInput{
		Identity: aws.String("Identity"), // Required
	}
	resp, err := svc.DeleteIdentity(params)

	if err != nil {
		// Print the error, cast err to awserr.Error to get the Code and
		// Message from an error.
		fmt.Println(err.Error())
		return
	}

	// Pretty-print the response data.
	fmt.Println(resp)
}

func ExampleSES_DeleteIdentityPolicy() {
	sess := session.Must(session.NewSession())

	svc := ses.New(sess)

	params := &ses.DeleteIdentityPolicyInput{
		Identity:   aws.String("Identity"),   // Required
		PolicyName: aws.String("PolicyName"), // Required
	}
	resp, err := svc.DeleteIdentityPolicy(params)

	if err != nil {
		// Print the error, cast err to awserr.Error to get the Code and
		// Message from an error.
		fmt.Println(err.Error())
		return
	}

	// Pretty-print the response data.
	fmt.Println(resp)
}

func ExampleSES_DeleteReceiptFilter() {
	sess := session.Must(session.NewSession())

	svc := ses.New(sess)

	params := &ses.DeleteReceiptFilterInput{
		FilterName: aws.String("ReceiptFilterName"), // Required
	}
	resp, err := svc.DeleteReceiptFilter(params)

	if err != nil {
		// Print the error, cast err to awserr.Error to get the Code and
		// Message from an error.
		fmt.Println(err.Error())
		return
	}

	// Pretty-print the response data.
	fmt.Println(resp)
}

func ExampleSES_DeleteReceiptRule() {
	sess := session.Must(session.NewSession())

	svc := ses.New(sess)

	params := &ses.DeleteReceiptRuleInput{
		RuleName:    aws.String("ReceiptRuleName"),    // Required
		RuleSetName: aws.String("ReceiptRuleSetName"), // Required
	}
	resp, err := svc.DeleteReceiptRule(params)

	if err != nil {
		// Print the error, cast err to awserr.Error to get the Code and
		// Message from an error.
		fmt.Println(err.Error())
		return
	}

	// Pretty-print the response data.
	fmt.Println(resp)
}

func ExampleSES_DeleteReceiptRuleSet() {
	sess := session.Must(session.NewSession())

	svc := ses.New(sess)

	params := &ses.DeleteReceiptRuleSetInput{
		RuleSetName: aws.String("ReceiptRuleSetName"), // Required
	}
	resp, err := svc.DeleteReceiptRuleSet(params)

	if err != nil {
		// Print the error, cast err to awserr.Error to get the Code and
		// Message from an error.
		fmt.Println(err.Error())
		return
	}

	// Pretty-print the response data.
	fmt.Println(resp)
}

func ExampleSES_DeleteVerifiedEmailAddress() {
	sess := session.Must(session.NewSession())

	svc := ses.New(sess)

	params := &ses.DeleteVerifiedEmailAddressInput{
		EmailAddress: aws.String("Address"), // Required
	}
	resp, err := svc.DeleteVerifiedEmailAddress(params)

	if err != nil {
		// Print the error, cast err to awserr.Error to get the Code and
		// Message from an error.
		fmt.Println(err.Error())
		return
	}

	// Pretty-print the response data.
	fmt.Println(resp)
}

func ExampleSES_DescribeActiveReceiptRuleSet() {
	sess := session.Must(session.NewSession())

	svc := ses.New(sess)

	var params *ses.DescribeActiveReceiptRuleSetInput
	resp, err := svc.DescribeActiveReceiptRuleSet(params)

	if err != nil {
		// Print the error, cast err to awserr.Error to get the Code and
		// Message from an error.
		fmt.Println(err.Error())
		return
	}

	// Pretty-print the response data.
	fmt.Println(resp)
}

func ExampleSES_DescribeConfigurationSet() {
	sess := session.Must(session.NewSession())

	svc := ses.New(sess)

	params := &ses.DescribeConfigurationSetInput{
		ConfigurationSetName: aws.String("ConfigurationSetName"), // Required
		ConfigurationSetAttributeNames: []*string{
			aws.String("ConfigurationSetAttribute"), // Required
			// More values...
		},
	}
	resp, err := svc.DescribeConfigurationSet(params)

	if err != nil {
		// Print the error, cast err to awserr.Error to get the Code and
		// Message from an error.
		fmt.Println(err.Error())
		return
	}

	// Pretty-print the response data.
	fmt.Println(resp)
}

func ExampleSES_DescribeReceiptRule() {
	sess := session.Must(session.NewSession())

	svc := ses.New(sess)

	params := &ses.DescribeReceiptRuleInput{
		RuleName:    aws.String("ReceiptRuleName"),    // Required
		RuleSetName: aws.String("ReceiptRuleSetName"), // Required
	}
	resp, err := svc.DescribeReceiptRule(params)

	if err != nil {
		// Print the error, cast err to awserr.Error to get the Code and
		// Message from an error.
		fmt.Println(err.Error())
		return
	}

	// Pretty-print the response data.
	fmt.Println(resp)
}

func ExampleSES_DescribeReceiptRuleSet() {
	sess := session.Must(session.NewSession())

	svc := ses.New(sess)

	params := &ses.DescribeReceiptRuleSetInput{
		RuleSetName: aws.String("ReceiptRuleSetName"), // Required
	}
	resp, err := svc.DescribeReceiptRuleSet(params)

	if err != nil {
		// Print the error, cast err to awserr.Error to get the Code and
		// Message from an error.
		fmt.Println(err.Error())
		return
	}

	// Pretty-print the response data.
	fmt.Println(resp)
}

func ExampleSES_GetIdentityDkimAttributes() {
	sess := session.Must(session.NewSession())

	svc := ses.New(sess)

	params := &ses.GetIdentityDkimAttributesInput{
		Identities: []*string{ // Required
			aws.String("Identity"), // Required
			// More values...
		},
	}
	resp, err := svc.GetIdentityDkimAttributes(params)

	if err != nil {
		// Print the error, cast err to awserr.Error to get the Code and
		// Message from an error.
		fmt.Println(err.Error())
		return
	}

	// Pretty-print the response data.
	fmt.Println(resp)
}

func ExampleSES_GetIdentityMailFromDomainAttributes() {
	sess := session.Must(session.NewSession())

	svc := ses.New(sess)

	params := &ses.GetIdentityMailFromDomainAttributesInput{
		Identities: []*string{ // Required
			aws.String("Identity"), // Required
			// More values...
		},
	}
	resp, err := svc.GetIdentityMailFromDomainAttributes(params)

	if err != nil {
		// Print the error, cast err to awserr.Error to get the Code and
		// Message from an error.
		fmt.Println(err.Error())
		return
	}

	// Pretty-print the response data.
	fmt.Println(resp)
}

func ExampleSES_GetIdentityNotificationAttributes() {
	sess := session.Must(session.NewSession())

	svc := ses.New(sess)

	params := &ses.GetIdentityNotificationAttributesInput{
		Identities: []*string{ // Required
			aws.String("Identity"), // Required
			// More values...
		},
	}
	resp, err := svc.GetIdentityNotificationAttributes(params)

	if err != nil {
		// Print the error, cast err to awserr.Error to get the Code and
		// Message from an error.
		fmt.Println(err.Error())
		return
	}

	// Pretty-print the response data.
	fmt.Println(resp)
}

func ExampleSES_GetIdentityPolicies() {
	sess := session.Must(session.NewSession())

	svc := ses.New(sess)

	params := &ses.GetIdentityPoliciesInput{
		Identity: aws.String("Identity"), // Required
		PolicyNames: []*string{ // Required
			aws.String("PolicyName"), // Required
			// More values...
		},
	}
	resp, err := svc.GetIdentityPolicies(params)

	if err != nil {
		// Print the error, cast err to awserr.Error to get the Code and
		// Message from an error.
		fmt.Println(err.Error())
		return
	}

	// Pretty-print the response data.
	fmt.Println(resp)
}

func ExampleSES_GetIdentityVerificationAttributes() {
	sess := session.Must(session.NewSession())

	svc := ses.New(sess)

	params := &ses.GetIdentityVerificationAttributesInput{
		Identities: []*string{ // Required
			aws.String("Identity"), // Required
			// More values...
		},
	}
	resp, err := svc.GetIdentityVerificationAttributes(params)

	if err != nil {
		// Print the error, cast err to awserr.Error to get the Code and
		// Message from an error.
		fmt.Println(err.Error())
		return
	}

	// Pretty-print the response data.
	fmt.Println(resp)
}

func ExampleSES_GetSendQuota() {
	sess := session.Must(session.NewSession())

	svc := ses.New(sess)

	var params *ses.GetSendQuotaInput
	resp, err := svc.GetSendQuota(params)

	if err != nil {
		// Print the error, cast err to awserr.Error to get the Code and
		// Message from an error.
		fmt.Println(err.Error())
		return
	}

	// Pretty-print the response data.
	fmt.Println(resp)
}

func ExampleSES_GetSendStatistics() {
	sess := session.Must(session.NewSession())

	svc := ses.New(sess)

	var params *ses.GetSendStatisticsInput
	resp, err := svc.GetSendStatistics(params)

	if err != nil {
		// Print the error, cast err to awserr.Error to get the Code and
		// Message from an error.
		fmt.Println(err.Error())
		return
	}

	// Pretty-print the response data.
	fmt.Println(resp)
}

func ExampleSES_ListConfigurationSets() {
	sess := session.Must(session.NewSession())

	svc := ses.New(sess)

	params := &ses.ListConfigurationSetsInput{
		MaxItems:  aws.Int64(1),
		NextToken: aws.String("NextToken"),
	}
	resp, err := svc.ListConfigurationSets(params)

	if err != nil {
		// Print the error, cast err to awserr.Error to get the Code and
		// Message from an error.
		fmt.Println(err.Error())
		return
	}

	// Pretty-print the response data.
	fmt.Println(resp)
}

func ExampleSES_ListIdentities() {
	sess := session.Must(session.NewSession())

	svc := ses.New(sess)

	params := &ses.ListIdentitiesInput{
		IdentityType: aws.String("IdentityType"),
		MaxItems:     aws.Int64(1),
		NextToken:    aws.String("NextToken"),
	}
	resp, err := svc.ListIdentities(params)

	if err != nil {
		// Print the error, cast err to awserr.Error to get the Code and
		// Message from an error.
		fmt.Println(err.Error())
		return
	}

	// Pretty-print the response data.
	fmt.Println(resp)
}

func ExampleSES_ListIdentityPolicies() {
	sess := session.Must(session.NewSession())

	svc := ses.New(sess)

	params := &ses.ListIdentityPoliciesInput{
		Identity: aws.String("Identity"), // Required
	}
	resp, err := svc.ListIdentityPolicies(params)

	if err != nil {
		// Print the error, cast err to awserr.Error to get the Code and
		// Message from an error.
		fmt.Println(err.Error())
		return
	}

	// Pretty-print the response data.
	fmt.Println(resp)
}

func ExampleSES_ListReceiptFilters() {
	sess := session.Must(session.NewSession())

	svc := ses.New(sess)

	var params *ses.ListReceiptFiltersInput
	resp, err := svc.ListReceiptFilters(params)

	if err != nil {
		// Print the error, cast err to awserr.Error to get the Code and
		// Message from an error.
		fmt.Println(err.Error())
		return
	}

	// Pretty-print the response data.
	fmt.Println(resp)
}

func ExampleSES_ListReceiptRuleSets() {
	sess := session.Must(session.NewSession())

	svc := ses.New(sess)

	params := &ses.ListReceiptRuleSetsInput{
		NextToken: aws.String("NextToken"),
	}
	resp, err := svc.ListReceiptRuleSets(params)

	if err != nil {
		// Print the error, cast err to awserr.Error to get the Code and
		// Message from an error.
		fmt.Println(err.Error())
		return
	}

	// Pretty-print the response data.
	fmt.Println(resp)
}

func ExampleSES_ListVerifiedEmailAddresses() {
	sess := session.Must(session.NewSession())

	svc := ses.New(sess)

	var params *ses.ListVerifiedEmailAddressesInput
	resp, err := svc.ListVerifiedEmailAddresses(params)

	if err != nil {
		// Print the error, cast err to awserr.Error to get the Code and
		// Message from an error.
		fmt.Println(err.Error())
		return
	}

	// Pretty-print the response data.
	fmt.Println(resp)
}

func ExampleSES_PutIdentityPolicy() {
	sess := session.Must(session.NewSession())

	svc := ses.New(sess)

	params := &ses.PutIdentityPolicyInput{
		Identity:   aws.String("Identity"),   // Required
		Policy:     aws.String("Policy"),     // Required
		PolicyName: aws.String("PolicyName"), // Required
	}
	resp, err := svc.PutIdentityPolicy(params)

	if err != nil {
		// Print the error, cast err to awserr.Error to get the Code and
		// Message from an error.
		fmt.Println(err.Error())
		return
	}

	// Pretty-print the response data.
	fmt.Println(resp)
}

func ExampleSES_ReorderReceiptRuleSet() {
	sess := session.Must(session.NewSession())

	svc := ses.New(sess)

	params := &ses.ReorderReceiptRuleSetInput{
		RuleNames: []*string{ // Required
			aws.String("ReceiptRuleName"), // Required
			// More values...
		},
		RuleSetName: aws.String("ReceiptRuleSetName"), // Required
	}
	resp, err := svc.ReorderReceiptRuleSet(params)

	if err != nil {
		// Print the error, cast err to awserr.Error to get the Code and
		// Message from an error.
		fmt.Println(err.Error())
		return
	}

	// Pretty-print the response data.
	fmt.Println(resp)
}

func ExampleSES_SendBounce() {
	sess := session.Must(session.NewSession())

	svc := ses.New(sess)

	params := &ses.SendBounceInput{
		BounceSender: aws.String("Address"), // Required
		BouncedRecipientInfoList: []*ses.BouncedRecipientInfo{ // Required
			{ // Required
				Recipient:    aws.String("Address"), // Required
				BounceType:   aws.String("BounceType"),
				RecipientArn: aws.String("AmazonResourceName"),
				RecipientDsnFields: &ses.RecipientDsnFields{
					Action:         aws.String("DsnAction"), // Required
					Status:         aws.String("DsnStatus"), // Required
					DiagnosticCode: aws.String("DiagnosticCode"),
					ExtensionFields: []*ses.ExtensionField{
						{ // Required
							Name:  aws.String("ExtensionFieldName"),  // Required
							Value: aws.String("ExtensionFieldValue"), // Required
						},
						// More values...
					},
					FinalRecipient:  aws.String("Address"),
					LastAttemptDate: aws.Time(time.Now()),
					RemoteMta:       aws.String("RemoteMta"),
				},
			},
			// More values...
		},
		OriginalMessageId: aws.String("MessageId"), // Required
		BounceSenderArn:   aws.String("AmazonResourceName"),
		Explanation:       aws.String("Explanation"),
		MessageDsn: &ses.MessageDsn{
			ReportingMta: aws.String("ReportingMta"), // Required
			ArrivalDate:  aws.Time(time.Now()),
			ExtensionFields: []*ses.ExtensionField{
				{ // Required
					Name:  aws.String("ExtensionFieldName"),  // Required
					Value: aws.String("ExtensionFieldValue"), // Required
				},
				// More values...
			},
		},
	}
	resp, err := svc.SendBounce(params)

	if err != nil {
		// Print the error, cast err to awserr.Error to get the Code and
		// Message from an error.
		fmt.Println(err.Error())
		return
	}

	// Pretty-print the response data.
	fmt.Println(resp)
}

func ExampleSES_SendEmail() {
	sess := session.Must(session.NewSession())

	svc := ses.New(sess)

	params := &ses.SendEmailInput{
		Destination: &ses.Destination{ // Required
			BccAddresses: []*string{
				aws.String("Address"), // Required
				// More values...
			},
			CcAddresses: []*string{
				aws.String("Address"), // Required
				// More values...
			},
			ToAddresses: []*string{
				aws.String("Address"), // Required
				// More values...
			},
		},
		Message: &ses.Message{ // Required
			Body: &ses.Body{ // Required
				Html: &ses.Content{
					Data:    aws.String("MessageData"), // Required
					Charset: aws.String("Charset"),
				},
				Text: &ses.Content{
					Data:    aws.String("MessageData"), // Required
					Charset: aws.String("Charset"),
				},
			},
			Subject: &ses.Content{ // Required
				Data:    aws.String("MessageData"), // Required
				Charset: aws.String("Charset"),
			},
		},
		Source:               aws.String("Address"), // Required
		ConfigurationSetName: aws.String("ConfigurationSetName"),
		ReplyToAddresses: []*string{
			aws.String("Address"), // Required
			// More values...
		},
		ReturnPath:    aws.String("Address"),
		ReturnPathArn: aws.String("AmazonResourceName"),
		SourceArn:     aws.String("AmazonResourceName"),
		Tags: []*ses.MessageTag{
			{ // Required
				Name:  aws.String("MessageTagName"),  // Required
				Value: aws.String("MessageTagValue"), // Required
			},
			// More values...
		},
	}
	resp, err := svc.SendEmail(params)

	if err != nil {
		// Print the error, cast err to awserr.Error to get the Code and
		// Message from an error.
		fmt.Println(err.Error())
		return
	}

	// Pretty-print the response data.
	fmt.Println(resp)
}

func ExampleSES_SendRawEmail() {
	sess := session.Must(session.NewSession())

	svc := ses.New(sess)

	params := &ses.SendRawEmailInput{
		RawMessage: &ses.RawMessage{ // Required
			Data: []byte("PAYLOAD"), // Required
		},
		ConfigurationSetName: aws.String("ConfigurationSetName"),
		Destinations: []*string{
			aws.String("Address"), // Required
			// More values...
		},
		FromArn:       aws.String("AmazonResourceName"),
		ReturnPathArn: aws.String("AmazonResourceName"),
		Source:        aws.String("Address"),
		SourceArn:     aws.String("AmazonResourceName"),
		Tags: []*ses.MessageTag{
			{ // Required
				Name:  aws.String("MessageTagName"),  // Required
				Value: aws.String("MessageTagValue"), // Required
			},
			// More values...
		},
	}
	resp, err := svc.SendRawEmail(params)

	if err != nil {
		// Print the error, cast err to awserr.Error to get the Code and
		// Message from an error.
		fmt.Println(err.Error())
		return
	}

	// Pretty-print the response data.
	fmt.Println(resp)
}

func ExampleSES_SetActiveReceiptRuleSet() {
	sess := session.Must(session.NewSession())

	svc := ses.New(sess)

	params := &ses.SetActiveReceiptRuleSetInput{
		RuleSetName: aws.String("ReceiptRuleSetName"),
	}
	resp, err := svc.SetActiveReceiptRuleSet(params)

	if err != nil {
		// Print the error, cast err to awserr.Error to get the Code and
		// Message from an error.
		fmt.Println(err.Error())
		return
	}

	// Pretty-print the response data.
	fmt.Println(resp)
}

func ExampleSES_SetIdentityDkimEnabled() {
	sess := session.Must(session.NewSession())

	svc := ses.New(sess)

	params := &ses.SetIdentityDkimEnabledInput{
		DkimEnabled: aws.Bool(true),         // Required
		Identity:    aws.String("Identity"), // Required
	}
	resp, err := svc.SetIdentityDkimEnabled(params)

	if err != nil {
		// Print the error, cast err to awserr.Error to get the Code and
		// Message from an error.
		fmt.Println(err.Error())
		return
	}

	// Pretty-print the response data.
	fmt.Println(resp)
}

func ExampleSES_SetIdentityFeedbackForwardingEnabled() {
	sess := session.Must(session.NewSession())

	svc := ses.New(sess)

	params := &ses.SetIdentityFeedbackForwardingEnabledInput{
		ForwardingEnabled: aws.Bool(true),         // Required
		Identity:          aws.String("Identity"), // Required
	}
	resp, err := svc.SetIdentityFeedbackForwardingEnabled(params)

	if err != nil {
		// Print the error, cast err to awserr.Error to get the Code and
		// Message from an error.
		fmt.Println(err.Error())
		return
	}

	// Pretty-print the response data.
	fmt.Println(resp)
}

func ExampleSES_SetIdentityHeadersInNotificationsEnabled() {
	sess := session.Must(session.NewSession())

	svc := ses.New(sess)

	params := &ses.SetIdentityHeadersInNotificationsEnabledInput{
		Enabled:          aws.Bool(true),                 // Required
		Identity:         aws.String("Identity"),         // Required
		NotificationType: aws.String("NotificationType"), // Required
	}
	resp, err := svc.SetIdentityHeadersInNotificationsEnabled(params)

	if err != nil {
		// Print the error, cast err to awserr.Error to get the Code and
		// Message from an error.
		fmt.Println(err.Error())
		return
	}

	// Pretty-print the response data.
	fmt.Println(resp)
}

func ExampleSES_SetIdentityMailFromDomain() {
	sess := session.Must(session.NewSession())

	svc := ses.New(sess)

	params := &ses.SetIdentityMailFromDomainInput{
		Identity:            aws.String("Identity"), // Required
		BehaviorOnMXFailure: aws.String("BehaviorOnMXFailure"),
		MailFromDomain:      aws.String("MailFromDomainName"),
	}
	resp, err := svc.SetIdentityMailFromDomain(params)

	if err != nil {
		// Print the error, cast err to awserr.Error to get the Code and
		// Message from an error.
		fmt.Println(err.Error())
		return
	}

	// Pretty-print the response data.
	fmt.Println(resp)
}

func ExampleSES_SetIdentityNotificationTopic() {
	sess := session.Must(session.NewSession())

	svc := ses.New(sess)

	params := &ses.SetIdentityNotificationTopicInput{
		Identity:         aws.String("Identity"),         // Required
		NotificationType: aws.String("NotificationType"), // Required
		SnsTopic:         aws.String("NotificationTopic"),
	}
	resp, err := svc.SetIdentityNotificationTopic(params)

	if err != nil {
		// Print the error, cast err to awserr.Error to get the Code and
		// Message from an error.
		fmt.Println(err.Error())
		return
	}

	// Pretty-print the response data.
	fmt.Println(resp)
}

func ExampleSES_SetReceiptRulePosition() {
	sess := session.Must(session.NewSession())

	svc := ses.New(sess)

	params := &ses.SetReceiptRulePositionInput{
		RuleName:    aws.String("ReceiptRuleName"),    // Required
		RuleSetName: aws.String("ReceiptRuleSetName"), // Required
		After:       aws.String("ReceiptRuleName"),
	}
	resp, err := svc.SetReceiptRulePosition(params)

	if err != nil {
		// Print the error, cast err to awserr.Error to get the Code and
		// Message from an error.
		fmt.Println(err.Error())
		return
	}

	// Pretty-print the response data.
	fmt.Println(resp)
}

func ExampleSES_UpdateConfigurationSetEventDestination() {
	sess := session.Must(session.NewSession())

	svc := ses.New(sess)

	params := &ses.UpdateConfigurationSetEventDestinationInput{
		ConfigurationSetName: aws.String("ConfigurationSetName"), // Required
		EventDestination: &ses.EventDestination{ // Required
			MatchingEventTypes: []*string{ // Required
				aws.String("EventType"), // Required
				// More values...
			},
			Name: aws.String("EventDestinationName"), // Required
			CloudWatchDestination: &ses.CloudWatchDestination{
				DimensionConfigurations: []*ses.CloudWatchDimensionConfiguration{ // Required
					{ // Required
						DefaultDimensionValue: aws.String("DefaultDimensionValue"), // Required
						DimensionName:         aws.String("DimensionName"),         // Required
						DimensionValueSource:  aws.String("DimensionValueSource"),  // Required
					},
					// More values...
				},
			},
			Enabled: aws.Bool(true),
			KinesisFirehoseDestination: &ses.KinesisFirehoseDestination{
				DeliveryStreamARN: aws.String("AmazonResourceName"), // Required
				IAMRoleARN:        aws.String("AmazonResourceName"), // Required
			},
		},
	}
	resp, err := svc.UpdateConfigurationSetEventDestination(params)

	if err != nil {
		// Print the error, cast err to awserr.Error to get the Code and
		// Message from an error.
		fmt.Println(err.Error())
		return
	}

	// Pretty-print the response data.
	fmt.Println(resp)
}

func ExampleSES_UpdateReceiptRule() {
	sess := session.Must(session.NewSession())

	svc := ses.New(sess)

	params := &ses.UpdateReceiptRuleInput{
		Rule: &ses.ReceiptRule{ // Required
			Name: aws.String("ReceiptRuleName"), // Required
			Actions: []*ses.ReceiptAction{
				{ // Required
					AddHeaderAction: &ses.AddHeaderAction{
						HeaderName:  aws.String("HeaderName"),  // Required
						HeaderValue: aws.String("HeaderValue"), // Required
					},
					BounceAction: &ses.BounceAction{
						Message:       aws.String("BounceMessage"),       // Required
						Sender:        aws.String("Address"),             // Required
						SmtpReplyCode: aws.String("BounceSmtpReplyCode"), // Required
						StatusCode:    aws.String("BounceStatusCode"),
						TopicArn:      aws.String("AmazonResourceName"),
					},
					LambdaAction: &ses.LambdaAction{
						FunctionArn:    aws.String("AmazonResourceName"), // Required
						InvocationType: aws.String("InvocationType"),
						TopicArn:       aws.String("AmazonResourceName"),
					},
					S3Action: &ses.S3Action{
						BucketName:      aws.String("S3BucketName"), // Required
						KmsKeyArn:       aws.String("AmazonResourceName"),
						ObjectKeyPrefix: aws.String("S3KeyPrefix"),
						TopicArn:        aws.String("AmazonResourceName"),
					},
					SNSAction: &ses.SNSAction{
						TopicArn: aws.String("AmazonResourceName"), // Required
						Encoding: aws.String("SNSActionEncoding"),
					},
					StopAction: &ses.StopAction{
						Scope:    aws.String("StopScope"), // Required
						TopicArn: aws.String("AmazonResourceName"),
					},
					WorkmailAction: &ses.WorkmailAction{
						OrganizationArn: aws.String("AmazonResourceName"), // Required
						TopicArn:        aws.String("AmazonResourceName"),
					},
				},
				// More values...
			},
			Enabled: aws.Bool(true),
			Recipients: []*string{
				aws.String("Recipient"), // Required
				// More values...
			},
			ScanEnabled: aws.Bool(true),
			TlsPolicy:   aws.String("TlsPolicy"),
		},
		RuleSetName: aws.String("ReceiptRuleSetName"), // Required
	}
	resp, err := svc.UpdateReceiptRule(params)

	if err != nil {
		// Print the error, cast err to awserr.Error to get the Code and
		// Message from an error.
		fmt.Println(err.Error())
		return
	}

	// Pretty-print the response data.
	fmt.Println(resp)
}

func ExampleSES_VerifyDomainDkim() {
	sess := session.Must(session.NewSession())

	svc := ses.New(sess)

	params := &ses.VerifyDomainDkimInput{
		Domain: aws.String("Domain"), // Required
	}
	resp, err := svc.VerifyDomainDkim(params)

	if err != nil {
		// Print the error, cast err to awserr.Error to get the Code and
		// Message from an error.
		fmt.Println(err.Error())
		return
	}

	// Pretty-print the response data.
	fmt.Println(resp)
}

func ExampleSES_VerifyDomainIdentity() {
	sess := session.Must(session.NewSession())

	svc := ses.New(sess)

	params := &ses.VerifyDomainIdentityInput{
		Domain: aws.String("Domain"), // Required
	}
	resp, err := svc.VerifyDomainIdentity(params)

	if err != nil {
		// Print the error, cast err to awserr.Error to get the Code and
		// Message from an error.
		fmt.Println(err.Error())
		return
	}

	// Pretty-print the response data.
	fmt.Println(resp)
}

func ExampleSES_VerifyEmailAddress() {
	sess := session.Must(session.NewSession())

	svc := ses.New(sess)

	params := &ses.VerifyEmailAddressInput{
		EmailAddress: aws.String("Address"), // Required
	}
	resp, err := svc.VerifyEmailAddress(params)

	if err != nil {
		// Print the error, cast err to awserr.Error to get the Code and
		// Message from an error.
		fmt.Println(err.Error())
		return
	}

	// Pretty-print the response data.
	fmt.Println(resp)
}

func ExampleSES_VerifyEmailIdentity() {
	sess := session.Must(session.NewSession())

	svc := ses.New(sess)

	params := &ses.VerifyEmailIdentityInput{
		EmailAddress: aws.String("Address"), // Required
	}
	resp, err := svc.VerifyEmailIdentity(params)

	if err != nil {
		// Print the error, cast err to awserr.Error to get the Code and
		// Message from an error.
		fmt.Println(err.Error())
		return
	}

	// Pretty-print the response data.
	fmt.Println(resp)
}
