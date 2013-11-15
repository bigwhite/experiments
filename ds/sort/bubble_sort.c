#include <stdio.h>
#include "utils.h"

/* 
    [3, 8, 2, 1, 5, 4, 6, 7]
*/
void
bubble_sort(int *arr, size_t arr_sz)
{
    int pass;
    int i;
    int swapped= 0;

    for (pass = 0; pass < arr_sz; pass++) {
        for (i = 1; i < arr_sz; i++) {
            if (arr[i-1] > arr[i]) {
                swap(arr, i-1, i);
                swapped = 1;
            }
        }

        if (swapped == 0) break;
    }

}

int 
main()
{
    int arr[] = {3, 8, 2, 1, 5, 4 , 6, 7};

    output_arr(arr, sizeof(arr)/sizeof(arr[0]));

    bubble_sort(arr, sizeof(arr)/sizeof(arr[0]));

    output_arr(arr, sizeof(arr)/sizeof(arr[0]));
}
