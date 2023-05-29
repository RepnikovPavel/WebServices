cmake_minimum_required(VERSION 3.24.0)
project(WebServices)

set(CMAKE_CXX_STANDARD 23)

set(libevent_include C:/vcpkg/vcpkg/packages/libevent_x64-windows/include)
set(libevent_lib C:/vcpkg/vcpkg/packages/libevent_x64-windows/lib)
set(libevent_bin C:/vcpkg/vcpkg/packages/libevent_x64-windows/bin)

set(boost_include C:/cpp_libs_source/boost/boost_1_82_0)
add_executable(EchoServ ChatRoom/main.cpp)
target_include_directories(EchoServ PUBLIC ${boost_include})
find_package(Libevent CONFIG REQUIRED)
target_include_directories(EchoServ PUBLIC ${libevent_include})
target_link_libraries(EchoServ PRIVATE libevent::core libevent::extra)

#target_include_directories(EchoServ PUBLIC ${libevent_include})
#target_link_libraries(EchoServ PUBLIC
#        "${libevent_lib}/event.lib"
#        "${libevent_lib}/event_core.lib"
#        "${libevent_lib}/event_extra.lib"
#        wsock32.lib)
