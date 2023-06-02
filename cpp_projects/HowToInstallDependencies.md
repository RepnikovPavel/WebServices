# platform support  
linux
wsl

# first of all
1. 
```console  
sudo apt update
sudo apt-get install -y pkg-config  
sudo apt install glibc-source -y
sudo apt install ninja-build
sudo apt-get install gdb
```  
2. install gcc  
```console  
tar xzf gcc-13.1.0.tar.gz
cd gcc-13.1.0
./contrib/download_prerequisites
cd ..
mkdir objdir
cd objdir
$PWD/../gcc-13.1.0/configure --prefix=$HOME/gcc-13.1.0 --enable-languages=c,c++
make -j 8
make install
```


# NOTE
check system path with
```console  
echo $PATH
echo $HOME
```
use includes of vcpkg in ide
1. for VScode  
```console  
${HOME}/vcpkg/packages/libevent_x64-linux/include
```  



# install OpenSSL
```console  
sudo apt install openssl
sudo apt-get install libssl-dev --fix-missing
```
# install cmake 
```console  
wget https://github.com/Kitware/CMake/releases/download/v3.26.4/cmake-3.26.4.tar.gz
tar -xf cmake-3.26.4.tar.gz
cd cmake-3.26.4
./bootstrap 
make
sudo make install
```  
**if cmake bin path not in PATH**:
```console
export PATH="$HOME/path/to/cmake/bin:$PATH"
for example:
export PATH="$HOME/cmake-3.26.4/bin:$PATH"
```


# install vcpkg  
```console
git clone https://github.com/Microsoft/vcpkg.git  
./vcpkg/bootstrap-vcpkg.sh
```  
**if vcpkg folder not in PATH**:
```console
export PATH="$HOME/path/to/vcpkgfolder:$PATH"
for example:
export PATH="$HOME/vcpkg:$PATH"
```

# dependencies installation  
```console  
vcpkg install libevent
vcpkg integrate install  
```  
# build project  
from folder with CMakeLists.txt  
```console  
mkdir build  
cd build  
cmake .. -DCMAKE_TOOLCHAIN_FILE=/home/user/vcpkg/scripts/buildsystems/vcpkg.cmake
cmake --build ./ --config Release  
```
NOTE:  
if y use vscode  
y can use cmake-tools extension and settings.json file 
