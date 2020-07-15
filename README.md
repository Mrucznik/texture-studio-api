# API for Texture Studio

API for managing Texture Studio server.

## Installation

You need to install ProtocolBuffers 3.0.0-beta-3 or later.
```shell script
mkdir tmp
cd tmp
git clone https://github.com/google/protobuf
cd protobuf
./autogen.sh
./configure
make
make check
sudo make install
```

## Build
```shell script
go generate
go build
```

