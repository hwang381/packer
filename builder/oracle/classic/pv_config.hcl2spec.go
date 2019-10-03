// Code generated by "hcl2-schema -type PVConfig"; DO NOT EDIT.\n

package classic

import (
	"github.com/hashicorp/hcl/v2/hcldec"
	"github.com/zclconf/go-cty/cty"
)

func (*PVConfig) HCL2Spec() map[string]hcldec.Spec {
	s := map[string]hcldec.Spec{
		"PersistentVolumeSize":      &hcldec.AttrSpec{Name: "persistent_volume_size", Type: cty.Number, Required: false},
		"BuilderUploadImageCommand": &hcldec.AttrSpec{Name: "builder_upload_image_command", Type: cty.String, Required: false},
		"BuilderShape":              &hcldec.AttrSpec{Name: "builder_shape", Type: cty.String, Required: false},
		"BuilderImageList":          &hcldec.AttrSpec{Name: "builder_image_list", Type: cty.String, Required: false},
		"BuilderImageListEntry":     &hcldec.AttrSpec{Name: "builder_image_list_entry", Type: cty.Number, Required: false},
	}
	return s
}
