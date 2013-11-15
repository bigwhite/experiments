#include <stdio.h>
#include "utils.h"

void
select_sort(int *arr, size_t arr_sz)
{
    int i;
    int j;
    int min;
    int min_index;

    for (i = 0; i < arr_sz; i++) { // n pass

        /* find the smallest item start from the current position */
        min = arr[i];
        min_index = i;
        for (j = i + 1; j < arr_sz; j++) {
            if (arr[j] < min) {
                min = arr[j];
                min_index = j;
            }
        }

        if (i != min_index)
            swap(arr, i, min_index);

    }

}

int 
main()
{
    int arr[] = {3, 8, 2, 1, 5, 4 , 6, 7};

    output_arr(arr, sizeof(arr)/sizeof(arr[0]));

    select_sort(arr, sizeof(arr)/sizeof(arr[0]));

    output_arr(arr, sizeof(arr)/sizeof(arr[0]));
}
