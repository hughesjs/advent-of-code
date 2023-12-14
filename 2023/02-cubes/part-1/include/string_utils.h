#pragma once

char** splitString(const char* sourceString, const char* delimiter, int* countOut);
int atoiSubstring(const char* sourceString, int startIndex, int endIndex);
void freeStrings(char** strings, int count);