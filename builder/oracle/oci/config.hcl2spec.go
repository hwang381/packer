// Code generated by "mapstructure-to-hcl2 -type Config"; DO NOT EDIT.
package oci

import (
	"github.com/hashicorp/hcl/v2/hcldec"
	"github.com/zclconf/go-cty/cty"
	"time"
)

type FlatConfig struct {
	PackerBuildName           string                            `mapstructure:"packer_build_name" cty:"packer_build_name"`
	PackerBuilderType         string                            `mapstructure:"packer_builder_type" cty:"packer_builder_type"`
	PackerDebug               bool                              `mapstructure:"packer_debug" cty:"packer_debug"`
	PackerForce               bool                              `mapstructure:"packer_force" cty:"packer_force"`
	PackerOnError             string                            `mapstructure:"packer_on_error" cty:"packer_on_error"`
	PackerUserVars            map[string]string                 `mapstructure:"packer_user_variables" cty:"packer_user_variables"`
	PackerSensitiveVars       []string                          `mapstructure:"packer_sensitive_variables" cty:"packer_sensitive_variables"`
	Type                      string                            `mapstructure:"communicator" cty:"communicator"`
	PauseBeforeConnect        time.Duration                     `mapstructure:"pause_before_connecting" cty:"pause_before_connecting"`
	SSHHost                   string                            `mapstructure:"ssh_host" cty:"ssh_host"`
	SSHPort                   int                               `mapstructure:"ssh_port" cty:"ssh_port"`
	SSHUsername               string                            `mapstructure:"ssh_username" cty:"ssh_username"`
	SSHPassword               string                            `mapstructure:"ssh_password" cty:"ssh_password"`
	SSHKeyPairName            string                            `mapstructure:"ssh_keypair_name" cty:"ssh_keypair_name"`
	SSHTemporaryKeyPairName   string                            `mapstructure:"temporary_key_pair_name" cty:"temporary_key_pair_name"`
	SSHClearAuthorizedKeys    bool                              `mapstructure:"ssh_clear_authorized_keys" cty:"ssh_clear_authorized_keys"`
	SSHPrivateKeyFile         string                            `mapstructure:"ssh_private_key_file" cty:"ssh_private_key_file"`
	SSHPty                    bool                              `mapstructure:"ssh_pty" cty:"ssh_pty"`
	SSHTimeout                time.Duration                     `mapstructure:"ssh_timeout" cty:"ssh_timeout"`
	SSHAgentAuth              bool                              `mapstructure:"ssh_agent_auth" cty:"ssh_agent_auth"`
	SSHDisableAgentForwarding bool                              `mapstructure:"ssh_disable_agent_forwarding" cty:"ssh_disable_agent_forwarding"`
	SSHHandshakeAttempts      int                               `mapstructure:"ssh_handshake_attempts" cty:"ssh_handshake_attempts"`
	SSHBastionHost            string                            `mapstructure:"ssh_bastion_host" cty:"ssh_bastion_host"`
	SSHBastionPort            int                               `mapstructure:"ssh_bastion_port" cty:"ssh_bastion_port"`
	SSHBastionAgentAuth       bool                              `mapstructure:"ssh_bastion_agent_auth" cty:"ssh_bastion_agent_auth"`
	SSHBastionUsername        string                            `mapstructure:"ssh_bastion_username" cty:"ssh_bastion_username"`
	SSHBastionPassword        string                            `mapstructure:"ssh_bastion_password" cty:"ssh_bastion_password"`
	SSHBastionPrivateKeyFile  string                            `mapstructure:"ssh_bastion_private_key_file" cty:"ssh_bastion_private_key_file"`
	SSHFileTransferMethod     string                            `mapstructure:"ssh_file_transfer_method" cty:"ssh_file_transfer_method"`
	SSHProxyHost              string                            `mapstructure:"ssh_proxy_host" cty:"ssh_proxy_host"`
	SSHProxyPort              int                               `mapstructure:"ssh_proxy_port" cty:"ssh_proxy_port"`
	SSHProxyUsername          string                            `mapstructure:"ssh_proxy_username" cty:"ssh_proxy_username"`
	SSHProxyPassword          string                            `mapstructure:"ssh_proxy_password" cty:"ssh_proxy_password"`
	SSHKeepAliveInterval      time.Duration                     `mapstructure:"ssh_keep_alive_interval" cty:"ssh_keep_alive_interval"`
	SSHReadWriteTimeout       time.Duration                     `mapstructure:"ssh_read_write_timeout" cty:"ssh_read_write_timeout"`
	SSHRemoteTunnels          []string                          `mapstructure:"ssh_remote_tunnels" cty:"ssh_remote_tunnels"`
	SSHLocalTunnels           []string                          `mapstructure:"ssh_local_tunnels" cty:"ssh_local_tunnels"`
	SSHPublicKey              []byte                            `cty:"ssh_public_key"`
	SSHPrivateKey             []byte                            `cty:"ssh_private_key"`
	WinRMUser                 string                            `mapstructure:"winrm_username" cty:"winrm_username"`
	WinRMPassword             string                            `mapstructure:"winrm_password" cty:"winrm_password"`
	WinRMHost                 string                            `mapstructure:"winrm_host" cty:"winrm_host"`
	WinRMPort                 int                               `mapstructure:"winrm_port" cty:"winrm_port"`
	WinRMTimeout              time.Duration                     `mapstructure:"winrm_timeout" cty:"winrm_timeout"`
	WinRMUseSSL               bool                              `mapstructure:"winrm_use_ssl" cty:"winrm_use_ssl"`
	WinRMInsecure             bool                              `mapstructure:"winrm_insecure" cty:"winrm_insecure"`
	WinRMUseNTLM              bool                              `mapstructure:"winrm_use_ntlm" cty:"winrm_use_ntlm"`
	AccessCfgFile             string                            `mapstructure:"access_cfg_file" cty:"access_cfg_file"`
	AccessCfgFileAccount      string                            `mapstructure:"access_cfg_file_account" cty:"access_cfg_file_account"`
	UserID                    string                            `mapstructure:"user_ocid" cty:"user_ocid"`
	TenancyID                 string                            `mapstructure:"tenancy_ocid" cty:"tenancy_ocid"`
	Region                    string                            `mapstructure:"region" cty:"region"`
	Fingerprint               string                            `mapstructure:"fingerprint" cty:"fingerprint"`
	KeyFile                   string                            `mapstructure:"key_file" cty:"key_file"`
	PassPhrase                string                            `mapstructure:"pass_phrase" cty:"pass_phrase"`
	UsePrivateIP              bool                              `mapstructure:"use_private_ip" cty:"use_private_ip"`
	AvailabilityDomain        string                            `mapstructure:"availability_domain" cty:"availability_domain"`
	CompartmentID             string                            `mapstructure:"compartment_ocid" cty:"compartment_ocid"`
	BaseImageID               string                            `mapstructure:"base_image_ocid" cty:"base_image_ocid"`
	Shape                     string                            `mapstructure:"shape" cty:"shape"`
	ImageName                 string                            `mapstructure:"image_name" cty:"image_name"`
	InstanceName              string                            `mapstructure:"instance_name" cty:"instance_name"`
	Metadata                  map[string]string                 `mapstructure:"metadata" cty:"metadata"`
	UserData                  string                            `mapstructure:"user_data" cty:"user_data"`
	UserDataFile              string                            `mapstructure:"user_data_file" cty:"user_data_file"`
	SubnetID                  string                            `mapstructure:"subnet_ocid" cty:"subnet_ocid"`
	Tags                      map[string]string                 `mapstructure:"tags" cty:"tags"`
	DefinedTags               map[string]map[string]interface{} `mapstructure:"defined_tags" cty:"defined_tags"`
}

