// Code generated by "mapstructure-to-hcl2 -type Config"; DO NOT EDIT.
package cvm

import (
	"github.com/hashicorp/hcl/v2/hcldec"
	"github.com/zclconf/go-cty/cty"
	"time"
)

// FlatConfig is an auto-generated flat version of Config.
// Where the contents of a field with a `mapstructure:,squash` tag are bubbled up.
type FlatConfig struct {
	PackerBuildName           string                 `mapstructure:"packer_build_name" cty:"packer_build_name"`
	PackerBuilderType         string                 `mapstructure:"packer_builder_type" cty:"packer_builder_type"`
	PackerDebug               bool                   `mapstructure:"packer_debug" cty:"packer_debug"`
	PackerForce               bool                   `mapstructure:"packer_force" cty:"packer_force"`
	PackerOnError             string                 `mapstructure:"packer_on_error" cty:"packer_on_error"`
	PackerUserVars            map[string]string      `mapstructure:"packer_user_variables" cty:"packer_user_variables"`
	PackerSensitiveVars       []string               `mapstructure:"packer_sensitive_variables" cty:"packer_sensitive_variables"`
	SecretId                  string                 `mapstructure:"secret_id" required:"true" cty:"secret_id"`
	SecretKey                 string                 `mapstructure:"secret_key" required:"true" cty:"secret_key"`
	Region                    string                 `mapstructure:"region" required:"true" cty:"region"`
	Zone                      string                 `mapstructure:"zone" required:"true" cty:"zone"`
	SkipValidation            bool                   `mapstructure:"skip_region_validation" required:"false" cty:"skip_region_validation"`
	ImageName                 string                 `mapstructure:"image_name" required:"true" cty:"image_name"`
	ImageDescription          string                 `mapstructure:"image_description" required:"false" cty:"image_description"`
	Reboot                    bool                   `mapstructure:"reboot" required:"false" cty:"reboot"`
	ForcePoweroff             bool                   `mapstructure:"force_poweroff" required:"false" cty:"force_poweroff"`
	Sysprep                   bool                   `mapstructure:"sysprep" required:"false" cty:"sysprep"`
	ImageForceDelete          bool                   `mapstructure:"image_force_delete" cty:"image_force_delete"`
	ImageCopyRegions          []string               `mapstructure:"image_copy_regions" required:"false" cty:"image_copy_regions"`
	ImageShareAccounts        []string               `mapstructure:"image_share_accounts" required:"false" cty:"image_share_accounts"`
	AssociatePublicIpAddress  bool                   `mapstructure:"associate_public_ip_address" required:"false" cty:"associate_public_ip_address"`
	SourceImageId             string                 `mapstructure:"source_image_id" required:"true" cty:"source_image_id"`
	InstanceType              string                 `mapstructure:"instance_type" required:"true" cty:"instance_type"`
	InstanceName              string                 `mapstructure:"instance_name" required:"false" cty:"instance_name"`
	DiskType                  string                 `mapstructure:"disk_type" required:"false" cty:"disk_type"`
	DiskSize                  int64                  `mapstructure:"disk_size" required:"false" cty:"disk_size"`
	DataDisks                 []tencentCloudDataDisk `mapstructure:"data_disks" cty:"data_disks"`
	VpcId                     string                 `mapstructure:"vpc_id" required:"false" cty:"vpc_id"`
	VpcName                   string                 `mapstructure:"vpc_name" required:"false" cty:"vpc_name"`
	VpcIp                     string                 `mapstructure:"vpc_ip" cty:"vpc_ip"`
	SubnetId                  string                 `mapstructure:"subnet_id" required:"false" cty:"subnet_id"`
	SubnetName                string                 `mapstructure:"subnet_name" required:"false" cty:"subnet_name"`
	CidrBlock                 string                 `mapstructure:"cidr_block" required:"false" cty:"cidr_block"`
	SubnectCidrBlock          string                 `mapstructure:"subnect_cidr_block" required:"false" cty:"subnect_cidr_block"`
	InternetChargeType        string                 `mapstructure:"internet_charge_type" cty:"internet_charge_type"`
	InternetMaxBandwidthOut   int64                  `mapstructure:"internet_max_bandwidth_out" required:"false" cty:"internet_max_bandwidth_out"`
	SecurityGroupId           string                 `mapstructure:"security_group_id" required:"false" cty:"security_group_id"`
	SecurityGroupName         string                 `mapstructure:"security_group_name" required:"false" cty:"security_group_name"`
	UserData                  string                 `mapstructure:"user_data" required:"false" cty:"user_data"`
	UserDataFile              string                 `mapstructure:"user_data_file" required:"false" cty:"user_data_file"`
	HostName                  string                 `mapstructure:"host_name" required:"false" cty:"host_name"`
	RunTags                   map[string]string      `mapstructure:"run_tags" required:"false" cty:"run_tags"`
	Type                      string                 `mapstructure:"communicator" cty:"communicator"`
	PauseBeforeConnect        time.Duration          `mapstructure:"pause_before_connecting" cty:"pause_before_connecting"`
	SSHHost                   string                 `mapstructure:"ssh_host" cty:"ssh_host"`
	SSHPort                   int                    `mapstructure:"ssh_port" cty:"ssh_port"`
	SSHUsername               string                 `mapstructure:"ssh_username" cty:"ssh_username"`
	SSHPassword               string                 `mapstructure:"ssh_password" cty:"ssh_password"`
	SSHKeyPairName            string                 `mapstructure:"ssh_keypair_name" cty:"ssh_keypair_name"`
	SSHTemporaryKeyPairName   string                 `mapstructure:"temporary_key_pair_name" cty:"temporary_key_pair_name"`
	SSHClearAuthorizedKeys    bool                   `mapstructure:"ssh_clear_authorized_keys" cty:"ssh_clear_authorized_keys"`
	SSHPrivateKeyFile         string                 `mapstructure:"ssh_private_key_file" cty:"ssh_private_key_file"`
	SSHPty                    bool                   `mapstructure:"ssh_pty" cty:"ssh_pty"`
	SSHTimeout                time.Duration          `mapstructure:"ssh_timeout" cty:"ssh_timeout"`
	SSHAgentAuth              bool                   `mapstructure:"ssh_agent_auth" cty:"ssh_agent_auth"`
	SSHDisableAgentForwarding bool                   `mapstructure:"ssh_disable_agent_forwarding" cty:"ssh_disable_agent_forwarding"`
	SSHHandshakeAttempts      int                    `mapstructure:"ssh_handshake_attempts" cty:"ssh_handshake_attempts"`
	SSHBastionHost            string                 `mapstructure:"ssh_bastion_host" cty:"ssh_bastion_host"`
	SSHBastionPort            int                    `mapstructure:"ssh_bastion_port" cty:"ssh_bastion_port"`
	SSHBastionAgentAuth       bool                   `mapstructure:"ssh_bastion_agent_auth" cty:"ssh_bastion_agent_auth"`
	SSHBastionUsername        string                 `mapstructure:"ssh_bastion_username" cty:"ssh_bastion_username"`
	SSHBastionPassword        string                 `mapstructure:"ssh_bastion_password" cty:"ssh_bastion_password"`
	SSHBastionPrivateKeyFile  string                 `mapstructure:"ssh_bastion_private_key_file" cty:"ssh_bastion_private_key_file"`
	SSHFileTransferMethod     string                 `mapstructure:"ssh_file_transfer_method" cty:"ssh_file_transfer_method"`
	SSHProxyHost              string                 `mapstructure:"ssh_proxy_host" cty:"ssh_proxy_host"`
	SSHProxyPort              int                    `mapstructure:"ssh_proxy_port" cty:"ssh_proxy_port"`
	SSHProxyUsername          string                 `mapstructure:"ssh_proxy_username" cty:"ssh_proxy_username"`
	SSHProxyPassword          string                 `mapstructure:"ssh_proxy_password" cty:"ssh_proxy_password"`
	SSHKeepAliveInterval      time.Duration          `mapstructure:"ssh_keep_alive_interval" cty:"ssh_keep_alive_interval"`
	SSHReadWriteTimeout       time.Duration          `mapstructure:"ssh_read_write_timeout" cty:"ssh_read_write_timeout"`
	SSHRemoteTunnels          []string               `mapstructure:"ssh_remote_tunnels" cty:"ssh_remote_tunnels"`
	SSHLocalTunnels           []string               `mapstructure:"ssh_local_tunnels" cty:"ssh_local_tunnels"`
	SSHPublicKey              []byte                 `cty:"ssh_public_key"`
	SSHPrivateKey             []byte                 `cty:"ssh_private_key"`
	WinRMUser                 string                 `mapstructure:"winrm_username" cty:"winrm_username"`
	WinRMPassword             string                 `mapstructure:"winrm_password" cty:"winrm_password"`
	WinRMHost                 string                 `mapstructure:"winrm_host" cty:"winrm_host"`
	WinRMPort                 int                    `mapstructure:"winrm_port" cty:"winrm_port"`
	WinRMTimeout              time.Duration          `mapstructure:"winrm_timeout" cty:"winrm_timeout"`
	WinRMUseSSL               bool                   `mapstructure:"winrm_use_ssl" cty:"winrm_use_ssl"`
	WinRMInsecure             bool                   `mapstructure:"winrm_insecure" cty:"winrm_insecure"`
	WinRMUseNTLM              bool                   `mapstructure:"winrm_use_ntlm" cty:"winrm_use_ntlm"`
	SSHPrivateIp              bool                   `mapstructure:"ssh_private_ip" cty:"ssh_private_ip"`
}

