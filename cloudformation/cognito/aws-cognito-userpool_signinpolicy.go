// Code generated by "go generate". Please don't change this file directly.

package cognito

import (
	"github.com/awslabs/goformation/v7/cloudformation/policies"
)

// UserPool_SignInPolicy AWS CloudFormation Resource (AWS::Cognito::UserPool.SignInPolicy)
// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-cognito-userpool-signinpolicy.html
type UserPool_SignInPolicy struct {

	// AllowedFirstAuthFactors AWS CloudFormation Property
	// Required: false
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-cognito-userpool-signinpolicy.html#cfn-cognito-userpool-signinpolicy-allowedfirstauthfactors
	AllowedFirstAuthFactors []string `json:"AllowedFirstAuthFactors,omitempty"`

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
func (r *UserPool_SignInPolicy) AWSCloudFormationType() string {
	return "AWS::Cognito::UserPool.SignInPolicy"
}
