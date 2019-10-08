// Code generated by "mapstructure-to-hcl2 -type OutputConfig"; DO NOT EDIT.
package common

import (
	"github.com/hashicorp/hcl/v2/hcldec"
	"github.com/zclconf/go-cty/cty"
)

type FlatOutputConfig struct {
	OutputDir string `mapstructure:"output_directory" required:"false" cty:"output_directory"`
}

func (*OutputConfig) HCL2Spec() map[string]hcldec.Spec {
	s := map[string]hcldec.Spec{
		"OutputDir": &hcldec.AttrSpec{Name: "output_directory", Type: cty.String, Required: false},
	}
	return s
}
