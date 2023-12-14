#include <unity.h>

#include "string_utils.h"

void can_split_string(void) {
    const char* test_data = "123;456;789";
    int resCount;
    const char** results = splitString(test_data, ";", &resCount);

    TEST_ASSERT_EQUAL_INT(3, resCount);
    TEST_ASSERT_EQUAL_STRING("123", results[0]);
    TEST_ASSERT_EQUAL_STRING("456", results[1]);
    TEST_ASSERT_EQUAL_STRING("789", results[2]);
}

void source_string_unaffected_after_split(void) {
    const char* test_data = "123;456;789";
    int resCount;
    splitString(test_data, ";", &resCount);

    TEST_ASSERT_EQUAL_STRING("123;456;789", test_data);
}

void works_for_zero_splits(void) {
    const char* test_data = "123";
    int resCount;
    const char** results = splitString(test_data, ";", &resCount);

    TEST_ASSERT_EQUAL_INT(1, resCount);
    TEST_ASSERT_EQUAL_STRING("123", results[0]);
}

void can_get_int_from_substring(void) {
    const char* test_data = "1234567890";
    const int res = atoiSubstring(test_data, 2, 4);
    TEST_ASSERT_EQUAL_INT(345, res);
}

void can_get_int_from_single_digit_substring(void) {
    const char* test_data = "1234567890";
    const int res = atoiSubstring(test_data, 0, 0);
    TEST_ASSERT_EQUAL_INT(1, res);
}

void run_string_tests(void) {
    RUN_TEST(can_split_string);
    RUN_TEST(source_string_unaffected_after_split);
    RUN_TEST(works_for_zero_splits);
    RUN_TEST(can_get_int_from_substring);
    RUN_TEST(can_get_int_from_single_digit_substring);
}

