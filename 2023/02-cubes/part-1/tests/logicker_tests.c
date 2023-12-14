#include <unity.h>
#include "logicker_tests.h"
#include "logicker.h"


void provided_test_case(void) {
    const char* testData = R"(Game 1: 3 blue, 4 red; 1 red, 2 green, 6 blue; 2 green
Game 2: 1 blue, 2 green; 3 green, 4 blue, 1 red; 1 green, 1 blue
Game 3: 8 green, 6 blue, 20 red; 5 blue, 4 red, 13 green; 5 green, 1 red
Game 4: 1 green, 3 red, 6 blue; 3 green, 6 red; 3 green, 15 blue, 14 red
Game 5: 6 red, 1 blue, 3 green; 2 blue, 1 red, 2 green)";

    const int expected = 8;

    const int answer = deduceAnswer(testData, 12, 13, 14);

    TEST_ASSERT_EQUAL_INT(expected, answer);
}

void run_logicker_tests(void) {
    RUN_TEST(provided_test_case);
}