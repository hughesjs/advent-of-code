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

void can_parse_every_game_id_in_range(void) {
    for (int i = 1; i <= 100; ++i) {
        char* idStringBuf = malloc(sizeof(char) * (GAME_ID_PREFIX_MAX_LEN));
        sprintf(idStringBuf, "Game %d:", i);

        const unsigned short res = getGameId(idStringBuf);
        free(idStringBuf);

        TEST_ASSERT_EQUAL_INT(i, res);
    }
}

void run_parser_tests(void) {
    RUN_TEST(can_parse_simple_single_line);
    RUN_TEST(can_parse_every_game_id_in_range);
}

