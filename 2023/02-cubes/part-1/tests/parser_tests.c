#include "parser_tests.h"
#include "parser.h"

#include "unity.h"

void can_parse_simple_single_line(void) {
    const char* testData = "Game 1: 3 blue, 4 red; 1 red, 2 green, 6 blue; 2 green";

    const GameResult* result = parseAllGames(testData);

    ;
//TODO
}

void run_parser_tests(void) {
    RUN_TEST(can_parse_simple_single_line);
}

