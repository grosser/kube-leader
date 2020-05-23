Kubernetes leader election via `ConfigMap` as `ENTRYPOINT`.
- single binary
- no sidecars
- no endpoints
- no runtime overhead
- noop when running outside a cluster
- reuses battle-tested operator-sdk leader election

```
pod-A {"level":"info","message":"Trying to become the leader."}
pod-A {"level":"info","message":"No pre-existing lock was found."}
pod-A {"level":"info","message":"Became the leader."}
pod-B {"level":"info","message":"Trying to become the leader."}
pod-B {"level":"info","message":"Found existing lock","LockOwner":"pod-A"}
```

# How it works

## Under the hood

- each `Pod` tries to create the same `ConfigMap` in it's `metadata.namespace`
- whoever manages to create it, is the leader
- when leader is deleted, `ConfigMap` is cleaned up via `OwnerReference` by kubernetes
- when `ConfigMap` `OwnerReference` points to a non-existent `Pod`, the `ConfigMap` is deleted by the next `Pod` that starts

## When to use

- want to make sure there are never 2 copies of your `Pod` running
- booting a new `Pod` takes a long time because of the image / cluster setup etc

## When not to use

- you `Pod` gets stuck and you need someone else to take over
- booting a new `Pod` takes a long time because it's executable is slow to start (in that case you need inline leader election)

# Install

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
  verbs: ["get", "create"]
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
- use lease api [example](https://carlosbecker.com/posts/k8s-leader-election), this requires it to spawn the child and relay signals like [consul lock](https://www.consul.io/docs/commands/lock.html#usage) does, so it can shut down the process when lease is lost


# Author
[Michael Grosser](http://grosser.it)<br/>
michael@grosser.it<br/>
License: MIT<br/>
