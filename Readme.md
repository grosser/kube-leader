Kubernetes leader election via `ConfigMap` as `ENTRYPOINT`.
- single binary
- no sidecars
- no endpoints
- no runtime overhead
- noop when running outside a cluster
- reuses battle-tested operator-sdk leader election

Install [latest binary](https://github.com/grosser/kube-leader/releases) in `Dockerfile`:
```
RUN curl -sfL <PICK URL FOR LATEST BINARY> | tar -zx && chmod +x kube-leader
ENTRYPOINT ["./kube-leader", "my-lock"]
CMD ["my", "stuff"]
```

Add to `Deployment` container:
```yaml
env:
- name: POD_NAME
  valueFrom:
    fieldRef:
      fieldPath: metadata.name
```

Add permissions to `Role`:
```
- apiGroups: [""]
  resources: ["configmaps"]
  verbs: ["create"]
- apiGroups: [""]
  resources: ["pods"]
  verbs: ["get", "delete"]
```

# Development

## Test

- install `stern` + `ruby`
- `rake server`
- delete pod
- check that it elects a new leader

## Release

Create a new release via github UI, workflow will automatically build a new binary.


# TODO

- support flags like `--help` or log/interval options
- reduce binary size by not relying on operator-sdk directly

# Author
[Michael Grosser](http://grosser.it)<br/>
michael@grosser.it<br/>
License: MIT<br/>
