FROM golang:1.15.6-alpine AS build

ARG release
COPY . /root

RUN /root/test.bash $release
