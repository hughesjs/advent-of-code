#include "string_utils.h"
#include "string.h"
#include <stdlib.h>

#define MAX_TOKENS 100

char** splitString(const char* sourceString, const char* delimiter, int* countOut) {
    const int srcLen = strlen(sourceString) + 1; // +1 for null terminator
    char* stringToTokenize = malloc(sizeof(char *) * srcLen);
    strncpy(stringToTokenize, sourceString, srcLen);

    char** substrings = malloc(MAX_TOKENS * sizeof(char*));

    int i = 0;
    const char* token = strtok(stringToTokenize, delimiter);

    while (token != NULL) {
        substrings[i] = malloc(strlen(token) + 1);
        strcpy(substrings[i], token);
        i++;
        token = strtok(NULL, delimiter);
    }
    *countOut = i;
    free(stringToTokenize);
    return substrings;
}


void freeStrings(char** strings, int count) {
    for (int i = 0; i < count; ++i) {
        free(strings[i]);
    }
    free(strings);
}
