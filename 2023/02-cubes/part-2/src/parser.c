#include <stdlib.h>
#include "parser.h"

#include <stdbool.h>
#include <string.h>

#include "string_utils.h"


unsigned short getGameId(const char* line) {
    char* idBuf = malloc(sizeof(char *) * GAME_ID_MAX_LEN);

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

char* extractDataPortionFromGameLine(const char* line) {
    unsigned int dataStartIndex = 0;
    for (unsigned short i = GAME_ID_NUM_START_INDEX; i < GAME_ID_PREFIX_MAX_LEN; ++i) {
        if (line[i] == ' ') {
            dataStartIndex = i + 1;
            break;
        }
    }

    const unsigned int substringLength = strlen(line) - dataStartIndex;
    char* gameDataSubstring = malloc(substringLength + 1);
    memcpy(gameDataSubstring, line + sizeof(char) * dataStartIndex, substringLength + 1);
    return gameDataSubstring;
}

bool isNumeric(const char* character) {
    return *character > 47 && *character < 58;
}

void processRevealSubstring(const char* revealSubstring, unsigned short* rMax, unsigned short* gMax, unsigned short* bMax) {
    int firstDigitIndex = 0;
    int lastDigitIndex = 0;
    bool justWritten = false;
    bool firstDigitFound = false;
    for (int k = 0; k < strlen(revealSubstring); ++k) {
        char carat = revealSubstring[k];
        if (justWritten || (carat == ' ' && !firstDigitIndex)) {
            justWritten = false;
            continue;
        }

        if (isNumeric(&carat)) {
            if (!firstDigitFound) {
                firstDigitFound = true;
                firstDigitIndex = k;
            }
            lastDigitIndex = k;
            continue;
        }

        if (carat == 'r') {
            justWritten = true;
            const unsigned int r = atoiSubstring(revealSubstring, firstDigitIndex, lastDigitIndex);
            if (r > *rMax) {
                *rMax = r;
            }
            continue;
        }

        if (carat == 'g') {
            justWritten = true;
            const unsigned int g = atoiSubstring(revealSubstring, firstDigitIndex, lastDigitIndex);
            if (g > *gMax) {
                *gMax = g;
            }
            continue;
        }

        if (carat == 'b') {
            justWritten = true;
            const unsigned int b = atoiSubstring(revealSubstring, firstDigitIndex, lastDigitIndex);
            if (b > *bMax) {
                *bMax = b;
            }
            continue;
        }
    }
}

// "Game 1: 3 blue, 4 red; 1 red, 2 green, 6 blue; 2 green"
void parseSingleGameLine(const char* line, GameResult* destination) {
    destination->game_id = getGameId(line);
    destination->red_max = 0;
    destination->green_max = 0;
    destination->blue_max = 0;

    const char* dataSubstring = extractDataPortionFromGameLine(line);
    int rounds;
    char** roundSubstrings = splitString(dataSubstring, ";", &rounds);
    for (int i = 0; i < rounds; ++i) {
        const char* round = roundSubstrings[i];
        int numReveals;
        char** reveals = splitString(round, ",", &numReveals);
        for (int j = 0; j < numReveals; ++j) {
            processRevealSubstring(reveals[j],
                &destination->red_max,
                &destination->green_max,
                &destination->blue_max);
        }
    }
}

GameResult* parseLinesToGames(const char** lines, const int numLines) {
    GameResult* results = malloc(sizeof(GameResult) * numLines);
    for (int i = 0; i < numLines; ++i) {
        parseSingleGameLine(lines[i], &results[i]);
    }
    return results;
}

GameResult* parseAllGames(const char* inputData, int* outNumGames) {
    int numParts;
    const char** gameLines = splitString(inputData, "\n", &numParts);
    GameResult* games = parseLinesToGames(gameLines, numParts);
    freeStrings(gameLines, numParts);
    *outNumGames = numParts;
    return games;
}
