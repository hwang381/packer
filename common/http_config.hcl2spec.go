// Code generated by "mapstructure-to-hcl2 -type HTTPConfig"; DO NOT EDIT.
package common

import (
	"github.com/hashicorp/hcl/v2/hcldec"
	"github.com/zclconf/go-cty/cty"
)

type FlatHTTPConfig struct {
	HTTPDir     string `mapstructure:"http_directory" cty:"http_directory"`
	HTTPPortMin int    `mapstructure:"http_port_min" cty:"http_port_min"`
	HTTPPortMax int    `mapstructure:"http_port_max" cty:"http_port_max"`
}

func (*HTTPConfig) HCL2Spec() map[string]hcldec.Spec {
	s := map[string]hcldec.Spec{
		"HTTPDir":     &hcldec.AttrSpec{Name: "http_directory", Type: cty.String, Required: false},
		"HTTPPortMin": &hcldec.AttrSpec{Name: "http_port_min", Type: cty.Number, Required: false},
		"HTTPPortMax": &hcldec.AttrSpec{Name: "http_port_max", Type: cty.Number, Required: false},
	}
	return s
}
