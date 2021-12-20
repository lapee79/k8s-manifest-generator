package templates

var Deployment = `apiVersion: apps/v1
kind: Deployment
metadata:
  name: {{.Name}}
  labels:
    app.kubernetes.io/name: {{.Name}}
spec:
  selector:
    matchLabels:
      app.kubernetes.io/name: {{.Name}}
  template:
    metadata:
      labels:
        app.kubernetes.io/name: {{.Name}}
    spec:
      containers:
      - name: app
        image: private-image
        imagePullPolicy: IfNotPresent
        lifecycle:
          preStop:
            exec:
              command:
              - /bin/sleep
              - "20"
        ports:
        - containerPort: {{.ContainerPort}}
        {{- if .Secret}}
        volumeMounts:
        - name: {{.Name}}
          mountPath: "/mnt/secrets-store"
          readOnly: true
        {{- end}}
        {{- if .Config}}
        envFrom:
        - configMapRef:
            name: {{.Name}}
        {{- end}}
        {{- if .Secret}}
        - secretRef:
            name: {{.Name}}
        {{- end}}
        readinessProbe:
          failureThreshold: 3
          httpGet:
            path: {{.HealthCheckURL}}
            port: {{.ContainerPort}}
            scheme: HTTP
          initialDelaySeconds: 20
          periodSeconds: 10
          successThreshold: 1
          timeoutSeconds: 10
        livenessProbe:
          failureThreshold: 3
          httpGet:
            path: {{.HealthCheckURL}}
            port: {{.ContainerPort}}
            scheme: HTTP
          initialDelaySeconds: 20
          periodSeconds: 10
          successThreshold: 1
          timeoutSeconds: 10
        resources:
          requests:
            cpu: {{.Resources.Requests.CPU}}
            memory: {{.Resources.Requests.Memory}}
          {{- if .Resources.Limits}}
          limits:
            cpu: {{.Resources.Limits.CPU}}
            memory: {{.Resources.Limits.Memory}}
          {{- end}}
      {{- if .Secret}}
      volumes:
        - name: {{.Name}}
          csi:
            driver: secrets-store.csi.k8s.io
            readOnly: true
            volumeAttributes:
              secretProviderClass: "{{.Name}}"
            nodePublishSecretRef:
              name: {{.AzKvSpSecret}}
      {{- end}}
`
