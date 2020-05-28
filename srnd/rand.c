#include "rand.h"

static bool initialized = false;

int get_random_number()
{
    if (!initialized)
    {
        srand(time(NULL));
        initialized = true;
        printf("init rand seed\n");
    }

    int r = rand() % 1000000;
    return r;
}
