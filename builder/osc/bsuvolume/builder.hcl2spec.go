// Code generated by "hcl2-schema -type Config"; DO NOT EDIT.\n

package bsuvolume

import (
	"github.com/hashicorp/hcl/v2/hcldec"
)

func (*Config) HCL2Spec() map[string]hcldec.Spec {
	s := map[string]hcldec.Spec{
		"bsu_volumes": &hcldec.BlockListSpec{TypeName: "[]BlockDevice", Nested: hcldec.ObjectSpec((&Config{}).VolumeMappings[0].HCL2Spec())},
	}
	for k, v := range (&Config{}).AccessConfig.HCL2Spec() {
		s[k] = v
	}
	for k, v := range (&Config{}).RunConfig.HCL2Spec() {
		s[k] = v
	}
	return s
}
