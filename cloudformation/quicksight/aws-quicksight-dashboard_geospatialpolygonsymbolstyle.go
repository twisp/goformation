// Code generated by "go generate". Please don't change this file directly.

package quicksight

import (
	"github.com/awslabs/goformation/v7/cloudformation/policies"
)

// Dashboard_GeospatialPolygonSymbolStyle AWS CloudFormation Resource (AWS::QuickSight::Dashboard.GeospatialPolygonSymbolStyle)
// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-quicksight-dashboard-geospatialpolygonsymbolstyle.html
type Dashboard_GeospatialPolygonSymbolStyle struct {

	// FillColor AWS CloudFormation Property
	// Required: false
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-quicksight-dashboard-geospatialpolygonsymbolstyle.html#cfn-quicksight-dashboard-geospatialpolygonsymbolstyle-fillcolor
	FillColor *Dashboard_GeospatialColor `json:"FillColor,omitempty"`

	// StrokeColor AWS CloudFormation Property
	// Required: false
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-quicksight-dashboard-geospatialpolygonsymbolstyle.html#cfn-quicksight-dashboard-geospatialpolygonsymbolstyle-strokecolor
	StrokeColor *Dashboard_GeospatialColor `json:"StrokeColor,omitempty"`

	// StrokeWidth AWS CloudFormation Property
	// Required: false
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-quicksight-dashboard-geospatialpolygonsymbolstyle.html#cfn-quicksight-dashboard-geospatialpolygonsymbolstyle-strokewidth
	StrokeWidth *Dashboard_GeospatialLineWidth `json:"StrokeWidth,omitempty"`

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
func (r *Dashboard_GeospatialPolygonSymbolStyle) AWSCloudFormationType() string {
	return "AWS::QuickSight::Dashboard.GeospatialPolygonSymbolStyle"
}
