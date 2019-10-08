// Code generated by "mapstructure-to-hcl2 -type GuestAdditionsConfig"; DO NOT EDIT.
package common

import (
	"github.com/hashicorp/hcl/v2/hcldec"
	"github.com/zclconf/go-cty/cty"
)

type FlatGuestAdditionsConfig struct {
	Communicator       string `mapstructure:"communicator" cty:"communicator"`
	GuestAdditionsMode string `mapstructure:"guest_additions_mode" required:"false" cty:"guest_additions_mode"`
}

func (*GuestAdditionsConfig) HCL2Spec() map[string]hcldec.Spec {
	s := map[string]hcldec.Spec{
		"Communicator":       &hcldec.AttrSpec{Name: "communicator", Type: cty.String, Required: false},
		"GuestAdditionsMode": &hcldec.AttrSpec{Name: "guest_additions_mode", Type: cty.String, Required: false},
	}
	return s
}
