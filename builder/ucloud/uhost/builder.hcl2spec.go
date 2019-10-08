// Code generated by "mapstructure-to-hcl2 -type Config"; DO NOT EDIT.
package uhost

import (
	"github.com/hashicorp/hcl/v2/hcldec"
	"github.com/zclconf/go-cty/cty"
	"time"
)

type FlatConfig struct {
	PackerBuildName           string             `mapstructure:"packer_build_name" cty:"packer_build_name"`
	PackerBuilderType         string             `mapstructure:"packer_builder_type" cty:"packer_builder_type"`
	PackerDebug               bool               `mapstructure:"packer_debug" cty:"packer_debug"`
	PackerForce               bool               `mapstructure:"packer_force" cty:"packer_force"`
	PackerOnError             string             `mapstructure:"packer_on_error" cty:"packer_on_error"`
	PackerUserVars            map[string]string  `mapstructure:"packer_user_variables" cty:"packer_user_variables"`
	PackerSensitiveVars       []string           `mapstructure:"packer_sensitive_variables" cty:"packer_sensitive_variables"`
	PublicKey                 string             `mapstructure:"public_key" cty:"public_key"`
	PrivateKey                string             `mapstructure:"private_key" cty:"private_key"`
	Region                    string             `mapstructure:"region" cty:"region"`
	ProjectId                 string             `mapstructure:"project_id" cty:"project_id"`
	BaseUrl                   string             `mapstructure:"base_url" cty:"base_url"`
	ImageName                 string             `mapstructure:"image_name" cty:"image_name"`
	ImageDescription          string             `mapstructure:"image_description" cty:"image_description"`
	ImageDestinations         []ImageDestination `mapstructure:"image_copy_to_mappings" cty:"image_copy_to_mappings"`
	Zone                      string             `mapstructure:"availability_zone" cty:"availability_zone"`
	SourceImageId             string             `mapstructure:"source_image_id" cty:"source_image_id"`
	InstanceType              string             `mapstructure:"instance_type" cty:"instance_type"`
	InstanceName              string             `mapstructure:"instance_name" cty:"instance_name"`
	BootDiskType              string             `mapstructure:"boot_disk_type" cty:"boot_disk_type"`
	VPCId                     string             `mapstructure:"vpc_id" cty:"vpc_id"`
	SubnetId                  string             `mapstructure:"subnet_id" cty:"subnet_id"`
	SecurityGroupId           string             `mapstructure:"security_group_id" cty:"security_group_id"`
	Type                      string             `mapstructure:"communicator" cty:"communicator"`
	PauseBeforeConnect        time.Duration      `mapstructure:"pause_before_connecting" cty:"pause_before_connecting"`
	SSHHost                   string             `mapstructure:"ssh_host" cty:"ssh_host"`
	SSHPort                   int                `mapstructure:"ssh_port" cty:"ssh_port"`
	SSHUsername               string             `mapstructure:"ssh_username" cty:"ssh_username"`
	SSHPassword               string             `mapstructure:"ssh_password" cty:"ssh_password"`
	SSHKeyPairName            string             `mapstructure:"ssh_keypair_name" cty:"ssh_keypair_name"`
	SSHTemporaryKeyPairName   string             `mapstructure:"temporary_key_pair_name" cty:"temporary_key_pair_name"`
	SSHClearAuthorizedKeys    bool               `mapstructure:"ssh_clear_authorized_keys" cty:"ssh_clear_authorized_keys"`
	SSHPrivateKeyFile         string             `mapstructure:"ssh_private_key_file" cty:"ssh_private_key_file"`
	SSHPty                    bool               `mapstructure:"ssh_pty" cty:"ssh_pty"`
	SSHTimeout                time.Duration      `mapstructure:"ssh_timeout" cty:"ssh_timeout"`
	SSHAgentAuth              bool               `mapstructure:"ssh_agent_auth" cty:"ssh_agent_auth"`
	SSHDisableAgentForwarding bool               `mapstructure:"ssh_disable_agent_forwarding" cty:"ssh_disable_agent_forwarding"`
	SSHHandshakeAttempts      int                `mapstructure:"ssh_handshake_attempts" cty:"ssh_handshake_attempts"`
	SSHBastionHost            string             `mapstructure:"ssh_bastion_host" cty:"ssh_bastion_host"`
	SSHBastionPort            int                `mapstructure:"ssh_bastion_port" cty:"ssh_bastion_port"`
	SSHBastionAgentAuth       bool               `mapstructure:"ssh_bastion_agent_auth" cty:"ssh_bastion_agent_auth"`
	SSHBastionUsername        string             `mapstructure:"ssh_bastion_username" cty:"ssh_bastion_username"`
	SSHBastionPassword        string             `mapstructure:"ssh_bastion_password" cty:"ssh_bastion_password"`
	SSHBastionPrivateKeyFile  string             `mapstructure:"ssh_bastion_private_key_file" cty:"ssh_bastion_private_key_file"`
	SSHFileTransferMethod     string             `mapstructure:"ssh_file_transfer_method" cty:"ssh_file_transfer_method"`
	SSHProxyHost              string             `mapstructure:"ssh_proxy_host" cty:"ssh_proxy_host"`
	SSHProxyPort              int                `mapstructure:"ssh_proxy_port" cty:"ssh_proxy_port"`
	SSHProxyUsername          string             `mapstructure:"ssh_proxy_username" cty:"ssh_proxy_username"`
	SSHProxyPassword          string             `mapstructure:"ssh_proxy_password" cty:"ssh_proxy_password"`
	SSHKeepAliveInterval      time.Duration      `mapstructure:"ssh_keep_alive_interval" cty:"ssh_keep_alive_interval"`
	SSHReadWriteTimeout       time.Duration      `mapstructure:"ssh_read_write_timeout" cty:"ssh_read_write_timeout"`
	SSHRemoteTunnels          []string           `mapstructure:"ssh_remote_tunnels" cty:"ssh_remote_tunnels"`
	SSHLocalTunnels           []string           `mapstructure:"ssh_local_tunnels" cty:"ssh_local_tunnels"`
	SSHPublicKey              []byte             `cty:"ssh_public_key"`
	SSHPrivateKey             []byte             `cty:"ssh_private_key"`
	WinRMUser                 string             `mapstructure:"winrm_username" cty:"winrm_username"`
	WinRMPassword             string             `mapstructure:"winrm_password" cty:"winrm_password"`
	WinRMHost                 string             `mapstructure:"winrm_host" cty:"winrm_host"`
	WinRMPort                 int                `mapstructure:"winrm_port" cty:"winrm_port"`
	WinRMTimeout              time.Duration      `mapstructure:"winrm_timeout" cty:"winrm_timeout"`
	WinRMUseSSL               bool               `mapstructure:"winrm_use_ssl" cty:"winrm_use_ssl"`
	WinRMInsecure             bool               `mapstructure:"winrm_insecure" cty:"winrm_insecure"`
	WinRMUseNTLM              bool               `mapstructure:"winrm_use_ntlm" cty:"winrm_use_ntlm"`
	UseSSHPrivateIp           bool               `mapstructure:"use_ssh_private_ip" cty:"use_ssh_private_ip"`
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
		"PublicKey":                 &hcldec.AttrSpec{Name: "public_key", Type: cty.String, Required: false},
		"PrivateKey":                &hcldec.AttrSpec{Name: "private_key", Type: cty.String, Required: false},
		"Region":                    &hcldec.AttrSpec{Name: "region", Type: cty.String, Required: false},
		"ProjectId":                 &hcldec.AttrSpec{Name: "project_id", Type: cty.String, Required: false},
		"BaseUrl":                   &hcldec.AttrSpec{Name: "base_url", Type: cty.String, Required: false},
		"ImageName":                 &hcldec.AttrSpec{Name: "image_name", Type: cty.String, Required: false},
		"ImageDescription":          &hcldec.AttrSpec{Name: "image_description", Type: cty.String, Required: false},
		"ImageDestinations":         hcldec.BlockListSpec{TypeName: "[]ImageDestination", Nested: &hcldec.BlockObjectSpec{TypeName: "ImageDestination", Nested: hcldec.ObjectSpec((*ImageDestination)(nil).HCL2Spec())}},
		"Zone":                      &hcldec.AttrSpec{Name: "availability_zone", Type: cty.String, Required: false},
		"SourceImageId":             &hcldec.AttrSpec{Name: "source_image_id", Type: cty.String, Required: false},
		"InstanceType":              &hcldec.AttrSpec{Name: "instance_type", Type: cty.String, Required: false},
		"InstanceName":              &hcldec.AttrSpec{Name: "instance_name", Type: cty.String, Required: false},
		"BootDiskType":              &hcldec.AttrSpec{Name: "boot_disk_type", Type: cty.String, Required: false},
		"VPCId":                     &hcldec.AttrSpec{Name: "vpc_id", Type: cty.String, Required: false},
		"SubnetId":                  &hcldec.AttrSpec{Name: "subnet_id", Type: cty.String, Required: false},
		"SecurityGroupId":           &hcldec.AttrSpec{Name: "security_group_id", Type: cty.String, Required: false},
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
		"UseSSHPrivateIp":           &hcldec.AttrSpec{Name: "use_ssh_private_ip", Type: cty.Bool, Required: false},
	}
	return s
}
