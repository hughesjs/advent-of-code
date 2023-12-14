#include <stdio.h>
#include <stdlib.h>
#include <sys/stat.h>

#include "logicker.h"

int main(const int argc, const char** argv) {
    if (!argc == 1) {
        printf("Ensure you're passing a file-path to your input data");
        return -1;
    }

    const char* dataPath = argv[1];

    FILE* file = fopen(dataPath, "r");

    struct stat fileStatBuf;
    fstat(fileno(file), &fileStatBuf);

    char* fileContents = malloc(fileStatBuf.st_size + 1);
    fread(fileContents, fileStatBuf.st_size, fileStatBuf.st_size, file);
    fclose(file);
    const int answer = deduceAnswer(fileContents, 12, 13, 14);
    printf("Answer: %d", answer);

    free(fileContents);
    return 0;
}
