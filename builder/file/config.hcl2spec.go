// Code generated by "hcl2-schema -type Config"; DO NOT EDIT.\n

package file

import (
	"github.com/hashicorp/hcl/v2/hcldec"
	"github.com/zclconf/go-cty/cty"
)

func (*Config) HCL2Spec() map[string]hcldec.Spec {
	s := map[string]hcldec.Spec{
		"Source":  &hcldec.AttrSpec{Name: "source", Type: cty.String, Required: false},
		"Target":  &hcldec.AttrSpec{Name: "target", Type: cty.String, Required: false},
		"Content": &hcldec.AttrSpec{Name: "content", Type: cty.String, Required: false},
	}
	return s
}
