// Code generated by "mapstructure-to-hcl2 -type Config"; DO NOT EDIT.
package docker

import (
	"github.com/hashicorp/hcl/v2/hcldec"
	"github.com/zclconf/go-cty/cty"
	"time"
)

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
	Author                    string            `mapstructure:"author" cty:"author"`
	Changes                   []string          `mapstructure:"changes" cty:"changes"`
	Commit                    bool              `mapstructure:"commit" required:"true" cty:"commit"`
	ContainerDir              string            `mapstructure:"container_dir" required:"false" cty:"container_dir"`
	Discard                   bool              `mapstructure:"discard" required:"true" cty:"discard"`
	ExecUser                  string            `mapstructure:"exec_user" required:"false" cty:"exec_user"`
	ExportPath                string            `mapstructure:"export_path" required:"true" cty:"export_path"`
	Image                     string            `mapstructure:"image" required:"true" cty:"image"`
	Message                   string            `mapstructure:"message" required:"true" cty:"message"`
	Privileged                bool              `mapstructure:"privileged" required:"false" cty:"privileged"`
	Pty                       bool              `cty:"pty"`
	Pull                      bool              `mapstructure:"pull" required:"false" cty:"pull"`
	RunCommand                []string          `mapstructure:"run_command" required:"false" cty:"run_command"`
	Volumes                   map[string]string `mapstructure:"volumes" required:"false" cty:"volumes"`
	FixUploadOwner            bool              `mapstructure:"fix_upload_owner" required:"false" cty:"fix_upload_owner"`
	WindowsContainer          bool              `mapstructure:"windows_container" required:"false" cty:"windows_container"`
	Login                     bool              `mapstructure:"login" required:"false" cty:"login"`
	LoginPassword             string            `mapstructure:"login_password" required:"false" cty:"login_password"`
	LoginServer               string            `mapstructure:"login_server" required:"false" cty:"login_server"`
	LoginUsername             string            `mapstructure:"login_username" required:"false" cty:"login_username"`
	EcrLogin                  bool              `mapstructure:"ecr_login" required:"false" cty:"ecr_login"`
	AccessKey                 string            `mapstructure:"aws_access_key" required:"false" cty:"aws_access_key"`
	SecretKey                 string            `mapstructure:"aws_secret_key" required:"false" cty:"aws_secret_key"`
	Token                     string            `mapstructure:"aws_token" required:"false" cty:"aws_token"`
	Profile                   string            `mapstructure:"aws_profile" required:"false" cty:"aws_profile"`
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
		"Author":                    &hcldec.AttrSpec{Name: "author", Type: cty.String, Required: false},
		"Changes":                   &hcldec.AttrSpec{Name: "changes", Type: cty.List(cty.String), Required: false},
		"Commit":                    &hcldec.AttrSpec{Name: "commit", Type: cty.Bool, Required: false},
		"ContainerDir":              &hcldec.AttrSpec{Name: "container_dir", Type: cty.String, Required: false},
		"Discard":                   &hcldec.AttrSpec{Name: "discard", Type: cty.Bool, Required: false},
		"ExecUser":                  &hcldec.AttrSpec{Name: "exec_user", Type: cty.String, Required: false},
		"ExportPath":                &hcldec.AttrSpec{Name: "export_path", Type: cty.String, Required: false},
		"Image":                     &hcldec.AttrSpec{Name: "image", Type: cty.String, Required: false},
		"Message":                   &hcldec.AttrSpec{Name: "message", Type: cty.String, Required: false},
		"Privileged":                &hcldec.AttrSpec{Name: "privileged", Type: cty.Bool, Required: false},
		"Pty":                       &hcldec.AttrSpec{Name: "pty", Type: cty.Bool, Required: false},
		"Pull":                      &hcldec.AttrSpec{Name: "pull", Type: cty.Bool, Required: false},
		"RunCommand":                &hcldec.AttrSpec{Name: "run_command", Type: cty.List(cty.String), Required: false},
		"Volumes":                   &hcldec.BlockAttrsSpec{TypeName: "volumes", ElementType: cty.String, Required: false},
		"FixUploadOwner":            &hcldec.AttrSpec{Name: "fix_upload_owner", Type: cty.Bool, Required: false},
		"WindowsContainer":          &hcldec.AttrSpec{Name: "windows_container", Type: cty.Bool, Required: false},
		"Login":                     &hcldec.AttrSpec{Name: "login", Type: cty.Bool, Required: false},
		"LoginPassword":             &hcldec.AttrSpec{Name: "login_password", Type: cty.String, Required: false},
		"LoginServer":               &hcldec.AttrSpec{Name: "login_server", Type: cty.String, Required: false},
		"LoginUsername":             &hcldec.AttrSpec{Name: "login_username", Type: cty.String, Required: false},
		"EcrLogin":                  &hcldec.AttrSpec{Name: "ecr_login", Type: cty.Bool, Required: false},
		"AccessKey":                 &hcldec.AttrSpec{Name: "aws_access_key", Type: cty.String, Required: false},
		"SecretKey":                 &hcldec.AttrSpec{Name: "aws_secret_key", Type: cty.String, Required: false},
		"Token":                     &hcldec.AttrSpec{Name: "aws_token", Type: cty.String, Required: false},
		"Profile":                   &hcldec.AttrSpec{Name: "aws_profile", Type: cty.String, Required: false},
	}
	return s
}
