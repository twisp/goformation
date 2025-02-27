// Code generated by "go generate". Please don't change this file directly.

package quicksight

import (
	"github.com/awslabs/goformation/v7/cloudformation/policies"
)

// Dashboard_WaterfallChartGroupColorConfiguration AWS CloudFormation Resource (AWS::QuickSight::Dashboard.WaterfallChartGroupColorConfiguration)
// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-quicksight-dashboard-waterfallchartgroupcolorconfiguration.html
type Dashboard_WaterfallChartGroupColorConfiguration struct {

	// NegativeBarColor AWS CloudFormation Property
	// Required: false
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-quicksight-dashboard-waterfallchartgroupcolorconfiguration.html#cfn-quicksight-dashboard-waterfallchartgroupcolorconfiguration-negativebarcolor
	NegativeBarColor *string `json:"NegativeBarColor,omitempty"`

	// PositiveBarColor AWS CloudFormation Property
	// Required: false
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-quicksight-dashboard-waterfallchartgroupcolorconfiguration.html#cfn-quicksight-dashboard-waterfallchartgroupcolorconfiguration-positivebarcolor
	PositiveBarColor *string `json:"PositiveBarColor,omitempty"`

	// TotalBarColor AWS CloudFormation Property
	// Required: false
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-quicksight-dashboard-waterfallchartgroupcolorconfiguration.html#cfn-quicksight-dashboard-waterfallchartgroupcolorconfiguration-totalbarcolor
	TotalBarColor *string `json:"TotalBarColor,omitempty"`

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
func (r *Dashboard_WaterfallChartGroupColorConfiguration) AWSCloudFormationType() string {
	return "AWS::QuickSight::Dashboard.WaterfallChartGroupColorConfiguration"
}
