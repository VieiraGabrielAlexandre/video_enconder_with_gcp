FROM golang:1.21
ENV PATH="$PATH:/bin/bash" \
    BENTO4_BIN="/opt/bento4/bin" \
    PATH="$PATH:/opt/bento4/bin"

# FFMPEG
RUN apt-get update && apt-get install -y ffmpeg bash curl make

# Install Bento
WORKDIR /tmp/bento4
ENV BENTO4_BASE_URL="http://zebulon.bok.net/Bento4/source/" \
    BENTO4_VERSION="1-5-0-615" \
    BENTO4_CHECKSUM="5378dbb374343bc274981d6e2ef93bce0851bda1" \
    BENTO4_TARGET="" \
    BENTO4_PATH="/opt/bento4" \
    BENTO4_TYPE="SRC"
    # download and unzip bento4
RUN apt-get update && apt-get install -y curl python3 unzip bash gcc g++ scons zip && \
    curl -O -s ${BENTO4_BASE_URL}/Bento4-${BENTO4_TYPE}-${BENTO4_VERSION}${BENTO4_TARGET}.zip && \
    sha1sum -b Bento4-SRC-1-5-0-615.zip | grep -o "^5378dbb374343bc274981d6e2ef93bce0851bda1" && \
#
     mkdir -p ${BENTO4_PATH} && \
     unzip Bento4-${BENTO4_TYPE}-${BENTO4_VERSION}${BENTO4_TARGET}.zip -d ${BENTO4_PATH} && \
     rm -rf Bento4-${BENTO4_TYPE}-${BENTO4_VERSION}${BENTO4_TARGET}.zip && \
     apt-get remove -y unzip && \
     apt-get clean && \
     # Don't do these steps if using binary install
     cd ${BENTO4_PATH} && scons -u build_config=Release target=x86_64-unknown-linux && \
     cp -R ${BENTO4_PATH}/Build/Targets/x86_64-unknown-linux/Release ${BENTO4_PATH}/bin && \
     cp -R ${BENTO4_PATH}/Source/Python/utils ${BENTO4_PATH}/utils && \
     cp -a ${BENTO4_PATH}/Source/Python/wrappers/. ${BENTO4_PATH}/bin

WORKDIR /go/src

ENTRYPOINT ["tail", "-f", "/dev/null"]