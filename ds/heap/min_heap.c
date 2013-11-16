/*
 * min_heap.c
 *
 * brief description of this file
 */

#include "utils.h"
#include <stdio.h>
#include <stdlib.h>
#include <string.h>

#define MIN_HEAP_EMPTY   -11
#define MIN_HEAP_FULL    -12
#define MIN_HEAP_SUCCESS  0

struct min_heap_t {
    size_t heap_sz;
    int items_count; /* start from 0 */
    int *arr;
};

struct min_heap_t* 
min_heap_new(size_t sz)
{
    struct min_heap_t *p = NULL;

    p = malloc(sizeof(*p));
    if (!p) 
        return NULL;

    p->heap_sz = sz;
    p->items_count = 0;
    p->arr = malloc(sz * sizeof(int));
    if (p->arr == NULL) {
        free(p);
        return NULL;
    }
    memset(p->arr, 0, sizeof(sz * sizeof(int)));

    return p;
}

int 
min_heap_insert(struct min_heap_t *hp, int item)
{
    if (hp->items_count == hp->heap_sz)
        return MIN_HEAP_FULL;

    /* first, place the item to the end
     * then, adjust the heap */
    int index = hp->items_count;
    hp->arr[index] = item;

    int parent_index;
    while (index > 0) {
        parent_index = (index - 1)/2;
        if (hp->arr[parent_index] > item) {
            swap(hp->arr, parent_index, index);
            index = parent_index;
        } else {
            break;
        }
    }

    hp->items_count++;
    return MIN_HEAP_SUCCESS;
}

int 
min_heap_remove_min(struct min_heap_t *hp)
{
    if (hp->items_count == 0)
        return MIN_HEAP_EMPTY;

    int min_item = hp->arr[0];
    hp->arr[0] = hp->arr[hp->items_count - 1];
    hp->items_count--;

    int right_child_index;
    int left_child_index;
    int cur_index = 0;
    int min_child_index;

    while (cur_index < hp->items_count) {
        left_child_index = (cur_index  * 2) + 1;
        right_child_index = (cur_index  * 2) + 2;

        if (left_child_index >= hp->items_count) break;
        if (right_child_index >= hp->items_count) {
            min_child_index = right_child_index;
        } else if (hp->arr[right_child_index] < hp->arr[left_child_index]) {
            min_child_index = right_child_index;
        } else {
            min_child_index = left_child_index;
        }

        if (hp->arr[cur_index] > hp->arr[min_child_index]) {
            swap(hp->arr, cur_index, min_child_index);
            cur_index = min_child_index;
        } else {
            break;
        }
    }

    return min_item;
}

int 
min_heap_items_count(const struct min_heap_t *hp)
{
    return hp->items_count;
}

void
min_heap_free(struct min_heap_t **php)
{
    struct min_heap_t *hp = (*php);
    free(hp->arr);
    free(hp);
    (*php) = NULL;
}

void 
min_heap_dump(const struct min_heap_t *hp)
{
    int i;
    printf("min heap dump ==>\n");
    for (i = 0; i < hp->items_count; i++) {
        printf("%d ", hp->arr[i]);
    }
    printf("\n");
    printf("<=== min heap dump end\n");
}

int
main()
{
    int arr[] = {3, 8, 2, 1, 5, 4, 6, 7};
    size_t sz = sizeof(arr)/sizeof(arr[0]);
    output_arr(arr, sizeof(arr)/sizeof(arr[0]));

    struct min_heap_t *hp = min_heap_new(sz);
    if (!hp) {
        printf("min_heap_new error\n");
        return -1;
    }

    int i;
    for (i = 0; i < sz; i++)
        min_heap_insert(hp, arr[i]);
    min_heap_dump(hp);  /* 1 2 3 7 5 4 6 8 */

    int new_arr[8] = {0};
    for (i = 0; i < sz; i++) {
        new_arr[i] = min_heap_remove_min(hp);
        min_heap_dump(hp);
    }
    
    output_arr(new_arr, sizeof(new_arr)/sizeof(new_arr[0]));
    min_heap_free(&hp);
}
