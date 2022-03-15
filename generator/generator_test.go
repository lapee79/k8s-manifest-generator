package generator

import (
	"encoding/json"
	"github.com/lapee79/k8s-manifest-generator/templates"
	"github.com/stretchr/testify/assert"
	"io/ioutil"
	"log"
	"os"
	"path/filepath"
	"testing"
)

var sampleAppJSON01 = `
{
  "Name": "webSvc1",
  "NameSpace": "test",
  "ContainerPort": 80,
  "ServicePort": 80,
  "CommonLabels": [
    {"Key": "app.kubernetes.io/instance", "Value": "webSvc1"},
    {"Key": "app.kubernetes.io/environment", "Value": "dev2"}
  ],
  "CommonAnnotations": [
    {"Key": "commitAuther", "Value": "lapee79"},
    {"Key": "buildId", "Value": "6776f266"}
  ],
  "ImageUrl": "artifactory-dev.nowcom.io/docker/nowcom.services.bookingwfs",
  "ImageTag": "6776f266",
  "Config": [
    {"Key": "ConfKey1", "Value": "ConfVal1"},
    {"Key": "ConfKey2", "Value": "ConfVal2"}
  ],
  "Secret": [
    {"Key": "SecKey1", "Value": "SecVal1"},
    {"Key": "SecKey2", "Value": "SecVal2"}
  ],
  "ReadinessProbe": {
    "InitialDelaySeconds": 20,
    "PeriodSeconds": 10,
    "TimeoutSeconds": 10,
    "SuccessThreshold": 1,
    "FailureThreshold": 3,
    "HttpGet": {
      "Path": "/healthz",
      "Port": 80,
      "HttpHeader": [
        {"Name": "Custom-Header-1", "Value": "Awesome1"},
        {"Name": "Custom-Header-2", "Value": "Awesome2"}
      ]
    }
  },
  "LivenessProbe": {
    "InitialDelaySeconds": 20,
    "PeriodSeconds": 10,
    "TimeoutSeconds": 10,
    "SuccessThreshold": 1,
    "FailureThreshold": 3,
    "HttpGet": {
      "Path": "/healthz",
      "Port": 80,
      "HttpHeader": [
        {"Name": "Custom-Header-3", "Value": "Awesome3"}
      ]
    }
  },
  "Resources": {
    "Requests": {
      "CPU": "100m",
      "Memory": "128Mi"
    },
    "Limits": {
      "CPU": "200m",
      "Memory": "256Mi"
    }
  },
  "AutoScale": {
    "MinPodNum": 1,
    "MaxPodNum": 10,
    "CpuUsage": 40,
    "MemUsage": 90
  },
  "AzKV":  "az-kv-01",
  "AzTid": "1234-12345678-00000000",
  "AzKvSpSecret": "secrets-store-creds"
}
`
var sampleAppJSON02 = `
{
  "Name": "webSvc1",
  "NameSpace": "test",
  "ContainerPort": 80,
  "ServicePort": 80,
  "CommonLabels": [
    {"Key": "app.kubernetes.io/instance", "Value": "webSvc1"},
    {"Key": "app.kubernetes.io/environment", "Value": "dev2"}
  ],
  "CommonAnnotations": [
    {"Key": "commitAuther", "Value": "lapee79"},
    {"Key": "buildId", "Value": "6776f266"}
  ],
  "ImageUrl": "artifactory-dev.nowcom.io/docker/nowcom.services.bookingwfs",
  "ImageTag": "6776f266",
  "LivenessProbe": {
    "Exec": {
      "Command": [
        "cat",
        "/tmp/test"
      ]
    }
  },
  "Resources": {
    "Requests": {
      "CPU": "100m",
      "Memory": "128Mi"
    }
  },
  "AutoScale": {
    "MinPodNum": 1,
    "MaxPodNum": 10,
    "CpuUsage": 40
  }
}
`

