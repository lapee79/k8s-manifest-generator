package generator

import (
	"encoding/json"
	"github.com/lapee79/k8s-manifest-generator/templates"
	"github.com/stretchr/testify/assert"
	"log"
	"os"
	"testing"
)

func TestGenerator(t *testing.T) {
	webDeploymentWithAzkvJson, err := os.ReadFile("./test/data/input-json/web-deployment-with-azkv.json")
	if err != nil {
		log.Fatalln(err)
	}
	var webDeploymentWithAzkv Application
	err = json.Unmarshal([]byte(webDeploymentWithAzkvJson), &webDeploymentWithAzkv)
	if err != nil {
		log.Fatalln(err)
	}

	webDeploymentWithoutAzkvJson, err := os.ReadFile("./test/data/input-json/web-deployment-without-azkv.json")
	if err != nil {
		log.Fatalln(err)
	}
	var webDeploymentWithoutAzkv Application
	err = json.Unmarshal([]byte(webDeploymentWithoutAzkvJson), &webDeploymentWithoutAzkv)
	if err != nil {
		log.Fatalln(err)
	}

	webDeploymentWithoutHpaJson, err := os.ReadFile("./test/data/input-json/web-deployment-without-hpa.json")
	if err != nil {
		log.Fatalln(err)
	}
	var webDeploymentWithoutHpa Application
	err = json.Unmarshal([]byte(webDeploymentWithoutHpaJson), &webDeploymentWithoutHpa)
	if err != nil {
		log.Fatalln(err)
	}

	webDeploymentWithAzkvUserAssignedidentityIDJson, err := os.ReadFile("./test/data/input-json/web-deployment-with-azkv-userassignedidentityid.json")
	if err != nil {
		log.Fatalln(err)
	}
	var webDeploymentWithAzkvUserAssignedidentityID Application
	err = json.Unmarshal([]byte(webDeploymentWithAzkvUserAssignedidentityIDJson), &webDeploymentWithAzkvUserAssignedidentityID)
	if err != nil {
		log.Fatalln(err)
	}

	wantConfigmapResult, err := os.ReadFile("./test/data/output-yaml/want-configmap.yaml")
	if err != nil {
		log.Fatalln(err)
	}

	wantSecretproviderclassAzKvSpSecretResult, err := os.ReadFile("./test/data/output-yaml/want-secretproviderclass-azkv-sp-secret.yaml")
	if err != nil {
		log.Fatalln(err)
	}

	wantSecretproviderclassAzKvUserAssignedIdentityIDResult, err := os.ReadFile("./test/data/output-yaml/want-secretproviderclass-azkv-userassignedidentity.yaml")
	if err != nil {
		log.Fatalln(err)
	}

	wantDeploymentWithAzkv, err := os.ReadFile("./test/data/output-yaml/want-deployment-with-azkv.yaml")
	if err != nil {
		log.Fatalln(err)
	}

	wantDeploymentExecHealthcheck, err := os.ReadFile("./test/data/output-yaml/want-deployment-exec-healthcheck.yaml")
	if err != nil {
		log.Fatalln(err)
	}

	wantDeploymentTcpHealthcheck, err := os.ReadFile("./test/data/output-yaml/want-deployment-tcp-healthcheck.yaml")
	if err != nil {
		log.Fatalln(err)
	}

	wantServiceResult, err := os.ReadFile("./test/data/output-yaml/want-service.yaml")
	if err != nil {
		log.Fatalln(err)
	}

	wantHpaCpuMemory, err := os.ReadFile("./test/data/output-yaml/want-hpa-cpu-memory.yaml")
	if err != nil {
		log.Fatalln(err)
	}

	wantHpaCpuOnly, err := os.ReadFile("./test/data/output-yaml/want-hpa-cpu-only.yaml")
	if err != nil {
		log.Fatalln(err)
	}

	wantKustomizationDeploymentWithSecret, err := os.ReadFile("./test/data/output-yaml/want-kustomization-deployment-with-secret.yaml")
	if err != nil {
		log.Fatalln(err)
	}
	wantKustomizationResult02 := `apiVersion: kustomize.config.k8s.io/v1beta1
kind: Kustomization
commonLabels:
  app.kubernetes.io/instance: "webSvc1"
  app.kubernetes.io/environment: "dev2"
commonAnnotations:
  commitAuther: "lapee79"
  buildId: "6776f266"
resources:
- service.yaml
- deployment.yaml
- hpa.yaml
images:
- name: private-image
  newName: "artifactory-dev.nowcom.io/docker/nowcom.services.bookingwfs"
  newTag: "6776f266"
`

	var tests = []struct {
		testName   string
		app        Application
		tmpl       string
		wantResult string
	}{
		{testName: "GenerateConfigMap", app: webDeploymentWithAzkv, tmpl: templates.ConfigMap, wantResult: string(wantConfigmapResult)},
		{testName: "GenerateSecretProviderClassAzKvSpSecret", app: webDeploymentWithAzkv, tmpl: templates.SecretProviderClass, wantResult: string(wantSecretproviderclassAzKvSpSecretResult)},
		{testName: "GenerateSecretProviderClassAzKvUserAssignedIdentityID", app: webDeploymentWithAzkvUserAssignedidentityID, tmpl: templates.SecretProviderClass, wantResult: string(wantSecretproviderclassAzKvUserAssignedIdentityIDResult)},
		{testName: "GenerateDeployment01", app: webDeploymentWithAzkv, tmpl: templates.Deployment, wantResult: string(wantDeploymentWithAzkv)},
		{testName: "GenerateDeployment02", app: webDeploymentWithoutAzkv, tmpl: templates.Deployment, wantResult: string(wantDeploymentExecHealthcheck)},
		{testName: "GenerateDeployment03", app: webDeploymentWithoutHpa, tmpl: templates.Deployment, wantResult: string(wantDeploymentTcpHealthcheck)},
		{testName: "GenerateService", app: webDeploymentWithAzkv, tmpl: templates.Service, wantResult: string(wantServiceResult)},
		{testName: "GenerateHPA01", app: webDeploymentWithAzkv, tmpl: templates.Hpa, wantResult: string(wantHpaCpuMemory)},
		{testName: "GenerateHPA02", app: webDeploymentWithoutAzkv, tmpl: templates.Hpa, wantResult: string(wantHpaCpuOnly)},
		{testName: "GenerateKustomization01", app: webDeploymentWithAzkv, tmpl: templates.Kustomization, wantResult: string(wantKustomizationDeploymentWithSecret)},
		{testName: "GenerateKustomization02", app: webDeploymentWithoutAzkv, tmpl: templates.Kustomization, wantResult: wantKustomizationResult02},
	}

	for _, tt := range tests {
		t.Run(tt.testName, func(t *testing.T) {
			r, err := Generator(tt.app, tt.tmpl, tt.testName)
			if err != nil {
				log.Fatalln(err)
			}
			assert.Equal(t, tt.wantResult, string(r))
		})
	}
}
