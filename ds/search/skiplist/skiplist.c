/*
 * skiplist.c
 *
 * brief description of this file
 */

#include "utils.h"
#include <stdio.h>
#include <stdlib.h>
#include <string.h>

#define SKIPLIST_MAXLEVEL 8

#define SL_SUCCESS 0

struct skiplist_node_t {
    int value;
    struct skiplist_node_t* levels[];
};

struct skiplist_t {
    int count;
    struct skiplist_node_t *sentinel; 
};

/* source from redis */
static int 
rand_level(void) {
    int level = 1;
    while ((rand() & 0xFFFF) < (0.5 * 0xFFFF))
        level += 1;
    return (level < SKIPLIST_MAXLEVEL) ? level : SKIPLIST_MAXLEVEL;
}

static struct skiplist_node_t*
create_node(int level, int value)
{
    struct skiplist_node_t *p = NULL;
    p = malloc(sizeof(*p) + level * sizeof(struct skiplist_node_t*));
    if (!p)
        return NULL;

    memset(p, 0, (sizeof(*p) + level * sizeof(struct skiplist_node_t*)));
    p->value = value;

    return p;
}

static void
free_node(struct skiplist_node_t **nd) 
{
    free(*nd);
    (*nd) = NULL;
}

struct skiplist_t* 
skiplist_new()
{
    struct skiplist_t *p = NULL;

    p = malloc(sizeof(*p));
    if (!p) 
        return NULL;

    p->count = 0;
    p->sentinel = create_node(SKIPLIST_MAXLEVEL, 0);
    if (!p->sentinel) {
        free(p);
        return NULL;
    }

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
skiplist_search(const struct skiplist_t *sl, const int item)
{
    return SL_SUCCESS;
}

int 
skiplist_items_count(const struct skiplist_t *sl)
{
    return sl->count;
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

    int ret;
    ret = skiplist_search(sl, 5);
    if (ret == SL_SUCCESS)
        printf("found %d\n", 5);

    ret = skiplist_search(sl, 10);
    if (ret == SL_SUCCESS)
        printf("found %d\n", 10);
    else
        printf("do not found %d\n", 10);

    ret = skiplist_search(sl, 3);
    if (ret == SL_SUCCESS)
        printf("found %d\n", 3);

    ret = skiplist_search(sl, 20);
    if (ret == SL_SUCCESS)
        printf("found %d\n", 20);
    else
        printf("do not found %d\n", 20);

    skiplist_free(&sl);
}
