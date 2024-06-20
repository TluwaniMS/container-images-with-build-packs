# container-images-with-build-packs

## What is a buildpack?

A buildpack is a piece of software that converts application source code into executable artifacts. It does this by analyzing the source code, identifying its dependencies, and determining the most appropriate way to build and package the application for deployment.

## Why buildpacks?

Buildpacks enable application developers to concentrate on writing code without having to worry about image security, optimizing container images, or devising container build strategies.

#### Each buildpack has two jobs to do

##### Detect 

The buildpack determines if it is needed or not.

`e.g`
* A Node buildpack may look for a package-lock.json file.

##### Build 

A buildpack transforms application source code by analyzing it and determining the necessary steps to convert it into an executable format.

* Setting build-time and run-time environment variables.
* Downloading dependencies.
* Running source code compilation (if needed).
* Configuring the application entrypoint and any startup scripts.

## What happens during the build process?

The build process executes one or more buildpacks against an application’s source code to produce a runnable OCI image. During this process, the build-time base image serves as the environment in which buildpacks are executed, while the runtime base image serves as the foundation for the final application image. Buildpacks can be bundled with a specific build-time base image, resulting in a builder image. Builders offer a convenient way to distribute buildpacks.
