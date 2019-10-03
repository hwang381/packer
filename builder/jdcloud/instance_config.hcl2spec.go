// Code generated by "hcl2-schema -type JDCloudInstanceSpecConfig"; DO NOT EDIT.\n

package jdcloud

import (
	"github.com/hashicorp/hcl/v2/hcldec"
	"github.com/zclconf/go-cty/cty"
)

func (*JDCloudInstanceSpecConfig) HCL2Spec() map[string]hcldec.Spec {
	s := map[string]hcldec.Spec{
		"ImageId":         &hcldec.AttrSpec{Name: "image_id", Type: cty.String, Required: false},
		"InstanceName":    &hcldec.AttrSpec{Name: "instance_name", Type: cty.String, Required: false},
		"InstanceType":    &hcldec.AttrSpec{Name: "instance_type", Type: cty.String, Required: false},
		"ImageName":       &hcldec.AttrSpec{Name: "image_name", Type: cty.String, Required: false},
		"SubnetId":        &hcldec.AttrSpec{Name: "subnet_id", Type: cty.String, Required: false},
		"InstanceId":      &hcldec.AttrSpec{Name: "instanceid", Type: cty.String, Required: false},
		"ArtifactId":      &hcldec.AttrSpec{Name: "artifactid", Type: cty.String, Required: false},
		"PublicIpAddress": &hcldec.AttrSpec{Name: "publicipaddress", Type: cty.String, Required: false},
		"PublicIpId":      &hcldec.AttrSpec{Name: "publicipid", Type: cty.String, Required: false},
	}
	return s
}
