// Code generated by "go generate". Please don't change this file directly.

package opensearchservice

import (
	"github.com/awslabs/goformation/v7/cloudformation/policies"
)

// Domain_JWTOptions AWS CloudFormation Resource (AWS::OpenSearchService::Domain.JWTOptions)
// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-opensearchservice-domain-jwtoptions.html
type Domain_JWTOptions struct {

	// Enabled AWS CloudFormation Property
	// Required: false
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-opensearchservice-domain-jwtoptions.html#cfn-opensearchservice-domain-jwtoptions-enabled
	Enabled *bool `json:"Enabled,omitempty"`

	// PublicKey AWS CloudFormation Property
	// Required: false
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-opensearchservice-domain-jwtoptions.html#cfn-opensearchservice-domain-jwtoptions-publickey
	PublicKey *string `json:"PublicKey,omitempty"`

	// RolesKey AWS CloudFormation Property
	// Required: false
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-opensearchservice-domain-jwtoptions.html#cfn-opensearchservice-domain-jwtoptions-roleskey
	RolesKey *string `json:"RolesKey,omitempty"`

	// SubjectKey AWS CloudFormation Property
	// Required: false
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-opensearchservice-domain-jwtoptions.html#cfn-opensearchservice-domain-jwtoptions-subjectkey
	SubjectKey *string `json:"SubjectKey,omitempty"`

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
func (r *Domain_JWTOptions) AWSCloudFormationType() string {
	return "AWS::OpenSearchService::Domain.JWTOptions"
}
