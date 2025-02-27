// Code generated by "go generate". Please don't change this file directly.

package customerprofiles

import (
	"github.com/awslabs/goformation/v7/cloudformation/policies"
)

// EventTrigger_EventTriggerDimension AWS CloudFormation Resource (AWS::CustomerProfiles::EventTrigger.EventTriggerDimension)
// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-customerprofiles-eventtrigger-eventtriggerdimension.html
type EventTrigger_EventTriggerDimension struct {

	// ObjectAttributes AWS CloudFormation Property
	// Required: true
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-customerprofiles-eventtrigger-eventtriggerdimension.html#cfn-customerprofiles-eventtrigger-eventtriggerdimension-objectattributes
	ObjectAttributes []EventTrigger_ObjectAttribute `json:"ObjectAttributes"`

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
func (r *EventTrigger_EventTriggerDimension) AWSCloudFormationType() string {
	return "AWS::CustomerProfiles::EventTrigger.EventTriggerDimension"
}
