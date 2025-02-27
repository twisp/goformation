// Code generated by "go generate". Please don't change this file directly.

package logs

import (
	"github.com/awslabs/goformation/v7/cloudformation/policies"
)

// Transformer_SubstituteStringEntry AWS CloudFormation Resource (AWS::Logs::Transformer.SubstituteStringEntry)
// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-logs-transformer-substitutestringentry.html
type Transformer_SubstituteStringEntry struct {

	// From AWS CloudFormation Property
	// Required: true
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-logs-transformer-substitutestringentry.html#cfn-logs-transformer-substitutestringentry-from
	From string `json:"From"`

	// Source AWS CloudFormation Property
	// Required: true
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-logs-transformer-substitutestringentry.html#cfn-logs-transformer-substitutestringentry-source
	Source string `json:"Source"`

	// To AWS CloudFormation Property
	// Required: true
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-logs-transformer-substitutestringentry.html#cfn-logs-transformer-substitutestringentry-to
	To string `json:"To"`

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
func (r *Transformer_SubstituteStringEntry) AWSCloudFormationType() string {
	return "AWS::Logs::Transformer.SubstituteStringEntry"
}
