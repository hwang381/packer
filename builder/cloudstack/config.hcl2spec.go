// Code generated by "mapstructure-to-hcl2 -type Config"; DO NOT EDIT.
package cloudstack

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
	HTTPDir                   string            `mapstructure:"http_directory" cty:"http_directory"`
	HTTPPortMin               int               `mapstructure:"http_port_min" cty:"http_port_min"`
	HTTPPortMax               int               `mapstructure:"http_port_max" cty:"http_port_max"`
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
	APIURL                    string            `mapstructure:"api_url" required:"true" cty:"api_url"`
	APIKey                    string            `mapstructure:"api_key" required:"true" cty:"api_key"`
	SecretKey                 string            `mapstructure:"secret_key" required:"true" cty:"secret_key"`
	AsyncTimeout              time.Duration     `mapstructure:"async_timeout" required:"false" cty:"async_timeout"`
	HTTPGetOnly               bool              `mapstructure:"http_get_only" required:"false" cty:"http_get_only"`
	SSLNoVerify               bool              `mapstructure:"ssl_no_verify" required:"false" cty:"ssl_no_verify"`
	CIDRList                  []string          `mapstructure:"cidr_list" required:"false" cty:"cidr_list"`
	CreateSecurityGroup       bool              `mapstructure:"create_security_group" required:"false" cty:"create_security_group"`
	DiskOffering              string            `mapstructure:"disk_offering" required:"false" cty:"disk_offering"`
	DiskSize                  int64             `mapstructure:"disk_size" required:"false" cty:"disk_size"`
	EjectISO                  bool              `mapstructure:"eject_iso" cty:"eject_iso"`
	EjectISODelay             time.Duration     `mapstructure:"eject_iso_delay" cty:"eject_iso_delay"`
	Expunge                   bool              `mapstructure:"expunge" required:"false" cty:"expunge"`
	Hypervisor                string            `mapstructure:"hypervisor" required:"false" cty:"hypervisor"`
	InstanceName              string            `mapstructure:"instance_name" required:"false" cty:"instance_name"`
	Network                   string            `mapstructure:"network" required:"true" cty:"network"`
	Project                   string            `mapstructure:"project" required:"false" cty:"project"`
	PublicIPAddress           string            `mapstructure:"public_ip_address" required:"false" cty:"public_ip_address"`
	PublicPort                int               `mapstructure:"public_port" required:"false" cty:"public_port"`
	SecurityGroups            []string          `mapstructure:"security_groups" required:"false" cty:"security_groups"`
	ServiceOffering           string            `mapstructure:"service_offering" required:"true" cty:"service_offering"`
	PreventFirewallChanges    bool              `mapstructure:"prevent_firewall_changes" required:"false" cty:"prevent_firewall_changes"`
	SourceISO                 string            `mapstructure:"source_iso" required:"true" cty:"source_iso"`
	SourceTemplate            string            `mapstructure:"source_template" required:"true" cty:"source_template"`
	TemporaryKeypairName      string            `mapstructure:"temporary_keypair_name" required:"false" cty:"temporary_keypair_name"`
	UseLocalIPAddress         bool              `mapstructure:"use_local_ip_address" required:"false" cty:"use_local_ip_address"`
	UserData                  string            `mapstructure:"user_data" required:"false" cty:"user_data"`
	UserDataFile              string            `mapstructure:"user_data_file" required:"false" cty:"user_data_file"`
	Zone                      string            `mapstructure:"zone" required:"true" cty:"zone"`
	TemplateName              string            `mapstructure:"template_name" required:"false" cty:"template_name"`
	TemplateDisplayText       string            `mapstructure:"template_display_text" required:"false" cty:"template_display_text"`
	TemplateOS                string            `mapstructure:"template_os" required:"true" cty:"template_os"`
	TemplateFeatured          bool              `mapstructure:"template_featured" required:"false" cty:"template_featured"`
	TemplatePublic            bool              `mapstructure:"template_public" required:"false" cty:"template_public"`
	TemplatePasswordEnabled   bool              `mapstructure:"template_password_enabled" required:"false" cty:"template_password_enabled"`
	TemplateRequiresHVM       bool              `mapstructure:"template_requires_hvm" required:"false" cty:"template_requires_hvm"`
	TemplateScalable          bool              `mapstructure:"template_scalable" required:"false" cty:"template_scalable"`
	TemplateTag               string            `mapstructure:"template_tag" cty:"template_tag"`
	Tags                      map[string]string `mapstructure:"tags" cty:"tags"`
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
		"HTTPDir":                   &hcldec.AttrSpec{Name: "http_directory", Type: cty.String, Required: false},
		"HTTPPortMin":               &hcldec.AttrSpec{Name: "http_port_min", Type: cty.Number, Required: false},
		"HTTPPortMax":               &hcldec.AttrSpec{Name: "http_port_max", Type: cty.Number, Required: false},
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
		"APIURL":                    &hcldec.AttrSpec{Name: "api_url", Type: cty.String, Required: false},
		"APIKey":                    &hcldec.AttrSpec{Name: "api_key", Type: cty.String, Required: false},
		"SecretKey":                 &hcldec.AttrSpec{Name: "secret_key", Type: cty.String, Required: false},
		"AsyncTimeout":              &hcldec.AttrSpec{Name: "async_timeout", Type: cty.String, Required: false},
		"HTTPGetOnly":               &hcldec.AttrSpec{Name: "http_get_only", Type: cty.Bool, Required: false},
		"SSLNoVerify":               &hcldec.AttrSpec{Name: "ssl_no_verify", Type: cty.Bool, Required: false},
		"CIDRList":                  &hcldec.AttrSpec{Name: "cidr_list", Type: cty.List(cty.String), Required: false},
		"CreateSecurityGroup":       &hcldec.AttrSpec{Name: "create_security_group", Type: cty.Bool, Required: false},
		"DiskOffering":              &hcldec.AttrSpec{Name: "disk_offering", Type: cty.String, Required: false},
		"DiskSize":                  &hcldec.AttrSpec{Name: "disk_size", Type: cty.Number, Required: false},
		"EjectISO":                  &hcldec.AttrSpec{Name: "eject_iso", Type: cty.Bool, Required: false},
		"EjectISODelay":             &hcldec.AttrSpec{Name: "eject_iso_delay", Type: cty.String, Required: false},
		"Expunge":                   &hcldec.AttrSpec{Name: "expunge", Type: cty.Bool, Required: false},
		"Hypervisor":                &hcldec.AttrSpec{Name: "hypervisor", Type: cty.String, Required: false},
		"InstanceName":              &hcldec.AttrSpec{Name: "instance_name", Type: cty.String, Required: false},
		"Network":                   &hcldec.AttrSpec{Name: "network", Type: cty.String, Required: false},
		"Project":                   &hcldec.AttrSpec{Name: "project", Type: cty.String, Required: false},
		"PublicIPAddress":           &hcldec.AttrSpec{Name: "public_ip_address", Type: cty.String, Required: false},
		"PublicPort":                &hcldec.AttrSpec{Name: "public_port", Type: cty.Number, Required: false},
		"SecurityGroups":            &hcldec.AttrSpec{Name: "security_groups", Type: cty.List(cty.String), Required: false},
		"ServiceOffering":           &hcldec.AttrSpec{Name: "service_offering", Type: cty.String, Required: false},
		"PreventFirewallChanges":    &hcldec.AttrSpec{Name: "prevent_firewall_changes", Type: cty.Bool, Required: false},
		"SourceISO":                 &hcldec.AttrSpec{Name: "source_iso", Type: cty.String, Required: false},
		"SourceTemplate":            &hcldec.AttrSpec{Name: "source_template", Type: cty.String, Required: false},
		"TemporaryKeypairName":      &hcldec.AttrSpec{Name: "temporary_keypair_name", Type: cty.String, Required: false},
		"UseLocalIPAddress":         &hcldec.AttrSpec{Name: "use_local_ip_address", Type: cty.Bool, Required: false},
		"UserData":                  &hcldec.AttrSpec{Name: "user_data", Type: cty.String, Required: false},
		"UserDataFile":              &hcldec.AttrSpec{Name: "user_data_file", Type: cty.String, Required: false},
		"Zone":                      &hcldec.AttrSpec{Name: "zone", Type: cty.String, Required: false},
		"TemplateName":              &hcldec.AttrSpec{Name: "template_name", Type: cty.String, Required: false},
		"TemplateDisplayText":       &hcldec.AttrSpec{Name: "template_display_text", Type: cty.String, Required: false},
		"TemplateOS":                &hcldec.AttrSpec{Name: "template_os", Type: cty.String, Required: false},
		"TemplateFeatured":          &hcldec.AttrSpec{Name: "template_featured", Type: cty.Bool, Required: false},
		"TemplatePublic":            &hcldec.AttrSpec{Name: "template_public", Type: cty.Bool, Required: false},
		"TemplatePasswordEnabled":   &hcldec.AttrSpec{Name: "template_password_enabled", Type: cty.Bool, Required: false},
		"TemplateRequiresHVM":       &hcldec.AttrSpec{Name: "template_requires_hvm", Type: cty.Bool, Required: false},
		"TemplateScalable":          &hcldec.AttrSpec{Name: "template_scalable", Type: cty.Bool, Required: false},
		"TemplateTag":               &hcldec.AttrSpec{Name: "template_tag", Type: cty.String, Required: false},
		"Tags":                      &hcldec.BlockAttrsSpec{TypeName: "tags", ElementType: cty.String, Required: false},
	}
	return s
}
