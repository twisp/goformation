// Code generated by "go generate". Please don't change this file directly.

package qbusiness

import (
	"github.com/awslabs/goformation/v7/cloudformation/policies"
)

// DataSource_MediaExtractionConfiguration AWS CloudFormation Resource (AWS::QBusiness::DataSource.MediaExtractionConfiguration)
// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-qbusiness-datasource-mediaextractionconfiguration.html
type DataSource_MediaExtractionConfiguration struct {

	// ImageExtractionConfiguration AWS CloudFormation Property
	// Required: false
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-qbusiness-datasource-mediaextractionconfiguration.html#cfn-qbusiness-datasource-mediaextractionconfiguration-imageextractionconfiguration
	ImageExtractionConfiguration *DataSource_ImageExtractionConfiguration `json:"ImageExtractionConfiguration,omitempty"`

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
func (r *DataSource_MediaExtractionConfiguration) AWSCloudFormationType() string {
	return "AWS::QBusiness::DataSource.MediaExtractionConfiguration"
}
