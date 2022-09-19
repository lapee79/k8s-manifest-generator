package templates

var Deployment = `apiVersion: apps/v1
kind: Deployment
metadata:
  name: {{.Name}}
  labels:
    app.kubernetes.io/name: {{.Name}}
spec:
  {{- if .Replicas}}
  replicas: {{.Replicas}}
  {{- end}}
  selector:
    matchLabels:
      app.kubernetes.io/name: {{.Name}}
  template:
    metadata:
      labels:
        app.kubernetes.io/name: {{.Name}}
    spec:
      {{ template "PodSpec" . | indent 6 }}
`
