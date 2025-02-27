// Code generated by "go generate". Please don't change this file directly.

package bedrock

import (
	"github.com/awslabs/goformation/v7/cloudformation/policies"
)

// Prompt_TextPromptTemplateConfiguration AWS CloudFormation Resource (AWS::Bedrock::Prompt.TextPromptTemplateConfiguration)
// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-bedrock-prompt-textprompttemplateconfiguration.html
type Prompt_TextPromptTemplateConfiguration struct {

	// CachePoint AWS CloudFormation Property
	// Required: false
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-bedrock-prompt-textprompttemplateconfiguration.html#cfn-bedrock-prompt-textprompttemplateconfiguration-cachepoint
	CachePoint *Prompt_CachePointBlock `json:"CachePoint,omitempty"`

	// InputVariables AWS CloudFormation Property
	// Required: false
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-bedrock-prompt-textprompttemplateconfiguration.html#cfn-bedrock-prompt-textprompttemplateconfiguration-inputvariables
	InputVariables []Prompt_PromptInputVariable `json:"InputVariables,omitempty"`

	// Text AWS CloudFormation Property
	// Required: false
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-bedrock-prompt-textprompttemplateconfiguration.html#cfn-bedrock-prompt-textprompttemplateconfiguration-text
	Text *string `json:"Text,omitempty"`

	// TextS3Location AWS CloudFormation Property
	// Required: false
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-bedrock-prompt-textprompttemplateconfiguration.html#cfn-bedrock-prompt-textprompttemplateconfiguration-texts3location
	TextS3Location *Prompt_TextS3Location `json:"TextS3Location,omitempty"`

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
func (r *Prompt_TextPromptTemplateConfiguration) AWSCloudFormationType() string {
	return "AWS::Bedrock::Prompt.TextPromptTemplateConfiguration"
}
