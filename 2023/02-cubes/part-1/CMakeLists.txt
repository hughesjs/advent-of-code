cmake_minimum_required(VERSION 3.22)
project(cube_conundrum_part_1)

set(CMAKE_CXX_STANDARD 17)

# Main exe
add_executable(cube_conundrum
        src/main.c
        src/parser.c
        src/logicker.c)

target_include_directories(cube_conundrum PUBLIC include)


# Test exe
add_executable(tests
        src/parser.c
        src/logicker.c
        #vendor/Unity/src/unity.c
        tests/tests.c
        )
target_include_directories(tests PRIVATE include)

add_subdirectory(vendor/Unity)
target_link_libraries(tests unity)




enable_testing()
add_test(NAME example_case_is_correct COMMAND tests)