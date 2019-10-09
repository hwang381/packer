// Code generated by "mapstructure-to-hcl2 -type Config"; DO NOT EDIT.
package ovf

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
	HTTPDir                   string            `mapstructure:"http_directory" cty:"http_directory"`
	HTTPPortMin               int               `mapstructure:"http_port_min" cty:"http_port_min"`
	HTTPPortMax               int               `mapstructure:"http_port_max" cty:"http_port_max"`
	FloppyFiles               []string          `mapstructure:"floppy_files" cty:"floppy_files"`
	FloppyDirectories         []string          `mapstructure:"floppy_dirs" cty:"floppy_dirs"`
	FloppyLabel               string            `mapstructure:"floppy_label" cty:"floppy_label"`
	RawBootGroupInterval      string            `mapstructure:"boot_keygroup_interval" cty:"boot_keygroup_interval"`
	RawBootWait               string            `mapstructure:"boot_wait" cty:"boot_wait"`
	BootCommand               []string          `mapstructure:"boot_command" cty:"boot_command"`
	BootGroupInterval         time.Duration     `cty:"boot_group_interval"`
	BootWait                  time.Duration     `cty:"boot_wait"`
	Format                    string            `mapstructure:"format" required:"false" cty:"format"`
	ExportOpts                []string          `mapstructure:"export_opts" required:"false" cty:"export_opts"`
	OutputDir                 string            `mapstructure:"output_directory" required:"false" cty:"output_directory"`
	Headless                  bool              `mapstructure:"headless" required:"false" cty:"headless"`
	VRDPBindAddress           string            `mapstructure:"vrdp_bind_address" required:"false" cty:"vrdp_bind_address"`
	VRDPPortMin               int               `mapstructure:"vrdp_port_min" required:"false" cty:"vrdp_port_min"`
	VRDPPortMax               int               `mapstructure:"vrdp_port_max" cty:"vrdp_port_max"`
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
	SSHHostPortMin            int               `mapstructure:"ssh_host_port_min" required:"false" cty:"ssh_host_port_min"`
	SSHHostPortMax            int               `mapstructure:"ssh_host_port_max" cty:"ssh_host_port_max"`
	SSHSkipNatMapping         bool              `mapstructure:"ssh_skip_nat_mapping" required:"false" cty:"ssh_skip_nat_mapping"`
	SSHWaitTimeout            time.Duration     `mapstructure:"ssh_wait_timeout" cty:"ssh_wait_timeout"`
	ShutdownCommand           string            `mapstructure:"shutdown_command" required:"false" cty:"shutdown_command"`
	RawShutdownTimeout        string            `mapstructure:"shutdown_timeout" required:"false" cty:"shutdown_timeout"`
	RawPostShutdownDelay      string            `mapstructure:"post_shutdown_delay" required:"false" cty:"post_shutdown_delay"`
	ShutdownTimeout           time.Duration     `cty:"shutdown_timeout"`
	PostShutdownDelay         time.Duration     `cty:"post_shutdown_delay"`
	VBoxManage                [][]string        `mapstructure:"vboxmanage" required:"false" cty:"vboxmanage"`
	VBoxManagePost            [][]string        `mapstructure:"vboxmanage_post" required:"false" cty:"vboxmanage_post"`
	Communicator              string            `mapstructure:"communicator" cty:"communicator"`
	VBoxVersionFile           *string           `mapstructure:"virtualbox_version_file" required:"false" cty:"virtualbox_version_file"`
	GuestAdditionsMode        string            `mapstructure:"guest_additions_mode" required:"false" cty:"guest_additions_mode"`
	Checksum                  string            `mapstructure:"checksum" required:"true" cty:"checksum"`
	ChecksumType              string            `mapstructure:"checksum_type" required:"false" cty:"checksum_type"`
	GuestAdditionsPath        string            `mapstructure:"guest_additions_path" required:"false" cty:"guest_additions_path"`
	GuestAdditionsInterface   string            `mapstructure:"guest_additions_interface" required:"false" cty:"guest_additions_interface"`
	GuestAdditionsSHA256      string            `mapstructure:"guest_additions_sha256" required:"false" cty:"guest_additions_sha256"`
	GuestAdditionsURL         string            `mapstructure:"guest_additions_url" required:"false" cty:"guest_additions_url"`
	ImportFlags               []string          `mapstructure:"import_flags" required:"false" cty:"import_flags"`
	ImportOpts                string            `mapstructure:"import_opts" required:"false" cty:"import_opts"`
	SourcePath                string            `mapstructure:"source_path" required:"true" cty:"source_path"`
	TargetPath                string            `mapstructure:"target_path" required:"false" cty:"target_path"`
	VMName                    string            `mapstructure:"vm_name" required:"false" cty:"vm_name"`
	KeepRegistered            bool              `mapstructure:"keep_registered" required:"false" cty:"keep_registered"`
	SkipExport                bool              `mapstructure:"skip_export" required:"false" cty:"skip_export"`
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
		"HTTPDir":                   &hcldec.AttrSpec{Name: "http_directory", Type: cty.String, Required: false},
		"HTTPPortMin":               &hcldec.AttrSpec{Name: "http_port_min", Type: cty.Number, Required: false},
		"HTTPPortMax":               &hcldec.AttrSpec{Name: "http_port_max", Type: cty.Number, Required: false},
		"FloppyFiles":               &hcldec.AttrSpec{Name: "floppy_files", Type: cty.List(cty.String), Required: false},
		"FloppyDirectories":         &hcldec.AttrSpec{Name: "floppy_dirs", Type: cty.List(cty.String), Required: false},
		"FloppyLabel":               &hcldec.AttrSpec{Name: "floppy_label", Type: cty.String, Required: false},
		"RawBootGroupInterval":      &hcldec.AttrSpec{Name: "boot_keygroup_interval", Type: cty.String, Required: false},
		"RawBootWait":               &hcldec.AttrSpec{Name: "boot_wait", Type: cty.String, Required: false},
		"BootCommand":               &hcldec.AttrSpec{Name: "boot_command", Type: cty.List(cty.String), Required: false},
		"BootGroupInterval":         &hcldec.AttrSpec{Name: "boot_group_interval", Type: cty.String, Required: false},
		"BootWait":                  &hcldec.AttrSpec{Name: "boot_wait", Type: cty.String, Required: false},
		"Format":                    &hcldec.AttrSpec{Name: "format", Type: cty.String, Required: false},
		"ExportOpts":                &hcldec.AttrSpec{Name: "export_opts", Type: cty.List(cty.String), Required: false},
		"OutputDir":                 &hcldec.AttrSpec{Name: "output_directory", Type: cty.String, Required: false},
		"Headless":                  &hcldec.AttrSpec{Name: "headless", Type: cty.Bool, Required: false},
		"VRDPBindAddress":           &hcldec.AttrSpec{Name: "vrdp_bind_address", Type: cty.String, Required: false},
		"VRDPPortMin":               &hcldec.AttrSpec{Name: "vrdp_port_min", Type: cty.Number, Required: false},
		"VRDPPortMax":               &hcldec.AttrSpec{Name: "vrdp_port_max", Type: cty.Number, Required: false},
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
		"SSHHostPortMin":            &hcldec.AttrSpec{Name: "ssh_host_port_min", Type: cty.Number, Required: false},
		"SSHHostPortMax":            &hcldec.AttrSpec{Name: "ssh_host_port_max", Type: cty.Number, Required: false},
		"SSHSkipNatMapping":         &hcldec.AttrSpec{Name: "ssh_skip_nat_mapping", Type: cty.Bool, Required: false},
		"SSHWaitTimeout":            &hcldec.AttrSpec{Name: "ssh_wait_timeout", Type: cty.String, Required: false},
		"ShutdownCommand":           &hcldec.AttrSpec{Name: "shutdown_command", Type: cty.String, Required: false},
		"RawShutdownTimeout":        &hcldec.AttrSpec{Name: "shutdown_timeout", Type: cty.String, Required: false},
		"RawPostShutdownDelay":      &hcldec.AttrSpec{Name: "post_shutdown_delay", Type: cty.String, Required: false},
		"ShutdownTimeout":           &hcldec.AttrSpec{Name: "shutdown_timeout", Type: cty.String, Required: false},
		"PostShutdownDelay":         &hcldec.AttrSpec{Name: "post_shutdown_delay", Type: cty.String, Required: false},
		"VBoxManage":                nil, // slice ([][]string),
		"VBoxManagePost":            nil, // slice ([][]string),
		"Communicator":              &hcldec.AttrSpec{Name: "communicator", Type: cty.String, Required: false},
		"VBoxVersionFile":           nil, /* not basic */
		"GuestAdditionsMode":        &hcldec.AttrSpec{Name: "guest_additions_mode", Type: cty.String, Required: false},
		"Checksum":                  &hcldec.AttrSpec{Name: "checksum", Type: cty.String, Required: false},
		"ChecksumType":              &hcldec.AttrSpec{Name: "checksum_type", Type: cty.String, Required: false},
		"GuestAdditionsPath":        &hcldec.AttrSpec{Name: "guest_additions_path", Type: cty.String, Required: false},
		"GuestAdditionsInterface":   &hcldec.AttrSpec{Name: "guest_additions_interface", Type: cty.String, Required: false},
		"GuestAdditionsSHA256":      &hcldec.AttrSpec{Name: "guest_additions_sha256", Type: cty.String, Required: false},
		"GuestAdditionsURL":         &hcldec.AttrSpec{Name: "guest_additions_url", Type: cty.String, Required: false},
		"ImportFlags":               &hcldec.AttrSpec{Name: "import_flags", Type: cty.List(cty.String), Required: false},
		"ImportOpts":                &hcldec.AttrSpec{Name: "import_opts", Type: cty.String, Required: false},
		"SourcePath":                &hcldec.AttrSpec{Name: "source_path", Type: cty.String, Required: false},
		"TargetPath":                &hcldec.AttrSpec{Name: "target_path", Type: cty.String, Required: false},
		"VMName":                    &hcldec.AttrSpec{Name: "vm_name", Type: cty.String, Required: false},
		"KeepRegistered":            &hcldec.AttrSpec{Name: "keep_registered", Type: cty.Bool, Required: false},
		"SkipExport":                &hcldec.AttrSpec{Name: "skip_export", Type: cty.Bool, Required: false},
	}
	return s
}
