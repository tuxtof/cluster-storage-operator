package csioperatorclient

import (
	"os"
	"strings"

	configv1 "github.com/openshift/api/config/v1"
)

const (
	NutanixCSIDriverName          = "csi.nutanix.com"
	envNutanixDriverOperatorImage = "NUTANIX_DRIVER_OPERATOR_IMAGE"
	envNutanixDriverImage         = "NUTANIX_DRIVER_IMAGE"
)

func GetNutanixCSIOperatorConfig() CSIOperatorConfig {
	pairs := []string{
		"${OPERATOR_IMAGE}", os.Getenv(envNutanixDriverOperatorImage),
		"${DRIVER_IMAGE}", os.Getenv(envNutanixDriverImage),
	}

	return CSIOperatorConfig{
		CSIDriverName:   NutanixCSIDriverName,
		ConditionPrefix: "NUTANIX",
		Platform:        configv1.NutanixPlatformType,
		StaticAssets: []string{
			"csidriveroperators/nutanix/02_sa.yaml",
			"csidriveroperators/nutanix/03_role.yaml",
			"csidriveroperators/nutanix/04_rolebinding.yaml",
			"csidriveroperators/nutanix/05_clusterrole.yaml",
			"csidriveroperators/nutanix/06_clusterrolebinding.yaml",
		},
		CRAsset:         "csidriveroperators/nutanix/08_cr.yaml",
		DeploymentAsset: "csidriveroperators/nutanix/07_deployment.yaml",
		ImageReplacer:   strings.NewReplacer(pairs...),
		AllowDisabled:   false,
	}
}
