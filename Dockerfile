FROM scratch
LABEL MAINTAINER=embano1@live.com
COPY runtime /
USER 10001
ENTRYPOINT ["/runtime"]
CMD [ "-c", "0" ]