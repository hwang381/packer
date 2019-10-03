// Code generated by "hcl2-schema -type Config"; DO NOT EDIT.\n

package googlecomputeimport

import (
	"github.com/hashicorp/hcl/v2/hcldec"
	"github.com/zclconf/go-cty/cty"
)

func (*Config) HCL2Spec() map[string]hcldec.Spec {
	s := map[string]hcldec.Spec{
		"AccountFile":          &hcldec.AttrSpec{Name: "account_file", Type: cty.String, Required: false},
		"ProjectId":            &hcldec.AttrSpec{Name: "project_id", Type: cty.String, Required: false},
		"Bucket":               &hcldec.AttrSpec{Name: "bucket", Type: cty.String, Required: false},
		"GCSObjectName":        &hcldec.AttrSpec{Name: "gcs_object_name", Type: cty.String, Required: false},
		"ImageDescription":     &hcldec.AttrSpec{Name: "image_description", Type: cty.String, Required: false},
		"ImageFamily":          &hcldec.AttrSpec{Name: "image_family", Type: cty.String, Required: false},
		"ImageGuestOsFeatures": &hcldec.AttrSpec{Name: "image_guest_os_features", Type: cty.List(cty.String), Required: false},
		"ImageLabels":          &hcldec.BlockAttrsSpec{TypeName: "image_labels", ElementType: cty.String, Required: false},
		"ImageName":            &hcldec.AttrSpec{Name: "image_name", Type: cty.String, Required: false},
		"SkipClean":            &hcldec.AttrSpec{Name: "skip_clean", Type: cty.Bool, Required: false},
		"VaultGCPOauthEngine":  &hcldec.AttrSpec{Name: "vault_gcp_oauth_engine", Type: cty.String, Required: false},
	}
	return s
}
