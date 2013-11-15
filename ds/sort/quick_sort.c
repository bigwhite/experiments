#include <stdio.h>
#include "utils.h"

/* 
    [3, 8, 2, 1, 5, 4, 6, 7]
*/

int
partition(int *arr, int left, int right, int pivot)
{
    int i, j;
    swap(arr, pivot, right);
    
    i = left;
    j = right - 1;

    while (1) {
        while (arr[i] < arr[right]) i++;
        while (arr[j] > arr[right]) j--;
        if (i <= j) 
            swap(arr, i, j);
        else
            break;
    }

    swap(arr, i, right);

    return i;
}


void
quick_sort(int *arr, int left, int right)
{
    if (left >= right) {
        return; /* already sorted */
    }

    int pivot = left + (((right - left)%2 == 0) ? (right - left)/2: ((right - left)/2 + 1));

    pivot = partition(arr, left, right, pivot);

    quick_sort(arr, left, pivot - 1);
    quick_sort(arr, pivot + 1, right);
}

int 
main()
{
    int arr[] = {3, 8, 2, 1, 5, 4 , 6, 7};

    output_arr(arr, sizeof(arr)/sizeof(arr[0]));

    quick_sort(arr, 0, sizeof(arr)/sizeof(arr[0]) - 1);

    output_arr(arr, sizeof(arr)/sizeof(arr[0]));
}
