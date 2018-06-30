cmake_minimum_required(VERSION 3.5)
project(Alpha)

set(CMAKE_CXX_STANDARD 11)

set(SOURCE_FILES
#        hdoj-1000.cpp
#        rollgame.cpp
#        sprial.c
        foo.cpp foo.h test/fooTest.cpp)

add_executable(alpha ${SOURCE_FILES})
target_link_libraries(alpha  gtestd)