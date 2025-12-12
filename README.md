# yaml2go

## Quickstart

> [💾 Download Here](https://github.com/lawrenson-siimpl/yaml2go/releases)

```console
~$ yaml2go test-values.yaml
values := map[string]interface{}{
  "image": map[string]interface{}{
    "repository": "dhi/cilium",
    "tag": "1.18.4",
    "useDigest": false,
  },
  "imagePullSecrets": []interface{}{
    map[string]interface{}{
      "name": "regcred",
    },
  },
  "operator": map[string]interface{}{
    "replicas": 1,
  },
  "hubble": map[string]interface{}{
    "enabled": false,
  },
  "extraHostPathMounts": []interface{}{
    map[string]interface{}{
      "readOnly": true,
      "mountPropagation": "HostToContainer",
      "name": "cilium-cgroup-override",
      "mountPath": "/run/cilium/cgroupv2",
      "hostPath": "/sys/fs/cgroup",
      "hostPathType": "Directory",
    },
  },
  "kubeProxyReplacement": false,
  "ipam": map[string]interface{}{
    "mode": "kubernetes",
  },
}
```
