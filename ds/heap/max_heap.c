/*
 * max_heap.c
 *
 * brief description of this file
 */

#include "utils.h"
#include <stdio.h>
#include <stdlib.h>
#include <string.h>

#define MAX_HEAP_EMPTY   -11
#define MAX_HEAP_FULL    -12
#define MAX_HEAP_SUCCESS  0

struct max_heap_t {
    size_t heap_sz;
    int items_count; /* start from 0 */
    int *arr;
};

struct max_heap_t* 
max_heap_new(size_t sz)
{
    struct max_heap_t *p = NULL;

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
max_heap_insert(struct max_heap_t *hp, int item)
{
    if (hp->items_count == hp->heap_sz)
        return MAX_HEAP_FULL;

    /* first, place the item to the end
     * then, adjust the heap */
    int index = hp->items_count;
    hp->arr[index] = item;

    int parent_index;
    while (index > 0) {
        parent_index = (index - 1)/2;
        if (hp->arr[parent_index] < item) {
            swap(hp->arr, parent_index, index);
            index = parent_index;
        } else {
            break;
        }
    }

    hp->items_count++;
    return MAX_HEAP_SUCCESS;
}

int 
max_heap_remove_max(struct max_heap_t *hp)
{
    if (hp->items_count == 0)
        return MAX_HEAP_EMPTY;

    int max_item = hp->arr[0];
    hp->arr[0] = hp->arr[hp->items_count - 1];
    hp->items_count--;

    int right_child_index;
    int left_child_index;
    int cur_index = 0;
    int max_child_index;

    while (cur_index < hp->items_count) {
        left_child_index = (cur_index  * 2) + 1;
        right_child_index = (cur_index  * 2) + 2;

        if (left_child_index >= hp->items_count) break;
        if (right_child_index >= hp->items_count) {
            max_child_index = right_child_index;
        } else if (hp->arr[right_child_index] > hp->arr[left_child_index]) {
            max_child_index = right_child_index;
        } else {
            max_child_index = left_child_index;
        }

        if (hp->arr[cur_index] < hp->arr[max_child_index]) {
            swap(hp->arr, cur_index, max_child_index);
            cur_index = max_child_index;
        } else {
            break;
        }
    }

    return max_item;
}

int 
max_heap_items_count(const struct max_heap_t *hp)
{
    return hp->items_count;
}

void
max_heap_free(struct max_heap_t **php)
{
    struct max_heap_t *hp = (*php);
    free(hp->arr);
    free(hp);
    (*php) = NULL;
}

void 
max_heap_dump(const struct max_heap_t *hp)
{
    int i;
    printf("max heap dump ==>\n");
    for (i = 0; i < hp->items_count; i++) {
        printf("%d ", hp->arr[i]);
    }
    printf("\n");
    printf("<=== max heap dump end\n");
}

int
main()
{
    int arr[] = {3, 8, 2, 1, 5, 4, 6, 7};
    size_t sz = sizeof(arr)/sizeof(arr[0]);
    output_arr(arr, sizeof(arr)/sizeof(arr[0]));

    struct max_heap_t *hp = max_heap_new(sz);
    if (!hp) {
        printf("max_heap_new error\n");
        return -1;
    }

    int i;
    for (i = 0; i < sz; i++)
        max_heap_insert(hp, arr[i]);
    max_heap_dump(hp); /*  8 7 6 5 3 2 4 1 */

    int new_arr[8] = {0};
    for (i = 0; i < sz; i++) {
        new_arr[i] = max_heap_remove_max(hp);
    }
    
    output_arr(new_arr, sizeof(new_arr)/sizeof(new_arr[0]));
    max_heap_free(&hp);
}
