// Code generated by "hcl2-schema -type RunConfig,ImageFilter,ImageFilterOptions"; DO NOT EDIT.\n

package openstack

import (
	"github.com/hashicorp/hcl/v2/hcldec"
	"github.com/zclconf/go-cty/cty"
)

func (*RunConfig) HCL2Spec() map[string]hcldec.Spec {
	s := map[string]hcldec.Spec{
		"SSHInterface":           &hcldec.AttrSpec{Name: "ssh_interface", Type: cty.String, Required: false},
		"SSHIPVersion":           &hcldec.AttrSpec{Name: "ssh_ip_version", Type: cty.String, Required: false},
		"SourceImage":            &hcldec.AttrSpec{Name: "source_image", Type: cty.String, Required: false},
		"SourceImageName":        &hcldec.AttrSpec{Name: "source_image_name", Type: cty.String, Required: false},
		"Flavor":                 &hcldec.AttrSpec{Name: "flavor", Type: cty.String, Required: false},
		"AvailabilityZone":       &hcldec.AttrSpec{Name: "availability_zone", Type: cty.String, Required: false},
		"RackconnectWait":        &hcldec.AttrSpec{Name: "rackconnect_wait", Type: cty.Bool, Required: false},
		"FloatingIPNetwork":      &hcldec.AttrSpec{Name: "floating_ip_network", Type: cty.String, Required: false},
		"InstanceFloatingIPNet":  &hcldec.AttrSpec{Name: "instance_floating_ip_net", Type: cty.String, Required: false},
		"FloatingIP":             &hcldec.AttrSpec{Name: "floating_ip", Type: cty.String, Required: false},
		"ReuseIPs":               &hcldec.AttrSpec{Name: "reuse_ips", Type: cty.Bool, Required: false},
		"SecurityGroups":         &hcldec.AttrSpec{Name: "security_groups", Type: cty.List(cty.String), Required: false},
		"Networks":               &hcldec.AttrSpec{Name: "networks", Type: cty.List(cty.String), Required: false},
		"Ports":                  &hcldec.AttrSpec{Name: "ports", Type: cty.List(cty.String), Required: false},
		"UserData":               &hcldec.AttrSpec{Name: "user_data", Type: cty.String, Required: false},
		"UserDataFile":           &hcldec.AttrSpec{Name: "user_data_file", Type: cty.String, Required: false},
		"InstanceName":           &hcldec.AttrSpec{Name: "instance_name", Type: cty.String, Required: false},
		"InstanceMetadata":       &hcldec.BlockAttrsSpec{TypeName: "instance_metadata", ElementType: cty.String, Required: false},
		"ForceDelete":            &hcldec.AttrSpec{Name: "force_delete", Type: cty.Bool, Required: false},
		"ConfigDrive":            &hcldec.AttrSpec{Name: "config_drive", Type: cty.Bool, Required: false},
		"FloatingIPPool":         &hcldec.AttrSpec{Name: "floating_ip_pool", Type: cty.String, Required: false},
		"UseBlockStorageVolume":  &hcldec.AttrSpec{Name: "use_blockstorage_volume", Type: cty.Bool, Required: false},
		"VolumeName":             &hcldec.AttrSpec{Name: "volume_name", Type: cty.String, Required: false},
		"VolumeType":             &hcldec.AttrSpec{Name: "volume_type", Type: cty.String, Required: false},
		"VolumeSize":             &hcldec.AttrSpec{Name: "volume_size", Type: cty.Number, Required: false},
		"VolumeAvailabilityZone": &hcldec.AttrSpec{Name: "volume_availability_zone", Type: cty.String, Required: false},
		"OpenstackProvider":      &hcldec.AttrSpec{Name: "openstack_provider", Type: cty.String, Required: false},
		"UseFloatingIp":          &hcldec.AttrSpec{Name: "use_floating_ip", Type: cty.Bool, Required: false},
		"source_image_filter":    &hcldec.BlockObjectSpec{TypeName: "ImageFilter", LabelNames: []string(nil), Nested: hcldec.ObjectSpec((&RunConfig{}).SourceImageFilters.HCL2Spec())},
	}
	return s
}

func (*ImageFilter) HCL2Spec() map[string]hcldec.Spec {
	s := map[string]hcldec.Spec{
		"MostRecent": &hcldec.AttrSpec{Name: "most_recent", Type: cty.Bool, Required: false},
		"filters":    &hcldec.BlockObjectSpec{TypeName: "ImageFilterOptions", LabelNames: []string(nil), Nested: hcldec.ObjectSpec((&ImageFilter{}).Filters.HCL2Spec())},
	}
	return s
}

func (*ImageFilterOptions) HCL2Spec() map[string]hcldec.Spec {
	s := map[string]hcldec.Spec{
		"Name":       &hcldec.AttrSpec{Name: "name", Type: cty.String, Required: false},
		"Owner":      &hcldec.AttrSpec{Name: "owner", Type: cty.String, Required: false},
		"Tags":       &hcldec.AttrSpec{Name: "tags", Type: cty.List(cty.String), Required: false},
		"Visibility": &hcldec.AttrSpec{Name: "visibility", Type: cty.String, Required: false},
		"Properties": &hcldec.BlockAttrsSpec{TypeName: "properties", ElementType: cty.String, Required: false},
	}
	return s
}