var sampleAppJSON03 = `
{
  "Name": "webSvc1",
  "NameSpace": "test",
  "ContainerPort": 80,
  "ServicePort": 80,
  "CommonLabels": [
    {"Key": "app.kubernetes.io/instance", "Value": "webSvc1"},
    {"Key": "app.kubernetes.io/environment", "Value": "dev2"}
  ],
  "CommonAnnotations": [
    {"Key": "commitAuther", "Value": "lapee79"},
    {"Key": "buildId", "Value": "6776f266"}
  ],
  "ImageUrl": "artifactory-dev.nowcom.io/docker/nowcom.services.bookingwfs",
  "ImageTag": "6776f266",
  "LivenessProbe": {
    "TcpSocket": {
      "Port": 80
    }
  },
  "Resources": {
    "Requests": {
      "CPU": "100m",
      "Memory": "128Mi"
    }
  },
  "Replicas": 3
}
`

var sampleAppJSON04 = `
{
  "Name": "webSvc1",
  "NameSpace": "test",
  "ContainerPort": 80,
  "ServicePort": 80,
  "CommonLabels": [
    {"Key": "app.kubernetes.io/instance", "Value": "webSvc1"},
    {"Key": "app.kubernetes.io/environment", "Value": "dev2"}
  ],
  "CommonAnnotations": [
    {"Key": "commitAuther", "Value": "lapee79"},
    {"Key": "buildId", "Value": "6776f266"}
  ],
  "ImageUrl": "artifactory-dev.nowcom.io/docker/nowcom.services.bookingwfs",
  "ImageTag": "6776f266",
  "Config": [
    {"Key": "ConfKey1", "Value": "ConfVal1"},
    {"Key": "ConfKey2", "Value": "ConfVal2"}
  ],
  "Secret": [
    {"Key": "SecKey1", "Value": "SecVal1"},
    {"Key": "SecKey2", "Value": "SecVal2"}
  ],
  "Resources": {
    "Requests": {
      "CPU": "100m",
      "Memory": "128Mi"
    },
    "Limits": {
      "CPU": "200m",
      "Memory": "256Mi"
    }
  },
  "AutoScale": {
    "MinPodNum": 1,
    "MaxPodNum": 10,
    "CpuUsage": 40,
    "MemUsage": 90
  },
  "AzKV":  "az-kv-01",
  "AzTid": "1234-12345678-00000000",
  "AzKvUserAssignedIdentityID": "12345678-1234-1234-1234-000000000000"
}
`

func TestRun(t *testing.T) {
	t.Run("Run", func(t *testing.T) {
		cwd, err := os.Getwd()
		if err != nil {
			log.Fatalln(err)
		}
		testAppJson := "testAppJson.json"
		testAppJsonPath := filepath.Join(cwd, testAppJson)
		err = ioutil.WriteFile(testAppJsonPath, []byte(sampleAppJSON01), os.FileMode(0644))
		if err != nil {
			log.Fatalln(err)
		}

		defer func() {
			var sampleApp01 Application

			err = json.Unmarshal([]byte(sampleAppJSON01), &sampleApp01)
			if err != nil {
				log.Fatalln(err)
			}

			err = os.Remove(testAppJsonPath)
			if err != nil {
				log.Fatalln(err)
			}

			namespacePath := filepath.Join(cwd, sampleApp01.NameSpace)
			err = os.RemoveAll(namespacePath)
			if err != nil {
				log.Fatalln(err)
			}
		}()

		err = Run(testAppJsonPath)
		assert.Nil(t, err)
	})
}

