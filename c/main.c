#include <stdio.h>
#include <stdlib.h>
#include <time.h>
#include "arr.h"
#include "sorts.h"

#define MAX 100000

static int nums[MAX];
static int nums_bubble[MAX];
static int nums_selection[MAX];
static int nums_insertion[MAX];
static int nums_merge[MAX];
static int nums_heap[MAX];
static int nums_quick[MAX];

static int array_size = 0;

void main(int argc, char *argv[])
{
    if (argc < 2)
    {
        printf("input file missing");
        exit(1);
    }

    char *input_file = argv[1];

    FILE *fptr = fopen(input_file, "r");
    if (fptr == NULL)
    {
        printf("Error opening file: %s\n", input_file);
        exit(1);
    }

    int num = 0;
    for (int i = 0; i < MAX; i++)
    {
        if (fscanf(fptr, "%d,", &num) < 0)
        {
            break;
        }

        nums[i] = num;
        array_size++;
    }

    copy_array(nums, nums_bubble, array_size * sizeof(int));
    copy_array(nums, nums_selection, array_size * sizeof(int));
    copy_array(nums, nums_insertion, array_size * sizeof(int));
    copy_array(nums, nums_merge, array_size * sizeof(int));
    copy_array(nums, nums_heap, array_size * sizeof(int));
    copy_array(nums, nums_quick, array_size * sizeof(int));

    clock_t t;
    double diff;

    // bubble sort
    t = clock();
    bubble_sort(nums_bubble, array_size);
    t = clock() - t;
    diff = ((double)t / CLOCKS_PER_SEC);
    printf("bubble sort:    %.6lfs\n", diff);

    // selection sort
    t = clock();
    selection_sort(nums_selection, array_size);
    t = clock() - t;
    diff = ((double)t / CLOCKS_PER_SEC);
    printf("selection sort: %.6lfs\n", diff);

    // insertion sort
    t = clock();
    insertion_sort(nums_insertion, array_size);
    t = clock() - t;
    diff = ((double)t / CLOCKS_PER_SEC);
    printf("insertion sort: %.6lfs\n", diff);

    // merge sort
    t = clock();
    merge_sort(nums_merge, 0, array_size - 1);
    t = clock() - t;
    diff = ((double)t / CLOCKS_PER_SEC);
    printf("merge sort:     %.6lfs\n", diff);

    // heap sort
    t = clock();
    heap_sort(nums_heap, array_size);
    t = clock() - t;
    diff = ((double)t / CLOCKS_PER_SEC);
    printf("heap sort:      %.6lfs\n", diff);

    // quick sort
    t = clock();
    quick_sort(nums_quick, 0, array_size - 1);
    t = clock() - t;
    diff = ((double)t / CLOCKS_PER_SEC);
    printf("quick sort:     %.6lfs\n", diff);

    if (fptr != NULL)
    {
        fclose(fptr);
    }
}
