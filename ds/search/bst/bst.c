/*
 * bst.c
 *
 * an implmentation of binary search tree 
 */

#include <stdio.h>
#include <stdlib.h>
#include <string.h>
#include "bst.h"
#include "queue.h"

struct bst_node_t {
    struct bst_node_t *pn; /*parent node */
    struct bst_node_t *ln; /*left node */
    struct bst_node_t *rn; /*right node */
    int value;
};

struct bst_t {
    struct bst_node_t *root; /* root node */
};

int 
is_bst_empty(const struct bst_t *t)
{
    return (t->root == NULL);
}

static struct bst_node_t*
search_node(const struct bst_t *t, int v) 
{
    struct bst_node_t *node = t->root;
    
    while (node != NULL) {
        if (v < node->value) {
            node = node->ln;
        } else if (v > node->value) {
            node = node->rn;
        } else {
            return node; /* found it */ 
        }
    }

    return NULL; /* not founded */
}

int 
bst_search_node(const struct bst_t *t, int v)
{
    return (search_node(t, v) != NULL);
}

int
bst_insert_node(struct bst_t *t, int v) 
{
    struct bst_node_t *node = NULL;
    node = malloc(sizeof(*node));
    if (!node) return -1;
    node->pn = NULL;    
    node->ln = NULL;    
    node->rn = NULL;    
    node->value = v;    

    /* case 1: empty tree, new node is treated as root node */
    if (is_bst_empty(t)) {
        t->root = node;
        return 0;
    }

    /* case 2: non-empty tree */
    struct bst_node_t *tn = t->root; /* temp node */
    
    while (1) {
        if (v < tn->value) {
            if (tn->ln) {
                tn = tn->ln;
            } else {
                tn->ln = node;
                node->pn = tn;
                return 0;
            }
        } else if (v > tn->value) {
            if (tn->rn) {
                tn = tn->rn;
            } else {
                tn->rn = node;
                node->pn = tn;
                return 0;
            }
        } else { /* equal */
            free(node);
            return -2; /* exist */
        }
    }

    return 0;
}

static struct bst_node_t*
successor(const struct bst_node_t *node)
{
    if (node->rn == NULL) return NULL;

    node = node->rn;
    while (node->ln != NULL) {
        node = node->ln;
    }
    return (struct bst_node_t*)node;
}

int
bst_remove_node(struct bst_t *t, int v) 
{
    struct bst_node_t *node, *pn = NULL;

    node = search_node(t, v);
    if (node == NULL)
        return -1; /* not found the value */

    /*
     * leaf node case
     */
    if ((node->ln == NULL) && (node->rn == NULL)) {
        if (node->pn == NULL) { /* root node */
            free(t->root);
            t->root = NULL;
        } else {
            pn = node->pn;
            if (node->value < pn->value) {
                pn->ln = NULL;
            } else {
                pn->rn = NULL;
            }
            free(node);
        }
        return 0;
    }

    /*
     * both left and right subtree are not empty
     */
    if ((node->ln != NULL) && (node->rn != NULL)) {
        struct bst_node_t *sn = successor(node);
        if (sn->value < sn->pn->value) {
            sn->pn->ln = NULL; /* remove the successor node */
        } else {
            sn->pn->rn = NULL; /* remove the successor node */
        }
        node->value = sn->value;
        free(sn);
        return 0;
    }

    /*
     * either left subtree or right subtree
     */
    struct bst_node_t *sub_n = NULL;
    if (node->ln != NULL) sub_n = node->ln;
    if (node->rn != NULL) sub_n = node->rn;

    if (node->value < node->pn->value) { /* node is left subtree of parent node */
        node->pn->ln = sub_n;
    } else {
        node->pn->rn = sub_n;
    }
    sub_n->pn = node->pn;
    free(node);
    return 0;
}

struct bst_t*
bst_new()
{
    struct bst_t *t = NULL;

    t = malloc(sizeof(*t));
    if (t == NULL)
        return NULL;

    memset(t, 0, sizeof(*t));
    return t;
}

static void
free_node(struct bst_node_t *node)
{
    if (node == NULL)
        return;

    if (node->ln != NULL)
        free_node(node->ln);

    if (node->rn != NULL)
        free_node(node->rn);

    free(node);
}

void 
bst_free(struct bst_t **t)
{
    free_node((*t)->root);
    free(*t);
    (*t) = NULL;
}

static int 
bst_node_height(const struct bst_node_t *node) 
{
    if (node == NULL)
        return 0;

    int h1 = bst_node_height((const struct bst_node_t*)node->ln);
    int h2 = bst_node_height((const struct bst_node_t*)node->rn);
    return 1 + ((h1 >= h2) ? h1 : h2);
}

int
bst_height(const struct bst_t *t)
{
    return bst_node_height((const struct bst_node_t*)t->root);
}

/*
 * according to level order
 * output nil for empty subtree node 
 */
void 
bst_levelorder_traverse(const struct bst_t *t) 
{
    struct queue_t *q1 = queue_new();
    if (q1 == NULL) return;
    struct queue_t *q2 = queue_new();
    if (q2 == NULL) {
        queue_free(&q1); 
        return;
    }

    struct bst_node_t *tn = NULL;

    tn = t->root;
    if (tn == NULL) return;

    enqueue(q1, tn);
    while(!is_queue_empty(q1)) {
        while (!is_queue_empty(q1)) {
            tn = dequeue(q1);
            enqueue(q2, tn);
        }
        while (!is_queue_empty(q2)) {
            tn = dequeue(q2);

            if (tn->pn != NULL)
                printf("%d(%d) ", tn->value, tn->pn->value);
            else
                printf("%d ", tn->value);
                
            
            if (tn->ln != NULL) {
                enqueue(q1, tn->ln);
            }

            if (tn->rn != NULL) {
                enqueue(q1, tn->rn);
            }
        }
        printf("\n");
    }

    queue_free(&q1);
    queue_free(&q2);
}

static void inorder_traverse_node(const struct bst_node_t *nd) 
{
    printf("%d ", nd->value);
    if (nd->ln != NULL) inorder_traverse_node(nd->ln);
    if (nd->rn != NULL) inorder_traverse_node(nd->rn);
}

void 
bst_inorder_traverse(const struct bst_t *t) 
{
    if (t->root != NULL) 
        inorder_traverse_node(t->root);

    printf("\n");

}

int 
main()
{
    struct bst_t *t = NULL;
    t = bst_new();
    if (!t) {
        printf("create bst tree error\n");
        return -1;
    }

    int arr[] = {8, 3, 10, 1, 6, 14, 4, 7, 13};
    int retv, i = 0;
    for (i = 0; i < sizeof(arr)/sizeof(arr[0]); i++) {
        if ((retv = bst_insert_node(t, arr[i])) != 0) {
            printf("err insert %d, err = %d\n", arr[i], retv);
            return -1;
        }
    }

    bst_inorder_traverse(t);
    bst_levelorder_traverse(t);
    printf("search 1 = %d\n", bst_search_node(t, 1));
    printf("search 14 = %d\n", bst_search_node(t, 14));
    printf("search 8 = %d\n", bst_search_node(t, 8));
    printf("search 23 = %d\n", bst_search_node(t, 23));

    bst_remove_node(t, 1);
    bst_levelorder_traverse(t);
    bst_insert_node(t, 1);
    bst_levelorder_traverse(t);
    bst_remove_node(t, 6);
    bst_levelorder_traverse(t);
    bst_remove_node(t, 14);
    bst_levelorder_traverse(t);
    bst_remove_node(t, 10);
    bst_levelorder_traverse(t);

    bst_free(&t);
    return 0;
}
