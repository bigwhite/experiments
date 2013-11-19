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
#define SL_FAIL    1
#define SL_NOTFOUND 2

struct skiplist_node_t {
    int value;
    int level;
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
    while ((rand() & 0xFFFF) < (0.5 * 0xFFFF)) /* simulating the coin tossing */
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
    p->level = level;

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
    int level = rand_level(); /* level starts from 1 to MAX_LEVEL */

    struct skiplist_node_t *new_node = NULL;
    new_node = create_node(level, item);
    if (!new_node)
        return SL_FAIL;

    struct skiplist_node_t *cur, *next;
    int i;

    for (i = level - 1; i >= 0; i--) {
        cur = sl->sentinel;
        next = cur->levels[i];
        while (next != NULL) {
            if (next->value < item) {
                cur = next;
                next = cur->levels[i];
            } else {
                break;
            }
        }

        new_node->levels[i] = cur->levels[i];
        cur->levels[i] = new_node;
    }

    sl->count++;
    return SL_SUCCESS;
}

int
skiplist_remove(struct skiplist_t *sl, int item)
{
    struct skiplist_node_t *pre_nodes[SKIPLIST_MAXLEVEL];
    struct skiplist_node_t *cur, *next, *node = NULL;
    int level;

    for (level = SKIPLIST_MAXLEVEL - 1; level >= 0; level--) {
        pre_nodes[level] = NULL;

        cur = sl->sentinel;
        next = cur->levels[level];
        while (next != NULL && next->value < item) {
            cur = next;
            next = cur->levels[level];
        }
        pre_nodes[level] = cur;
    }

    if (pre_nodes[0] != NULL) {
        node = pre_nodes[0]->levels[0];
        if (node != NULL && node->value == item) {
            for (level = node->level - 1; level >= 0; level--) {
                if (pre_nodes[level] != NULL) {
                    pre_nodes[level]->levels[level] = node->levels[level];
                }
            }

            free_node(&node);
            sl->count--;
            return SL_SUCCESS;
        }
    }

    return SL_NOTFOUND;
}

int
skiplist_search(const struct skiplist_t *sl, int item)
{
    struct skiplist_node_t *cur, *next;
    int found = 0;

    cur = sl->sentinel;
    int level = SKIPLIST_MAXLEVEL - 1;

    while (level >= 0) {
        next = cur->levels[level];

        if (next != NULL) {
            if (next->value == item) {
                found = 1;
                break;
            } else if (next->value < item) {
                cur = next;
            } else {
                level--;
            }
        } else  {
            level--;
        }
    }

    if (found)
        return SL_SUCCESS;
    else
        return SL_NOTFOUND;
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
    struct skiplist_node_t *node = NULL;
    struct skiplist_node_t *next = NULL;

    node = sl->sentinel;
    while (node != NULL) {
        next = node->levels[0];
        free(node);
        node = next;
    };

    free(sl);
    (*psl) = NULL;
}

void
skiplist_dump(const struct skiplist_t *sl)
{
    int i, j;
    struct skiplist_node_t *next = NULL;

    printf("skip list dump ==>\n");

    for (i = SKIPLIST_MAXLEVEL - 1; i >= 0; i--) {
        next = sl->sentinel->levels[0];
        if (next != NULL) {
            for (j = 0; j < sl->count; j++) {
                if (i <= next->level - 1) {
                    printf("[%d ]  ", next->value);
                } else {
                    //printf("[ ] - ");
                    printf("      ");
                }
                next = next->levels[0];
            }
        }
        printf("\n");
    }

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
    printf("skiplist_new ok\n");

    int i;
    for (i = 0; i < sz; i++) {
        skiplist_insert(sl, arr[i]);
        printf("sl insert [%d] ok\n", arr[i]);
    }
    skiplist_dump(sl);

    int ret;
    ret = skiplist_search(sl, 1);
    if (ret == SL_SUCCESS)
        printf("found %d\n", 1);

    ret = skiplist_search(sl, 8);
    if (ret == SL_SUCCESS)
        printf("found %d\n", 8);

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

    ret = skiplist_remove(sl, 1);
    if (ret == SL_SUCCESS)
        printf("delete %d\n", 1);
    else
        printf("do not delete %d\n", 1);
    skiplist_dump(sl);

    ret = skiplist_remove(sl, 8);
    if (ret == SL_SUCCESS)
        printf("delete %d\n", 8);
    else
        printf("do not delete %d\n", 8);
    skiplist_dump(sl);

    ret = skiplist_remove(sl, 4);
    if (ret == SL_SUCCESS)
        printf("delete %d\n", 4);
    else
        printf("do not delete %d\n", 4);
    skiplist_dump(sl);

    ret = skiplist_remove(sl, 20);
    if (ret == SL_SUCCESS)
        printf("delete %d\n", 20);
    else
        printf("do not delete %d\n", 20);


    skiplist_free(&sl);
}
