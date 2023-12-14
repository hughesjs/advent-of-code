#include "unity.h"
#include "logicker.h"
#include "parser.h"

void provided_test_case(void) {
    const char* testData = R"(Game 1: 3 blue, 4 red; 1 red, 2 green, 6 blue; 2 green
Game 2: 1 blue, 2 green; 3 green, 4 blue, 1 red; 1 green, 1 blue
Game 3: 8 green, 6 blue, 20 red; 5 blue, 4 red, 13 green; 5 green, 1 red
Game 4: 1 green, 3 red, 6 blue; 3 green, 6 red; 3 green, 15 blue, 14 red
Game 5: 6 red, 1 blue, 3 green; 2 blue, 1 red, 2 green)";

    const int expected = 8;

    const int answer = deduceAnswer(testData);

    TEST_ASSERT_EQUAL_INT(expected, answer);
}

void unity_self_test(void) {
    TEST_PASS();
}

void can_parse_simple_single_line(void) {
    const char* testData = "Game 1: 3 blue, 4 red; 1 red, 2 green, 6 blue; 2 green";

    const GameResult* result = parseAllGames(testData);

    ;
}

void setUp(void) {}
void tearDown(void) {}

int main(void) {
    UNITY_BEGIN();

    //RUN_TEST(provided_test_case);
    RUN_TEST(can_parse_simple_single_line);

    RUN_TEST(unity_self_test);


    return UNITY_END();
}

