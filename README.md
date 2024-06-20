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

## What is a builder?

A builder is an OCI image that contains an ordered combination of buildpacks and a build-time base image. A builder uses the lifecycle to run the detect phase for all the buildpacks it contains, in sequence, and then proceeds to run the build phase for all the buildpacks that passed detection. This allows a single builder to automatically detect and build various kinds of applications.

`NB!`

The `pack build` command has a useful flag called `--publish` that builds your image directly onto a Docker registry.

Builder Suggestion:

you can run `pack builder suggest` for a list of suggested builders.

`pack builder suggest`

## What happens during rebase?

The rebase command allows app developers or operators to quickly update an app image when its runtime base image has changed. By using image layer rebasing, this command eliminates the need to fully rebuild the app.

The rebase command can check if a newer version of the app’s base image exists (either locally or in a registry). If a newer version is found, rebase updates the app image’s layer metadata to reference the newer base image version.

The `pack rebase` command has a --publish flag that can be used to publish the updated app image directly to a registry.

## What is a platform?

A platform orchestrates builds by invoking the lifecycle binary along with buildpacks and application source code to produce a runnable OCI image.

Examples of a platform might include:
* A local CLI tool that uses buildpacks to create OCI images.
* A plugin for a continuous integration service that uses buildpacks to create OCI images.

## What are base images?

#### Build image

The build image provides the base image from which the build environment is constructed. The build environment is the containerized environment in which the lifecycle (and thereby buildpacks) are executed.

#### Run image

The run image provides the base image for application images.