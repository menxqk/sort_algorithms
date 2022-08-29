#include <stdio.h>
#include <string.h>

void copy_array(int src[], int dst[], int size)
{
    memcpy(dst, src, size);
}

void print_array(int arr[], int size)
{
    printf("[");
    for (int i = 0; i < size; i++)
    {
        if (i < size - 1)
        {
            printf("%d ", arr[i]);
        }
        else
        {
            printf("%d", arr[i]);
        }
    }
    printf("]");
}