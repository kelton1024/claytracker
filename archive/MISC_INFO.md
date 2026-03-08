# Misc Info
This page has information that I learned while building RocksDB/Glue code.

## Info About Build Command
```bash
# fPIC - Position Independent Code tells the compiler to generate machine code that can run at any memory address
# Shared tells the compiler to make a shared library
... -fPIC -shared ...
```

## Testing Info
The following commands are some that I used for testing
```bash
# Compile and assemble, but don't link
g++ -std=c++20 -I../rocksdb/include -c glue.cpp -o glue.o

# Build static library
ar rcs libglue.a glue.o

# For testing
g++ --std=c++20 -I../rocksdb/include glue.cpp -o glue -L../rocksdb -lrocksdb 
```

## Debug Info
```bash
# Debugging
gdb ./glue core

# Check Core file size limit
ulimit -c 

# Set to unlimited
ulimit -c unlimited

# List coredumps
coredumpctl list

# Open coredump
coredumpctl gdb <PID>

# File Location
/var/lib/systemd/coredump/

# GDB Commands
bt # backtrace
info local # local variables
print myVar # inspect variable
up # Move up frame
down # Move down a frame
frame <Number> # Go to frame <Number>

# List symbols
nm # List symbols in object files or libraries
c++filt # Demangle C++ symbol names
```

## TODO: Read about the following
```bash
-lstdc++
-lpthread
```