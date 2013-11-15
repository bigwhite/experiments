#include <stdio.h>
#include "utils.h"

void
insert_sort(int *arr, size_t arr_sz)
{
    int i;
    int j;

    for (i = 0; i < arr_sz; i++) {
        if (i == 0) 
            continue; /* one element, already sorted */

        for (j = i; j > 0; j--) {
            if (arr[j] < arr[j - 1]) {
                swap(arr, j, j - 1); 
            }
        } 
    }
}

int 
main()
{
    int arr[] = {3, 8, 2, 1, 5, 4 , 6, 7};

    output_arr(arr, sizeof(arr)/sizeof(arr[0]));

    insert_sort(arr, sizeof(arr)/sizeof(arr[0]));

    output_arr(arr, sizeof(arr)/sizeof(arr[0]));
}
