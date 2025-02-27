// Code generated by "go generate". Please don't change this file directly.

package kinesisfirehose

import (
	"github.com/awslabs/goformation/v7/cloudformation/policies"
)

// DeliveryStream_DestinationTableConfiguration AWS CloudFormation Resource (AWS::KinesisFirehose::DeliveryStream.DestinationTableConfiguration)
// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-kinesisfirehose-deliverystream-destinationtableconfiguration.html
type DeliveryStream_DestinationTableConfiguration struct {

	// DestinationDatabaseName AWS CloudFormation Property
	// Required: true
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-kinesisfirehose-deliverystream-destinationtableconfiguration.html#cfn-kinesisfirehose-deliverystream-destinationtableconfiguration-destinationdatabasename
	DestinationDatabaseName string `json:"DestinationDatabaseName"`

	// DestinationTableName AWS CloudFormation Property
	// Required: true
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-kinesisfirehose-deliverystream-destinationtableconfiguration.html#cfn-kinesisfirehose-deliverystream-destinationtableconfiguration-destinationtablename
	DestinationTableName string `json:"DestinationTableName"`

	// S3ErrorOutputPrefix AWS CloudFormation Property
	// Required: false
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-kinesisfirehose-deliverystream-destinationtableconfiguration.html#cfn-kinesisfirehose-deliverystream-destinationtableconfiguration-s3erroroutputprefix
	S3ErrorOutputPrefix *string `json:"S3ErrorOutputPrefix,omitempty"`

	// UniqueKeys AWS CloudFormation Property
	// Required: false
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-kinesisfirehose-deliverystream-destinationtableconfiguration.html#cfn-kinesisfirehose-deliverystream-destinationtableconfiguration-uniquekeys
	UniqueKeys []string `json:"UniqueKeys,omitempty"`

	// AWSCloudFormationDeletionPolicy represents a CloudFormation DeletionPolicy
	AWSCloudFormationDeletionPolicy policies.DeletionPolicy `json:"-"`

	// AWSCloudFormationUpdateReplacePolicy represents a CloudFormation UpdateReplacePolicy
	AWSCloudFormationUpdateReplacePolicy policies.UpdateReplacePolicy `json:"-"`

	// AWSCloudFormationDependsOn stores the logical ID of the resources to be created before this resource
	AWSCloudFormationDependsOn []string `json:"-"`

	// AWSCloudFormationMetadata stores structured data associated with this resource
	AWSCloudFormationMetadata map[string]interface{} `json:"-"`

	// AWSCloudFormationCondition stores the logical ID of the condition that must be satisfied for this resource to be created
	AWSCloudFormationCondition string `json:"-"`
}

// AWSCloudFormationType returns the AWS CloudFormation resource type
func (r *DeliveryStream_DestinationTableConfiguration) AWSCloudFormationType() string {
	return "AWS::KinesisFirehose::DeliveryStream.DestinationTableConfiguration"
}
