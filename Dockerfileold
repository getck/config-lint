FROM ubuntu:latest
MAINTAINER Chida K "ck@mphasis.com"
ENV DEPLOY_DIR /root/
ADD ./config-lint_0.0.35_Linux_x86_64.tar.gz $DEPLOY_DIR
RUN mkdir -p ${DEPLOY_DIR}/example-files
COPY ./example-files $DEPLOY_DIR/example-files
CMD ["./root/config-lint", "-terraform", "example-files/config",">","file1.txt"]
