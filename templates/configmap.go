package templates

var ConfigMap = `apiVersion: v1
kind: ConfigMap
metadata:
  name: {{.Name}}
data:
  {{- range $value := .Config}}
  {{$value.Key}}: {{$value.Value}}
  {{- end}}
`
