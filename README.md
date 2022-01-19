![example workflow](https://github.com/lapee79/k8s-manifest-generator/actions/workflows/verify-pr.yml/badge.svg)
# k8s-manifest-generator

The `k8s-manifest-generator` CLI is developed to generate the Kustomize manifests. It helps to integrate generating the Kustomize manifests into a CD pipeline.

This is the sample application spec JSON file.

```json
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
```