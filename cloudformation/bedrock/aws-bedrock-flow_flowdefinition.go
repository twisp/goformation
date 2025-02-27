// Code generated by "go generate". Please don't change this file directly.

package bedrock

import (
	"github.com/awslabs/goformation/v7/cloudformation/policies"
)

// Flow_FlowDefinition AWS CloudFormation Resource (AWS::Bedrock::Flow.FlowDefinition)
// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-bedrock-flow-flowdefinition.html
type Flow_FlowDefinition struct {

	// Connections AWS CloudFormation Property
	// Required: false
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-bedrock-flow-flowdefinition.html#cfn-bedrock-flow-flowdefinition-connections
	Connections []Flow_FlowConnection `json:"Connections,omitempty"`

	// Nodes AWS CloudFormation Property
	// Required: false
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-bedrock-flow-flowdefinition.html#cfn-bedrock-flow-flowdefinition-nodes
	Nodes []Flow_FlowNode `json:"Nodes,omitempty"`

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
func (r *Flow_FlowDefinition) AWSCloudFormationType() string {
	return "AWS::Bedrock::Flow.FlowDefinition"
}
