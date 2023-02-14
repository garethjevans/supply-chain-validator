[![Go Report Card](https://goreportcard.com/badge/github.com/garethjevans/supply-chain-validator)](https://goreportcard.com/report/github.com/garethjevans/supply-chain-validator)

# supply-chain-validator

a small CLI that can validate a supply chain

## To Install

TODO

## Usage

TODO 

## Documentation

TODO

### A Supply Chain to Validate

```yaml
apiVersion: carto.run/v1alpha1
kind: ClusterSupplyChain
metadata:
  name: supply-chain-supply-chain
spec:
  selectorMatchExpressions:
  - key: apps.tanzu.vmware.com/workload-type
    operator: In
    values:
    - supply-chain
  resources:
  - name: source-provider
    templateRef:
      kind: ClusterSourceTemplate
      name: source-template
    params:
    - name: serviceAccount
      default: default
    - name: gitImplementation
      value: go-git
  - name: source-tester
    templateRef:
      kind: ClusterSourceTemplate
      name: testing-pipeline
    sources:
    - resource: source-provider
      name: source
  - name: source-scanner
    templateRef:
      kind: ClusterSourceTemplate
      name: source-scanner-template
    params:
    - name: scanning_source_policy
      default: scan-policy
    - name: scanning_source_template
      default: blob-source-scan-template
    sources:
    - resource: source-tester
      name: source
```



