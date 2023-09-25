![example workflow](https://github.com/lapee79/k8s-manifest-generator/actions/workflows/verify-pr.yml/badge.svg)
# k8s-manifest-generator

The `k8s-manifest-generator` CLI is developed to generate the Kustomize manifests. It helps to integrate generating the Kustomize manifests into a CD pipeline.

This is the sample application spec JSON file.

```json
{
  "Name": "aspcoretest",  // service name on Kubernetes(allow lower cases and hyphen(-) and numbers only)
  "NameSpace": "nowcom",  // namespace on Kubernetes(allow lower cases and hyphen(-) and numbers only)
  "Kind": "Deployment",
  "ContainerPort": 80,  // application port
  "ServicePort": 80,  // the port to expose
  "CommonLabels": [
    {"Key": "app.kubernetes.io/instance", "Value": "aspcoretest"},
    {"Key": "app.kubernetes.io/environment", "Value": "dev2"}
  ],
  "CommonAnnotations": [
    {"Key": "commitAuther", "Value": "AUTHER"},  // Do not change this line.
    {"Key": "buildId", "Value": "BUILDID"},  // Do not change this line.
    {"Key": "commitAutherEmail", "Value": "EMAIL"}  // Do not change this line.
  ],
  "ImageUrl": "IMAGEURL",  // Do not change this line.
  "ImageTag": "IMAGETAG",  // Do not change this line.
  "Config": [    // The value of "Key" is used as the name of an environment variable. The value of "Value" should have a actual value to use.
    {"Key": "ASPNETCORE_ENVIRONMENT", "Value": "dev2"},
    {"Key": "DOTNET_ENVIRONMENT", "Value": "dev2"}
  ],
  "Secret": [    // The value of "Key" is used as the name of an environment variable. The value of "Value" has to use the name of a secret on the Azure Key Vault.(It can be omitted.)
    {"Key": "Book_BlackBookAccounts__0__Password", "Value": "Book-BlackBookAccounts--0--Password"},
    {"Key": "Book_KBBAccounts__0__ApiKey", "Value": "Book-KBBAccounts--0--ApiKey"}
  ],
  "ReadinessProbe": {
    "HttpGet": {
      "Path": "/",    // Update for Health check URL
      "Port": 80
    }
  },
  "LivenessProbe": {
    "HttpGet": {
      "Path": "/",    // Update for Health check URL
      "Port": 80
    }
  },
  "Resources": {
    "Requests": {
      "CPU": "100m",          // "1000m" = "1" CPU
      "Memory": "128Mi"       // "Mi" = Megabyte, "Gi" = Gigabyte
    },
    "Limits": {    // Set it not to exceed the limits.(It can be omitted.)
      "CPU": "200m",
      "Memory": "256Mi"
    }
  },
  "AutoScale": {    // You can configure the settings for the autoscaler.(It can be omitted.)
    "MinPodNum": 1,
    "MaxPodNum": 10,
    "CpuUsage": 40,
    "MemUsage": 90      // (It can be omitted.)
  },
  "AzKV":  "az-kv-01",
  "AzTid": "1234-12345678-00000000",
  "AzKvSpSecret": "secrets-store-creds"
}
```