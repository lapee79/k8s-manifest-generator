package templates

var Kustomization = `apiVersion: kustomize.config.k8s.io/v1beta1
kind: Kustomization
commonLabels:
{{- range $value := .CommonLabels}}
  {{$value.Key}}: "{{$value.Value}}"
{{- end}}
commonAnnotations:
{{- range $value := .CommonAnnotations}}
  {{$value.Key}}: "{{$value.Value}}"
{{- end}}
resources:
{{- if eq .Kind "Deployment"}}
- service.yaml
- deployment.yaml
{{- end}}
{{- if eq .Kind "CronJob"}}
- cronjob.yaml
{{- end}}
{{- if .AutoScale}}
- hpa.yaml
{{- end}}
{{- if .Config}}
- configmap.yaml
{{- end}}
{{- if .Secret}}
- secretproviderclass.yaml
{{- end}}
images:
- name: private-image
  newName: "{{.ImageUrl}}"
  newTag: "{{.ImageTag}}"
`
