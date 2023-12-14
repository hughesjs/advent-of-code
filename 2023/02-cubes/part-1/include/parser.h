#pragma once

typedef struct {
    unsigned short id;
    unsigned short blue;
    unsigned short red;
    unsigned short green;
} GameResult;

GameResult* parseAllGames(const char* inputData);

