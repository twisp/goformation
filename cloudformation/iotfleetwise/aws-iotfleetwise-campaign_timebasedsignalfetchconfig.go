// Code generated by "go generate". Please don't change this file directly.

package iotfleetwise

import (
	"github.com/awslabs/goformation/v7/cloudformation/policies"
)

// Campaign_TimeBasedSignalFetchConfig AWS CloudFormation Resource (AWS::IoTFleetWise::Campaign.TimeBasedSignalFetchConfig)
// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-iotfleetwise-campaign-timebasedsignalfetchconfig.html
type Campaign_TimeBasedSignalFetchConfig struct {

	// ExecutionFrequencyMs AWS CloudFormation Property
	// Required: true
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-iotfleetwise-campaign-timebasedsignalfetchconfig.html#cfn-iotfleetwise-campaign-timebasedsignalfetchconfig-executionfrequencyms
	ExecutionFrequencyMs float64 `json:"ExecutionFrequencyMs"`

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
func (r *Campaign_TimeBasedSignalFetchConfig) AWSCloudFormationType() string {
	return "AWS::IoTFleetWise::Campaign.TimeBasedSignalFetchConfig"
}
