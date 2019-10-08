// Code generated by "mapstructure-to-hcl2 -type Config"; DO NOT EDIT.
package file

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
	Source              string            `cty:"source"`
	Sources             []string          `cty:"sources"`
	Destination         string            `cty:"destination"`
	Direction           string            `cty:"direction"`
	Generated           bool              `cty:"generated"`
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
		"Source":              &hcldec.AttrSpec{Name: "source", Type: cty.String, Required: false},
		"Sources":             &hcldec.AttrSpec{Name: "sources", Type: cty.List(cty.String), Required: false},
		"Destination":         &hcldec.AttrSpec{Name: "destination", Type: cty.String, Required: false},
		"Direction":           &hcldec.AttrSpec{Name: "direction", Type: cty.String, Required: false},
		"Generated":           &hcldec.AttrSpec{Name: "generated", Type: cty.Bool, Required: false},
	}
	return s
}
