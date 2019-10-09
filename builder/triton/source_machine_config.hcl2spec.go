// Code generated by "mapstructure-to-hcl2 -type MachineImageFilter"; DO NOT EDIT.
package triton

import (
	"github.com/hashicorp/hcl/v2/hcldec"
	"github.com/zclconf/go-cty/cty"
)

type FlatMachineImageFilter struct {
	MostRecent bool   `mapstructure:"most_recent" cty:"most_recent"`
	Name       string `cty:"name"`
	OS         string `cty:"os"`
	Version    string `cty:"version"`
	Public     bool   `cty:"public"`
	State      string `cty:"state"`
	Owner      string `cty:"owner"`
	Type       string `cty:"type"`
}

func (*MachineImageFilter) HCL2Spec() map[string]hcldec.Spec {
	s := map[string]hcldec.Spec{
		"MostRecent": &hcldec.AttrSpec{Name: "most_recent", Type: cty.Bool, Required: false},
		"Name":       &hcldec.AttrSpec{Name: "name", Type: cty.String, Required: false},
		"OS":         &hcldec.AttrSpec{Name: "os", Type: cty.String, Required: false},
		"Version":    &hcldec.AttrSpec{Name: "version", Type: cty.String, Required: false},
		"Public":     &hcldec.AttrSpec{Name: "public", Type: cty.Bool, Required: false},
		"State":      &hcldec.AttrSpec{Name: "state", Type: cty.String, Required: false},
		"Owner":      &hcldec.AttrSpec{Name: "owner", Type: cty.String, Required: false},
		"Type":       &hcldec.AttrSpec{Name: "type", Type: cty.String, Required: false},
	}
	return s
}
