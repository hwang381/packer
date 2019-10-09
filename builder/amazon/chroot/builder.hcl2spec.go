// Code generated by "mapstructure-to-hcl2 -type Config,BlockDevices,BlockDevice"; DO NOT EDIT.
package chroot

import (
	"github.com/hashicorp/hcl/v2/hcldec"
	"github.com/hashicorp/packer/builder/amazon/common"
	"github.com/hashicorp/packer/helper/config"
	"github.com/zclconf/go-cty/cty"
)

type FlatConfig struct {
	PackerBuildName         string                       `mapstructure:"packer_build_name" cty:"packer_build_name"`
	PackerBuilderType       string                       `mapstructure:"packer_builder_type" cty:"packer_builder_type"`
	PackerDebug             bool                         `mapstructure:"packer_debug" cty:"packer_debug"`
	PackerForce             bool                         `mapstructure:"packer_force" cty:"packer_force"`
	PackerOnError           string                       `mapstructure:"packer_on_error" cty:"packer_on_error"`
	PackerUserVars          map[string]string            `mapstructure:"packer_user_variables" cty:"packer_user_variables"`
	PackerSensitiveVars     []string                     `mapstructure:"packer_sensitive_variables" cty:"packer_sensitive_variables"`
	AMIName                 string                       `mapstructure:"ami_name" required:"true" cty:"ami_name"`
	AMIDescription          string                       `mapstructure:"ami_description" required:"false" cty:"ami_description"`
	AMIVirtType             string                       `mapstructure:"ami_virtualization_type" required:"false" cty:"ami_virtualization_type"`
	AMIUsers                []string                     `mapstructure:"ami_users" required:"false" cty:"ami_users"`
	AMIGroups               []string                     `mapstructure:"ami_groups" required:"false" cty:"ami_groups"`
	AMIProductCodes         []string                     `mapstructure:"ami_product_codes" required:"false" cty:"ami_product_codes"`
	AMIRegions              []string                     `mapstructure:"ami_regions" required:"false" cty:"ami_regions"`
	AMISkipRegionValidation bool                         `mapstructure:"skip_region_validation" required:"false" cty:"skip_region_validation"`
	AMITags                 common.TagMap                `mapstructure:"tags" required:"false" cty:"tags"`
	AMIENASupport           config.Trilean               `mapstructure:"ena_support" required:"false" cty:"ena_support"`
	AMISriovNetSupport      bool                         `mapstructure:"sriov_support" required:"false" cty:"sriov_support"`
	AMIForceDeregister      bool                         `mapstructure:"force_deregister" required:"false" cty:"force_deregister"`
	AMIForceDeleteSnapshot  bool                         `mapstructure:"force_delete_snapshot" required:"false" cty:"force_delete_snapshot"`
	AMIEncryptBootVolume    config.Trilean               `mapstructure:"encrypt_boot" required:"false" cty:"encrypt_boot"`
	AMIKmsKeyId             string                       `mapstructure:"kms_key_id" required:"false" cty:"kms_key_id"`
	AMIRegionKMSKeyIDs      map[string]string            `mapstructure:"region_kms_key_ids" required:"false" cty:"region_kms_key_ids"`
	AMISkipBuildRegion      bool                         `mapstructure:"skip_save_build_region" cty:"skip_save_build_region"`
	SnapshotTags            common.TagMap                `mapstructure:"snapshot_tags" required:"false" cty:"snapshot_tags"`
	SnapshotUsers           []string                     `mapstructure:"snapshot_users" required:"false" cty:"snapshot_users"`
	SnapshotGroups          []string                     `mapstructure:"snapshot_groups" required:"false" cty:"snapshot_groups"`
	AccessKey               string                       `mapstructure:"access_key" required:"true" cty:"access_key"`
	CustomEndpointEc2       string                       `mapstructure:"custom_endpoint_ec2" required:"false" cty:"custom_endpoint_ec2"`
	DecodeAuthZMessages     bool                         `mapstructure:"decode_authorization_messages" required:"false" cty:"decode_authorization_messages"`
	InsecureSkipTLSVerify   bool                         `mapstructure:"insecure_skip_tls_verify" required:"false" cty:"insecure_skip_tls_verify"`
	MFACode                 string                       `mapstructure:"mfa_code" required:"false" cty:"mfa_code"`
	ProfileName             string                       `mapstructure:"profile" required:"false" cty:"profile"`
	RawRegion               string                       `mapstructure:"region" required:"true" cty:"region"`
	SecretKey               string                       `mapstructure:"secret_key" required:"true" cty:"secret_key"`
	SkipValidation          bool                         `mapstructure:"skip_region_validation" required:"false" cty:"skip_region_validation"`
	SkipMetadataApiCheck    bool                         `mapstructure:"skip_metadata_api_check" cty:"skip_metadata_api_check"`
	Token                   string                       `mapstructure:"token" required:"false" cty:"token"`
	VaultAWSEngine          common.VaultAWSEngineOptions `mapstructure:"vault_aws_engine" required:"false" cty:"vault_aws_engine"`
	AMIMappings             common.BlockDevices          `mapstructure:"ami_block_device_mappings" hcl2-schema-generator:"ami_block_device_mappings,direct" required:"false" cty:"ami_block_device_mappings"`
	ChrootMounts            [][]string                   `mapstructure:"chroot_mounts" required:"false" cty:"chroot_mounts"`
	CommandWrapper          string                       `mapstructure:"command_wrapper" required:"false" cty:"command_wrapper"`
	CopyFiles               []string                     `mapstructure:"copy_files" required:"false" cty:"copy_files"`
	DevicePath              string                       `mapstructure:"device_path" required:"false" cty:"device_path"`
	NVMEDevicePath          string                       `mapstructure:"nvme_device_path" required:"false" cty:"nvme_device_path"`
	FromScratch             bool                         `mapstructure:"from_scratch" required:"false" cty:"from_scratch"`
	MountOptions            []string                     `mapstructure:"mount_options" required:"false" cty:"mount_options"`
	MountPartition          string                       `mapstructure:"mount_partition" required:"false" cty:"mount_partition"`
	MountPath               string                       `mapstructure:"mount_path" required:"false" cty:"mount_path"`
	PostMountCommands       []string                     `mapstructure:"post_mount_commands" required:"false" cty:"post_mount_commands"`
	PreMountCommands        []string                     `mapstructure:"pre_mount_commands" required:"false" cty:"pre_mount_commands"`
	RootDeviceName          string                       `mapstructure:"root_device_name" required:"false" cty:"root_device_name"`
	RootVolumeSize          int64                        `mapstructure:"root_volume_size" required:"false" cty:"root_volume_size"`
	RootVolumeType          string                       `mapstructure:"root_volume_type" required:"false" cty:"root_volume_type"`
	SourceAmi               string                       `mapstructure:"source_ami" required:"true" cty:"source_ami"`
	SourceAmiFilter         common.AmiFilterOptions      `mapstructure:"source_ami_filter" required:"false" cty:"source_ami_filter"`
	RootVolumeTags          common.TagMap                `mapstructure:"root_volume_tags" required:"false" cty:"root_volume_tags"`
	Architecture            string                       `mapstructure:"ami_architecture" required:"false" cty:"ami_architecture"`
}

