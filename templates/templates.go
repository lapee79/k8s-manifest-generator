package templates

import (
	_ "embed"
)

var (
	//go:embed configmap.goyaml
	ConfigmapYAML []byte

	//go:embed cronjob.goyaml
	CronjobYAML []byte

	//go:embed deployment.goyaml
	DeploymentYAML []byte

	//go:embed hpa.goyaml
	HpaYAML []byte

	//go:embed vpa.goyaml
	VpaYAML []byte

	//go:embed kustomization.goyaml
	KustomizationYAML []byte

	//go:embed secretproviderclass.goyaml
	SecretProviderClassYAML []byte

	//go:embed service.goyaml
	ServiceYAML []byte
)
