# The line below states we will base our new image on the Latest Official Ubuntu
FROM ubuntu:20.04

ENV DEBIAN_FRONTEND=noninteractive

# Update the image to the latest packages
RUN apt-get update && apt-get upgrade -y
RUN apt install -y unzip
RUN apt install -y wget
RUN apt install -y xz-utils
RUN apt install -y git
RUN apt install -y nano
RUN apt install -y make
RUN apt install -y patchelf
RUN apt-get install -y gcc-aarch64-linux-gnu
RUN apt-get install -y build-essential


WORKDIR /home

# Download compiler repository
RUN git clone --single-branch --branch main git@github.com:Freesoul1982/go-cross-compile-library.git

# Download ndk
RUN wget https://dl.google.com/android/repository/android-ndk-r23c-linux.zip
RUN unzip android-ndk-r23c-linux.zip && rm android-ndk-r23c-linux.zip
RUN mv android-ndk-r23c/ ndk

# Download go
RUN wget https://go.dev/dl/go1.19.3.linux-amd64.tar.gz
RUN tar -C /usr/local -xzf go1.19.3.linux-amd64.tar.gz

# Envirement variables
ENV PATH=$PATH:/home/ndk
ENV GOPATH=$HOME/go
ENV PATH=$PATH:/usr/local/go/bin:$GOPATH/bin

RUN exec $SHELL