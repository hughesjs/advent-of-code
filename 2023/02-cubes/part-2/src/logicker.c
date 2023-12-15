#include "logicker.h"

#include <stdio.h>

#include "parser.h"

unsigned int getPower(const GameResult* gameResult) {
    const int power = (int)gameResult->red_max * (int)gameResult->green_max * (int)gameResult->blue_max;
    printf("GameId: %d, R: %d, G: %d, B: %d, POW: %d\n", gameResult->game_id, gameResult->red_max, gameResult->green_max, gameResult->blue_max, power);
    return power;
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



