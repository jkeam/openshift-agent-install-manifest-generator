# OpenShift Agent Install Manifest Generator

Generates all  of the required manifests for an OpenShift agent installation.

## OpenShift Setup

```shell
oc new-project oaimg
oc apply -k ./pipeline
```

### Performing Builds and Deployment

```shell
oc create -f ./pipeline/pipeline-run.yaml
```


## References

1. [OCP Local Agent Install Guide](https://github.com/mathianasj/ocp-local-agent-install-guide)
2. [Go with Tests](https://quii.gitbook.io/learn-go-with-tests)
