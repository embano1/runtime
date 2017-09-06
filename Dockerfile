FROM scratch
MAINTAINER embano1@live.com
COPY runtime /
ENTRYPOINT ["/runtime"]