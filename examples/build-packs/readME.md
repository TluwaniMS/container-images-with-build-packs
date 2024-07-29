What is a builder?
A builder is an OCI image containing an ordered combination of buildpacks and a build-time base image, a lifecycle binary, and a reference to a runtime base image.

Build image
The build image provides the base image from which the build environment is constructed.
The build environment is the containerized environment in which the lifecycle (and thereby buildpacks) are executed.

Run image
The run image provides the base image for application images.

What is the lifecycle?
The lifecycle orchestrates buildpacks, then assembles the resulting artifacts into an OCI image. It does its work in a series of distinct “phases”.

Analyze
The analyzer restores files that buildpacks may use to optimize the build and export phases.
responsible for analyzing the metadata from the cache and the previously built image (if available) to determine what layers can or cannot be reused.
This information is used during the export phase in order to avoid re-uploading unchanged layers.

Detect
The detector finds an ordered group of buildpacks to use during the build phase.

Restore
The restorer copies layers from the cache into the build container.

Build
The builder transforms application source code into runnable artifacts that can be packaged into a container.

Export
The exporter creates the final OCI image.




//////////

```
pack buildpack new ben/node-js \
    --api 0.10 \
    --path ben-node-js-buildpack \
    --version 0.1.0
```

```
pack config default-builder cnbs/sample-builder:jammy
```

```
pack config trusted-builders add cnbs/sample-builder:jammy
```

```
pack build ben-pack-nodejs --path ../apps/nodejs/app --buildpack ./ben-node-js-buildpack --no-color
```