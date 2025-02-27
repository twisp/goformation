// Code generated by "go generate". Please don't change this file directly.

package bedrock

import (
	"github.com/awslabs/goformation/v7/cloudformation/policies"
)

// DataSource_SemanticChunkingConfiguration AWS CloudFormation Resource (AWS::Bedrock::DataSource.SemanticChunkingConfiguration)
// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-bedrock-datasource-semanticchunkingconfiguration.html
type DataSource_SemanticChunkingConfiguration struct {

	// BreakpointPercentileThreshold AWS CloudFormation Property
	// Required: true
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-bedrock-datasource-semanticchunkingconfiguration.html#cfn-bedrock-datasource-semanticchunkingconfiguration-breakpointpercentilethreshold
	BreakpointPercentileThreshold int `json:"BreakpointPercentileThreshold"`

	// BufferSize AWS CloudFormation Property
	// Required: true
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-bedrock-datasource-semanticchunkingconfiguration.html#cfn-bedrock-datasource-semanticchunkingconfiguration-buffersize
	BufferSize int `json:"BufferSize"`

	// MaxTokens AWS CloudFormation Property
	// Required: true
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-bedrock-datasource-semanticchunkingconfiguration.html#cfn-bedrock-datasource-semanticchunkingconfiguration-maxtokens
	MaxTokens int `json:"MaxTokens"`

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
func (r *DataSource_SemanticChunkingConfiguration) AWSCloudFormationType() string {
	return "AWS::Bedrock::DataSource.SemanticChunkingConfiguration"
}
