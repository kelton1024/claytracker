FROM golang:1.25-alpine

#RUN echo "@testing http://nl.alpinelinux.org/alpine/edge/testing" >>/etc/apk/repositories
#RUN apk add --update --no-cache build-base linux-headers git cmake bash perl #wget mercurial g++ autoconf libgflags-dev cmake bash
#RUN apk add --update --no-cache zlib zlib-dev bzip2 bzip2-dev snappy snappy-dev lz4 lz4-dev zstd #@testing zstd-dev@testing libtbb-dev@testing libtbb@testing

# installing latest gflags
#RUN cd /tmp && \
#    git clone https://github.com/gflags/gflags.git && \
#    cd gflags && \
#    mkdir build && \
#    cd build && \
#    cmake -DBUILD_SHARED_LIBS=1 -DGFLAGS_INSTALL_SHARED_LIBS=1 .. && \
#    make install && \
#    cd /tmp && \
#    rm -R /tmp/gflags/

RUN apk add --update --no-cache build-base linux-headers git cmake bash perl wget g++ autoconf cmake #libgflags-dev bash mercurial 
# Install Rocksdb
RUN cd /tmp && \
    git clone https://github.com/facebook/rocksdb.git && \
    cd rocksdb && \
    git checkout v10.7.5 && \
    export EXTRA_CXXFLAGS="-Wno-error=restrict -Wno-error=unused-parameter -fPIC" && \
    export EXTRA_CFLAGS="-fPIC" && \
    make static_lib

COPY frontend /frontend
COPY glue /glue
COPY go /go
