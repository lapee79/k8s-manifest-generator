package generator

import (
	"encoding/json"
	"github.com/stretchr/testify/assert"
	"log"
	"os"
	"path/filepath"
	"testing"
)

func TestGenerator(t *testing.T) {
	var tests = []struct {
		testName   string
		appFile    string
		tmpl       string
		wantResult string
	}{
		{
			testName:   "GenerateConfigMap",
			appFile:    "test/data/input-json/web-deployment-with-azkv.json",
			tmpl:       "templates/configmap.goyaml",
			wantResult: "test/data/output-yaml/want-configmap.yaml",
		},
		{
			testName:   "GenerateSecretProviderClassWithAzureKVSPSecret",
			appFile:    "test/data/input-json/web-deployment-with-azkv.json",
			tmpl:       "templates/secretproviderclass.goyaml",
			wantResult: "test/data/output-yaml/want-secretproviderclass-azkv-sp-secret.yaml",
		},
		{
			testName:   "GenerateSecretProviderClassWithAzureKVUserAssignedIdentityID",
			appFile:    "/test/data/input-json/web-deployment-with-azkv-userassignedidentityid.json",
			tmpl:       "templates/secretproviderclass.goyaml",
			wantResult: "test/data/output-yaml/want-secretproviderclass-azkv-userassignedidentity.yaml",
		},
		{
			testName:   "GenerateDeploymentWithAzureKV",
			appFile:    "test/data/input-json/web-deployment-with-azkv.json",
			tmpl:       "templates/deployment.goyaml",
			wantResult: "test/data/output-yaml/want-deployment-with-azkv.yaml",
		},
		{
			testName:   "GenerateDeploymentWithExecHealthcheck",
			appFile:    "test/data/input-json/web-deployment-without-azkv.json",
			tmpl:       "templates/deployment.goyaml",
			wantResult: "test/data/output-yaml/want-deployment-exec-healthcheck.yaml",
		},
		{
			testName:   "GenerateDeploymentWithTCPHealthcheck",
			appFile:    "test/data/input-json/web-deployment-without-hpa.json",
			tmpl:       "templates/deployment.goyaml",
			wantResult: "test/data/output-yaml/want-deployment-tcp-healthcheck.yaml",
		},
		{
			testName:   "GenerateService",
			appFile:    "test/data/input-json/web-deployment-with-azkv.json",
			tmpl:       "templates/service.goyaml",
			wantResult: "test/data/output-yaml/want-service.yaml",
		},
		{
			testName:   "GenerateHpaWithCpuAndMemory",
			appFile:    "test/data/input-json/web-deployment-with-azkv.json",
			tmpl:       "templates/hpa.goyaml",
			wantResult: "test/data/output-yaml/want-hpa-cpu-memory.yaml",
		},
		{
			testName:   "GenerateHpaWithCpuOnly",
			appFile:    "test/data/input-json/web-deployment-without-azkv.json",
			tmpl:       "templates/hpa.goyaml",
			wantResult: "test/data/output-yaml/want-hpa-cpu-only.yaml",
		},
		{
			testName:   "GenerateKustomizationWithSecret",
			appFile:    "test/data/input-json/web-deployment-with-azkv.json",
			tmpl:       "templates/kustomization.goyaml",
			wantResult: "test/data/output-yaml/want-kustomization-deployment-with-secret.yaml",
		},
		{
			testName:   "GenerateKustomizationWithoutSecret",
			appFile:    "test/data/input-json/web-deployment-without-azkv.json",
			tmpl:       "templates/kustomization.goyaml",
			wantResult: "test/data/output-yaml/want-kustomiaztion-deployment-without-secret.yaml",
		},
		{
			testName:   "GenerateCronjobWithAzureKV",
			appFile:    "test/data/input-json/cronjob-with-azkv.json",
			tmpl:       "templates/cronjob.goyaml",
			wantResult: "test/data/output-yaml/want-cronjob-with-secret.yaml",
		},
		{
			testName:   "GenerateCronjobWithoutAzureKV",
			appFile:    "test/data/input-json/cronjob-without-azkv.json",
			tmpl:       "templates/cronjob.goyaml",
			wantResult: "test/data/output-yaml/want-cronjob-without-secret.yaml",
		},
	}

	for _, tt := range tests {
		t.Run(tt.testName, func(t *testing.T) {
			abs, err := filepath.Abs("..")
			if err != nil {
				log.Fatalln(err)
			}
			appJson, err := os.ReadFile(filepath.Join(abs, tt.appFile))
			if err != nil {
				log.Fatalln(err)
			}
			var app Application
			err = json.Unmarshal(appJson, &app)
			if err != nil {
				log.Fatalln(err)
			}

			template, err := os.ReadFile(filepath.Join(abs, tt.tmpl))
			if err != nil {
				log.Fatalln(err)
			}

			r, err := Generator(app, string(template), tt.testName)
			if err != nil {
				log.Fatalln(err)
			}

			wantResult, err := os.ReadFile(filepath.Join(abs, tt.wantResult))
			if err != nil {
				log.Fatalln(err)
			}
			assert.Equal(t, string(wantResult), string(r))
		})
	}
}
