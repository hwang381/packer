// Code generated by "mapstructure-to-hcl2 -type Config"; DO NOT EDIT.
package vagrant

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
	ISOChecksum               string            `mapstructure:"iso_checksum" required:"true" cty:"iso_checksum"`
	ISOChecksumURL            string            `mapstructure:"iso_checksum_url" cty:"iso_checksum_url"`
	ISOChecksumType           string            `mapstructure:"iso_checksum_type" cty:"iso_checksum_type"`
	RawSingleISOUrl           string            `mapstructure:"iso_url" required:"true" cty:"iso_url"`
	ISOUrls                   []string          `mapstructure:"iso_urls" cty:"iso_urls"`
	TargetPath                string            `mapstructure:"iso_target_path" cty:"iso_target_path"`
	TargetExtension           string            `mapstructure:"iso_target_extension" cty:"iso_target_extension"`
	FloppyFiles               []string          `mapstructure:"floppy_files" cty:"floppy_files"`
	FloppyDirectories         []string          `mapstructure:"floppy_dirs" cty:"floppy_dirs"`
	FloppyLabel               string            `mapstructure:"floppy_label" cty:"floppy_label"`
	RawBootGroupInterval      string            `mapstructure:"boot_keygroup_interval" cty:"boot_keygroup_interval"`
	RawBootWait               string            `mapstructure:"boot_wait" cty:"boot_wait"`
	BootCommand               []string          `mapstructure:"boot_command" cty:"boot_command"`
	BootGroupInterval         time.Duration     `cty:"boot_group_interval"`
	BootWait                  time.Duration     `cty:"boot_wait"`
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
	OutputDir                 string            `mapstructure:"output_dir" required:"false" cty:"output_dir"`
	SourceBox                 string            `mapstructure:"source_path" required:"true" cty:"source_path"`
	GlobalID                  string            `mapstructure:"global_id" required:"true" cty:"global_id"`
	Checksum                  string            `mapstructure:"checksum" required:"false" cty:"checksum"`
	ChecksumType              string            `mapstructure:"checksum_type" required:"false" cty:"checksum_type"`
	BoxName                   string            `mapstructure:"box_name" required:"false" cty:"box_name"`
	Provider                  string            `mapstructure:"provider" required:"false" cty:"provider"`
	Communicator              string            `mapstructure:"communicator" cty:"communicator"`
	VagrantfileTpl            string            `mapstructure:"vagrantfile_template" cty:"vagrantfile_template"`
	TeardownMethod            string            `mapstructure:"teardown_method" required:"false" cty:"teardown_method"`
	BoxVersion                string            `mapstructure:"box_version" required:"false" cty:"box_version"`
	Template                  string            `mapstructure:"template" required:"false" cty:"template"`
	SyncedFolder              string            `mapstructure:"synced_folder" cty:"synced_folder"`
	SkipAdd                   bool              `mapstructure:"skip_add" required:"false" cty:"skip_add"`
	AddCACert                 string            `mapstructure:"add_cacert" required:"false" cty:"add_cacert"`
	AddCAPath                 string            `mapstructure:"add_capath" required:"false" cty:"add_capath"`
	AddCert                   string            `mapstructure:"add_cert" required:"false" cty:"add_cert"`
	AddClean                  bool              `mapstructure:"add_clean" required:"false" cty:"add_clean"`
	AddForce                  bool              `mapstructure:"add_force" required:"false" cty:"add_force"`
	AddInsecure               bool              `mapstructure:"add_insecure" required:"false" cty:"add_insecure"`
	SkipPackage               bool              `mapstructure:"skip_package" required:"false" cty:"skip_package"`
	OutputVagrantfile         string            `mapstructure:"output_vagrantfile" cty:"output_vagrantfile"`
	PackageInclude            []string          `mapstructure:"package_include" cty:"package_include"`
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
		"ISOChecksum":               &hcldec.AttrSpec{Name: "iso_checksum", Type: cty.String, Required: false},
		"ISOChecksumURL":            &hcldec.AttrSpec{Name: "iso_checksum_url", Type: cty.String, Required: false},
		"ISOChecksumType":           &hcldec.AttrSpec{Name: "iso_checksum_type", Type: cty.String, Required: false},
		"RawSingleISOUrl":           &hcldec.AttrSpec{Name: "iso_url", Type: cty.String, Required: false},
		"ISOUrls":                   &hcldec.AttrSpec{Name: "iso_urls", Type: cty.List(cty.String), Required: false},
		"TargetPath":                &hcldec.AttrSpec{Name: "iso_target_path", Type: cty.String, Required: false},
		"TargetExtension":           &hcldec.AttrSpec{Name: "iso_target_extension", Type: cty.String, Required: false},
		"FloppyFiles":               &hcldec.AttrSpec{Name: "floppy_files", Type: cty.List(cty.String), Required: false},
		"FloppyDirectories":         &hcldec.AttrSpec{Name: "floppy_dirs", Type: cty.List(cty.String), Required: false},
		"FloppyLabel":               &hcldec.AttrSpec{Name: "floppy_label", Type: cty.String, Required: false},
		"RawBootGroupInterval":      &hcldec.AttrSpec{Name: "boot_keygroup_interval", Type: cty.String, Required: false},
		"RawBootWait":               &hcldec.AttrSpec{Name: "boot_wait", Type: cty.String, Required: false},
		"BootCommand":               &hcldec.AttrSpec{Name: "boot_command", Type: cty.List(cty.String), Required: false},
		"BootGroupInterval":         &hcldec.AttrSpec{Name: "boot_group_interval", Type: cty.String, Required: false},
		"BootWait":                  &hcldec.AttrSpec{Name: "boot_wait", Type: cty.String, Required: false},
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
		"OutputDir":                 &hcldec.AttrSpec{Name: "output_dir", Type: cty.String, Required: false},
		"SourceBox":                 &hcldec.AttrSpec{Name: "source_path", Type: cty.String, Required: false},
		"GlobalID":                  &hcldec.AttrSpec{Name: "global_id", Type: cty.String, Required: false},
		"Checksum":                  &hcldec.AttrSpec{Name: "checksum", Type: cty.String, Required: false},
		"ChecksumType":              &hcldec.AttrSpec{Name: "checksum_type", Type: cty.String, Required: false},
		"BoxName":                   &hcldec.AttrSpec{Name: "box_name", Type: cty.String, Required: false},
		"Provider":                  &hcldec.AttrSpec{Name: "provider", Type: cty.String, Required: false},
		"Communicator":              &hcldec.AttrSpec{Name: "communicator", Type: cty.String, Required: false},
		"VagrantfileTpl":            &hcldec.AttrSpec{Name: "vagrantfile_template", Type: cty.String, Required: false},
		"TeardownMethod":            &hcldec.AttrSpec{Name: "teardown_method", Type: cty.String, Required: false},
		"BoxVersion":                &hcldec.AttrSpec{Name: "box_version", Type: cty.String, Required: false},
		"Template":                  &hcldec.AttrSpec{Name: "template", Type: cty.String, Required: false},
		"SyncedFolder":              &hcldec.AttrSpec{Name: "synced_folder", Type: cty.String, Required: false},
		"SkipAdd":                   &hcldec.AttrSpec{Name: "skip_add", Type: cty.Bool, Required: false},
		"AddCACert":                 &hcldec.AttrSpec{Name: "add_cacert", Type: cty.String, Required: false},
		"AddCAPath":                 &hcldec.AttrSpec{Name: "add_capath", Type: cty.String, Required: false},
		"AddCert":                   &hcldec.AttrSpec{Name: "add_cert", Type: cty.String, Required: false},
		"AddClean":                  &hcldec.AttrSpec{Name: "add_clean", Type: cty.Bool, Required: false},
		"AddForce":                  &hcldec.AttrSpec{Name: "add_force", Type: cty.Bool, Required: false},
		"AddInsecure":               &hcldec.AttrSpec{Name: "add_insecure", Type: cty.Bool, Required: false},
		"SkipPackage":               &hcldec.AttrSpec{Name: "skip_package", Type: cty.Bool, Required: false},
		"OutputVagrantfile":         &hcldec.AttrSpec{Name: "output_vagrantfile", Type: cty.String, Required: false},
		"PackageInclude":            &hcldec.AttrSpec{Name: "package_include", Type: cty.List(cty.String), Required: false},
	}
	return s
}
