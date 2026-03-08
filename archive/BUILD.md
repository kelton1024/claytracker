# Build
We have to build RocksDB and the REST API before we are able to use them. 

## RocksDB
1. Pull RocksDB
```bash
git clone https://github.com/facebook/rocksdb.git
```

2. Compile static library 
```bash
make static_lib # TODO: Add extra flags argument
``` 
**Note:** A future enhancement could be linking RocksDB with a compression library

### GCC Issues

**Issue #1:** GCC was throwing an error when it encountered the following code because it thinks that a read/write is happening to overlapping memory despite the literal. 
```bash
str.insert(i, "-")
```

I disabled the warning being promoted to an error using the following:
```bash
CXXFLAGS += -Wno-error=restrict
```
**Issue #2:** The compiler also returned an unused parameter error, which I silenced using the following. This error is related to note #1 where I mentioned using snappy compression.
```bash
CXXFLAGS += -Wno-error=unused-parameter
```

**Issue #3:**
I was having a hard time building the Go binary without using PIC. Once I enabled it, Go didn't have any problems building the binary. 
```bash
CFLAGS += "-fPIC"
CXXFLAGS += "-fPIC"
```

## Glue Code
I was only able to get CGO to work with the glue code after building using the following command. I think it is something to do with PIC and not the shared library. I'm going to try to compile without the shared library later.
```bash
g++ -std=c++20 -fPIC -shared -I../rocksdb/include glue.cpp -o libglue.so -L../rocksdb/ -lrocksdb
```

## Go
Building the Go code is pretty easy, but it took some time to get it to compile with the C code. 
```bash
go build .
```

I had to set the following prior to running the code 
```bash
LD_LIBRARY_PATH=<Path to .so directory>
```

## Frontend
Install dependencies
```bash
npm install
```

Run development server
```bash
npm run dev
```
