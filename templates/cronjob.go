package templates

var CronJob = `apiVersion: batch/v1
kind: CronJob
metadata:
  name: {{.Name}}
  labels:
    app.kubernetes.io/name: {{.Name}}
spec:
  schedule: "{{.JobSpec.Schedule}}"
  jobTemplate:
    spec:
      template:
        spec:
          {{ template "PodSpec" . | indent 10 }}
          restartPolicy: {{.JobSpec.RestartPolicy}}
`
