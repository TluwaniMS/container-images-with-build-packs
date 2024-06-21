# Pack CLI

Pack is a CLI tool maintained by the CNB project to support the use of buildpacks.

It enables the following functionality:

- Building an application using buildpacks.
- Rebasing application images created using buildpacks.
- Creating various components used within the ecosystem.

Pack works as both a Command Line Interface (CLI) and a Go library.

[Install Pack CLI](https://buildpacks.io/docs/for-platform-operators/how-to/integrate-ci/pack/)

## Nodejs

build basic nodejs webserver image:

```
pack build <image-name>:<tag> --buildpack paketo-buildpacks/nodejs \
  --builder paketobuildpacks/builder-jammy-base \
  --path examples/apps/nodejs/app
```

```
pack build <account>/<image-name>:<tag> --buildpack paketo-buildpacks/nodejs \
  --builder paketobuildpacks/builder-jammy-base \
  --path examples/apps/nodejs/app \
  --publish
```

## Golang

```
pack build <image-name>:<tag> --buildpack paketo-buildpacks/go \
  --builder paketobuildpacks/builder-jammy-base \
  --path examples/apps/golang/app
```

```
pack build <account>/<image-name>:<tag>  --buildpack paketo-buildpacks/go \
  --builder paketobuildpacks/builder-jammy-base \
  --path examples/apps/golang/app \
  --publish
```