// Code generated by "hcl2-schema -type Config"; DO NOT EDIT.\n

package googlecompute

import (
	"github.com/hashicorp/hcl/v2/hcldec"
	"github.com/zclconf/go-cty/cty"
)

func (*Config) HCL2Spec() map[string]hcldec.Spec {
	s := map[string]hcldec.Spec{
		"AccountFile":                  &hcldec.AttrSpec{Name: "account_file", Type: cty.String, Required: false},
		"ProjectId":                    &hcldec.AttrSpec{Name: "project_id", Type: cty.String, Required: false},
		"AcceleratorType":              &hcldec.AttrSpec{Name: "accelerator_type", Type: cty.String, Required: false},
		"AcceleratorCount":             &hcldec.AttrSpec{Name: "accelerator_count", Type: cty.Number, Required: false},
		"Address":                      &hcldec.AttrSpec{Name: "address", Type: cty.String, Required: false},
		"DisableDefaultServiceAccount": &hcldec.AttrSpec{Name: "disable_default_service_account", Type: cty.Bool, Required: false},
		"DiskName":                     &hcldec.AttrSpec{Name: "disk_name", Type: cty.String, Required: false},
		"DiskSizeGb":                   &hcldec.AttrSpec{Name: "disk_size", Type: cty.Number, Required: false},
		"DiskType":                     &hcldec.AttrSpec{Name: "disk_type", Type: cty.String, Required: false},
		"ImageName":                    &hcldec.AttrSpec{Name: "image_name", Type: cty.String, Required: false},
		"ImageDescription":             &hcldec.AttrSpec{Name: "image_description", Type: cty.String, Required: false},
		"ImageFamily":                  &hcldec.AttrSpec{Name: "image_family", Type: cty.String, Required: false},
		"ImageLabels":                  &hcldec.BlockAttrsSpec{TypeName: "image_labels", ElementType: cty.String, Required: false},
		"ImageLicenses":                &hcldec.AttrSpec{Name: "image_licenses", Type: cty.List(cty.String), Required: false},
		"InstanceName":                 &hcldec.AttrSpec{Name: "instance_name", Type: cty.String, Required: false},
		"Labels":                       &hcldec.BlockAttrsSpec{TypeName: "labels", ElementType: cty.String, Required: false},
		"MachineType":                  &hcldec.AttrSpec{Name: "machine_type", Type: cty.String, Required: false},
		"Metadata":                     &hcldec.BlockAttrsSpec{TypeName: "metadata", ElementType: cty.String, Required: false},
		"MetadataFiles":                &hcldec.BlockAttrsSpec{TypeName: "metadata_files", ElementType: cty.String, Required: false},
		"MinCpuPlatform":               &hcldec.AttrSpec{Name: "min_cpu_platform", Type: cty.String, Required: false},
		"Network":                      &hcldec.AttrSpec{Name: "network", Type: cty.String, Required: false},
		"NetworkProjectId":             &hcldec.AttrSpec{Name: "network_project_id", Type: cty.String, Required: false},
		"OmitExternalIP":               &hcldec.AttrSpec{Name: "omit_external_ip", Type: cty.Bool, Required: false},
		"OnHostMaintenance":            &hcldec.AttrSpec{Name: "on_host_maintenance", Type: cty.String, Required: false},
		"Preemptible":                  &hcldec.AttrSpec{Name: "preemptible", Type: cty.Bool, Required: false},
		"RawStateTimeout":              &hcldec.AttrSpec{Name: "state_timeout", Type: cty.String, Required: false},
		"Region":                       &hcldec.AttrSpec{Name: "region", Type: cty.String, Required: false},
		"Scopes":                       &hcldec.AttrSpec{Name: "scopes", Type: cty.List(cty.String), Required: false},
		"ServiceAccountEmail":          &hcldec.AttrSpec{Name: "service_account_email", Type: cty.String, Required: false},
		"SourceImage":                  &hcldec.AttrSpec{Name: "source_image", Type: cty.String, Required: false},
		"SourceImageFamily":            &hcldec.AttrSpec{Name: "source_image_family", Type: cty.String, Required: false},
		"SourceImageProjectId":         &hcldec.AttrSpec{Name: "source_image_project_id", Type: cty.String, Required: false},
		"StartupScriptFile":            &hcldec.AttrSpec{Name: "startup_script_file", Type: cty.String, Required: false},
		"Subnetwork":                   &hcldec.AttrSpec{Name: "subnetwork", Type: cty.String, Required: false},
		"Tags":                         &hcldec.AttrSpec{Name: "tags", Type: cty.List(cty.String), Required: false},
		"UseInternalIP":                &hcldec.AttrSpec{Name: "use_internal_ip", Type: cty.Bool, Required: false},
		"VaultGCPOauthEngine":          &hcldec.AttrSpec{Name: "vault_gcp_oauth_engine", Type: cty.String, Required: false},
		"Zone":                         &hcldec.AttrSpec{Name: "zone", Type: cty.String, Required: false},
		"image_encryption_key":         &hcldec.BlockObjectSpec{TypeName: "CustomerEncryptionKey", LabelNames: []string(nil), Nested: hcldec.ObjectSpec((&Config{}).ImageEncryptionKey.HCL2Spec())},
	}
	return s
}
