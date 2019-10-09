// Code generated by "mapstructure-to-hcl2 -type ImageDestination"; DO NOT EDIT.
package uhost

import (
	"github.com/hashicorp/hcl/v2/hcldec"
	"github.com/zclconf/go-cty/cty"
)

type FlatImageDestination struct {
	ProjectId   string `mapstructure:"project_id" cty:"project_id"`
	Region      string `mapstructure:"region" cty:"region"`
	Name        string `mapstructure:"name" cty:"name"`
	Description string `mapstructure:"description" cty:"description"`
}

func (*ImageDestination) HCL2Spec() map[string]hcldec.Spec {
	s := map[string]hcldec.Spec{
		"ProjectId":   &hcldec.AttrSpec{Name: "project_id", Type: cty.String, Required: false},
		"Region":      &hcldec.AttrSpec{Name: "region", Type: cty.String, Required: false},
		"Name":        &hcldec.AttrSpec{Name: "name", Type: cty.String, Required: false},
		"Description": &hcldec.AttrSpec{Name: "description", Type: cty.String, Required: false},
	}
	return s
}
