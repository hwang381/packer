// Code generated by "mapstructure-to-hcl2 -type SSH,WinRM"; DO NOT EDIT.
package communicator

import (
	"github.com/hashicorp/hcl/v2/hcldec"
	"github.com/zclconf/go-cty/cty"
	"time"
)

type FlatSSH struct {
	SSHHost                   string        `mapstructure:"ssh_host" cty:"ssh_host"`
	SSHPort                   int           `mapstructure:"ssh_port" cty:"ssh_port"`
	SSHUsername               string        `mapstructure:"ssh_username" cty:"ssh_username"`
	SSHPassword               string        `mapstructure:"ssh_password" cty:"ssh_password"`
	SSHKeyPairName            string        `mapstructure:"ssh_keypair_name" cty:"ssh_keypair_name"`
	SSHTemporaryKeyPairName   string        `mapstructure:"temporary_key_pair_name" cty:"temporary_key_pair_name"`
	SSHClearAuthorizedKeys    bool          `mapstructure:"ssh_clear_authorized_keys" cty:"ssh_clear_authorized_keys"`
	SSHPrivateKeyFile         string        `mapstructure:"ssh_private_key_file" cty:"ssh_private_key_file"`
	SSHPty                    bool          `mapstructure:"ssh_pty" cty:"ssh_pty"`
	SSHTimeout                time.Duration `mapstructure:"ssh_timeout" cty:"ssh_timeout"`
	SSHAgentAuth              bool          `mapstructure:"ssh_agent_auth" cty:"ssh_agent_auth"`
	SSHDisableAgentForwarding bool          `mapstructure:"ssh_disable_agent_forwarding" cty:"ssh_disable_agent_forwarding"`
	SSHHandshakeAttempts      int           `mapstructure:"ssh_handshake_attempts" cty:"ssh_handshake_attempts"`
	SSHBastionHost            string        `mapstructure:"ssh_bastion_host" cty:"ssh_bastion_host"`
	SSHBastionPort            int           `mapstructure:"ssh_bastion_port" cty:"ssh_bastion_port"`
	SSHBastionAgentAuth       bool          `mapstructure:"ssh_bastion_agent_auth" cty:"ssh_bastion_agent_auth"`
	SSHBastionUsername        string        `mapstructure:"ssh_bastion_username" cty:"ssh_bastion_username"`
	SSHBastionPassword        string        `mapstructure:"ssh_bastion_password" cty:"ssh_bastion_password"`
	SSHBastionPrivateKeyFile  string        `mapstructure:"ssh_bastion_private_key_file" cty:"ssh_bastion_private_key_file"`
	SSHFileTransferMethod     string        `mapstructure:"ssh_file_transfer_method" cty:"ssh_file_transfer_method"`
	SSHProxyHost              string        `mapstructure:"ssh_proxy_host" cty:"ssh_proxy_host"`
	SSHProxyPort              int           `mapstructure:"ssh_proxy_port" cty:"ssh_proxy_port"`
	SSHProxyUsername          string        `mapstructure:"ssh_proxy_username" cty:"ssh_proxy_username"`
	SSHProxyPassword          string        `mapstructure:"ssh_proxy_password" cty:"ssh_proxy_password"`
	SSHKeepAliveInterval      time.Duration `mapstructure:"ssh_keep_alive_interval" cty:"ssh_keep_alive_interval"`
	SSHReadWriteTimeout       time.Duration `mapstructure:"ssh_read_write_timeout" cty:"ssh_read_write_timeout"`
	SSHRemoteTunnels          []string      `mapstructure:"ssh_remote_tunnels" cty:"ssh_remote_tunnels"`
	SSHLocalTunnels           []string      `mapstructure:"ssh_local_tunnels" cty:"ssh_local_tunnels"`
	SSHPublicKey              []byte        `cty:"ssh_public_key"`
	SSHPrivateKey             []byte        `cty:"ssh_private_key"`
}

