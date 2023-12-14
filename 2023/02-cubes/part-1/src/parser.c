#include <stdlib.h>

#include "parser.h"
#include "string_utils.h"

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



