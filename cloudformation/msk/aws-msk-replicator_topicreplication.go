// Code generated by "go generate". Please don't change this file directly.

package msk

import (
	"github.com/awslabs/goformation/v7/cloudformation/policies"
)

// Replicator_TopicReplication AWS CloudFormation Resource (AWS::MSK::Replicator.TopicReplication)
// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-msk-replicator-topicreplication.html
type Replicator_TopicReplication struct {

	// CopyAccessControlListsForTopics AWS CloudFormation Property
	// Required: false
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-msk-replicator-topicreplication.html#cfn-msk-replicator-topicreplication-copyaccesscontrollistsfortopics
	CopyAccessControlListsForTopics *bool `json:"CopyAccessControlListsForTopics,omitempty"`

	// CopyTopicConfigurations AWS CloudFormation Property
	// Required: false
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-msk-replicator-topicreplication.html#cfn-msk-replicator-topicreplication-copytopicconfigurations
	CopyTopicConfigurations *bool `json:"CopyTopicConfigurations,omitempty"`

	// DetectAndCopyNewTopics AWS CloudFormation Property
	// Required: false
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-msk-replicator-topicreplication.html#cfn-msk-replicator-topicreplication-detectandcopynewtopics
	DetectAndCopyNewTopics *bool `json:"DetectAndCopyNewTopics,omitempty"`

	// StartingPosition AWS CloudFormation Property
	// Required: false
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-msk-replicator-topicreplication.html#cfn-msk-replicator-topicreplication-startingposition
	StartingPosition *Replicator_ReplicationStartingPosition `json:"StartingPosition,omitempty"`

	// TopicNameConfiguration AWS CloudFormation Property
	// Required: false
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-msk-replicator-topicreplication.html#cfn-msk-replicator-topicreplication-topicnameconfiguration
	TopicNameConfiguration *Replicator_ReplicationTopicNameConfiguration `json:"TopicNameConfiguration,omitempty"`

	// TopicsToExclude AWS CloudFormation Property
	// Required: false
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-msk-replicator-topicreplication.html#cfn-msk-replicator-topicreplication-topicstoexclude
	TopicsToExclude []string `json:"TopicsToExclude,omitempty"`

	// TopicsToReplicate AWS CloudFormation Property
	// Required: true
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-msk-replicator-topicreplication.html#cfn-msk-replicator-topicreplication-topicstoreplicate
	TopicsToReplicate []string `json:"TopicsToReplicate"`

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
func (r *Replicator_TopicReplication) AWSCloudFormationType() string {
	return "AWS::MSK::Replicator.TopicReplication"
}