func (*Config) HCL2Spec() map[string]hcldec.Spec {
	s := map[string]hcldec.Spec{
		"PackerBuildName":           &hcldec.AttrSpec{Name: "packer_build_name", Type: cty.String, Required: false},
		"PackerBuilderType":         &hcldec.AttrSpec{Name: "packer_builder_type", Type: cty.String, Required: false},
		"PackerDebug":               &hcldec.AttrSpec{Name: "packer_debug", Type: cty.Bool, Required: false},
		"PackerForce":               &hcldec.AttrSpec{Name: "packer_force", Type: cty.Bool, Required: false},
		"PackerOnError":             &hcldec.AttrSpec{Name: "packer_on_error", Type: cty.String, Required: false},
		"PackerUserVars":            &hcldec.BlockAttrsSpec{TypeName: "packer_user_variables", ElementType: cty.String, Required: false},
		"PackerSensitiveVars":       &hcldec.AttrSpec{Name: "packer_sensitive_variables", Type: cty.List(cty.String), Required: false},
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
		"AccessCfgFile":             &hcldec.AttrSpec{Name: "access_cfg_file", Type: cty.String, Required: false},
		"AccessCfgFileAccount":      &hcldec.AttrSpec{Name: "access_cfg_file_account", Type: cty.String, Required: false},
		"UserID":                    &hcldec.AttrSpec{Name: "user_ocid", Type: cty.String, Required: false},
		"TenancyID":                 &hcldec.AttrSpec{Name: "tenancy_ocid", Type: cty.String, Required: false},
		"Region":                    &hcldec.AttrSpec{Name: "region", Type: cty.String, Required: false},
		"Fingerprint":               &hcldec.AttrSpec{Name: "fingerprint", Type: cty.String, Required: false},
		"KeyFile":                   &hcldec.AttrSpec{Name: "key_file", Type: cty.String, Required: false},
		"PassPhrase":                &hcldec.AttrSpec{Name: "pass_phrase", Type: cty.String, Required: false},
		"UsePrivateIP":              &hcldec.AttrSpec{Name: "use_private_ip", Type: cty.Bool, Required: false},
		"AvailabilityDomain":        &hcldec.AttrSpec{Name: "availability_domain", Type: cty.String, Required: false},
		"CompartmentID":             &hcldec.AttrSpec{Name: "compartment_ocid", Type: cty.String, Required: false},
		"BaseImageID":               &hcldec.AttrSpec{Name: "base_image_ocid", Type: cty.String, Required: false},
		"Shape":                     &hcldec.AttrSpec{Name: "shape", Type: cty.String, Required: false},
		"ImageName":                 &hcldec.AttrSpec{Name: "image_name", Type: cty.String, Required: false},
		"InstanceName":              &hcldec.AttrSpec{Name: "instance_name", Type: cty.String, Required: false},
		"Metadata":                  &hcldec.BlockAttrsSpec{TypeName: "metadata", ElementType: cty.String, Required: false},
		"UserData":                  &hcldec.AttrSpec{Name: "user_data", Type: cty.String, Required: false},
		"UserDataFile":              &hcldec.AttrSpec{Name: "user_data_file", Type: cty.String, Required: false},
		"SubnetID":                  &hcldec.AttrSpec{Name: "subnet_ocid", Type: cty.String, Required: false},
		"Tags":                      &hcldec.BlockAttrsSpec{TypeName: "tags", ElementType: cty.String, Required: false},
		"DefinedTags":               &hcldec.BlockAttrsSpec{TypeName: "defined_tags", ElementType: cty.String, Required: false},
	}
	return s
}
