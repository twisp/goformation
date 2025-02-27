// Code generated by "go generate". Please don't change this file directly.

package wisdom

import (
	"github.com/awslabs/goformation/v7/cloudformation/policies"
)

// AIAgent_KnowledgeBaseAssociationConfigurationData AWS CloudFormation Resource (AWS::Wisdom::AIAgent.KnowledgeBaseAssociationConfigurationData)
// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-wisdom-aiagent-knowledgebaseassociationconfigurationdata.html
type AIAgent_KnowledgeBaseAssociationConfigurationData struct {

	// ContentTagFilter AWS CloudFormation Property
	// Required: false
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-wisdom-aiagent-knowledgebaseassociationconfigurationdata.html#cfn-wisdom-aiagent-knowledgebaseassociationconfigurationdata-contenttagfilter
	ContentTagFilter *AIAgent_TagFilter `json:"ContentTagFilter,omitempty"`

	// MaxResults AWS CloudFormation Property
	// Required: false
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-wisdom-aiagent-knowledgebaseassociationconfigurationdata.html#cfn-wisdom-aiagent-knowledgebaseassociationconfigurationdata-maxresults
	MaxResults *float64 `json:"MaxResults,omitempty"`

	// OverrideKnowledgeBaseSearchType AWS CloudFormation Property
	// Required: false
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-wisdom-aiagent-knowledgebaseassociationconfigurationdata.html#cfn-wisdom-aiagent-knowledgebaseassociationconfigurationdata-overrideknowledgebasesearchtype
	OverrideKnowledgeBaseSearchType *string `json:"OverrideKnowledgeBaseSearchType,omitempty"`

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
func (r *AIAgent_KnowledgeBaseAssociationConfigurationData) AWSCloudFormationType() string {
	return "AWS::Wisdom::AIAgent.KnowledgeBaseAssociationConfigurationData"
}
