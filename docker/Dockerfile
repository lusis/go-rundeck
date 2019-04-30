FROM debian:stretch
LABEL authors="John E. Vincent <lusis.org+github.com@gmail.com>"
ENV DEBIAN_FRONTEND noninteractive
ARG RDECK_VER
EXPOSE 4440:4440

RUN apt-get update
RUN apt-get install -yq openssh-client curl uuid-runtime python3 git gnupg2 apt-transport-https

RUN echo "deb https://rundeck.bintray.com/rundeck-deb /" | tee -a /etc/apt/sources.list.d/rundeck.list 
RUN curl 'https://bintray.com/user/downloadSubjectPublicKey?username=bintray' | apt-key add -
RUN apt-get update
RUN apt-get install -yq rundeck=${RDECK_VER}

ADD rundeckd.init /etc/init.d/rundeckd
ADD admin.aclpolicy /etc/rundeck/admin.aclpolicy
ADD apitoken.aclpolicy /etc/rundeck/apitoken.aclpolicy
ADD realm.properties /etc/rundeck/realm.properties
ADD token.properties /etc/rundeck/token.properties
RUN sed -ie "s/-Dserver.http.port/-Dfile.encoding=UTF-8 -Dserver.http.port/g" /etc/rundeck/profile
RUN sed -ie "s/grails.serverURL=http:\/\/localhost:4440/grails.serverURL=http:\/\/127.0.0.1:4440/g" /etc/rundeck/rundeck-config.properties
RUN echo "rundeck.tokens.file=/etc/rundeck/token.properties" >> /etc/rundeck/framework.properties
ADD gitrepo.sh /gitrepo.sh
RUN chmod +x /gitrepo.sh
RUN /gitrepo.sh
RUN chmod +x /etc/init.d/rundeckd

CMD /etc/init.d/rundeckd foreground
