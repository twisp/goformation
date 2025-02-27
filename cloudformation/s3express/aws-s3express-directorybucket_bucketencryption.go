// Code generated by "go generate". Please don't change this file directly.

package s3express

import (
	"github.com/awslabs/goformation/v7/cloudformation/policies"
)

// DirectoryBucket_BucketEncryption AWS CloudFormation Resource (AWS::S3Express::DirectoryBucket.BucketEncryption)
// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-s3express-directorybucket-bucketencryption.html
type DirectoryBucket_BucketEncryption struct {

	// ServerSideEncryptionConfiguration AWS CloudFormation Property
	// Required: true
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-s3express-directorybucket-bucketencryption.html#cfn-s3express-directorybucket-bucketencryption-serversideencryptionconfiguration
	ServerSideEncryptionConfiguration []DirectoryBucket_ServerSideEncryptionRule `json:"ServerSideEncryptionConfiguration"`

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
func (r *DirectoryBucket_BucketEncryption) AWSCloudFormationType() string {
	return "AWS::S3Express::DirectoryBucket.BucketEncryption"
}
