FROM ubuntu:latest
MAINTAINER Chida K "ck@mphasis.com"
ENV DEPLOY_DIR /root/
ADD ./config-lint_0.0.35_Linux_x86_64.tar.gz $DEPLOY_DIR
RUN mkdir -p ${DEPLOY_DIR}/config-input
COPY ./config-input $DEPLOY_DIR/config-input
CMD ["./root/config-lint", "-terraform", "config-input",">","file1.txt"]
