FROM debian:jessie
LABEL authors="John E. Vincent <lusis.org+github.com@gmail.com>"
ENV DEBIAN_FRONTEND noninteractive

RUN echo "deb http://ftp.de.debian.org/debian jessie-backports main" > /etc/apt/sources.list.d/backports.list
RUN apt-get update
RUN apt-get install -t jessie-backports -yq openjdk-8-jdk
RUN apt-get install -yq openssh-client curl uuid-runtime python3 git
RUN groupadd rundeck
RUN useradd -g rundeck rundeck
RUN curl -L http://dl.bintray.com/rundeck/rundeck-deb/rundeck_2.10.6-1-GA_all.deb -o /tmp/rundeck.deb; dpkg -i /tmp/rundeck.deb; rm -rf /tmp/rundeck.deb
ADD rundeckd.init /etc/init.d/rundeckd
ADD admin.aclpolicy /etc/rundeck/admin.aclpolicy
ADD apitoken.aclpolicy /etc/rundeck/apitoken.aclpolicy
ADD realm.properties /etc/rundeck/realm.properties
ADD token.properties /etc/rundeck/token.properties
RUN sed -ie "s/-Dserver.http.port/-Dfile.encoding=UTF-8 -Dserver.http.port/g" /etc/rundeck/profile
RUN echo "rundeck.tokens.file=/etc/rundeck/token.properties" >> /etc/rundeck/framework.properties
#ADD rundeckdb.mv.db /var/lib/rundeck/data/rundeckdb.mv.db
#RUN chown rundeck:rundeck /var/lib/rundeck/data/rundeckdb.mv.db
ADD gitrepo.sh /gitrepo.sh
RUN chmod +x /gitrepo.sh
RUN /gitrepo.sh
RUN chmod +x /etc/init.d/rundeckd

CMD /etc/init.d/rundeckd foreground


EXPOSE 4440:4440
