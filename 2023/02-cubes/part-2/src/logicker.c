#include <stdbool.h>
#include "logicker.h"

#include "parser.h"

bool isGamePossible(const GameResult* gameResult, const int totalRed, const int totalGreen, const int totalBlue) {
    return  gameResult->red_max <= totalRed &&
           gameResult->blue_max <= totalBlue &&
           gameResult->green_max <= totalGreen;
}

int deduceAnswer(const char* inputData, const int totalRed, const int totalGreen, const int totalBlue) {
    int numGames;
    const GameResult* gameResults = parseAllGames(inputData, &numGames);

    int acc = 0;
    for (int i = 0; i < numGames; ++i) {
        if (isGamePossible(&gameResults[i], totalRed, totalGreen, totalBlue)) {
            acc += gameResults[i].game_id;
        }
    }
    return acc;
}


