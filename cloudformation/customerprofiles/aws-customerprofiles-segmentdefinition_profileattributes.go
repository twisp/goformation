// Code generated by "go generate". Please don't change this file directly.

package customerprofiles

import (
	"github.com/awslabs/goformation/v7/cloudformation/policies"
)

// SegmentDefinition_ProfileAttributes AWS CloudFormation Resource (AWS::CustomerProfiles::SegmentDefinition.ProfileAttributes)
// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-customerprofiles-segmentdefinition-profileattributes.html
type SegmentDefinition_ProfileAttributes struct {

	// AccountNumber AWS CloudFormation Property
	// Required: false
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-customerprofiles-segmentdefinition-profileattributes.html#cfn-customerprofiles-segmentdefinition-profileattributes-accountnumber
	AccountNumber *SegmentDefinition_ProfileDimension `json:"AccountNumber,omitempty"`

	// AdditionalInformation AWS CloudFormation Property
	// Required: false
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-customerprofiles-segmentdefinition-profileattributes.html#cfn-customerprofiles-segmentdefinition-profileattributes-additionalinformation
	AdditionalInformation *SegmentDefinition_ExtraLengthValueProfileDimension `json:"AdditionalInformation,omitempty"`

	// Address AWS CloudFormation Property
	// Required: false
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-customerprofiles-segmentdefinition-profileattributes.html#cfn-customerprofiles-segmentdefinition-profileattributes-address
	Address *SegmentDefinition_AddressDimension `json:"Address,omitempty"`

	// Attributes AWS CloudFormation Property
	// Required: false
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-customerprofiles-segmentdefinition-profileattributes.html#cfn-customerprofiles-segmentdefinition-profileattributes-attributes
	Attributes map[string]SegmentDefinition_AttributeDimension `json:"Attributes,omitempty"`

	// BillingAddress AWS CloudFormation Property
	// Required: false
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-customerprofiles-segmentdefinition-profileattributes.html#cfn-customerprofiles-segmentdefinition-profileattributes-billingaddress
	BillingAddress *SegmentDefinition_AddressDimension `json:"BillingAddress,omitempty"`

	// BirthDate AWS CloudFormation Property
	// Required: false
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-customerprofiles-segmentdefinition-profileattributes.html#cfn-customerprofiles-segmentdefinition-profileattributes-birthdate
	BirthDate *SegmentDefinition_DateDimension `json:"BirthDate,omitempty"`

	// BusinessEmailAddress AWS CloudFormation Property
	// Required: false
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-customerprofiles-segmentdefinition-profileattributes.html#cfn-customerprofiles-segmentdefinition-profileattributes-businessemailaddress
	BusinessEmailAddress *SegmentDefinition_ProfileDimension `json:"BusinessEmailAddress,omitempty"`

	// BusinessName AWS CloudFormation Property
	// Required: false
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-customerprofiles-segmentdefinition-profileattributes.html#cfn-customerprofiles-segmentdefinition-profileattributes-businessname
	BusinessName *SegmentDefinition_ProfileDimension `json:"BusinessName,omitempty"`

	// BusinessPhoneNumber AWS CloudFormation Property
	// Required: false
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-customerprofiles-segmentdefinition-profileattributes.html#cfn-customerprofiles-segmentdefinition-profileattributes-businessphonenumber
	BusinessPhoneNumber *SegmentDefinition_ProfileDimension `json:"BusinessPhoneNumber,omitempty"`

	// EmailAddress AWS CloudFormation Property
	// Required: false
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-customerprofiles-segmentdefinition-profileattributes.html#cfn-customerprofiles-segmentdefinition-profileattributes-emailaddress
	EmailAddress *SegmentDefinition_ProfileDimension `json:"EmailAddress,omitempty"`

	// FirstName AWS CloudFormation Property
	// Required: false
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-customerprofiles-segmentdefinition-profileattributes.html#cfn-customerprofiles-segmentdefinition-profileattributes-firstname
	FirstName *SegmentDefinition_ProfileDimension `json:"FirstName,omitempty"`

	// GenderString AWS CloudFormation Property
	// Required: false
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-customerprofiles-segmentdefinition-profileattributes.html#cfn-customerprofiles-segmentdefinition-profileattributes-genderstring
	GenderString *SegmentDefinition_ProfileDimension `json:"GenderString,omitempty"`

	// HomePhoneNumber AWS CloudFormation Property
	// Required: false
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-customerprofiles-segmentdefinition-profileattributes.html#cfn-customerprofiles-segmentdefinition-profileattributes-homephonenumber
	HomePhoneNumber *SegmentDefinition_ProfileDimension `json:"HomePhoneNumber,omitempty"`

	// LastName AWS CloudFormation Property
	// Required: false
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-customerprofiles-segmentdefinition-profileattributes.html#cfn-customerprofiles-segmentdefinition-profileattributes-lastname
	LastName *SegmentDefinition_ProfileDimension `json:"LastName,omitempty"`

	// MailingAddress AWS CloudFormation Property
	// Required: false
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-customerprofiles-segmentdefinition-profileattributes.html#cfn-customerprofiles-segmentdefinition-profileattributes-mailingaddress
	MailingAddress *SegmentDefinition_AddressDimension `json:"MailingAddress,omitempty"`

	// MiddleName AWS CloudFormation Property
	// Required: false
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-customerprofiles-segmentdefinition-profileattributes.html#cfn-customerprofiles-segmentdefinition-profileattributes-middlename
	MiddleName *SegmentDefinition_ProfileDimension `json:"MiddleName,omitempty"`

	// MobilePhoneNumber AWS CloudFormation Property
	// Required: false
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-customerprofiles-segmentdefinition-profileattributes.html#cfn-customerprofiles-segmentdefinition-profileattributes-mobilephonenumber
	MobilePhoneNumber *SegmentDefinition_ProfileDimension `json:"MobilePhoneNumber,omitempty"`

	// PartyTypeString AWS CloudFormation Property
	// Required: false
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-customerprofiles-segmentdefinition-profileattributes.html#cfn-customerprofiles-segmentdefinition-profileattributes-partytypestring
	PartyTypeString *SegmentDefinition_ProfileDimension `json:"PartyTypeString,omitempty"`

	// PersonalEmailAddress AWS CloudFormation Property
	// Required: false
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-customerprofiles-segmentdefinition-profileattributes.html#cfn-customerprofiles-segmentdefinition-profileattributes-personalemailaddress
	PersonalEmailAddress *SegmentDefinition_ProfileDimension `json:"PersonalEmailAddress,omitempty"`

	// PhoneNumber AWS CloudFormation Property
	// Required: false
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-customerprofiles-segmentdefinition-profileattributes.html#cfn-customerprofiles-segmentdefinition-profileattributes-phonenumber
	PhoneNumber *SegmentDefinition_ProfileDimension `json:"PhoneNumber,omitempty"`

	// ShippingAddress AWS CloudFormation Property
	// Required: false
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-customerprofiles-segmentdefinition-profileattributes.html#cfn-customerprofiles-segmentdefinition-profileattributes-shippingaddress
	ShippingAddress *SegmentDefinition_AddressDimension `json:"ShippingAddress,omitempty"`

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
func (r *SegmentDefinition_ProfileAttributes) AWSCloudFormationType() string {
	return "AWS::CustomerProfiles::SegmentDefinition.ProfileAttributes"
}
