FROM golang

# Golang Setup
ENV PATH $PATH:$GOPATH/bin:$GOROOT/bin
COPY . /echo-server
WORKDIR /echo-server
RUN sh ./getfile
RUN apt-get update && apt-get install -y \
      apt-transport-https \
      ca-certificates \
      curl \
      gnupg2 \
      software-properties-common

# DockerSetup
RUN curl -fsSL https://download.docker.com/linux/debian/gpg | apt-key add -
RUN apt-key fingerprint 0EBFCD88
RUN add-apt-repository "deb [arch=amd64] https://download.docker.com/linux/debian $(lsb_release -cs) stable"
RUN apt-get update && apt-get install -y docker-ce

EXPOSE 1323
CMD ["gin"]
