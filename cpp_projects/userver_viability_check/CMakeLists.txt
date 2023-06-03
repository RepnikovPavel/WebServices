cmake_minimum_required(VERSION 3.26.4)
project(check_userver)

set(CMAKE_CXX_STANDARD 23)

include(/home/user/userver/cmake/SetupEnvironment.cmake)
include(GNUInstallDirs)
add_subdirectory(/home/user/userver)


add_executable(uservtmp main.cpp hello.cpp hello.hpp)
# set(userver_include 
# /home/user/userver/third_party/amqp-cpp/include
# /home/user/userver/third_party/spdlog/include
# /home/user/userver/third_party/pfr/include
# /home/user/userver/third_party/date/include
# /home/user/userver/third_party/clickhouse-cpp/contrib/gtest/include
# /home/user/userver/third_party/boost_stacktrace/include
# /home/user/userver/third_party/c-ares/include
# /home/user/userver/third_party/moodycamel/include
# /home/user/userver/third_party/rapidjson/include
# /home/user/userver/third_party/fmt/include
# /home/user/userver/build_release/third_party/c-ares/include
# /home/user/userver/grpc/include
# /home/user/userver/grpc/handlers/include
# /home/user/userver/rabbitmq/include
# /home/user/userver/universal/include
# /home/user/userver/mysql/include
# /home/user/userver/postgresql/include
# /home/user/userver/mongo/include
# /home/user/userver/redis/include
# /home/user/userver/redis/testing/include
# /home/user/userver/shared/include
# /home/user/userver/shared/internal/include
# /home/user/userver/core/include
# /home/user/userver/core/internal/include
# /home/user/userver/core/testing/include
# /home/user/userver/uboost_coro/include
# /home/user/userver/clickhouse/include
# )

# target_include_directories(uservtmp PUBLIC ${userver_include})


#  cmake -D CMAKE_C_COMPILER=/home/user/gcc-13.1.0/bin/gcc -D CMAKE_CXX_COMPILER=/home/user/gcc-13.1.0/bin/g++ ..
#  cmake --build ./ --config Release