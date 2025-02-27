// Code generated by "go generate". Please don't change this file directly.

package quicksight

import (
	"github.com/awslabs/goformation/v7/cloudformation/policies"
)

// Analysis_GeospatialCategoricalColor AWS CloudFormation Resource (AWS::QuickSight::Analysis.GeospatialCategoricalColor)
// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-quicksight-analysis-geospatialcategoricalcolor.html
type Analysis_GeospatialCategoricalColor struct {

	// CategoryDataColors AWS CloudFormation Property
	// Required: true
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-quicksight-analysis-geospatialcategoricalcolor.html#cfn-quicksight-analysis-geospatialcategoricalcolor-categorydatacolors
	CategoryDataColors []Analysis_GeospatialCategoricalDataColor `json:"CategoryDataColors"`

	// DefaultOpacity AWS CloudFormation Property
	// Required: false
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-quicksight-analysis-geospatialcategoricalcolor.html#cfn-quicksight-analysis-geospatialcategoricalcolor-defaultopacity
	DefaultOpacity *float64 `json:"DefaultOpacity,omitempty"`

	// NullDataSettings AWS CloudFormation Property
	// Required: false
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-quicksight-analysis-geospatialcategoricalcolor.html#cfn-quicksight-analysis-geospatialcategoricalcolor-nulldatasettings
	NullDataSettings *Analysis_GeospatialNullDataSettings `json:"NullDataSettings,omitempty"`

	// NullDataVisibility AWS CloudFormation Property
	// Required: false
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-quicksight-analysis-geospatialcategoricalcolor.html#cfn-quicksight-analysis-geospatialcategoricalcolor-nulldatavisibility
	NullDataVisibility *string `json:"NullDataVisibility,omitempty"`

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
func (r *Analysis_GeospatialCategoricalColor) AWSCloudFormationType() string {
	return "AWS::QuickSight::Analysis.GeospatialCategoricalColor"
}
