#include <stdlib.h>
#include "parser.h"
#include "string_utils.h"



unsigned short getGameId(const char* line) {
    char* idBuf = malloc(sizeof(char*) * GAME_ID_MAX_LEN);

    for (int i = 0; i < GAME_ID_MAX_LEN; ++i) {
        const int j = GAME_ID_NUM_START_INDEX + i;
        const char carat = line[j];
        if (carat == ':') {
            break;
        }
        idBuf[i] = carat;
    }

    return atoi(idBuf);
}

// "Game 1: 3 blue, 4 red; 1 red, 2 green, 6 blue; 2 green"
void parseSingleGameLine(const char* line, GameResult* destination) {
    unsigned short id = getGameId(line);
}

GameResult* parseLinesToGames(const char** lines, const int numLines) {
    GameResult* results = malloc(sizeof(GameResult) * numLines);
    for (int i = 0; i < numLines; ++i) {
        parseSingleGameLine(lines[i], &results[i]);
    }
    return results;
}

GameResult* parseAllGames(const char* inputData) {
    int numParts;
    const char** gameLines = splitString(inputData, "\n", &numParts);

    GameResult* games = parseLinesToGames(gameLines, numParts);

    freeStrings(gameLines, numParts);
    return games;
}



