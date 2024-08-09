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

Run docker container for Nodejs App
```
docker run --detach -p 3002:3002 --name node-server node-webserver-paketo:0.4.0
```

------------

Run builpack on Nodejs with scripts App

```
pack build node-webserver-with-script-paketo:0.3.0 --builder tlokweng-builder:0.4.0 --path examples/apps/nodejs/app-with-script --env BP_NODE_RUN_SCRIPTS="test:unit,test:integration"
```

Run docker container for Nodejs with scripts App
```
docker run --detach -p 3008:3008 --name node-server-with-script node-webserver-with-script-paketo:0.3.0
```

Check app logs

```
docker container logs <container-name>
```