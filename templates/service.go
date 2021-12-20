package templates

var Service = `apiVersion: v1
kind: Service
metadata:
  labels:
    app.kubernetes.io/name: {{.Name}}
  name: {{.Name}}
spec:
  type: ClusterIP
  selector:
    app.kubernetes.io/name: {{.Name}}
  ports:
  - protocol: TCP
    port: {{.ServicePort}}
    targetPort: {{.ContainerPort}}
`