func (*Config) HCL2Spec() map[string]hcldec.Spec {
	s := map[string]hcldec.Spec{
		"PackerBuildName":         &hcldec.AttrSpec{Name: "packer_build_name", Type: cty.String, Required: false},
		"PackerBuilderType":       &hcldec.AttrSpec{Name: "packer_builder_type", Type: cty.String, Required: false},
		"PackerDebug":             &hcldec.AttrSpec{Name: "packer_debug", Type: cty.Bool, Required: false},
		"PackerForce":             &hcldec.AttrSpec{Name: "packer_force", Type: cty.Bool, Required: false},
		"PackerOnError":           &hcldec.AttrSpec{Name: "packer_on_error", Type: cty.String, Required: false},
		"PackerUserVars":          &hcldec.BlockAttrsSpec{TypeName: "packer_user_variables", ElementType: cty.String, Required: false},
		"PackerSensitiveVars":     &hcldec.AttrSpec{Name: "packer_sensitive_variables", Type: cty.List(cty.String), Required: false},
		"AMIName":                 &hcldec.AttrSpec{Name: "ami_name", Type: cty.String, Required: false},
		"AMIDescription":          &hcldec.AttrSpec{Name: "ami_description", Type: cty.String, Required: false},
		"AMIVirtType":             &hcldec.AttrSpec{Name: "ami_virtualization_type", Type: cty.String, Required: false},
		"AMIUsers":                &hcldec.AttrSpec{Name: "ami_users", Type: cty.List(cty.String), Required: false},
		"AMIGroups":               &hcldec.AttrSpec{Name: "ami_groups", Type: cty.List(cty.String), Required: false},
		"AMIProductCodes":         &hcldec.AttrSpec{Name: "ami_product_codes", Type: cty.List(cty.String), Required: false},
		"AMIRegions":              &hcldec.AttrSpec{Name: "ami_regions", Type: cty.List(cty.String), Required: false},
		"AMISkipRegionValidation": &hcldec.AttrSpec{Name: "skip_region_validation", Type: cty.Bool, Required: false},
		"AMITags":                 &hcldec.BlockAttrsSpec{TypeName: "common.TagMap", ElementType: cty.String, Required: false},
		"AMIENASupport":           &hcldec.AttrSpec{Name: "config.Trilean", Type: cty.Number, Required: false},
		"AMISriovNetSupport":      &hcldec.AttrSpec{Name: "sriov_support", Type: cty.Bool, Required: false},
		"AMIForceDeregister":      &hcldec.AttrSpec{Name: "force_deregister", Type: cty.Bool, Required: false},
		"AMIForceDeleteSnapshot":  &hcldec.AttrSpec{Name: "force_delete_snapshot", Type: cty.Bool, Required: false},
		"AMIEncryptBootVolume":    &hcldec.AttrSpec{Name: "config.Trilean", Type: cty.Number, Required: false},
		"AMIKmsKeyId":             &hcldec.AttrSpec{Name: "kms_key_id", Type: cty.String, Required: false},
		"AMIRegionKMSKeyIDs":      &hcldec.BlockAttrsSpec{TypeName: "region_kms_key_ids", ElementType: cty.String, Required: false},
		"AMISkipBuildRegion":      &hcldec.AttrSpec{Name: "skip_save_build_region", Type: cty.Bool, Required: false},
		"SnapshotTags":            &hcldec.BlockAttrsSpec{TypeName: "common.TagMap", ElementType: cty.String, Required: false},
		"SnapshotUsers":           &hcldec.AttrSpec{Name: "snapshot_users", Type: cty.List(cty.String), Required: false},
		"SnapshotGroups":          &hcldec.AttrSpec{Name: "snapshot_groups", Type: cty.List(cty.String), Required: false},
		"AccessKey":               &hcldec.AttrSpec{Name: "access_key", Type: cty.String, Required: false},
		"CustomEndpointEc2":       &hcldec.AttrSpec{Name: "custom_endpoint_ec2", Type: cty.String, Required: false},
		"DecodeAuthZMessages":     &hcldec.AttrSpec{Name: "decode_authorization_messages", Type: cty.Bool, Required: false},
		"InsecureSkipTLSVerify":   &hcldec.AttrSpec{Name: "insecure_skip_tls_verify", Type: cty.Bool, Required: false},
		"MFACode":                 &hcldec.AttrSpec{Name: "mfa_code", Type: cty.String, Required: false},
		"ProfileName":             &hcldec.AttrSpec{Name: "profile", Type: cty.String, Required: false},
		"RawRegion":               &hcldec.AttrSpec{Name: "region", Type: cty.String, Required: false},
		"SecretKey":               &hcldec.AttrSpec{Name: "secret_key", Type: cty.String, Required: false},
		"SkipValidation":          &hcldec.AttrSpec{Name: "skip_region_validation", Type: cty.Bool, Required: false},
		"SkipMetadataApiCheck":    &hcldec.AttrSpec{Name: "skip_metadata_api_check", Type: cty.Bool, Required: false},
		"Token":                   &hcldec.AttrSpec{Name: "token", Type: cty.String, Required: false},
		"VaultAWSEngine":          &hcldec.BlockObjectSpec{TypeName: "common.VaultAWSEngineOptions", Nested: hcldec.ObjectSpec((*common.VaultAWSEngineOptions)(nil).HCL2Spec())},
		"AMIMappings":             &hcldec.BlockListSpec{TypeName: "[]common.BlockDevice", Nested: &hcldec.BlockObjectSpec{TypeName: "common.BlockDevice", Nested: hcldec.ObjectSpec((*common.BlockDevice)(nil).HCL2Spec())}},
		"ChrootMounts":            nil, // slice ([][]string),
		"CommandWrapper":          &hcldec.AttrSpec{Name: "command_wrapper", Type: cty.String, Required: false},
		"CopyFiles":               &hcldec.AttrSpec{Name: "copy_files", Type: cty.List(cty.String), Required: false},
		"DevicePath":              &hcldec.AttrSpec{Name: "device_path", Type: cty.String, Required: false},
		"NVMEDevicePath":          &hcldec.AttrSpec{Name: "nvme_device_path", Type: cty.String, Required: false},
		"FromScratch":             &hcldec.AttrSpec{Name: "from_scratch", Type: cty.Bool, Required: false},
		"MountOptions":            &hcldec.AttrSpec{Name: "mount_options", Type: cty.List(cty.String), Required: false},
		"MountPartition":          &hcldec.AttrSpec{Name: "mount_partition", Type: cty.String, Required: false},
		"MountPath":               &hcldec.AttrSpec{Name: "mount_path", Type: cty.String, Required: false},
		"PostMountCommands":       &hcldec.AttrSpec{Name: "post_mount_commands", Type: cty.List(cty.String), Required: false},
		"PreMountCommands":        &hcldec.AttrSpec{Name: "pre_mount_commands", Type: cty.List(cty.String), Required: false},
		"RootDeviceName":          &hcldec.AttrSpec{Name: "root_device_name", Type: cty.String, Required: false},
		"RootVolumeSize":          &hcldec.AttrSpec{Name: "root_volume_size", Type: cty.Number, Required: false},
		"RootVolumeType":          &hcldec.AttrSpec{Name: "root_volume_type", Type: cty.String, Required: false},
		"SourceAmi":               &hcldec.AttrSpec{Name: "source_ami", Type: cty.String, Required: false},
		"SourceAmiFilter":         &hcldec.BlockObjectSpec{TypeName: "common.AmiFilterOptions", Nested: hcldec.ObjectSpec((*common.AmiFilterOptions)(nil).HCL2Spec())},
		"RootVolumeTags":          &hcldec.BlockAttrsSpec{TypeName: "common.TagMap", ElementType: cty.String, Required: false},
		"Architecture":            &hcldec.AttrSpec{Name: "ami_architecture", Type: cty.String, Required: false},
	}
	return s
}
