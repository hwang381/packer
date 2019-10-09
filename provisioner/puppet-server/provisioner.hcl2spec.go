// Code generated by "mapstructure-to-hcl2 -type Config"; DO NOT EDIT.
package puppetserver

import (
	"github.com/hashicorp/hcl/v2/hcldec"
	"github.com/zclconf/go-cty/cty"
)

// FlatConfig is an auto-generated flat version of Config.
// Where the contents of a field with a `mapstructure:,squash` tag are bubbled up.
type FlatConfig struct {
	PackerBuildName      string            `mapstructure:"packer_build_name" cty:"packer_build_name"`
	PackerBuilderType    string            `mapstructure:"packer_builder_type" cty:"packer_builder_type"`
	PackerDebug          bool              `mapstructure:"packer_debug" cty:"packer_debug"`
	PackerForce          bool              `mapstructure:"packer_force" cty:"packer_force"`
	PackerOnError        string            `mapstructure:"packer_on_error" cty:"packer_on_error"`
	PackerUserVars       map[string]string `mapstructure:"packer_user_variables" cty:"packer_user_variables"`
	PackerSensitiveVars  []string          `mapstructure:"packer_sensitive_variables" cty:"packer_sensitive_variables"`
	CleanStagingDir      bool              `mapstructure:"clean_staging_directory" cty:"clean_staging_directory"`
	ClientCertPath       string            `mapstructure:"client_cert_path" cty:"client_cert_path"`
	ClientPrivateKeyPath string            `mapstructure:"client_private_key_path" cty:"client_private_key_path"`
	ExecuteCommand       string            `mapstructure:"execute_command" cty:"execute_command"`
	ExtraArguments       []string          `mapstructure:"extra_arguments" cty:"extra_arguments"`
	Facter               map[string]string `cty:"facter"`
	GuestOSType          string            `mapstructure:"guest_os_type" cty:"guest_os_type"`
	IgnoreExitCodes      bool              `mapstructure:"ignore_exit_codes" cty:"ignore_exit_codes"`
	PreventSudo          bool              `mapstructure:"prevent_sudo" cty:"prevent_sudo"`
	PuppetBinDir         string            `mapstructure:"puppet_bin_dir" cty:"puppet_bin_dir"`
	PuppetNode           string            `mapstructure:"puppet_node" cty:"puppet_node"`
	PuppetServer         string            `mapstructure:"puppet_server" cty:"puppet_server"`
	StagingDir           string            `mapstructure:"staging_dir" cty:"staging_dir"`
	WorkingDir           string            `mapstructure:"working_directory" cty:"working_directory"`
	ElevatedUser         string            `mapstructure:"elevated_user" cty:"elevated_user"`
	ElevatedPassword     string            `mapstructure:"elevated_password" cty:"elevated_password"`
}

// FlatMapstructure returns a new FlatConfig.
// FlatConfig is an auto-generated flat version of Config.
// Where the contents a fields with a `mapstructure:,squash` tag are bubbled up.
func (*Config) FlatMapstructure() interface{} { return new(FlatConfig) }

// HCL2Spec returns the hcldec.Spec of a Config.
// This spec is used by HCL to read the fields of Config.
func (*Config) HCL2Spec() map[string]hcldec.Spec {
	s := map[string]hcldec.Spec{
		"PackerBuildName":      &hcldec.AttrSpec{Name: "packer_build_name", Type: cty.String, Required: false},
		"PackerBuilderType":    &hcldec.AttrSpec{Name: "packer_builder_type", Type: cty.String, Required: false},
		"PackerDebug":          &hcldec.AttrSpec{Name: "packer_debug", Type: cty.Bool, Required: false},
		"PackerForce":          &hcldec.AttrSpec{Name: "packer_force", Type: cty.Bool, Required: false},
		"PackerOnError":        &hcldec.AttrSpec{Name: "packer_on_error", Type: cty.String, Required: false},
		"PackerUserVars":       &hcldec.BlockAttrsSpec{TypeName: "packer_user_variables", ElementType: cty.String, Required: false},
		"PackerSensitiveVars":  &hcldec.AttrSpec{Name: "packer_sensitive_variables", Type: cty.List(cty.String), Required: false},
		"CleanStagingDir":      &hcldec.AttrSpec{Name: "clean_staging_directory", Type: cty.Bool, Required: false},
		"ClientCertPath":       &hcldec.AttrSpec{Name: "client_cert_path", Type: cty.String, Required: false},
		"ClientPrivateKeyPath": &hcldec.AttrSpec{Name: "client_private_key_path", Type: cty.String, Required: false},
		"ExecuteCommand":       &hcldec.AttrSpec{Name: "execute_command", Type: cty.String, Required: false},
		"ExtraArguments":       &hcldec.AttrSpec{Name: "extra_arguments", Type: cty.List(cty.String), Required: false},
		"Facter":               &hcldec.BlockAttrsSpec{TypeName: "facter", ElementType: cty.String, Required: false},
		"GuestOSType":          &hcldec.AttrSpec{Name: "guest_os_type", Type: cty.String, Required: false},
		"IgnoreExitCodes":      &hcldec.AttrSpec{Name: "ignore_exit_codes", Type: cty.Bool, Required: false},
		"PreventSudo":          &hcldec.AttrSpec{Name: "prevent_sudo", Type: cty.Bool, Required: false},
		"PuppetBinDir":         &hcldec.AttrSpec{Name: "puppet_bin_dir", Type: cty.String, Required: false},
		"PuppetNode":           &hcldec.AttrSpec{Name: "puppet_node", Type: cty.String, Required: false},
		"PuppetServer":         &hcldec.AttrSpec{Name: "puppet_server", Type: cty.String, Required: false},
		"StagingDir":           &hcldec.AttrSpec{Name: "staging_dir", Type: cty.String, Required: false},
		"WorkingDir":           &hcldec.AttrSpec{Name: "working_directory", Type: cty.String, Required: false},
		"ElevatedUser":         &hcldec.AttrSpec{Name: "elevated_user", Type: cty.String, Required: false},
		"ElevatedPassword":     &hcldec.AttrSpec{Name: "elevated_password", Type: cty.String, Required: false},
	}
	return s
}
