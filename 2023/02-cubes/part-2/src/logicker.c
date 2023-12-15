#include "logicker.h"
#include "parser.h"

unsigned int getPower(const GameResult* gameResult) {
    return (int)gameResult->red_max * (int)gameResult->green_max * (int)gameResult->blue_max;
}

int deduceAnswer(const char* inputData) {
    int numGames;
    const GameResult* gameResults = parseAllGames(inputData, &numGames);

    unsigned int acc = 0;
    for (int i = 0; i < numGames; ++i) {
        acc += getPower(&gameResults[i]);
    }
    return acc;
}



