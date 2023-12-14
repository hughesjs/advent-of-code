#include "parser.h"

#include <stdlib.h>
#include <string.h>

#define MAX_TOKENS 100

char** splitString(const char* sourceString, const char* delimiter, int* countOut) {
    const int srcLen = strlen(sourceString) + 1; // +1 for null terminator
    char* stringToTokenize = malloc(sizeof(char *) * srcLen);
    strncpy((char *) stringToTokenize, sourceString, srcLen);

    char** substrings = malloc(MAX_TOKENS * sizeof(char*));

    int i = 0;
    const char* token = strtok(stringToTokenize, delimiter);

    while (token != NULL) {
        substrings[i] = malloc(strlen(token) + 1);
        strcpy(substrings[i], token);
        token = strtok(NULL, delimiter);
    }
    *countOut = i;
    return substrings;
}


void freeStrings(char** strings, int count) {
    for (int i = 0; i < count; ++i) {
        free(strings[i]);
    }
    free(strings);
}

void parseSingleGameLine(const char* line, GameResult* destination) {
    destination->id = 69;
}

GameResult* parseLinesToGames(const char** lines, int numLines) {
    GameResult* results = malloc(sizeof(GameResult) * numLines);
    for (int i = 0; i < numLines; ++i) {
        parseSingleGameLine(lines[i], &results[i]);
    }
    return results;
}

GameResult* parseAllGames(const char* inputData) {
    int numParts;
    const char** gameLines = (const char**)splitString(inputData, (const char*)"\n", &numParts);

    GameResult* games = parseLinesToGames(gameLines, numParts);

    freeStrings((char**)gameLines, numParts);
    return games;
}



