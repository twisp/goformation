// Code generated by "go generate". Please don't change this file directly.

package customerprofiles

import (
	"github.com/awslabs/goformation/v7/cloudformation/policies"
)

// SegmentDefinition_ProfileDimension AWS CloudFormation Resource (AWS::CustomerProfiles::SegmentDefinition.ProfileDimension)
// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-customerprofiles-segmentdefinition-profiledimension.html
type SegmentDefinition_ProfileDimension struct {

	// DimensionType AWS CloudFormation Property
	// Required: true
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-customerprofiles-segmentdefinition-profiledimension.html#cfn-customerprofiles-segmentdefinition-profiledimension-dimensiontype
	DimensionType string `json:"DimensionType"`

	// Values AWS CloudFormation Property
	// Required: true
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-customerprofiles-segmentdefinition-profiledimension.html#cfn-customerprofiles-segmentdefinition-profiledimension-values
	Values []string `json:"Values"`

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
func (r *SegmentDefinition_ProfileDimension) AWSCloudFormationType() string {
	return "AWS::CustomerProfiles::SegmentDefinition.ProfileDimension"
}
