package templates

var SecretProviderClass = `apiVersion: secrets-store.csi.x-k8s.io/v1
kind: SecretProviderClass
metadata:
  name: {{.Name}}
spec:
  provider: azure
  secretObjects:
  - secretName: {{.Name}}
    type: Opaque
    data:
    {{- range $value := .Secret}}
    - objectName: {{$value.Value}}
      key: {{$value.Key}}
    {{- end}}
  parameters:
    usePodIdentity: "false"
    {{- if .AzKvUserAssignedIdentityID}}
    useVMManagedIdentity: "true"
    userAssignedIdentityID: "{{.AzKvUserAssignedIdentityID}}"
    {{- end}}
    keyvaultName: "{{.AzKV}}"
    objects: |
      array:
        {{- range $value := .Secret}}
        - |
          objectName: {{$value.Value}}
          objectType: secret
        {{- end}}
    tenantId: "{{.AzTid}}"
`
