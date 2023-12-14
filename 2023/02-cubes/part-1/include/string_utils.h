#pragma once

char** splitString(const char* sourceString, const char* delimiter, int* countOut);

void freeStrings(char** strings, int count);