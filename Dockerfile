# Please use this Dockerfile only if
# you want to build an image from source without
# pnpm and Go installed on your dev machine.

# You can build oCIS using this Dockerfile
# by running following command:
# `docker build -t owncloud/ocis:custom .`

# In most other cases you might want to run the
# following command instead:
# `make -C ocis dev-docker`
# It will build a `owncloud/ocis:dev` image for you
# and use your local pnpm and Go caches and therefore
# is a lot faster than the build steps below.


FROM owncloudci/nodejs:16 as generate

COPY ./ /ocis/

WORKDIR /ocis/ocis
RUN make ci-node-generate

FROM owncloudci/golang:1.22 as build

COPY --from=generate /ocis /ocis

WORKDIR /ocis/ocis
RUN make ci-go-generate build

FROM alpine:3.17

RUN apk add --no-cache ca-certificates mailcap tree attr curl && \
	echo 'hosts: files dns' >| /etc/nsswitch.conf

LABEL maintainer="ownCloud GmbH <devops@owncloud.com>" \
	org.label-schema.name="ownCloud Infinite Scale" \
	org.label-schema.vendor="ownCloud GmbH" \
	org.label-schema.schema-version="1.0"

ENTRYPOINT ["/usr/bin/ocis"]
CMD ["server"]

COPY --from=build /ocis/ocis/bin/ocis /usr/bin/ocis