func TestGenerator(t *testing.T) {
	var sampleApp01 Application

	err := json.Unmarshal([]byte(sampleAppJSON01), &sampleApp01)
	if err != nil {
		log.Fatalln(err)
	}

	var sampleApp02 Application

	err = json.Unmarshal([]byte(sampleAppJSON02), &sampleApp02)
	if err != nil {
		log.Fatalln(err)
	}

	var sampleApp03 Application

	err = json.Unmarshal([]byte(sampleAppJSON03), &sampleApp03)
	if err != nil {
		log.Fatalln(err)
	}

	var sampleApp04 Application

	err = json.Unmarshal([]byte(sampleAppJSON04), &sampleApp04)
	if err != nil {
		log.Fatalln(err)
	}

	wantConfigmapResult := `apiVersion: v1
kind: ConfigMap
metadata:
  name: webSvc1
data:
  ConfKey1: "ConfVal1"
  ConfKey2: "ConfVal2"
`
	wantSecretproviderclassAzKvSpSecretResult := `apiVersion: secrets-store.csi.x-k8s.io/v1
kind: SecretProviderClass
metadata:
  name: webSvc1
spec:
  provider: azure
  secretObjects:
  - secretName: webSvc1
    type: Opaque
    data:
    - objectName: SecVal1
      key: SecKey1
    - objectName: SecVal2
      key: SecKey2
  parameters:
    usePodIdentity: "false"
    keyvaultName: "az-kv-01"
    objects: |
      array:
        - |
          objectName: SecVal1
          objectType: secret
        - |
          objectName: SecVal2
          objectType: secret
    tenantId: "1234-12345678-00000000"
`
	wantSecretproviderclassAzKvUserAssignedIdentityIDResult := `apiVersion: secrets-store.csi.x-k8s.io/v1
kind: SecretProviderClass
metadata:
  name: webSvc1
spec:
  provider: azure
  secretObjects:
  - secretName: webSvc1
    type: Opaque
    data:
    - objectName: SecVal1
      key: SecKey1
    - objectName: SecVal2
      key: SecKey2
  parameters:
    usePodIdentity: "false"
    useVMManagedIdentity: "true"
    userAssignedIdentityID: "12345678-1234-1234-1234-000000000000"
    keyvaultName: "az-kv-01"
    objects: |
      array:
        - |
          objectName: SecVal1
          objectType: secret
        - |
          objectName: SecVal2
          objectType: secret
    tenantId: "1234-12345678-00000000"
`
	wantDeploymentResult01 := `apiVersion: apps/v1
kind: Deployment
metadata:
  name: webSvc1
  labels:
    app.kubernetes.io/name: webSvc1
spec:
  selector:
    matchLabels:
      app.kubernetes.io/name: webSvc1
  template:
    metadata:
      labels:
        app.kubernetes.io/name: webSvc1
    spec:
      containers:
      - name: app
        image: private-image
        imagePullPolicy: IfNotPresent
        lifecycle:
          preStop:
            exec:
              command:
              - /bin/sleep
              - "20"
        ports:
        - containerPort: 80
        volumeMounts:
        - name: webSvc1
          mountPath: "/mnt/secrets-store"
          readOnly: true
        envFrom:
        - configMapRef:
            name: webSvc1
        - secretRef:
            name: webSvc1
        readinessProbe:
          httpGet:
            path: /healthz
            port: 80
            httpHeaders:
            - name: Custom-Header-1
              value: Awesome1
            - name: Custom-Header-2
              value: Awesome2
          initialDelaySeconds: 20
          periodSeconds: 10
          timeoutSeconds: 10
          successThreshold: 1
          failureThreshold: 3
        livenessProbe:
          httpGet:
            path: /healthz
            port: 80
            httpHeaders:
            - name: Custom-Header-3
              value: Awesome3
          initialDelaySeconds: 20
          periodSeconds: 10
          timeoutSeconds: 10
          successThreshold: 1
          failureThreshold: 3
        resources:
          requests:
            cpu: 100m
            memory: 128Mi
          limits:
            cpu: 200m
            memory: 256Mi
      volumes:
        - name: webSvc1
          csi:
            driver: secrets-store.csi.k8s.io
            readOnly: true
            volumeAttributes:
              secretProviderClass: "webSvc1"
            nodePublishSecretRef:
              name: secrets-store-creds
`
	wantDeploymentResult02 := `apiVersion: apps/v1
kind: Deployment
metadata:
  name: webSvc1
  labels:
    app.kubernetes.io/name: webSvc1
spec:
  selector:
    matchLabels:
      app.kubernetes.io/name: webSvc1
  template:
    metadata:
      labels:
        app.kubernetes.io/name: webSvc1
    spec:
      containers:
      - name: app
        image: private-image
        imagePullPolicy: IfNotPresent
        lifecycle:
          preStop:
            exec:
              command:
              - /bin/sleep
              - "20"
        ports:
        - containerPort: 80
        livenessProbe:
          exec:
            command:
            - cat
            - /tmp/test
        resources:
          requests:
            cpu: 100m
            memory: 128Mi
`
	wantDeploymentResult03 := `apiVersion: apps/v1
kind: Deployment
metadata:
  name: webSvc1
  labels:
    app.kubernetes.io/name: webSvc1
spec:
  replicas: 3
  selector:
    matchLabels:
      app.kubernetes.io/name: webSvc1
  template:
    metadata:
      labels:
        app.kubernetes.io/name: webSvc1
    spec:
      containers:
      - name: app
        image: private-image
        imagePullPolicy: IfNotPresent
        lifecycle:
          preStop:
            exec:
              command:
              - /bin/sleep
              - "20"
        ports:
        - containerPort: 80
        livenessProbe:
          tcpSocket:
            port: 80
        resources:
          requests:
            cpu: 100m
            memory: 128Mi
`
	wantServiceResult := `apiVersion: v1
kind: Service
metadata:
  labels:
    app.kubernetes.io/name: webSvc1
  name: webSvc1
spec:
  type: ClusterIP
  selector:
    app.kubernetes.io/name: webSvc1
  ports:
  - protocol: TCP
    port: 80
    targetPort: 80
`
	wantHpaResult01 := `apiVersion: autoscaling/v2beta2
kind: HorizontalPodAutoscaler
metadata:
  labels:
    app.kubernetes.io/name: webSvc1
  name: webSvc1
spec:
  scaleTargetRef:
    apiVersion: apps/v1
    kind: Deployment
    name: webSvc1
  minReplicas: 1
  maxReplicas: 10
  metrics:
  - type: Resource
    resource:
      name: memory
      target:
        type: Utilization
        averageUtilization: 90
  - type: Resource
    resource:
      name: cpu
      target:
        type: Utilization
        averageUtilization: 40
`
	wantHpaResult02 := `apiVersion: autoscaling/v2beta2
kind: HorizontalPodAutoscaler
metadata:
  labels:
    app.kubernetes.io/name: webSvc1
  name: webSvc1
spec:
  scaleTargetRef:
    apiVersion: apps/v1
    kind: Deployment
    name: webSvc1
  minReplicas: 1
  maxReplicas: 10
  metrics:
  - type: Resource
    resource:
      name: cpu
      target:
        type: Utilization
        averageUtilization: 40
`
	wantKustomizationResult01 := `apiVersion: kustomize.config.k8s.io/v1beta1
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
- configmap.yaml
- secretproviderclass.yaml
images:
- name: private-image
  newName: "artifactory-dev.nowcom.io/docker/nowcom.services.bookingwfs"
  newTag: "6776f266"
`
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
		{testName: "GenerateConfigMap", app: sampleApp01, tmpl: templates.ConfigMap, wantResult: wantConfigmapResult},
		{testName: "GenerateSecretProviderClassAzKvSpSecret", app: sampleApp01, tmpl: templates.SecretProviderClass, wantResult: wantSecretproviderclassAzKvSpSecretResult},
		{testName: "GenerateSecretProviderClassAzKvUserAssignedIdentityID", app: sampleApp04, tmpl: templates.SecretProviderClass, wantResult: wantSecretproviderclassAzKvUserAssignedIdentityIDResult},
		{testName: "GenerateDeployment01", app: sampleApp01, tmpl: templates.Deployment, wantResult: wantDeploymentResult01},
		{testName: "GenerateDeployment02", app: sampleApp02, tmpl: templates.Deployment, wantResult: wantDeploymentResult02},
		{testName: "GenerateDeployment03", app: sampleApp03, tmpl: templates.Deployment, wantResult: wantDeploymentResult03},
		{testName: "GenerateService", app: sampleApp01, tmpl: templates.Service, wantResult: wantServiceResult},
		{testName: "GenerateHPA01", app: sampleApp01, tmpl: templates.Hpa, wantResult: wantHpaResult01},
		{testName: "GenerateHPA02", app: sampleApp02, tmpl: templates.Hpa, wantResult: wantHpaResult02},
		{testName: "GenerateKustomization01", app: sampleApp01, tmpl: templates.Kustomization, wantResult: wantKustomizationResult01},
		{testName: "GenerateKustomization02", app: sampleApp02, tmpl: templates.Kustomization, wantResult: wantKustomizationResult02},
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
