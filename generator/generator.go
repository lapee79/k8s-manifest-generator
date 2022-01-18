package generator

import (
	"bytes"
	"encoding/json"
	"github.com/lapee79/k8s-manifest-generator/logger"
	"github.com/lapee79/k8s-manifest-generator/templates"
	"io/ioutil"
	"log"
	"os"
	"sync"
	"text/template"
)

type Resource struct {
	CPU    string `json:"CPU"`
	Memory string `json:"Memory"`
}

type KeyValPair struct {
	Key   string `json:"Key"`
	Value string `json:"Value"`
}

type Application struct {
	Name              string        `json:"name"`
	NameSpace         string        `json:"nameSpace"`
	CommonLabels      *[]KeyValPair `json:"commonLabels"`
	CommonAnnotations *[]KeyValPair `json:"commonAnnotations"`
	ImageUrl          string        `json:"imageUrl"`
	ImageTag          string        `json:"imageTag"`
	ContainerPort     int           `json:"containerPort"`
	ServicePort       int           `json:"servicePort"`
	Config            *[]KeyValPair `json:"config,omitempty"`
	Secret            *[]KeyValPair `json:"secret,omitempty"`
	HealthCheckURL    string        `json:"healthCheckURL"`
	Resources         struct {
		Requests Resource  `json:"requests"`
		Limits   *Resource `json:"limits,omitempty"`
	} `json:"resources"`
	Replicas  *int
	AutoScale *struct {
		MinPodNum int `json:"minPodNum"`
		MaxPodNum int `json:"maxPodNum"`
		CpuUsage  int `json:"cpuUsage"`
		MemUsage  int `json:"memUsage,omitempty"`
	} `json:"autoScale"`
	AzKV         *string `json:"AzKV,omitempty"`
	AzTid        *string `json:"AzTid,omitempty"`
	AzKvSpSecret *string `json:"azKvSpSecret,omitempty"`
}

// Run generates the Kubernetes YAML manifests.
func Run(path string) {
	var app Application
	var tmpls map[string]string
	tmpls = make(map[string]string)

	log.Printf("Reading \"%s\"...\n", path)
	appData, err := ioutil.ReadFile(path)
	logger.Error(err)

	log.Println("Parsing the JSON data...")
	err = json.Unmarshal(appData, &app)
	logger.Error(err)

	if _, err = os.Stat(app.NameSpace); os.IsNotExist(err) {
		err = os.Mkdir(app.NameSpace, os.FileMode(0755))
		logger.Error(err)
	}
	err = os.Chdir(app.NameSpace)
	logger.Error(err)
	if _, err = os.Stat(app.Name); os.IsNotExist(err) {
		err = os.Mkdir(app.Name, os.FileMode(0755))
		logger.Error(err)
	}
	err = os.Chdir(app.Name)
	logger.Error(err)

	tmpls["kustomization.yaml"] = templates.Kustomization
	tmpls["deployment.yaml"] = templates.Deployment
	tmpls["service.yaml"] = templates.Service
	if app.AutoScale != nil && app.Replicas == nil {
		tmpls["hpa.yaml"] = templates.Hpa
	}
	if app.Config != nil {
		tmpls["configmap.yaml"] = templates.ConfigMap
	}
	if app.Secret != nil {
		tmpls["secretproviderclass.yaml"] = templates.SecretProviderClass
	}

	var wg sync.WaitGroup
	for yaml, tmpl := range tmpls {
		wg.Add(1)
		go func(y string, t string) {
			defer wg.Done()
			result, err := Generator(app, t, y)
			logger.Error(err)

			err = ioutil.WriteFile(y, result, os.FileMode(0644))
			logger.Error(err)
			log.Printf("Generated \"%s/%s/%s\".\n", app.NameSpace, app.Name, y)
		}(yaml, tmpl)
	}
	wg.Wait()
}

// Generator returns a parsed template.
func Generator(app Application, tmpl string, path string) ([]byte, error) {
	var tpl bytes.Buffer

	temp := template.Must(template.New(path).Parse(tmpl))
	err := temp.Execute(&tpl, app)
	if err != nil {
		return nil, err
	}

	return tpl.Bytes(), nil
}
