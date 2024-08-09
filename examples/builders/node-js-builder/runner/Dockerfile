# Specify the platform to ensure the image is built for linux/amd64
# Use --platform=linux/amd64 in the FROM instruction
FROM --platform=linux/amd64 ubuntu:jammy

# Install packages that we want to make available at build time
RUN apt-get update && \
  apt-get install -y xz-utils ca-certificates libc6 && \
  rm -rf /var/lib/apt/lists/*

# Set required CNB user information
ARG cnb_uid=1000
ARG cnb_gid=1000
ENV CNB_USER_ID=${cnb_uid}
ENV CNB_GROUP_ID=${cnb_gid}

# Create user and group
RUN groupadd cnb --gid ${CNB_GROUP_ID} && \
  useradd --uid ${CNB_USER_ID} --gid ${CNB_GROUP_ID} -m -s /bin/bash cnb

# Set user and group
USER ${CNB_USER_ID}:${CNB_GROUP_ID}

# Set required CNB target information
LABEL io.buildpacks.base.distro.name="vhatwanamba-runner"
LABEL io.buildpacks.base.distro.version="0.1.0"