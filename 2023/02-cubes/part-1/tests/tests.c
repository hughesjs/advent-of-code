#include "unity.h"

#include "parser_tests.h"
#include "string_utils_tests.h"
#include "logicker_tests.h"


void unity_self_test(void) {
    TEST_PASS();
}

void setUp(void) {}
void tearDown(void) {}

int main(void) {
    UNITY_BEGIN();

    RUN_TEST(unity_self_test);

    run_string_tests();
    run_parser_tests();
    run_logicker_tests();

    return UNITY_END();
}

