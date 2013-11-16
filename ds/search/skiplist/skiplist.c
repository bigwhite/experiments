/*
 * skiplist.c
 *
 * brief description of this file
 */

#include "utils.h"
#include <stdio.h>
#include <stdlib.h>
#include <string.h>

#define SL_SUCCESS 0

struct skiplist_t {

};

struct skiplist_t* 
skiplist_new(size_t sz)
{
    struct skiplist_t *p = NULL;

    p = malloc(sizeof(*p));
    if (!p) 
        return NULL;

    /* TODO */

    return p;
}

int 
skiplist_insert(struct skiplist_t *sl, int item)
{
    return SL_SUCCESS;
}

int 
skiplist_remove(struct skiplist_t *sl, int item)
{
    return SL_SUCCESS;
}

int 
skiplist_items_count(const struct skiplist_t *sl)
{
    return 0;
}

void
skiplist_free(struct skiplist_t **psl)
{
    struct skiplist_t *sl = (*psl);
    free(sl);
    (*psl) = NULL;
}

void 
skiplist_dump(const struct skiplist_t *sl)
{
    int i;
    printf("skip list dump ==>\n");

    printf("<=== skip list dump end\n");
}

int
main()
{
    int arr[] = {3, 8, 2, 1, 5, 4, 6, 7};
    size_t sz = sizeof(arr)/sizeof(arr[0]);
    output_arr(arr, sizeof(arr)/sizeof(arr[0]));

    struct skiplist_t *sl = skiplist_new(sz);
    if (!sl) {
        printf("skiplist_new error\n");
        return -1;
    }

    int i;
    for (i = 0; i < sz; i++)
        skiplist_insert(sl, arr[i]);
    skiplist_dump(sl);

    skiplist_free(&sl);
}
