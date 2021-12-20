package templates

var Hpa = `apiVersion: autoscaling/v2beta2
kind: HorizontalPodAutoscaler
metadata:
  labels:
    app.kubernetes.io/name: {{.Name}}
  name: {{.Name}}
spec:
  scaleTargetRef:
    apiVersion: apps/v1
    kind: Deployment
    name: {{.Name}}
  minReplicas: {{.AutoScale.MinPodNum}}
  maxReplicas: {{.AutoScale.MaxPodNum}}
  metrics:
  - type: Resource
    resource:
      name: cpu
      target:
        type: Utilization
        averageUtilization: {{.AutoScale.CpuUsage}}
  {{- if .AutoScale.MemUsage}}
  - type: Resource
    resource:
      name: memory
      target:
        type: Utilization
        averageUtilization: {{.AutoScale.MemUsage}}
  {{- end}}
`