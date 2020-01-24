# Build Stage
FROM lacion/alpine-golang-buildimage:1.13 AS build-stage

LABEL app="build-spread"
LABEL REPO="https://github.com/arnour/spread"

ENV PROJPATH=/go/src/github.com/arnour/spread

# Because of https://github.com/docker/docker/issues/14914
ENV PATH=$PATH:$GOROOT/bin:$GOPATH/bin

ADD . /go/src/github.com/arnour/spread
WORKDIR /go/src/github.com/arnour/spread

RUN make build-alpine

# Final Stage
FROM lacion/alpine-base-image:latest

ARG GIT_COMMIT
ARG VERSION
LABEL REPO="https://github.com/arnour/spread"
LABEL GIT_COMMIT=$GIT_COMMIT
LABEL VERSION=$VERSION

# Because of https://github.com/docker/docker/issues/14914
ENV PATH=$PATH:/opt/spread/bin

WORKDIR /opt/spread/bin

COPY --from=build-stage /go/src/github.com/arnour/spread/bin/spread /opt/spread/bin/
RUN chmod +x /opt/spread/bin/spread

# Create appuser
RUN adduser -D -g '' spread
USER spread

ENTRYPOINT ["/usr/bin/dumb-init", "--"]

CMD ["/opt/spread/bin/spread"]
