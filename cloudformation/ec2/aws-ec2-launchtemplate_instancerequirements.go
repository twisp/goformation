// Code generated by "go generate". Please don't change this file directly.

package ec2

import (
	"github.com/awslabs/goformation/v7/cloudformation/policies"
)

// LaunchTemplate_InstanceRequirements AWS CloudFormation Resource (AWS::EC2::LaunchTemplate.InstanceRequirements)
// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-ec2-launchtemplate-instancerequirements.html
type LaunchTemplate_InstanceRequirements struct {

	// AcceleratorCount AWS CloudFormation Property
	// Required: false
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-ec2-launchtemplate-instancerequirements.html#cfn-ec2-launchtemplate-instancerequirements-acceleratorcount
	AcceleratorCount *LaunchTemplate_AcceleratorCount `json:"AcceleratorCount,omitempty"`

	// AcceleratorManufacturers AWS CloudFormation Property
	// Required: false
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-ec2-launchtemplate-instancerequirements.html#cfn-ec2-launchtemplate-instancerequirements-acceleratormanufacturers
	AcceleratorManufacturers []string `json:"AcceleratorManufacturers,omitempty"`

	// AcceleratorNames AWS CloudFormation Property
	// Required: false
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-ec2-launchtemplate-instancerequirements.html#cfn-ec2-launchtemplate-instancerequirements-acceleratornames
	AcceleratorNames []string `json:"AcceleratorNames,omitempty"`

	// AcceleratorTotalMemoryMiB AWS CloudFormation Property
	// Required: false
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-ec2-launchtemplate-instancerequirements.html#cfn-ec2-launchtemplate-instancerequirements-acceleratortotalmemorymib
	AcceleratorTotalMemoryMiB *LaunchTemplate_AcceleratorTotalMemoryMiB `json:"AcceleratorTotalMemoryMiB,omitempty"`

	// AcceleratorTypes AWS CloudFormation Property
	// Required: false
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-ec2-launchtemplate-instancerequirements.html#cfn-ec2-launchtemplate-instancerequirements-acceleratortypes
	AcceleratorTypes []string `json:"AcceleratorTypes,omitempty"`

	// AllowedInstanceTypes AWS CloudFormation Property
	// Required: false
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-ec2-launchtemplate-instancerequirements.html#cfn-ec2-launchtemplate-instancerequirements-allowedinstancetypes
	AllowedInstanceTypes []string `json:"AllowedInstanceTypes,omitempty"`

	// BareMetal AWS CloudFormation Property
	// Required: false
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-ec2-launchtemplate-instancerequirements.html#cfn-ec2-launchtemplate-instancerequirements-baremetal
	BareMetal *string `json:"BareMetal,omitempty"`

	// BaselineEbsBandwidthMbps AWS CloudFormation Property
	// Required: false
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-ec2-launchtemplate-instancerequirements.html#cfn-ec2-launchtemplate-instancerequirements-baselineebsbandwidthmbps
	BaselineEbsBandwidthMbps *LaunchTemplate_BaselineEbsBandwidthMbps `json:"BaselineEbsBandwidthMbps,omitempty"`

	// BaselinePerformanceFactors AWS CloudFormation Property
	// Required: false
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-ec2-launchtemplate-instancerequirements.html#cfn-ec2-launchtemplate-instancerequirements-baselineperformancefactors
	BaselinePerformanceFactors *LaunchTemplate_BaselinePerformanceFactors `json:"BaselinePerformanceFactors,omitempty"`

	// BurstablePerformance AWS CloudFormation Property
	// Required: false
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-ec2-launchtemplate-instancerequirements.html#cfn-ec2-launchtemplate-instancerequirements-burstableperformance
	BurstablePerformance *string `json:"BurstablePerformance,omitempty"`

	// CpuManufacturers AWS CloudFormation Property
	// Required: false
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-ec2-launchtemplate-instancerequirements.html#cfn-ec2-launchtemplate-instancerequirements-cpumanufacturers
	CpuManufacturers []string `json:"CpuManufacturers,omitempty"`

	// ExcludedInstanceTypes AWS CloudFormation Property
	// Required: false
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-ec2-launchtemplate-instancerequirements.html#cfn-ec2-launchtemplate-instancerequirements-excludedinstancetypes
	ExcludedInstanceTypes []string `json:"ExcludedInstanceTypes,omitempty"`

	// InstanceGenerations AWS CloudFormation Property
	// Required: false
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-ec2-launchtemplate-instancerequirements.html#cfn-ec2-launchtemplate-instancerequirements-instancegenerations
	InstanceGenerations []string `json:"InstanceGenerations,omitempty"`

	// LocalStorage AWS CloudFormation Property
	// Required: false
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-ec2-launchtemplate-instancerequirements.html#cfn-ec2-launchtemplate-instancerequirements-localstorage
	LocalStorage *string `json:"LocalStorage,omitempty"`

	// LocalStorageTypes AWS CloudFormation Property
	// Required: false
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-ec2-launchtemplate-instancerequirements.html#cfn-ec2-launchtemplate-instancerequirements-localstoragetypes
	LocalStorageTypes []string `json:"LocalStorageTypes,omitempty"`

	// MaxSpotPriceAsPercentageOfOptimalOnDemandPrice AWS CloudFormation Property
	// Required: false
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-ec2-launchtemplate-instancerequirements.html#cfn-ec2-launchtemplate-instancerequirements-maxspotpriceaspercentageofoptimalondemandprice
	MaxSpotPriceAsPercentageOfOptimalOnDemandPrice *int `json:"MaxSpotPriceAsPercentageOfOptimalOnDemandPrice,omitempty"`

	// MemoryGiBPerVCpu AWS CloudFormation Property
	// Required: false
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-ec2-launchtemplate-instancerequirements.html#cfn-ec2-launchtemplate-instancerequirements-memorygibpervcpu
	MemoryGiBPerVCpu *LaunchTemplate_MemoryGiBPerVCpu `json:"MemoryGiBPerVCpu,omitempty"`

	// MemoryMiB AWS CloudFormation Property
	// Required: false
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-ec2-launchtemplate-instancerequirements.html#cfn-ec2-launchtemplate-instancerequirements-memorymib
	MemoryMiB *LaunchTemplate_MemoryMiB `json:"MemoryMiB,omitempty"`

	// NetworkBandwidthGbps AWS CloudFormation Property
	// Required: false
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-ec2-launchtemplate-instancerequirements.html#cfn-ec2-launchtemplate-instancerequirements-networkbandwidthgbps
	NetworkBandwidthGbps *LaunchTemplate_NetworkBandwidthGbps `json:"NetworkBandwidthGbps,omitempty"`

	// NetworkInterfaceCount AWS CloudFormation Property
	// Required: false
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-ec2-launchtemplate-instancerequirements.html#cfn-ec2-launchtemplate-instancerequirements-networkinterfacecount
	NetworkInterfaceCount *LaunchTemplate_NetworkInterfaceCount `json:"NetworkInterfaceCount,omitempty"`

	// OnDemandMaxPricePercentageOverLowestPrice AWS CloudFormation Property
	// Required: false
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-ec2-launchtemplate-instancerequirements.html#cfn-ec2-launchtemplate-instancerequirements-ondemandmaxpricepercentageoverlowestprice
	OnDemandMaxPricePercentageOverLowestPrice *int `json:"OnDemandMaxPricePercentageOverLowestPrice,omitempty"`

	// RequireHibernateSupport AWS CloudFormation Property
	// Required: false
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-ec2-launchtemplate-instancerequirements.html#cfn-ec2-launchtemplate-instancerequirements-requirehibernatesupport
	RequireHibernateSupport *bool `json:"RequireHibernateSupport,omitempty"`

	// SpotMaxPricePercentageOverLowestPrice AWS CloudFormation Property
	// Required: false
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-ec2-launchtemplate-instancerequirements.html#cfn-ec2-launchtemplate-instancerequirements-spotmaxpricepercentageoverlowestprice
	SpotMaxPricePercentageOverLowestPrice *int `json:"SpotMaxPricePercentageOverLowestPrice,omitempty"`

	// TotalLocalStorageGB AWS CloudFormation Property
	// Required: false
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-ec2-launchtemplate-instancerequirements.html#cfn-ec2-launchtemplate-instancerequirements-totallocalstoragegb
	TotalLocalStorageGB *LaunchTemplate_TotalLocalStorageGB `json:"TotalLocalStorageGB,omitempty"`

	// VCpuCount AWS CloudFormation Property
	// Required: false
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-ec2-launchtemplate-instancerequirements.html#cfn-ec2-launchtemplate-instancerequirements-vcpucount
	VCpuCount *LaunchTemplate_VCpuCount `json:"VCpuCount,omitempty"`

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
func (r *LaunchTemplate_InstanceRequirements) AWSCloudFormationType() string {
	return "AWS::EC2::LaunchTemplate.InstanceRequirements"
}
