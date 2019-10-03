// Code generated by "hcl2-schema -type AwsAccessConfig"; DO NOT EDIT.\n

package docker

import (
	"github.com/hashicorp/hcl/v2/hcldec"
	"github.com/zclconf/go-cty/cty"
)

func (*AwsAccessConfig) HCL2Spec() map[string]hcldec.Spec {
	s := map[string]hcldec.Spec{
		"AccessKey": &hcldec.AttrSpec{Name: "aws_access_key", Type: cty.String, Required: false},
		"SecretKey": &hcldec.AttrSpec{Name: "aws_secret_key", Type: cty.String, Required: false},
		"Token":     &hcldec.AttrSpec{Name: "aws_token", Type: cty.String, Required: false},
		"Profile":   &hcldec.AttrSpec{Name: "aws_profile", Type: cty.String, Required: false},
	}
	return s
}
