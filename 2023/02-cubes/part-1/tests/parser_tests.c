#include <unity.h>
#include <stdlib.h>
#include <string.h>

#include "parser_tests.h"
#include "parser.h"

void can_parse_simple_single_line(void) {
    const char* testData = "Game 1: 3 blue, 4 red; 1 red, 2 green, 6 blue; 2 green";

    int numRes;
    const GameResult* result = parseAllGames(testData, &numRes);

    TEST_ASSERT_EQUAL_INT(1, result->game_id);
    TEST_ASSERT_EQUAL_INT(6, result->blue_max);
    TEST_ASSERT_EQUAL_INT(4, result->red_max);
    TEST_ASSERT_EQUAL_INT(2, result->green_max);
    TEST_ASSERT_EQUAL_INT(1, numRes);
}

void can_parse_more_complex_single_line(void) {
    const char* testData = "Game 18: 4 blue, 1 red, 14 green; 8 red, 4 blue, 14 green; 6 red, 11 blue, 10 green; 5 blue, 2 green, 3 red; 16 green, 10 blue, 2 red; 6 red, 4 blue, 12 green";

    int numRes;
    const GameResult* result = parseAllGames(testData, &numRes);

    TEST_ASSERT_EQUAL_INT(18, result->game_id);
    TEST_ASSERT_EQUAL_INT(11, result->blue_max);
    TEST_ASSERT_EQUAL_INT(8, result->red_max);
    TEST_ASSERT_EQUAL_INT(16, result->green_max);
    TEST_ASSERT_EQUAL_INT(1, numRes);
}

void can_parse_every_game_id_in_range(void) {
    for (int i = 1; i <= 100; ++i) {
        char* idStringBuf = malloc(sizeof(char) * (GAME_ID_PREFIX_MAX_LEN));
        sprintf(idStringBuf, "Game %d:", i);

        const unsigned short res = getGameId(idStringBuf);
        free(idStringBuf);

        TEST_ASSERT_EQUAL_INT(i, res);
    }
}

void can_parse_green_without_affecting_red(void) {
    const char* test_data = "12 green";
    int r = 0, g = 0, b = 0;
    processRevealSubstring(test_data, &r, &g, &b);

    TEST_ASSERT_EQUAL_INT(0, r);
    TEST_ASSERT_EQUAL_INT(12, g);
}

void run_parser_tests(void) {
    RUN_TEST(can_parse_simple_single_line);
    RUN_TEST(can_parse_every_game_id_in_range);
    RUN_TEST(can_parse_more_complex_single_line);
    RUN_TEST(can_parse_green_without_affecting_red);
}

