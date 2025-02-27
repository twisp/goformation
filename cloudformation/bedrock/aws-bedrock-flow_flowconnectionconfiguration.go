// Code generated by "go generate". Please don't change this file directly.

package bedrock

import (
	"github.com/awslabs/goformation/v7/cloudformation/policies"
)

// Flow_FlowConnectionConfiguration AWS CloudFormation Resource (AWS::Bedrock::Flow.FlowConnectionConfiguration)
// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-bedrock-flow-flowconnectionconfiguration.html
type Flow_FlowConnectionConfiguration struct {

	// Conditional AWS CloudFormation Property
	// Required: false
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-bedrock-flow-flowconnectionconfiguration.html#cfn-bedrock-flow-flowconnectionconfiguration-conditional
	Conditional *Flow_FlowConditionalConnectionConfiguration `json:"Conditional,omitempty"`

	// Data AWS CloudFormation Property
	// Required: false
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-bedrock-flow-flowconnectionconfiguration.html#cfn-bedrock-flow-flowconnectionconfiguration-data
	Data *Flow_FlowDataConnectionConfiguration `json:"Data,omitempty"`

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
func (r *Flow_FlowConnectionConfiguration) AWSCloudFormationType() string {
	return "AWS::Bedrock::Flow.FlowConnectionConfiguration"
}