func (*SSH) HCL2Spec() map[string]hcldec.Spec {
	s := map[string]hcldec.Spec{
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
	}
	return s
}

type FlatWinRM struct {
	WinRMUser     string        `mapstructure:"winrm_username" cty:"winrm_username"`
	WinRMPassword string        `mapstructure:"winrm_password" cty:"winrm_password"`
	WinRMHost     string        `mapstructure:"winrm_host" cty:"winrm_host"`
	WinRMPort     int           `mapstructure:"winrm_port" cty:"winrm_port"`
	WinRMTimeout  time.Duration `mapstructure:"winrm_timeout" cty:"winrm_timeout"`
	WinRMUseSSL   bool          `mapstructure:"winrm_use_ssl" cty:"winrm_use_ssl"`
	WinRMInsecure bool          `mapstructure:"winrm_insecure" cty:"winrm_insecure"`
	WinRMUseNTLM  bool          `mapstructure:"winrm_use_ntlm" cty:"winrm_use_ntlm"`
}

func (*WinRM) HCL2Spec() map[string]hcldec.Spec {
	s := map[string]hcldec.Spec{
		"WinRMUser":     &hcldec.AttrSpec{Name: "winrm_username", Type: cty.String, Required: false},
		"WinRMPassword": &hcldec.AttrSpec{Name: "winrm_password", Type: cty.String, Required: false},
		"WinRMHost":     &hcldec.AttrSpec{Name: "winrm_host", Type: cty.String, Required: false},
		"WinRMPort":     &hcldec.AttrSpec{Name: "winrm_port", Type: cty.Number, Required: false},
		"WinRMTimeout":  &hcldec.AttrSpec{Name: "winrm_timeout", Type: cty.String, Required: false},
		"WinRMUseSSL":   &hcldec.AttrSpec{Name: "winrm_use_ssl", Type: cty.Bool, Required: false},
		"WinRMInsecure": &hcldec.AttrSpec{Name: "winrm_insecure", Type: cty.Bool, Required: false},
		"WinRMUseNTLM":  &hcldec.AttrSpec{Name: "winrm_use_ntlm", Type: cty.Bool, Required: false},
	}
	return s
}

type FlatSSH struct {
	SSHHost                   string        `mapstructure:"ssh_host" cty:"ssh_host"`
	SSHPort                   int           `mapstructure:"ssh_port" cty:"ssh_port"`
	SSHUsername               string        `mapstructure:"ssh_username" cty:"ssh_username"`
	SSHPassword               string        `mapstructure:"ssh_password" cty:"ssh_password"`
	SSHKeyPairName            string        `mapstructure:"ssh_keypair_name" cty:"ssh_keypair_name"`
	SSHTemporaryKeyPairName   string        `mapstructure:"temporary_key_pair_name" cty:"temporary_key_pair_name"`
	SSHClearAuthorizedKeys    bool          `mapstructure:"ssh_clear_authorized_keys" cty:"ssh_clear_authorized_keys"`
	SSHPrivateKeyFile         string        `mapstructure:"ssh_private_key_file" cty:"ssh_private_key_file"`
	SSHPty                    bool          `mapstructure:"ssh_pty" cty:"ssh_pty"`
	SSHTimeout                time.Duration `mapstructure:"ssh_timeout" cty:"ssh_timeout"`
	SSHAgentAuth              bool          `mapstructure:"ssh_agent_auth" cty:"ssh_agent_auth"`
	SSHDisableAgentForwarding bool          `mapstructure:"ssh_disable_agent_forwarding" cty:"ssh_disable_agent_forwarding"`
	SSHHandshakeAttempts      int           `mapstructure:"ssh_handshake_attempts" cty:"ssh_handshake_attempts"`
	SSHBastionHost            string        `mapstructure:"ssh_bastion_host" cty:"ssh_bastion_host"`
	SSHBastionPort            int           `mapstructure:"ssh_bastion_port" cty:"ssh_bastion_port"`
	SSHBastionAgentAuth       bool          `mapstructure:"ssh_bastion_agent_auth" cty:"ssh_bastion_agent_auth"`
	SSHBastionUsername        string        `mapstructure:"ssh_bastion_username" cty:"ssh_bastion_username"`
	SSHBastionPassword        string        `mapstructure:"ssh_bastion_password" cty:"ssh_bastion_password"`
	SSHBastionPrivateKeyFile  string        `mapstructure:"ssh_bastion_private_key_file" cty:"ssh_bastion_private_key_file"`
	SSHFileTransferMethod     string        `mapstructure:"ssh_file_transfer_method" cty:"ssh_file_transfer_method"`
	SSHProxyHost              string        `mapstructure:"ssh_proxy_host" cty:"ssh_proxy_host"`
	SSHProxyPort              int           `mapstructure:"ssh_proxy_port" cty:"ssh_proxy_port"`
	SSHProxyUsername          string        `mapstructure:"ssh_proxy_username" cty:"ssh_proxy_username"`
	SSHProxyPassword          string        `mapstructure:"ssh_proxy_password" cty:"ssh_proxy_password"`
	SSHKeepAliveInterval      time.Duration `mapstructure:"ssh_keep_alive_interval" cty:"ssh_keep_alive_interval"`
	SSHReadWriteTimeout       time.Duration `mapstructure:"ssh_read_write_timeout" cty:"ssh_read_write_timeout"`
	SSHRemoteTunnels          []string      `mapstructure:"ssh_remote_tunnels" cty:"ssh_remote_tunnels"`
	SSHLocalTunnels           []string      `mapstructure:"ssh_local_tunnels" cty:"ssh_local_tunnels"`
	SSHPublicKey              []byte        `cty:"ssh_public_key"`
	SSHPrivateKey             []byte        `cty:"ssh_private_key"`
}

