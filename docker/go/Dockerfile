FROM debian:10

# build tdlib
RUN apt-get update \
    && apt-get upgrade \
    && apt-get install make git zlib1g-dev libssl-dev gperf php cmake clang libc++-dev libc++abi-dev -y \
    && git clone https://github.com/tdlib/td.git \
    && cd td \
    && git checkout v1.6.0 \
    && rm -rf build \
    && mkdir build \
    && cd build \
    && export CXXFLAGS="-stdlib=libc++" \
    && CC=/usr/bin/clang CXX=/usr/bin/clang++ cmake -DCMAKE_BUILD_TYPE=Release -DCMAKE_INSTALL_PREFIX:PATH=/usr/local .. \
    && cmake --build . --target install \
    && cd .. \
    && cd .. \
    && ls -l /usr/local
RUN ldconfig

# install golang
RUN apt-get install curl
RUN curl -O https://dl.google.com/go/go1.14.7.linux-amd64.tar.gz
RUN tar -C /usr/local -xzf go1.14.7.linux-amd64.tar.gz
RUN mkdir /go
ENV GOPATH="/go"
ENV PATH="${PATH}:/usr/local/go/bin:${GOPATH}/bin"

WORKDIR /app
