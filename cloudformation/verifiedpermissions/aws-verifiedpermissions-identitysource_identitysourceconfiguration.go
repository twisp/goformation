// Code generated by "go generate". Please don't change this file directly.

package verifiedpermissions

import (
	"github.com/awslabs/goformation/v7/cloudformation/policies"
)

// IdentitySource_IdentitySourceConfiguration AWS CloudFormation Resource (AWS::VerifiedPermissions::IdentitySource.IdentitySourceConfiguration)
// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-verifiedpermissions-identitysource-identitysourceconfiguration.html
type IdentitySource_IdentitySourceConfiguration struct {

	// CognitoUserPoolConfiguration AWS CloudFormation Property
	// Required: false
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-verifiedpermissions-identitysource-identitysourceconfiguration.html#cfn-verifiedpermissions-identitysource-identitysourceconfiguration-cognitouserpoolconfiguration
	CognitoUserPoolConfiguration *IdentitySource_CognitoUserPoolConfiguration `json:"CognitoUserPoolConfiguration,omitempty"`

	// OpenIdConnectConfiguration AWS CloudFormation Property
	// Required: false
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-verifiedpermissions-identitysource-identitysourceconfiguration.html#cfn-verifiedpermissions-identitysource-identitysourceconfiguration-openidconnectconfiguration
	OpenIdConnectConfiguration *IdentitySource_OpenIdConnectConfiguration `json:"OpenIdConnectConfiguration,omitempty"`

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
func (r *IdentitySource_IdentitySourceConfiguration) AWSCloudFormationType() string {
	return "AWS::VerifiedPermissions::IdentitySource.IdentitySourceConfiguration"
}
