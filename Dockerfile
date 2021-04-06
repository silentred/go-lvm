##############################################
# This Dockerfile is created for the CI test #
##############################################
FROM centos:latest
RUN yum -y install golang device-mapper-devel lvm2-devel gcc git
RUN go get -u github.com/silentred/go-lvm
RUN cd ~/go/src/github.com/silentred/go-lvm && go build ./...
#RUN go run cmd/example.go