// FlatMapstructure returns a new FlatConfig.
// FlatConfig is an auto-generated flat version of Config.
// Where the contents a fields with a `mapstructure:,squash` tag are bubbled up.
func (*Config) FlatMapstructure() interface{} { return new(FlatConfig) }

// HCL2Spec returns the hcldec.Spec of a Config.
// This spec is used by HCL to read the fields of Config.
func (*Config) HCL2Spec() map[string]hcldec.Spec {
	s := map[string]hcldec.Spec{
		"PackerBuildName":           &hcldec.AttrSpec{Name: "packer_build_name", Type: cty.String, Required: false},
		"PackerBuilderType":         &hcldec.AttrSpec{Name: "packer_builder_type", Type: cty.String, Required: false},
		"PackerDebug":               &hcldec.AttrSpec{Name: "packer_debug", Type: cty.Bool, Required: false},
		"PackerForce":               &hcldec.AttrSpec{Name: "packer_force", Type: cty.Bool, Required: false},
		"PackerOnError":             &hcldec.AttrSpec{Name: "packer_on_error", Type: cty.String, Required: false},
		"PackerUserVars":            &hcldec.BlockAttrsSpec{TypeName: "packer_user_variables", ElementType: cty.String, Required: false},
		"PackerSensitiveVars":       &hcldec.AttrSpec{Name: "packer_sensitive_variables", Type: cty.List(cty.String), Required: false},
		"SecretId":                  &hcldec.AttrSpec{Name: "secret_id", Type: cty.String, Required: false},
		"SecretKey":                 &hcldec.AttrSpec{Name: "secret_key", Type: cty.String, Required: false},
		"Region":                    &hcldec.AttrSpec{Name: "region", Type: cty.String, Required: false},
		"Zone":                      &hcldec.AttrSpec{Name: "zone", Type: cty.String, Required: false},
		"SkipValidation":            &hcldec.AttrSpec{Name: "skip_region_validation", Type: cty.Bool, Required: false},
		"ImageName":                 &hcldec.AttrSpec{Name: "image_name", Type: cty.String, Required: false},
		"ImageDescription":          &hcldec.AttrSpec{Name: "image_description", Type: cty.String, Required: false},
		"Reboot":                    &hcldec.AttrSpec{Name: "reboot", Type: cty.Bool, Required: false},
		"ForcePoweroff":             &hcldec.AttrSpec{Name: "force_poweroff", Type: cty.Bool, Required: false},
		"Sysprep":                   &hcldec.AttrSpec{Name: "sysprep", Type: cty.Bool, Required: false},
		"ImageForceDelete":          &hcldec.AttrSpec{Name: "image_force_delete", Type: cty.Bool, Required: false},
		"ImageCopyRegions":          &hcldec.AttrSpec{Name: "image_copy_regions", Type: cty.List(cty.String), Required: false},
		"ImageShareAccounts":        &hcldec.AttrSpec{Name: "image_share_accounts", Type: cty.List(cty.String), Required: false},
		"AssociatePublicIpAddress":  &hcldec.AttrSpec{Name: "associate_public_ip_address", Type: cty.Bool, Required: false},
		"SourceImageId":             &hcldec.AttrSpec{Name: "source_image_id", Type: cty.String, Required: false},
		"InstanceType":              &hcldec.AttrSpec{Name: "instance_type", Type: cty.String, Required: false},
		"InstanceName":              &hcldec.AttrSpec{Name: "instance_name", Type: cty.String, Required: false},
		"DiskType":                  &hcldec.AttrSpec{Name: "disk_type", Type: cty.String, Required: false},
		"DiskSize":                  &hcldec.AttrSpec{Name: "disk_size", Type: cty.Number, Required: false},
		"DataDisks":                 &hcldec.BlockListSpec{TypeName: "[]tencentCloudDataDisk", Nested: &hcldec.BlockObjectSpec{TypeName: "tencentCloudDataDisk", Nested: hcldec.ObjectSpec((*tencentCloudDataDisk)(nil).HCL2Spec())}},
		"VpcId":                     &hcldec.AttrSpec{Name: "vpc_id", Type: cty.String, Required: false},
		"VpcName":                   &hcldec.AttrSpec{Name: "vpc_name", Type: cty.String, Required: false},
		"VpcIp":                     &hcldec.AttrSpec{Name: "vpc_ip", Type: cty.String, Required: false},
		"SubnetId":                  &hcldec.AttrSpec{Name: "subnet_id", Type: cty.String, Required: false},
		"SubnetName":                &hcldec.AttrSpec{Name: "subnet_name", Type: cty.String, Required: false},
		"CidrBlock":                 &hcldec.AttrSpec{Name: "cidr_block", Type: cty.String, Required: false},
		"SubnectCidrBlock":          &hcldec.AttrSpec{Name: "subnect_cidr_block", Type: cty.String, Required: false},
		"InternetChargeType":        &hcldec.AttrSpec{Name: "internet_charge_type", Type: cty.String, Required: false},
		"InternetMaxBandwidthOut":   &hcldec.AttrSpec{Name: "internet_max_bandwidth_out", Type: cty.Number, Required: false},
		"SecurityGroupId":           &hcldec.AttrSpec{Name: "security_group_id", Type: cty.String, Required: false},
		"SecurityGroupName":         &hcldec.AttrSpec{Name: "security_group_name", Type: cty.String, Required: false},
		"UserData":                  &hcldec.AttrSpec{Name: "user_data", Type: cty.String, Required: false},
		"UserDataFile":              &hcldec.AttrSpec{Name: "user_data_file", Type: cty.String, Required: false},
		"HostName":                  &hcldec.AttrSpec{Name: "host_name", Type: cty.String, Required: false},
		"RunTags":                   &hcldec.BlockAttrsSpec{TypeName: "run_tags", ElementType: cty.String, Required: false},
		"Type":                      &hcldec.AttrSpec{Name: "communicator", Type: cty.String, Required: false},
		"PauseBeforeConnect":        &hcldec.AttrSpec{Name: "pause_before_connecting", Type: cty.String, Required: false},
		"SSHHost":                   &hcldec.AttrSpec{Name: "ssh_host", Type: cty.String, Required: false},
		"SSHPort":                   &hcldec.AttrSpec{Name: "ssh_port", Type: cty.Number, Required: false},
		"SSHUsername":               &hcldec.AttrSpec{Name: "ssh_username", Type: cty.String, Required: false},
		"SSHPassword":               &hcldec.AttrSpec{Name: "ssh_password", Type: cty.String, Required: false},
		"SSHKeyPairName":            &hcldec.AttrSpec{Name: "ssh_keypair_name", Type: cty.String, Required: false},
		"SSHTemporaryKeyPairName":   &hcldec.AttrSpec{Name: "temporary_key_pair_name", Type: cty.String, Required: false},
		"SSHClearAuthorizedKeys":    &hcldec.AttrSpec{Name: "ssh_clear_authorized_keys", Type: cty.Bool, Required: false},
		"SSHPrivateKeyFile":         &hcldec.AttrSpec{Name: "ssh_private_key_file", Type: cty.String, Required: false},
		"SSHPty":                    &hcldec.AttrSpec{Name: "ssh_pty", Type: cty.Bool, Required: false},
		"SSHTimeout":                &hcldec.AttrSpec{Name: "ssh_timeout", Type: cty.String, Required: false},
		"SSHAgentAuth":              &hcldec.AttrSpec{Name: "ssh_agent_auth", Type: cty.Bool, Required: false},
		"SSHDisableAgentForwarding": &hcldec.AttrSpec{Name: "ssh_disable_agent_forwarding", Type: cty.Bool, Required: false},
		"SSHHandshakeAttempts":      &hcldec.AttrSpec{Name: "ssh_handshake_attempts", Type: cty.Number, Required: false},
		"SSHBastionHost":            &hcldec.AttrSpec{Name: "ssh_bastion_host", Type: cty.String, Required: false},
		"SSHBastionPort":            &hcldec.AttrSpec{Name: "ssh_bastion_port", Type: cty.Number, Required: false},
		"SSHBastionAgentAuth":       &hcldec.AttrSpec{Name: "ssh_bastion_agent_auth", Type: cty.Bool, Required: false},
		"SSHBastionUsername":        &hcldec.AttrSpec{Name: "ssh_bastion_username", Type: cty.String, Required: false},
		"SSHBastionPassword":        &hcldec.AttrSpec{Name: "ssh_bastion_password", Type: cty.String, Required: false},
		"SSHBastionPrivateKeyFile":  &hcldec.AttrSpec{Name: "ssh_bastion_private_key_file", Type: cty.String, Required: false},
		"SSHFileTransferMethod":     &hcldec.AttrSpec{Name: "ssh_file_transfer_method", Type: cty.String, Required: false},
		"SSHProxyHost":              &hcldec.AttrSpec{Name: "ssh_proxy_host", Type: cty.String, Required: false},
		"SSHProxyPort":              &hcldec.AttrSpec{Name: "ssh_proxy_port", Type: cty.Number, Required: false},
		"SSHProxyUsername":          &hcldec.AttrSpec{Name: "ssh_proxy_username", Type: cty.String, Required: false},
		"SSHProxyPassword":          &hcldec.AttrSpec{Name: "ssh_proxy_password", Type: cty.String, Required: false},
		"SSHKeepAliveInterval":      &hcldec.AttrSpec{Name: "ssh_keep_alive_interval", Type: cty.String, Required: false},
		"SSHReadWriteTimeout":       &hcldec.AttrSpec{Name: "ssh_read_write_timeout", Type: cty.String, Required: false},
		"SSHRemoteTunnels":          &hcldec.AttrSpec{Name: "ssh_remote_tunnels", Type: cty.List(cty.String), Required: false},
		"SSHLocalTunnels":           &hcldec.AttrSpec{Name: "ssh_local_tunnels", Type: cty.List(cty.String), Required: false},
		"SSHPublicKey":              &hcldec.AttrSpec{Name: "ssh_public_key", Type: cty.List(cty.Number), Required: false},
		"SSHPrivateKey":             &hcldec.AttrSpec{Name: "ssh_private_key", Type: cty.List(cty.Number), Required: false},
		"WinRMUser":                 &hcldec.AttrSpec{Name: "winrm_username", Type: cty.String, Required: false},
		"WinRMPassword":             &hcldec.AttrSpec{Name: "winrm_password", Type: cty.String, Required: false},
		"WinRMHost":                 &hcldec.AttrSpec{Name: "winrm_host", Type: cty.String, Required: false},
		"WinRMPort":                 &hcldec.AttrSpec{Name: "winrm_port", Type: cty.Number, Required: false},
		"WinRMTimeout":              &hcldec.AttrSpec{Name: "winrm_timeout", Type: cty.String, Required: false},
		"WinRMUseSSL":               &hcldec.AttrSpec{Name: "winrm_use_ssl", Type: cty.Bool, Required: false},
		"WinRMInsecure":             &hcldec.AttrSpec{Name: "winrm_insecure", Type: cty.Bool, Required: false},
		"WinRMUseNTLM":              &hcldec.AttrSpec{Name: "winrm_use_ntlm", Type: cty.Bool, Required: false},
		"SSHPrivateIp":              &hcldec.AttrSpec{Name: "ssh_private_ip", Type: cty.Bool, Required: false},
	}
	return s
}