func (*SSH) HCL2Spec() map[string]hcldec.Spec {
	s := map[string]hcldec.Spec{
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
	}
	return s
}

type FlatWinRM struct {
	WinRMUser     string        `mapstructure:"winrm_username" cty:"winrm_username"`
	WinRMPassword string        `mapstructure:"winrm_password" cty:"winrm_password"`
	WinRMHost     string        `mapstructure:"winrm_host" cty:"winrm_host"`
	WinRMPort     int           `mapstructure:"winrm_port" cty:"winrm_port"`
	WinRMTimeout  time.Duration `mapstructure:"winrm_timeout" cty:"winrm_timeout"`
	WinRMUseSSL   bool          `mapstructure:"winrm_use_ssl" cty:"winrm_use_ssl"`
	WinRMInsecure bool          `mapstructure:"winrm_insecure" cty:"winrm_insecure"`
	WinRMUseNTLM  bool          `mapstructure:"winrm_use_ntlm" cty:"winrm_use_ntlm"`
}

func (*WinRM) HCL2Spec() map[string]hcldec.Spec {
	s := map[string]hcldec.Spec{
		"WinRMUser":     &hcldec.AttrSpec{Name: "winrm_username", Type: cty.String, Required: false},
		"WinRMPassword": &hcldec.AttrSpec{Name: "winrm_password", Type: cty.String, Required: false},
		"WinRMHost":     &hcldec.AttrSpec{Name: "winrm_host", Type: cty.String, Required: false},
		"WinRMPort":     &hcldec.AttrSpec{Name: "winrm_port", Type: cty.Number, Required: false},
		"WinRMTimeout":  &hcldec.AttrSpec{Name: "winrm_timeout", Type: cty.String, Required: false},
		"WinRMUseSSL":   &hcldec.AttrSpec{Name: "winrm_use_ssl", Type: cty.Bool, Required: false},
		"WinRMInsecure": &hcldec.AttrSpec{Name: "winrm_insecure", Type: cty.Bool, Required: false},
		"WinRMUseNTLM":  &hcldec.AttrSpec{Name: "winrm_use_ntlm", Type: cty.Bool, Required: false},
	}
	return s
}
