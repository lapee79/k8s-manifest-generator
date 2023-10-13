package generator

import (
	"bytes"
	"encoding/json"
	"github.com/lapee79/k8s-manifest-generator/templates"
	"log"
	"os"
	"sync"
	"text/template"
)

type Resource struct {
	CPU    string `json:"CPU"`
	Memory string `json:"Memory"`
}

type Port struct {
	ContainerPort int `json:"containerPort"`
	ServicePort   int `json:"servicePort"`
}

type KeyValPair struct {
	Key   string `json:"Key"`
	Value string `json:"Value"`
}

type HealthCheck struct {
	InitialDelaySeconds *int `json:"initialDelaySeconds"`
	PeriodSeconds       *int `json:"periodSeconds"`
	TimeoutSeconds      *int `json:"timeoutSeconds"`
	SuccessThreshold    *int `json:"successThreshold"`
	FailureThreshold    *int `json:"failureThreshold"`
	Exec                *struct {
		Command []string `json:"command"`
	} `json:"exec"`
	HttpGet *struct {
		Path       string `json:"path"`
		Port       int    `json:"port"`
		HttpHeader *[]struct {
			Name  string `json:"name"`
			Value string `json:"value"`
		} `json:"httpHeader"`
	} `json:"httpGet"`
	TcpSocket *struct {
		Port int `json:"port"`
	} `json:"tcpSocket"`
}

type Scaler struct {
	Hpa *struct {
		MinPodNum int `json:"minPodNum"`
		MaxPodNum int `json:"maxPodNum"`
		CpuUsage  int `json:"cpuUsage"`
		MemUsage  int `json:"memUsage,omitempty"`
	} `json:"hpa"`
	Vpa *struct {
		UpdateMode string `json:"updateMode"`
	} `json:"vpa"`
}

type Ingress struct {
	IngressAnnotations *[]KeyValPair `json:"ingressAnnotations"`
	IngressClassName   string        `json:"ingressClassName"`
	HostName           string        `json:"hostName"`
	UrlPath            string        `json:"urlPath"`
	TlsSecretName      *string       `json:"tlsSecretName"`
	WebPort            int           `json:"webPort"`
}

type CronJobSpec struct {
	TimeZone                *string `json:"timeZone"`
	Schedule                string  `json:"schedule"`
	RestartPolicy           *string `json:"restartPolicy"`
	ActiveDeadlineSeconds   *int    `json:"activeDeadlineSeconds"`
	TtlSecondsAfterFinished *int    `json:"ttlSecondsAfterFinished"`
	BackoffLimit            *int    `json:"backoffLimit"`
}

type Application struct {
	Name              string        `json:"name"`
	NameSpace         string        `json:"nameSpace"`
	Kind              string        `json:"kind"`
	CommonLabels      *[]KeyValPair `json:"commonLabels"`
	CommonAnnotations *[]KeyValPair `json:"commonAnnotations"`
	ImageUrl          string        `json:"imageUrl"`
	ImageTag          string        `json:"imageTag"`
	CronJobSpec       *CronJobSpec  `json:"cronJobSpec"`
	FileList          *[]string     `json:"fileList"`
	Ports             *[]Port       `json:"ports"`
	Config            *[]KeyValPair `json:"config,omitempty"`
	Secret            *[]KeyValPair `json:"secret,omitempty"`
	ReadinessProbe    *HealthCheck  `json:"readinessProbe"`
	LivenessProbe     *HealthCheck  `json:"livenessProbe"`
	Resources         struct {
		Requests Resource  `json:"requests"`
		Limits   *Resource `json:"limits,omitempty"`
	} `json:"resources"`
	Replicas                   *int     `json:"replicas"`
	AutoScale                  *Scaler  `json:"autoScale"`
	AzKV                       *string  `json:"azKV,omitempty"`
	AzTid                      *string  `json:"azTid,omitempty"`
	AzKvSpSecret               *string  `json:"azKvSpSecret,omitempty"`
	AzKvUserAssignedIdentityID *string  `json:"azKvUserAssignedIdentityID,omitempty"`
	Ingress                    *Ingress `json:"ingress"`
}

// Run generates the Kubernetes YAML manifests.
func Run(path string) error {
	var app Application
	var tmpls map[string]string
	tmpls = make(map[string]string)

	log.Printf("Reading \"%s\"...\n", path)
	appData, err := os.ReadFile(path)
	if err != nil {
		return err
	}

	log.Println("Parsing the JSON data...")
	err = json.Unmarshal(appData, &app)
	if err != nil {
		return err
	}

	if app.Config != nil {
		tmpls["configmap.yaml"] = string(templates.ConfigmapYAML)
	}
	if app.Secret != nil {
		tmpls["secretproviderclass.yaml"] = string(templates.SecretProviderClassYAML)
	}

	switch app.Kind {
	case "Custom":
		tmpls["kustomization.yaml"] = string(templates.KustomizationYAML)
	case "Deployment":
		tmpls["deployment.yaml"] = string(templates.DeploymentYAML)
		if app.Ports != nil {
			tmpls["service.yaml"] = string(templates.ServiceYAML)
		}
		if app.AutoScale.Hpa != nil && app.Replicas == nil {
			tmpls["hpa.yaml"] = string(templates.HpaYAML)
		} else if app.AutoScale.Vpa != nil && app.Replicas != nil {
			tmpls["vpa.yaml"] = string(templates.VpaYAML)
		}
		if app.Ingress != nil {
			tmpls["ingress.yaml"] = string(templates.IngressYAML)
		}
		tmpls["kustomization.yaml"] = string(templates.KustomizationYAML)
	case "CronJob":
		tmpls["cronjob.yaml"] = string(templates.CronjobYAML)
		tmpls["kustomization.yaml"] = string(templates.KustomizationYAML)
	}

	var wg sync.WaitGroup
	c := make(chan error)
	wg.Add(len(tmpls))
	go func() {
		wg.Wait()
		close(c)
	}()
	for yaml, tmpl := range tmpls {
		go func(y string, t string) {
			defer wg.Done()
			result, err := Generator(app, t, y)
			if err != nil {
				c <- err
			}

			err = os.WriteFile(y, result, os.FileMode(0644))
			if err != nil {
				c <- err
			}
			log.Printf("Generated \"%s/%s/%s\".\n", app.NameSpace, app.Name, y)
		}(yaml, tmpl)
	}
	for err = range c {
		if err != nil {
			return err
		}
	}

	return nil
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
