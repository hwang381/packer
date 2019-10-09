// Code generated by "mapstructure-to-hcl2 -type Config"; DO NOT EDIT.
package linode

import (
	"github.com/hashicorp/hcl/v2/hcldec"
	"github.com/zclconf/go-cty/cty"
	"time"
)

// FlatConfig is an auto-generated flat version of Config.
// Where the contents of a field with a `mapstructure:,squash` tag are bubbled up.
type FlatConfig struct {
	PackerBuildName           string            `mapstructure:"packer_build_name" cty:"packer_build_name"`
	PackerBuilderType         string            `mapstructure:"packer_builder_type" cty:"packer_builder_type"`
	PackerDebug               bool              `mapstructure:"packer_debug" cty:"packer_debug"`
	PackerForce               bool              `mapstructure:"packer_force" cty:"packer_force"`
	PackerOnError             string            `mapstructure:"packer_on_error" cty:"packer_on_error"`
	PackerUserVars            map[string]string `mapstructure:"packer_user_variables" cty:"packer_user_variables"`
	PackerSensitiveVars       []string          `mapstructure:"packer_sensitive_variables" cty:"packer_sensitive_variables"`
	Type                      string            `mapstructure:"communicator" cty:"communicator"`
	PauseBeforeConnect        time.Duration     `mapstructure:"pause_before_connecting" cty:"pause_before_connecting"`
	SSHHost                   string            `mapstructure:"ssh_host" cty:"ssh_host"`
	SSHPort                   int               `mapstructure:"ssh_port" cty:"ssh_port"`
	SSHUsername               string            `mapstructure:"ssh_username" cty:"ssh_username"`
	SSHPassword               string            `mapstructure:"ssh_password" cty:"ssh_password"`
	SSHKeyPairName            string            `mapstructure:"ssh_keypair_name" cty:"ssh_keypair_name"`
	SSHTemporaryKeyPairName   string            `mapstructure:"temporary_key_pair_name" cty:"temporary_key_pair_name"`
	SSHClearAuthorizedKeys    bool              `mapstructure:"ssh_clear_authorized_keys" cty:"ssh_clear_authorized_keys"`
	SSHPrivateKeyFile         string            `mapstructure:"ssh_private_key_file" cty:"ssh_private_key_file"`
	SSHPty                    bool              `mapstructure:"ssh_pty" cty:"ssh_pty"`
	SSHTimeout                time.Duration     `mapstructure:"ssh_timeout" cty:"ssh_timeout"`
	SSHAgentAuth              bool              `mapstructure:"ssh_agent_auth" cty:"ssh_agent_auth"`
	SSHDisableAgentForwarding bool              `mapstructure:"ssh_disable_agent_forwarding" cty:"ssh_disable_agent_forwarding"`
	SSHHandshakeAttempts      int               `mapstructure:"ssh_handshake_attempts" cty:"ssh_handshake_attempts"`
	SSHBastionHost            string            `mapstructure:"ssh_bastion_host" cty:"ssh_bastion_host"`
	SSHBastionPort            int               `mapstructure:"ssh_bastion_port" cty:"ssh_bastion_port"`
	SSHBastionAgentAuth       bool              `mapstructure:"ssh_bastion_agent_auth" cty:"ssh_bastion_agent_auth"`
	SSHBastionUsername        string            `mapstructure:"ssh_bastion_username" cty:"ssh_bastion_username"`
	SSHBastionPassword        string            `mapstructure:"ssh_bastion_password" cty:"ssh_bastion_password"`
	SSHBastionPrivateKeyFile  string            `mapstructure:"ssh_bastion_private_key_file" cty:"ssh_bastion_private_key_file"`
	SSHFileTransferMethod     string            `mapstructure:"ssh_file_transfer_method" cty:"ssh_file_transfer_method"`
	SSHProxyHost              string            `mapstructure:"ssh_proxy_host" cty:"ssh_proxy_host"`
	SSHProxyPort              int               `mapstructure:"ssh_proxy_port" cty:"ssh_proxy_port"`
	SSHProxyUsername          string            `mapstructure:"ssh_proxy_username" cty:"ssh_proxy_username"`
	SSHProxyPassword          string            `mapstructure:"ssh_proxy_password" cty:"ssh_proxy_password"`
	SSHKeepAliveInterval      time.Duration     `mapstructure:"ssh_keep_alive_interval" cty:"ssh_keep_alive_interval"`
	SSHReadWriteTimeout       time.Duration     `mapstructure:"ssh_read_write_timeout" cty:"ssh_read_write_timeout"`
	SSHRemoteTunnels          []string          `mapstructure:"ssh_remote_tunnels" cty:"ssh_remote_tunnels"`
	SSHLocalTunnels           []string          `mapstructure:"ssh_local_tunnels" cty:"ssh_local_tunnels"`
	SSHPublicKey              []byte            `cty:"ssh_public_key"`
	SSHPrivateKey             []byte            `cty:"ssh_private_key"`
	WinRMUser                 string            `mapstructure:"winrm_username" cty:"winrm_username"`
	WinRMPassword             string            `mapstructure:"winrm_password" cty:"winrm_password"`
	WinRMHost                 string            `mapstructure:"winrm_host" cty:"winrm_host"`
	WinRMPort                 int               `mapstructure:"winrm_port" cty:"winrm_port"`
	WinRMTimeout              time.Duration     `mapstructure:"winrm_timeout" cty:"winrm_timeout"`
	WinRMUseSSL               bool              `mapstructure:"winrm_use_ssl" cty:"winrm_use_ssl"`
	WinRMInsecure             bool              `mapstructure:"winrm_insecure" cty:"winrm_insecure"`
	WinRMUseNTLM              bool              `mapstructure:"winrm_use_ntlm" cty:"winrm_use_ntlm"`
	PersonalAccessToken       string            `mapstructure:"linode_token" cty:"linode_token"`
	Region                    string            `mapstructure:"region" cty:"region"`
	InstanceType              string            `mapstructure:"instance_type" cty:"instance_type"`
	Label                     string            `mapstructure:"instance_label" cty:"instance_label"`
	Tags                      []string          `mapstructure:"instance_tags" cty:"instance_tags"`
	Image                     string            `mapstructure:"image" cty:"image"`
	SwapSize                  int               `mapstructure:"swap_size" cty:"swap_size"`
	RootPass                  string            `mapstructure:"root_pass" cty:"root_pass"`
	RootSSHKey                string            `mapstructure:"root_ssh_key" cty:"root_ssh_key"`
	ImageLabel                string            `mapstructure:"image_label" cty:"image_label"`
	Description               string            `mapstructure:"image_description" cty:"image_description"`
	RawStateTimeout           string            `mapstructure:"state_timeout" cty:"state_timeout"`
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
		"PersonalAccessToken":       &hcldec.AttrSpec{Name: "linode_token", Type: cty.String, Required: false},
		"Region":                    &hcldec.AttrSpec{Name: "region", Type: cty.String, Required: false},
		"InstanceType":              &hcldec.AttrSpec{Name: "instance_type", Type: cty.String, Required: false},
		"Label":                     &hcldec.AttrSpec{Name: "instance_label", Type: cty.String, Required: false},
		"Tags":                      &hcldec.AttrSpec{Name: "instance_tags", Type: cty.List(cty.String), Required: false},
		"Image":                     &hcldec.AttrSpec{Name: "image", Type: cty.String, Required: false},
		"SwapSize":                  &hcldec.AttrSpec{Name: "swap_size", Type: cty.Number, Required: false},
		"RootPass":                  &hcldec.AttrSpec{Name: "root_pass", Type: cty.String, Required: false},
		"RootSSHKey":                &hcldec.AttrSpec{Name: "root_ssh_key", Type: cty.String, Required: false},
		"ImageLabel":                &hcldec.AttrSpec{Name: "image_label", Type: cty.String, Required: false},
		"Description":               &hcldec.AttrSpec{Name: "image_description", Type: cty.String, Required: false},
		"RawStateTimeout":           &hcldec.AttrSpec{Name: "state_timeout", Type: cty.String, Required: false},
	}
	return s
}
