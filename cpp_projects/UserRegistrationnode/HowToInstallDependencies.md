# platform support  
linux
wsl

# first of all  
The general recommendation is to use the latest version of your operating system.  
Use a compiler that supports at least c++17 well enough(this is necessary for the successful compilation of some third-party libraries).  
stock up on disk space before installation (100 gigabytes for reliability).  
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
make -j$(nproc)
make install 
```  
add gcc to PATH:  
e.g. in file $HOME/.profile write line  
export PATH="$HOME/gcc-13.1.0/bin:$PATH"

if you are using wsl - restart it
check supported versions of c++  
```console  
gcc -v --help 2> /dev/null | sed -n '/^ *-std=\([^<][^ ]\+\).*/ {s//\1/p}'
```  


# NOTE
check system path with
```console  
echo $PATH
echo $HOME
```
use includes of vcpkg installed packages in ide
1. for VScode  
```console  
${HOME}/vcpkg/packages//* name of library*//include
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
vcpkg install /*name of library*/:x64-linux 
vcpkg integrate install  
```  
# install [crow](https://github.com/CrowCpp/Crow)  
```console  
vcpkg install crow:x64-linux  
vcpkg integrate install  
```
add the following lines to CMakeLists.txt
```console    
find_package(Crow CONFIG REQUIRED)  
target_link_libraries(/*your target name*/ PRIVATE Crow::Crow)  
```


# build project  
from folder with CMakeLists.txt  
```console  
mkdir build  
cd build  
cmake -DCMAKE_TOOLCHAIN_FILE=/home/user/vcpkg/scripts/buildsystems/vcpkg.cmake -DCMAKE_C_COMPILER=/home/user/gcc-13.1.0/bin/gcc -DCMAKE_CXX_COMPILER=/home/user/gcc-13.1.0/bin/g++ ..
cmake --build ./ --config Release  
```
NOTE:  
if y use vscode  
y can use cmake-tools extension and settings.json file 

