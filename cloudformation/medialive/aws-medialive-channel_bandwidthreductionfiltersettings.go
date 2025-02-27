// Code generated by "go generate". Please don't change this file directly.

package medialive

import (
	"github.com/awslabs/goformation/v7/cloudformation/policies"
)

// Channel_BandwidthReductionFilterSettings AWS CloudFormation Resource (AWS::MediaLive::Channel.BandwidthReductionFilterSettings)
// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-medialive-channel-bandwidthreductionfiltersettings.html
type Channel_BandwidthReductionFilterSettings struct {

	// PostFilterSharpening AWS CloudFormation Property
	// Required: false
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-medialive-channel-bandwidthreductionfiltersettings.html#cfn-medialive-channel-bandwidthreductionfiltersettings-postfiltersharpening
	PostFilterSharpening *string `json:"PostFilterSharpening,omitempty"`

	// Strength AWS CloudFormation Property
	// Required: false
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-medialive-channel-bandwidthreductionfiltersettings.html#cfn-medialive-channel-bandwidthreductionfiltersettings-strength
	Strength *string `json:"Strength,omitempty"`

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
func (r *Channel_BandwidthReductionFilterSettings) AWSCloudFormationType() string {
	return "AWS::MediaLive::Channel.BandwidthReductionFilterSettings"
}
