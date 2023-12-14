#pragma once

#define GAME_ID_MAX_LEN 3
#define GAME_ID_NUM_START_INDEX 5

typedef struct {
    unsigned short game_id;
    unsigned short blue_max;
    unsigned short red_max;
    unsigned short green_max;
} GameResult;

GameResult* parseAllGames(const char* inputData);

// Lazy for testing
unsigned short getGameId(const char* line);

