Create builder image

```
docker build examples/builders/node-js-builder/builder -t tlou/bjatladi-builder:0.4.0
```

Create runner image

```
docker build examples/builders/node-js-builder/runner -t tlou/bjatladi-runner:0.4.0
```

Create builder:

```
pack builder create tlokweng-builder:0.4.0 --config examples/builders/node-js-builder/builder.toml
```

Run builpack on Nodejs App

```
pack build node-webserver-paketo:0.4.0 --builder tlokweng-builder:0.4.0 --path examples/apps/nodejs/app
```

docker run --detach -p 3002:3002 --name node-server node-webserver-paketo:0.4.0

Check app logs

```
docker container logs <container-name>
```