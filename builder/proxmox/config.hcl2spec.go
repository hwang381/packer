// Code generated by "mapstructure-to-hcl2 -type Config,nicConfig,diskConfig"; DO NOT EDIT.
package proxmox

import (
	"time"

	"github.com/hashicorp/hcl/v2/hcldec"
	"github.com/zclconf/go-cty/cty"
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
	RawBootGroupInterval      string            `mapstructure:"boot_keygroup_interval" cty:"boot_keygroup_interval"`
	RawBootWait               string            `mapstructure:"boot_wait" cty:"boot_wait"`
	BootCommand               []string          `mapstructure:"boot_command" cty:"boot_command"`
	BootGroupInterval         time.Duration     `cty:"boot_group_interval"`
	BootWait                  time.Duration     `cty:"boot_wait"`
	RawBootKeyInterval        string            `mapstructure:"boot_key_interval" cty:"boot_key_interval"`
	BootKeyInterval           time.Duration     `cty:"boot_key_interval"`
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
	ProxmoxURLRaw             string            `mapstructure:"proxmox_url" cty:"proxmox_url"`
	SkipCertValidation        bool              `mapstructure:"insecure_skip_tls_verify" cty:"insecure_skip_tls_verify"`
	Username                  string            `mapstructure:"username" cty:"username"`
	Password                  string            `mapstructure:"password" cty:"password"`
	Node                      string            `mapstructure:"node" cty:"node"`
	Pool                      string            `mapstructure:"pool" cty:"pool"`
	VMName                    string            `mapstructure:"vm_name" cty:"vm_name"`
	VMID                      int               `mapstructure:"vm_id" cty:"vm_id"`
	Memory                    int               `mapstructure:"memory" cty:"memory"`
	Cores                     int               `mapstructure:"cores" cty:"cores"`
	Sockets                   int               `mapstructure:"sockets" cty:"sockets"`
	OS                        string            `mapstructure:"os" cty:"os"`
	NICs                      []nicConfig       `mapstructure:"network_adapters" cty:"network_adapters"`
	Disks                     []diskConfig      `mapstructure:"disks" cty:"disks"`
	ISOFile                   string            `mapstructure:"iso_file" cty:"iso_file"`
	Agent                     bool              `mapstructure:"qemu_agent" cty:"qemu_agent"`
	TemplateName              string            `mapstructure:"template_name" cty:"template_name"`
	TemplateDescription       string            `mapstructure:"template_description" cty:"template_description"`
	UnmountISO                bool              `mapstructure:"unmount_iso" cty:"unmount_iso"`
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
		"RawBootGroupInterval":      &hcldec.AttrSpec{Name: "boot_keygroup_interval", Type: cty.String, Required: false},
		"RawBootWait":               &hcldec.AttrSpec{Name: "boot_wait", Type: cty.String, Required: false},
		"BootCommand":               &hcldec.AttrSpec{Name: "boot_command", Type: cty.List(cty.String), Required: false},
		"BootGroupInterval":         &hcldec.AttrSpec{Name: "boot_group_interval", Type: cty.String, Required: false},
		"BootWait":                  &hcldec.AttrSpec{Name: "boot_wait", Type: cty.String, Required: false},
		"RawBootKeyInterval":        &hcldec.AttrSpec{Name: "boot_key_interval", Type: cty.String, Required: false},
		"BootKeyInterval":           &hcldec.AttrSpec{Name: "boot_key_interval", Type: cty.String, Required: false},
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
		"ProxmoxURLRaw":             &hcldec.AttrSpec{Name: "proxmox_url", Type: cty.String, Required: false},
		"ProxmoxURL":                nil, /* not basic */
		"SkipCertValidation":        &hcldec.AttrSpec{Name: "insecure_skip_tls_verify", Type: cty.Bool, Required: false},
		"Username":                  &hcldec.AttrSpec{Name: "username", Type: cty.String, Required: false},
		"Password":                  &hcldec.AttrSpec{Name: "password", Type: cty.String, Required: false},
		"Node":                      &hcldec.AttrSpec{Name: "node", Type: cty.String, Required: false},
		"Pool":                      &hcldec.AttrSpec{Name: "pool", Type: cty.String, Required: false},
		"VMName":                    &hcldec.AttrSpec{Name: "vm_name", Type: cty.String, Required: false},
		"VMID":                      &hcldec.AttrSpec{Name: "vm_id", Type: cty.Number, Required: false},
		"Memory":                    &hcldec.AttrSpec{Name: "memory", Type: cty.Number, Required: false},
		"Cores":                     &hcldec.AttrSpec{Name: "cores", Type: cty.Number, Required: false},
		"Sockets":                   &hcldec.AttrSpec{Name: "sockets", Type: cty.Number, Required: false},
		"OS":                        &hcldec.AttrSpec{Name: "os", Type: cty.String, Required: false},
		"NICs":                      hcldec.BlockListSpec{TypeName: "[]nicConfig", Nested: &hcldec.BlockObjectSpec{TypeName: "nicConfig", Nested: hcldec.ObjectSpec((*nicConfig)(nil).HCL2Spec())}},
		"Disks":                     hcldec.BlockListSpec{TypeName: "[]diskConfig", Nested: &hcldec.BlockObjectSpec{TypeName: "diskConfig", Nested: hcldec.ObjectSpec((*diskConfig)(nil).HCL2Spec())}},
		"ISOFile":                   &hcldec.AttrSpec{Name: "iso_file", Type: cty.String, Required: false},
		"Agent":                     &hcldec.AttrSpec{Name: "qemu_agent", Type: cty.Bool, Required: false},
		"TemplateName":              &hcldec.AttrSpec{Name: "template_name", Type: cty.String, Required: false},
		"TemplateDescription":       &hcldec.AttrSpec{Name: "template_description", Type: cty.String, Required: false},
		"UnmountISO":                &hcldec.AttrSpec{Name: "unmount_iso", Type: cty.Bool, Required: false},
	}
	return s
}

