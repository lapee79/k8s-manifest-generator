# k8s-manifest-generator

The `k8s-manifest-generator` CLI is developed to generate the Kustomize manifests. It helps to integrate generating the Kustomize manifests into a CD pipeline.

This is the sample application spec JSON file.

```json
{
  "Name": "webSvc1",
  "NameSpace": "test",
  "ContainerPort": 80,
  "Config": [
    {"Key": "ConfKey1", "Value": "ConfVal1"},
    {"Key": "ConfKey2", "Value": "ConfVal2"}
  ],
  "Secret": [
    {"Key": "SecKey1", "Value": "SecVal1"},
    {"Key": "SecKey2", "Value": "SecVal2"}
  ],
  "HealthCheckURL": "/ready",
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
  "AzKV":  "az-kv-01",
  "AzTid": "1234-12345678-00000000"
}
```