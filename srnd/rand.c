#include "rand.h"

static bool initialized = false;

int get_random_number()
{
    if (!initialized)
    {
        time_t t = time(NULL);
        srand(t);
        initialized = true;
        printf("init rand seed: %ld\n", t);
    }

    int r = rand() % 1000000;
    return r;
}