type FlatdiskConfig struct {
	Type            string `mapstructure:"type" cty:"type"`
	StoragePool     string `mapstructure:"storage_pool" cty:"storage_pool"`
	StoragePoolType string `mapstructure:"storage_pool_type" cty:"storage_pool_type"`
	Size            string `mapstructure:"disk_size" cty:"disk_size"`
	CacheMode       string `mapstructure:"cache_mode" cty:"cache_mode"`
	DiskFormat      string `mapstructure:"format" cty:"format"`
}

func (*diskConfig) HCL2Spec() map[string]hcldec.Spec {
	s := map[string]hcldec.Spec{
		"Type":            &hcldec.AttrSpec{Name: "type", Type: cty.String, Required: false},
		"StoragePool":     &hcldec.AttrSpec{Name: "storage_pool", Type: cty.String, Required: false},
		"StoragePoolType": &hcldec.AttrSpec{Name: "storage_pool_type", Type: cty.String, Required: false},
		"Size":            &hcldec.AttrSpec{Name: "disk_size", Type: cty.String, Required: false},
		"CacheMode":       &hcldec.AttrSpec{Name: "cache_mode", Type: cty.String, Required: false},
		"DiskFormat":      &hcldec.AttrSpec{Name: "format", Type: cty.String, Required: false},
	}
	return s
}

type FlatnicConfig struct {
	Model      string `mapstructure:"model" cty:"model"`
	MACAddress string `mapstructure:"mac_address" cty:"mac_address"`
	Bridge     string `mapstructure:"bridge" cty:"bridge"`
	VLANTag    string `mapstructure:"vlan_tag" cty:"vlan_tag"`
}

func (*nicConfig) HCL2Spec() map[string]hcldec.Spec {
	s := map[string]hcldec.Spec{
		"Model":      &hcldec.AttrSpec{Name: "model", Type: cty.String, Required: false},
		"MACAddress": &hcldec.AttrSpec{Name: "mac_address", Type: cty.String, Required: false},
		"Bridge":     &hcldec.AttrSpec{Name: "bridge", Type: cty.String, Required: false},
		"VLANTag":    &hcldec.AttrSpec{Name: "vlan_tag", Type: cty.String, Required: false},
	}
	return s
}
