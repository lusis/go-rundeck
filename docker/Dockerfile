FROM debian:jessie
MAINTAINER John E Vincent
ENV DEBIAN_FRONTEND noninteractive

RUN apt-get update; apt-get install -yq openssh-client openjdk-7-jre curl
RUN curl -L http://dl.bintray.com/rundeck/rundeck-deb/rundeck-2.6.9-1-GA.deb -o /tmp/rundeck.deb; dpkg -i /tmp/rundeck.deb; rm -rf /tmp/rundeck.deb
ADD rundeckd.init /etc/init.d/rundeckd
RUN chmod +x /etc/init.d/rundeckd
CMD /etc/init.d/rundeckd foreground


EXPOSE 4440:4440
