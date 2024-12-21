FROM ajaysc30/go-health-tracker:go-base-image-1.22 AS build

RUN mkdir -p /Go
WORKDIR /Go

# RUN /bin/sh -c set -eux; apk add --no-cache golangci-lint=1.57.2-r3 gosec=2.20.0-r3

COPY common /Go/common
COPY src /Go/src

ENV REPO_ROOT=/Go
# RUN make -C src clean

ENV GOCACHE=/root/.cache/go-build
RUN --mount=type=cache,target="/root/.cache/go-build" cd src/go-health-app && time go install -ldflags="-s -w" -gcflags="all=-N -l" -v -tags musl
RUN strip /root/go/bin/go-health-app

#
#------------ BUILD RUNTIME IMAGE ------------
#

FROM ajaysc30/go-health-tracker:alpine-3.14 AS runtime

ARG GID=30001
ARG UID=30003
ARG PORT=8332

STOPSIGNAL SIGKILL

# Create a 'user' user in the 'usergroup' group.
RUN addgroup --gid $GID --system usergroup \
  && adduser --uid $UID --system --ingroup usergroup user

# Create logging directories.
RUN mkdir -p /var/log/gohealth \
	&& chown -R $GID:$UID /var/log/gohealth \
    && chmod -R 755 /var/log/gohealth

ENV HOME=/home/user
USER $GID:$UID
COPY --from=build /root/go/bin/go-health-app $HOME/go-health-app

ENV CONFIG_FILE_PATH=infra/configs

EXPOSE $PORT/tcp

CMD ["bin/ash","-c","cd /home/user && /home/user/go-health-app"]
