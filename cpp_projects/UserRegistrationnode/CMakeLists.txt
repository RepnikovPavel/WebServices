cmake_minimum_required(VERSION 3.26.4)
project(UserRegistrationNode)
set(target_name "UserRegistrator")
message("target name: "${target_name})
set(CMAKE_CXX_STANDARD 23)

# set(vcpkgRoot "/home/user/vcpkg/installed/x64-linux")
# set(vcpkgInclude "${vcpkgRoot}/include")

add_executable(${target_name} src/main.cpp)
message("include path: "${vcpkgInclude})
# target_include_directories(${target_name} PUBLIC ${vcpkgInclude})

find_package(Crow CONFIG REQUIRED)
target_link_libraries(${target_name} PRIVATE Crow::Crow)

