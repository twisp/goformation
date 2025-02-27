// Code generated by "go generate". Please don't change this file directly.

package quicksight

import (
	"github.com/awslabs/goformation/v7/cloudformation/policies"
)

// Analysis_GeospatialLayerDefinition AWS CloudFormation Resource (AWS::QuickSight::Analysis.GeospatialLayerDefinition)
// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-quicksight-analysis-geospatiallayerdefinition.html
type Analysis_GeospatialLayerDefinition struct {

	// LineLayer AWS CloudFormation Property
	// Required: false
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-quicksight-analysis-geospatiallayerdefinition.html#cfn-quicksight-analysis-geospatiallayerdefinition-linelayer
	LineLayer *Analysis_GeospatialLineLayer `json:"LineLayer,omitempty"`

	// PointLayer AWS CloudFormation Property
	// Required: false
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-quicksight-analysis-geospatiallayerdefinition.html#cfn-quicksight-analysis-geospatiallayerdefinition-pointlayer
	PointLayer *Analysis_GeospatialPointLayer `json:"PointLayer,omitempty"`

	// PolygonLayer AWS CloudFormation Property
	// Required: false
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-quicksight-analysis-geospatiallayerdefinition.html#cfn-quicksight-analysis-geospatiallayerdefinition-polygonlayer
	PolygonLayer *Analysis_GeospatialPolygonLayer `json:"PolygonLayer,omitempty"`

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
func (r *Analysis_GeospatialLayerDefinition) AWSCloudFormationType() string {
	return "AWS::QuickSight::Analysis.GeospatialLayerDefinition"
}
