FROM debian:jessie
MAINTAINER John E Vincent
ENV DEBIAN_FRONTEND noninteractive

RUN echo "deb http://ftp.de.debian.org/debian jessie-backports main" > /etc/apt/sources.list.d/backports.list
RUN apt-get update
RUN apt-get install -t jessie-backports -yq openjdk-8-jdk
RUN apt-get install -yq openssh-client curl uuid-runtime
RUN groupadd rundeck
RUN useradd -g rundeck rundeck
RUN curl -L http://dl.bintray.com/rundeck/rundeck-deb/rundeck-2.10.1-1-GA.deb -o /tmp/rundeck.deb; dpkg -i /tmp/rundeck.deb; rm -rf /tmp/rundeck.deb
ADD rundeckd.init /etc/init.d/rundeckd
ADD admin.aclpolicy /etc/rundeck/admin.aclpolicy
ADD apitoken.aclpolicy /etc/rundeck/apitoken.aclpolicy
RUN chmod +x /etc/init.d/rundeckd

CMD /etc/init.d/rundeckd foreground


EXPOSE 4440:4440
