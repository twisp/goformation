// Code generated by "go generate". Please don't change this file directly.

package ec2

import (
	"github.com/awslabs/goformation/v7/cloudformation/policies"
)

// VPNConnection_VpnTunnelOptionsSpecification AWS CloudFormation Resource (AWS::EC2::VPNConnection.VpnTunnelOptionsSpecification)
// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-ec2-vpnconnection-vpntunneloptionsspecification.html
type VPNConnection_VpnTunnelOptionsSpecification struct {

	// DPDTimeoutAction AWS CloudFormation Property
	// Required: false
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-ec2-vpnconnection-vpntunneloptionsspecification.html#cfn-ec2-vpnconnection-vpntunneloptionsspecification-dpdtimeoutaction
	DPDTimeoutAction *string `json:"DPDTimeoutAction,omitempty"`

	// DPDTimeoutSeconds AWS CloudFormation Property
	// Required: false
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-ec2-vpnconnection-vpntunneloptionsspecification.html#cfn-ec2-vpnconnection-vpntunneloptionsspecification-dpdtimeoutseconds
	DPDTimeoutSeconds *int `json:"DPDTimeoutSeconds,omitempty"`

	// EnableTunnelLifecycleControl AWS CloudFormation Property
	// Required: false
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-ec2-vpnconnection-vpntunneloptionsspecification.html#cfn-ec2-vpnconnection-vpntunneloptionsspecification-enabletunnellifecyclecontrol
	EnableTunnelLifecycleControl *bool `json:"EnableTunnelLifecycleControl,omitempty"`

	// IKEVersions AWS CloudFormation Property
	// Required: false
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-ec2-vpnconnection-vpntunneloptionsspecification.html#cfn-ec2-vpnconnection-vpntunneloptionsspecification-ikeversions
	IKEVersions []VPNConnection_IKEVersionsRequestListValue `json:"IKEVersions,omitempty"`

	// LogOptions AWS CloudFormation Property
	// Required: false
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-ec2-vpnconnection-vpntunneloptionsspecification.html#cfn-ec2-vpnconnection-vpntunneloptionsspecification-logoptions
	LogOptions *VPNConnection_VpnTunnelLogOptionsSpecification `json:"LogOptions,omitempty"`

	// Phase1DHGroupNumbers AWS CloudFormation Property
	// Required: false
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-ec2-vpnconnection-vpntunneloptionsspecification.html#cfn-ec2-vpnconnection-vpntunneloptionsspecification-phase1dhgroupnumbers
	Phase1DHGroupNumbers []VPNConnection_Phase1DHGroupNumbersRequestListValue `json:"Phase1DHGroupNumbers,omitempty"`

	// Phase1EncryptionAlgorithms AWS CloudFormation Property
	// Required: false
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-ec2-vpnconnection-vpntunneloptionsspecification.html#cfn-ec2-vpnconnection-vpntunneloptionsspecification-phase1encryptionalgorithms
	Phase1EncryptionAlgorithms []VPNConnection_Phase1EncryptionAlgorithmsRequestListValue `json:"Phase1EncryptionAlgorithms,omitempty"`

	// Phase1IntegrityAlgorithms AWS CloudFormation Property
	// Required: false
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-ec2-vpnconnection-vpntunneloptionsspecification.html#cfn-ec2-vpnconnection-vpntunneloptionsspecification-phase1integrityalgorithms
	Phase1IntegrityAlgorithms []VPNConnection_Phase1IntegrityAlgorithmsRequestListValue `json:"Phase1IntegrityAlgorithms,omitempty"`

	// Phase1LifetimeSeconds AWS CloudFormation Property
	// Required: false
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-ec2-vpnconnection-vpntunneloptionsspecification.html#cfn-ec2-vpnconnection-vpntunneloptionsspecification-phase1lifetimeseconds
	Phase1LifetimeSeconds *int `json:"Phase1LifetimeSeconds,omitempty"`

	// Phase2DHGroupNumbers AWS CloudFormation Property
	// Required: false
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-ec2-vpnconnection-vpntunneloptionsspecification.html#cfn-ec2-vpnconnection-vpntunneloptionsspecification-phase2dhgroupnumbers
	Phase2DHGroupNumbers []VPNConnection_Phase2DHGroupNumbersRequestListValue `json:"Phase2DHGroupNumbers,omitempty"`

	// Phase2EncryptionAlgorithms AWS CloudFormation Property
	// Required: false
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-ec2-vpnconnection-vpntunneloptionsspecification.html#cfn-ec2-vpnconnection-vpntunneloptionsspecification-phase2encryptionalgorithms
	Phase2EncryptionAlgorithms []VPNConnection_Phase2EncryptionAlgorithmsRequestListValue `json:"Phase2EncryptionAlgorithms,omitempty"`

	// Phase2IntegrityAlgorithms AWS CloudFormation Property
	// Required: false
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-ec2-vpnconnection-vpntunneloptionsspecification.html#cfn-ec2-vpnconnection-vpntunneloptionsspecification-phase2integrityalgorithms
	Phase2IntegrityAlgorithms []VPNConnection_Phase2IntegrityAlgorithmsRequestListValue `json:"Phase2IntegrityAlgorithms,omitempty"`

	// Phase2LifetimeSeconds AWS CloudFormation Property
	// Required: false
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-ec2-vpnconnection-vpntunneloptionsspecification.html#cfn-ec2-vpnconnection-vpntunneloptionsspecification-phase2lifetimeseconds
	Phase2LifetimeSeconds *int `json:"Phase2LifetimeSeconds,omitempty"`

	// PreSharedKey AWS CloudFormation Property
	// Required: false
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-ec2-vpnconnection-vpntunneloptionsspecification.html#cfn-ec2-vpnconnection-vpntunneloptionsspecification-presharedkey
	PreSharedKey *string `json:"PreSharedKey,omitempty"`

	// RekeyFuzzPercentage AWS CloudFormation Property
	// Required: false
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-ec2-vpnconnection-vpntunneloptionsspecification.html#cfn-ec2-vpnconnection-vpntunneloptionsspecification-rekeyfuzzpercentage
	RekeyFuzzPercentage *int `json:"RekeyFuzzPercentage,omitempty"`

	// RekeyMarginTimeSeconds AWS CloudFormation Property
	// Required: false
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-ec2-vpnconnection-vpntunneloptionsspecification.html#cfn-ec2-vpnconnection-vpntunneloptionsspecification-rekeymargintimeseconds
	RekeyMarginTimeSeconds *int `json:"RekeyMarginTimeSeconds,omitempty"`

	// ReplayWindowSize AWS CloudFormation Property
	// Required: false
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-ec2-vpnconnection-vpntunneloptionsspecification.html#cfn-ec2-vpnconnection-vpntunneloptionsspecification-replaywindowsize
	ReplayWindowSize *int `json:"ReplayWindowSize,omitempty"`

	// StartupAction AWS CloudFormation Property
	// Required: false
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-ec2-vpnconnection-vpntunneloptionsspecification.html#cfn-ec2-vpnconnection-vpntunneloptionsspecification-startupaction
	StartupAction *string `json:"StartupAction,omitempty"`

	// TunnelInsideCidr AWS CloudFormation Property
	// Required: false
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-ec2-vpnconnection-vpntunneloptionsspecification.html#cfn-ec2-vpnconnection-vpntunneloptionsspecification-tunnelinsidecidr
	TunnelInsideCidr *string `json:"TunnelInsideCidr,omitempty"`

	// TunnelInsideIpv6Cidr AWS CloudFormation Property
	// Required: false
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-ec2-vpnconnection-vpntunneloptionsspecification.html#cfn-ec2-vpnconnection-vpntunneloptionsspecification-tunnelinsideipv6cidr
	TunnelInsideIpv6Cidr *string `json:"TunnelInsideIpv6Cidr,omitempty"`

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
func (r *VPNConnection_VpnTunnelOptionsSpecification) AWSCloudFormationType() string {
	return "AWS::EC2::VPNConnection.VpnTunnelOptionsSpecification"
}
