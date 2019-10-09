// Code generated by "mapstructure-to-hcl2 -type Config"; DO NOT EDIT.
package artifice

import (
	"github.com/hashicorp/hcl/v2/hcldec"
	"github.com/zclconf/go-cty/cty"
)

// FlatConfig is an auto-generated flat version of Config.
// Where the contents of a field with a `mapstructure:,squash` tag are bubbled up.
type FlatConfig struct {
	PackerBuildName     string            `mapstructure:"packer_build_name" cty:"packer_build_name"`
	PackerBuilderType   string            `mapstructure:"packer_builder_type" cty:"packer_builder_type"`
	PackerDebug         bool              `mapstructure:"packer_debug" cty:"packer_debug"`
	PackerForce         bool              `mapstructure:"packer_force" cty:"packer_force"`
	PackerOnError       string            `mapstructure:"packer_on_error" cty:"packer_on_error"`
	PackerUserVars      map[string]string `mapstructure:"packer_user_variables" cty:"packer_user_variables"`
	PackerSensitiveVars []string          `mapstructure:"packer_sensitive_variables" cty:"packer_sensitive_variables"`
	Files               []string          `mapstructure:"files" cty:"files"`
	Keep                bool              `mapstructure:"keep_input_artifact" cty:"keep_input_artifact"`
}

// FlatMapstructure returns a new FlatConfig.
// FlatConfig is an auto-generated flat version of Config.
// Where the contents a fields with a `mapstructure:,squash` tag are bubbled up.
func (*Config) FlatMapstructure() interface{} { return new(FlatConfig) }

// HCL2Spec returns the hcldec.Spec of a Config.
// This spec is used by HCL to read the fields of Config.
func (*Config) HCL2Spec() map[string]hcldec.Spec {
	s := map[string]hcldec.Spec{
		"PackerBuildName":     &hcldec.AttrSpec{Name: "packer_build_name", Type: cty.String, Required: false},
		"PackerBuilderType":   &hcldec.AttrSpec{Name: "packer_builder_type", Type: cty.String, Required: false},
		"PackerDebug":         &hcldec.AttrSpec{Name: "packer_debug", Type: cty.Bool, Required: false},
		"PackerForce":         &hcldec.AttrSpec{Name: "packer_force", Type: cty.Bool, Required: false},
		"PackerOnError":       &hcldec.AttrSpec{Name: "packer_on_error", Type: cty.String, Required: false},
		"PackerUserVars":      &hcldec.BlockAttrsSpec{TypeName: "packer_user_variables", ElementType: cty.String, Required: false},
		"PackerSensitiveVars": &hcldec.AttrSpec{Name: "packer_sensitive_variables", Type: cty.List(cty.String), Required: false},
		"Files":               &hcldec.AttrSpec{Name: "files", Type: cty.List(cty.String), Required: false},
		"Keep":                &hcldec.AttrSpec{Name: "keep_input_artifact", Type: cty.Bool, Required: false},
	}
	return s
}
