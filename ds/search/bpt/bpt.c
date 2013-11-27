/*
 * bpt.c
 *
 * an implementation of b+ tree - b plus tree
 */

#include <stdio.h>
#include <stdlib.h>
#include <stdint.h>

#include "bpt.h"

#define BPT_SUCCESS 0

/*
 



 */

struct record_t {
    int value;
};

struct bpt_node_t {
    struct bpt_node_t *parent;
    int *keys;
    int num_keys; /* for leaf, nodes count = num_keys; 
                     for internal node, nodes count = num_keys + 1 */
    void **nodes; /* for leaf, the nodes points to records,
                     for internal node, nodes points to sub nodes */
    uint8_t is_leaf;
};

struct bpt_t {
    int order;
    struct bpt_node_t *root;
};

struct bpt_t* 
bpt_new(int order)
{

    return NULL;
}

void 
bpt_free(struct bpt_t **t)
{

}

int 
bpt_search_node(const struct bpt_t *t, int key)
{
    return BPT_SUCCESS;
}

int 
bpt_insert_node(struct bpt_t *t, int v)
{
    return BPT_SUCCESS;
}

int 
bpt_remove_node(struct bpt_t *t, int v)
{
    return BPT_SUCCESS;
}

int
main()
{
    struct bpt_t *t;

    t = bpt_new(4);
    if (!t) 
        return -1;
    printf("bpt new ok\n");

    int ret;
    
    ret = bpt_insert_node(t, 1);
    if (ret != BPT_SUCCESS) {
        printf("insert %d error\n", 1);
        return -1;
    }
    printf("insert %d error\n", 1);

    ret = bpt_search_node(t, 1);
    printf("search %d = %d\n", 1, ret);
    
    bpt_remove_node(t, 1);

    bpt_free(&t);
}
