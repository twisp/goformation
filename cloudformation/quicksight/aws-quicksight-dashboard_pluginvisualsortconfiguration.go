// Code generated by "go generate". Please don't change this file directly.

package quicksight

import (
	"github.com/awslabs/goformation/v7/cloudformation/policies"
)

// Dashboard_PluginVisualSortConfiguration AWS CloudFormation Resource (AWS::QuickSight::Dashboard.PluginVisualSortConfiguration)
// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-quicksight-dashboard-pluginvisualsortconfiguration.html
type Dashboard_PluginVisualSortConfiguration struct {

	// PluginVisualTableQuerySort AWS CloudFormation Property
	// Required: false
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-quicksight-dashboard-pluginvisualsortconfiguration.html#cfn-quicksight-dashboard-pluginvisualsortconfiguration-pluginvisualtablequerysort
	PluginVisualTableQuerySort *Dashboard_PluginVisualTableQuerySort `json:"PluginVisualTableQuerySort,omitempty"`

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
func (r *Dashboard_PluginVisualSortConfiguration) AWSCloudFormationType() string {
	return "AWS::QuickSight::Dashboard.PluginVisualSortConfiguration"
}
