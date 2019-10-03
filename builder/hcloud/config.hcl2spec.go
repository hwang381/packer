// Code generated by "hcl2-schema -type Config,imageFilter"; DO NOT EDIT.\n

package hcloud

import (
	"github.com/hashicorp/hcl/v2/hcldec"
	"github.com/zclconf/go-cty/cty"
)

func (*Config) HCL2Spec() map[string]hcldec.Spec {
	s := map[string]hcldec.Spec{
		"HCloudToken":    &hcldec.AttrSpec{Name: "token", Type: cty.String, Required: false},
		"Endpoint":       &hcldec.AttrSpec{Name: "endpoint", Type: cty.String, Required: false},
		"PollInterval":   &hcldec.AttrSpec{Name: "poll_interval", Type: cty.String, Required: false},
		"ServerName":     &hcldec.AttrSpec{Name: "server_name", Type: cty.String, Required: false},
		"Location":       &hcldec.AttrSpec{Name: "location", Type: cty.String, Required: false},
		"ServerType":     &hcldec.AttrSpec{Name: "server_type", Type: cty.String, Required: false},
		"Image":          &hcldec.AttrSpec{Name: "image", Type: cty.String, Required: false},
		"SnapshotName":   &hcldec.AttrSpec{Name: "snapshot_name", Type: cty.String, Required: false},
		"SnapshotLabels": &hcldec.BlockAttrsSpec{TypeName: "snapshot_labels", ElementType: cty.String, Required: false},
		"UserData":       &hcldec.AttrSpec{Name: "user_data", Type: cty.String, Required: false},
		"UserDataFile":   &hcldec.AttrSpec{Name: "user_data_file", Type: cty.String, Required: false},
		"SSHKeys":        &hcldec.AttrSpec{Name: "ssh_keys", Type: cty.List(cty.String), Required: false},
		"RescueMode":     &hcldec.AttrSpec{Name: "rescue", Type: cty.String, Required: false},
		"image_filter":   &hcldec.BlockObjectSpec{TypeName: "*imageFilter", LabelNames: []string(nil), Nested: hcldec.ObjectSpec((&Config{}).ImageFilter.HCL2Spec())},
	}
	return s
}

func (*imageFilter) HCL2Spec() map[string]hcldec.Spec {
	s := map[string]hcldec.Spec{
		"WithSelector": &hcldec.AttrSpec{Name: "with_selector", Type: cty.List(cty.String), Required: false},
		"MostRecent":   &hcldec.AttrSpec{Name: "most_recent", Type: cty.Bool, Required: false},
	}
	return s
}
