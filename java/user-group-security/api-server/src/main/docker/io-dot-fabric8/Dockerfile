ROM openjdk:8-jre
LABEL maintainer="tangfeixiong <tangfx128@gmail.com>" \
      project="https://github.com/tangfeixiong/go-to-kubernetes" \
      name="api" \
      annotation='{"example.com/project":"user-group-api"}' \
      tag="centos java1.8 openjdk springboot"

ARG jarTgt
ARG javaOpt

COPY ${jarTgt:-/maven/api.jar} /api.jar

ENV JAVA_OPTIONS="${javaOpt:-'-Xms128m -Xmx512m -XX:PermSize=128m -XX:MaxPermSize=256m'}" \
    GO_TO_KUBERNETES="user-group-security-api"

ENTRYPOINT ["java","-Djava.security.egd=file:/dev/./urandom","-jar","/api.jar"]
