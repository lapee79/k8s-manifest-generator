package templates

var Kustomization = `apiVersion: kustomize.config.k8s.io/v1beta1
kind: Kustomization
commonLabels:
{{- range $value := .CommonLabels}}
  {{$value.Key}}: {{$value.Value}}
{{- end}}
commonAnnotations:
{{- range $value := .CommonAnnotations}}
  {{$value.Key}}: "{{$value.Value}}"
{{- end}}
resources:
- service.yaml
- deployment.yaml
- hpa.yaml
{{- if .Config}}
- configmap.yaml
{{- end}}
{{- if .Secret}}
- secretproviderclass.yaml
{{- end}}
images:
- name: private-image
  newName: {{.ImageUrl}}
  newTag: {{.ImageTag}}
`
