// Code generated by "mapstructure-to-hcl2 -type Config"; DO NOT EDIT.
package checksum

import (
	"github.com/hashicorp/hcl/v2/hcldec"
	"github.com/zclconf/go-cty/cty"
)

type FlatConfig struct {
	PackerBuildName     string            `mapstructure:"packer_build_name" cty:"packer_build_name"`
	PackerBuilderType   string            `mapstructure:"packer_builder_type" cty:"packer_builder_type"`
	PackerDebug         bool              `mapstructure:"packer_debug" cty:"packer_debug"`
	PackerForce         bool              `mapstructure:"packer_force" cty:"packer_force"`
	PackerOnError       string            `mapstructure:"packer_on_error" cty:"packer_on_error"`
	PackerUserVars      map[string]string `mapstructure:"packer_user_variables" cty:"packer_user_variables"`
	PackerSensitiveVars []string          `mapstructure:"packer_sensitive_variables" cty:"packer_sensitive_variables"`
	Keep                bool              `mapstructure:"keep_input_artifact" cty:"keep_input_artifact"`
	ChecksumTypes       []string          `mapstructure:"checksum_types" cty:"checksum_types"`
	OutputPath          string            `mapstructure:"output" cty:"output"`
}

func (*Config) HCL2Spec() map[string]hcldec.Spec {
	s := map[string]hcldec.Spec{
		"PackerBuildName":     &hcldec.AttrSpec{Name: "packer_build_name", Type: cty.String, Required: false},
		"PackerBuilderType":   &hcldec.AttrSpec{Name: "packer_builder_type", Type: cty.String, Required: false},
		"PackerDebug":         &hcldec.AttrSpec{Name: "packer_debug", Type: cty.Bool, Required: false},
		"PackerForce":         &hcldec.AttrSpec{Name: "packer_force", Type: cty.Bool, Required: false},
		"PackerOnError":       &hcldec.AttrSpec{Name: "packer_on_error", Type: cty.String, Required: false},
		"PackerUserVars":      &hcldec.BlockAttrsSpec{TypeName: "packer_user_variables", ElementType: cty.String, Required: false},
		"PackerSensitiveVars": &hcldec.AttrSpec{Name: "packer_sensitive_variables", Type: cty.List(cty.String), Required: false},
		"Keep":                &hcldec.AttrSpec{Name: "keep_input_artifact", Type: cty.Bool, Required: false},
		"ChecksumTypes":       &hcldec.AttrSpec{Name: "checksum_types", Type: cty.List(cty.String), Required: false},
		"OutputPath":          &hcldec.AttrSpec{Name: "output", Type: cty.String, Required: false},
	}
	return s
}
