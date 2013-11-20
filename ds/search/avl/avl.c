/*
 * avl.c
 *
 * an implemention of AVL tree
 */

#include <stdio.h>
#include <stdlib.h>
#include <string.h>
#include "avl.h"
#include "queue.h"

struct avl_tree_node_t {
    struct avl_tree_node_t *parent; 
    struct avl_tree_node_t *left; 
    struct avl_tree_node_t *right;
    int height; 
    int value;
};

struct avl_tree_t {
    struct avl_tree_node_t *root;
};

static int is_avl_empty(const struct avl_tree_t *t);
static void free_node(struct avl_tree_node_t *node);
static struct avl_tree_node_t* search_node(const struct avl_tree_t *t, int v);
static struct avl_tree_node_t* insert_node(struct avl_tree_t *t, int v);
static void balance();
static void left_rotate();
static void right_rotate();
static void left_right_rotate();
static void right_left_rotate();
static void inorder_traverse_node(const struct avl_tree_node_t *nd);
static int node_height(const struct avl_tree_node_t *node);
static void adjust_ancestor_height(const struct avl_tree_node_t *node);
static int balance_factor(const struct avl_tree_node_t *node);


struct avl_tree_t* 
avl_tree_new()
{
    struct avl_tree_t *t = NULL;

    t = malloc(sizeof(*t));
    if (t == NULL)
        return NULL;

    memset(t, 0, sizeof(*t));
    return t;
}

void 
avl_tree_free(struct avl_tree_t **t)
{
    free_node((*t)->root);
    free(*t);
    (*t) = NULL;
}

int 
avl_tree_search_node(const struct avl_tree_t *t, int v)
{
    return (search_node(t, v) != NULL);
}

int 
avl_tree_insert_node(struct avl_tree_t *t, int v)
{

}

int 
avl_tree_remove_node(struct avl_tree_t *t, int v)
{

}

void
avl_tree_levelorder_traverse(const struct avl_tree_t *t)
{
    struct queue_t *q1 = queue_new();
    if (q1 == NULL) return;
    struct queue_t *q2 = queue_new();
    if (q2 == NULL) {
        queue_free(&q1);
        return;
    }

    struct avl_tree_node_t *tn = NULL;

    tn = t->root;
    enqueue(q1, tn);
    while (!is_queue_empty(q1)) {
        while (!is_queue_empty(q1)) {
            tn = dequeue(q1);
            enqueue(q2, tn);
        }
        while (!is_queue_empty(q2)) {
            tn = dequeue(q2);

            if (tn->parent != NULL)
                printf("%d(%d) ", tn->value, tn->parent->value);
            else
                printf("%d ", tn->value);


            if (tn->left != NULL) {
                enqueue(q1, tn->left);
            }

            if (tn->right != NULL) {
                enqueue(q1, tn->right);
            }
        }
        printf("\n");
    }

    queue_free(&q1);
    queue_free(&q2);
}

void
avl_tree_inorder_traverse(const struct avl_tree_t *t)
{
    if (t->root != NULL)
        inorder_traverse_node(t->root);

    printf("\n");

}

static void 
inorder_traverse_node(const struct avl_tree_node_t *nd)
{
    printf("%d ", nd->value);
    if (nd->left != NULL) inorder_traverse_node(nd->left);
    if (nd->right != NULL) inorder_traverse_node(nd->right);
}

static int
is_avl_empty(const struct avl_tree_t *t) 
{
    return (t->root == NULL);
}

static void 
free_node(struct avl_tree_node_t *node)
{
    if (node == NULL)
        return;

    if (node->left != NULL)
        free_node(node->left);

    if (node->right != NULL)
        free_node(node->right);

    free(node);
}

static struct avl_tree_node_t*
search_node(const struct avl_tree_t *t, int v)
{
    struct avl_tree_node_t *node = t->root;

    while (node != NULL) {
        if (v < node->value) {
            node = node->left;
        } else if (v > node->value) {
            node = node->right;
        } else {
            return node; /* found it */
        }
    }

    return NULL; /* not founded */
}

static int 
node_height(const struct avl_tree_node_t *node) 
{
    /* 
     * caculate the height of the tree 
     * which view this node as root
     *
     * tree height starts from 1
     */ 
    if (node == NULL) return 0;

    int h1 = node_height(node->left);
    int h2 = node_height(node->right);

    return 1 + ((h1 >= h2) ? h1 : h2);
}

static void 
adjust_ancestor_height(const struct avl_tree_node_t *node)
{
    /*
     * recaculate the ancestor nodes' height 
     * of this node 
     */ 

    struct avl_tree_node_t *parent = node->parent;
    int height;

    while (parent != NULL) {
        height = node_height(parent);
        if (height == parent->height) {
            return;
        } 
        parent = parent->parent;
    }
}

static int 
balance_factor(const struct avl_tree_node_t *node)
{
    int left_subtree_height = 0;
    int right_subtree_height = 0;

    if (node->left != NULL)
        left_subtree_height = node->left->height;
    if (node->right != NULL)
        right_subtree_height = node->right->height;

    return left_subtree_height - right_subtree_height;
}

static struct avl_tree_node_t* 
insert_node(struct avl_tree_t *t, int v)
{
    struct avl_tree_node_t *node = NULL;
    node = malloc(sizeof(*node));
    if (!node) return NULL;
    node->parent = NULL;
    node->left = NULL;
    node->right = NULL;
    node->height = 1;
    node->value = v;

    /* case 1: empty tree, new node is treated as root node */
    if (is_avl_empty(t)) {
        t->root = node;
        return 0;
    }

    /* case 2: non-empty tree */
    struct avl_tree_node_t *tn = t->root; /* temp node */

    while (1) {
        if (v < tn->value) {
            if (tn->left) {
                tn = tn->left;
            } else {
                tn->left = node;
                node->parent = tn;
                adjust_ancestor_height(node);
                return node;
            }
        } else if (v > tn->value) {
            if (tn->right) {
                tn = tn->right;
            } else {
                tn->right = node;
                node->parent = tn;
                adjust_ancestor_height(node);
                return node;
            }
        } else { /* equal */
            free(node);
            return NULL; /* exist */
        }
    }

    return 0;
}

int 
main()
{
    struct avl_tree_t *t = NULL;
    t = avl_tree_new();
    if (!t) {
        printf("create avl tree error\n");
        return -1;
    }

    int arr[] = {8, 3, 10, 1, 6, 14, 4, 7, 13};
    int retv, i = 0;
    for (i = 0; i < sizeof(arr)/sizeof(arr[0]); i++) {
        if ((retv = avl_tree_insert_node(t, arr[i])) != 0) {
            printf("err insert %d, err = %d\n", arr[i], retv);
            return -1;
        }
    }

    avl_tree_inorder_traverse(t);
    avl_tree_levelorder_traverse(t);
    printf("search 1 = %d\n", avl_tree_search_node(t, 1));
    printf("search 14 = %d\n", avl_tree_search_node(t, 14));
    printf("search 8 = %d\n", avl_tree_search_node(t, 8));
    printf("search 23 = %d\n", avl_tree_search_node(t, 23));

    avl_tree_remove_node(t, 1);
    avl_tree_levelorder_traverse(t);
    avl_tree_insert_node(t, 1);
    avl_tree_levelorder_traverse(t);
    avl_tree_remove_node(t, 6);
    avl_tree_levelorder_traverse(t);
    avl_tree_remove_node(t, 14);
    avl_tree_levelorder_traverse(t);
    avl_tree_remove_node(t, 10);
    avl_tree_levelorder_traverse(t);

    avl_tree_free(&t);

    return 0;
}
