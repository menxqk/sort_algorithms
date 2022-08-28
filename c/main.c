#include <stdio.h>
#include <stdlib.h>

#define MAX 100000
const char *NUMBERS_FILE = "../numbers.txt";

static int nums[MAX];

void main(void)
{
    FILE *fptr = fopen(NUMBERS_FILE, "r");
    if (fptr == NULL)
    {
        printf("Error opening file: %s\n", NUMBERS_FILE);
        exit(1);
    }

    int count = 0;
    for (int i = 0; i < MAX; i++)
    {
        int num = 0;
        if (!fscanf(fptr, "%d,", &num))
        {
            break;
        }
        nums[i] = num;
        count++;
    }

    printf("[");
    for (int i = 0; i < count; i++)
    {
        if (nums[i] == -1)
        {
            break;
        }
        printf("%d ", nums[i]);
    }
    printf("]\n");
    printf("len(nums): %d\n", count);

    if (fptr != NULL)
    {
        fclose(fptr);
    }
}
