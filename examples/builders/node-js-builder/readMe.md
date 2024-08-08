Create builder image

```
docker build examples/builders/node-js-builder/builder -t tlou/bjatladi-builder:0.1.0
```

Create runner image

```
docker build examples/builders/node-js-builder/runner -t tlou/bjatladi-runner:0.1.0
```