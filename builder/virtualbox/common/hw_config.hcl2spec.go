// Code generated by "mapstructure-to-hcl2 -type HWConfig"; DO NOT EDIT.
package common

import (
	"github.com/hashicorp/hcl/v2/hcldec"
	"github.com/zclconf/go-cty/cty"
)

type FlatHWConfig struct {
	CpuCount   int    `mapstructure:"cpus" required:"false" cty:"cpus"`
	MemorySize int    `mapstructure:"memory" required:"false" cty:"memory"`
	Sound      string `mapstructure:"sound" required:"false" cty:"sound"`
	USB        bool   `mapstructure:"usb" required:"false" cty:"usb"`
}

func (*HWConfig) HCL2Spec() map[string]hcldec.Spec {
	s := map[string]hcldec.Spec{
		"CpuCount":   &hcldec.AttrSpec{Name: "cpus", Type: cty.Number, Required: false},
		"MemorySize": &hcldec.AttrSpec{Name: "memory", Type: cty.Number, Required: false},
		"Sound":      &hcldec.AttrSpec{Name: "sound", Type: cty.String, Required: false},
		"USB":        &hcldec.AttrSpec{Name: "usb", Type: cty.Bool, Required: false},
	}
	return s
}
