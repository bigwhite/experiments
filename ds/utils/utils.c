
#include "utils.h"

void
output_arr(int *arr, size_t arr_sz)
{

    int i = 0;
    printf("[");
    for (i = 0; i < arr_sz; i++) {
        if (i == arr_sz - 1)
            printf("%d", *(arr + i));
        else
            printf("%d ", *(arr + i));
    }
    printf("]\n");
}


void 
swap(int *arr, int left, int right)
{
    if (left != right) {
        int temp = arr[left];
        arr[left] = arr[right];
        arr[right] = temp;
    }
}

