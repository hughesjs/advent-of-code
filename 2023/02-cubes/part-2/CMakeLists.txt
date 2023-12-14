cmake_minimum_required(VERSION 3.22)
project(cube_conundrum_part_1)

set(CMAKE_CXX_STANDARD 17)

# Main exe
file(GLOB SOURCES src/*.c)
add_executable(cube_conundrum ${SOURCES})

target_include_directories(cube_conundrum PUBLIC include)


# Test exe
file(GLOB TEST_SOURCES tests/*.c)

# Remove main.c from src
file(GLOB SRC_MAIN src/main.c)
list(REMOVE_ITEM SOURCES ${SRC_MAIN})

add_executable(tests ${SOURCES} ${TEST_SOURCES}
        tests/string_utils_tests.c
        tests/string_utils_tests.h)
target_include_directories(tests PRIVATE include)

add_subdirectory(vendor/Unity)
target_link_libraries(tests unity)




enable_testing()
add_test(NAME example_case_is_correct COMMAND tests)