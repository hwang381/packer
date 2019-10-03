// Code generated by "hcl2-schema -type BlockDevice,BlockDevices,OMIBlockDevices,LaunchBlockDevices"; DO NOT EDIT.\n

package common

import (
	"github.com/hashicorp/hcl/v2/hcldec"
	"github.com/zclconf/go-cty/cty"
)

func (*BlockDevice) HCL2Spec() map[string]hcldec.Spec {
	s := map[string]hcldec.Spec{
		"DeleteOnVmDeletion": &hcldec.AttrSpec{Name: "delete_on_vm_deletion", Type: cty.Bool, Required: false},
		"DeviceName":         &hcldec.AttrSpec{Name: "device_name", Type: cty.String, Required: false},
		"IOPS":               &hcldec.AttrSpec{Name: "iops", Type: cty.Number, Required: false},
		"NoDevice":           &hcldec.AttrSpec{Name: "no_device", Type: cty.Bool, Required: false},
		"SnapshotId":         &hcldec.AttrSpec{Name: "snapshot_id", Type: cty.String, Required: false},
		"VirtualName":        &hcldec.AttrSpec{Name: "virtual_name", Type: cty.String, Required: false},
		"VolumeType":         &hcldec.AttrSpec{Name: "volume_type", Type: cty.String, Required: false},
		"VolumeSize":         &hcldec.AttrSpec{Name: "volume_size", Type: cty.Number, Required: false},
	}
	return s
}

func (*BlockDevices) HCL2Spec() map[string]hcldec.Spec {
	s := map[string]hcldec.Spec{}
	for k, v := range (&BlockDevices{}).OMIBlockDevices.HCL2Spec() {
		s[k] = v
	}
	for k, v := range (&BlockDevices{}).LaunchBlockDevices.HCL2Spec() {
		s[k] = v
	}
	return s
}

func (*OMIBlockDevices) HCL2Spec() map[string]hcldec.Spec {
	s := map[string]hcldec.Spec{
		"omi_block_device_mappings": &hcldec.BlockListSpec{TypeName: "[]BlockDevice", Nested: hcldec.ObjectSpec((&OMIBlockDevices{}).OMIMappings[0].HCL2Spec())},
	}
	return s
}

func (*LaunchBlockDevices) HCL2Spec() map[string]hcldec.Spec {
	s := map[string]hcldec.Spec{
		"launch_block_device_mappings": &hcldec.BlockListSpec{TypeName: "[]BlockDevice", Nested: hcldec.ObjectSpec((&LaunchBlockDevices{}).LaunchMappings[0].HCL2Spec())},
	}
	return s
}
