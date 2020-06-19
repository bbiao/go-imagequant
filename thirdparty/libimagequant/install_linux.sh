#!/usr/bin/env sh

set -ex

INSTALL_PREFIX=/usr
LIB_VERSION=2.12.6
DOWNLOAD_URL=https://github.com/ImageOptim/libimagequant/archive/${LIB_VERSION}.tar.gz
ARCHIVE_NAME=libimagequant-${LIB_VERSION}.tar.gz

InstallLibImageQuant() {
    curl -Ss -L ${DOWNLOAD_URL} --output ${ARCHIVE_NAME}
    tar xzvf ${ARCHIVE_NAME} -C .
    cd libimagequant-${LIB_VERSION} && ./configure --prefix=/usr1 && make && sudo make install
}

InstallLibImageQuant