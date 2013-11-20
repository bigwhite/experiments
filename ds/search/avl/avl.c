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

#define AVL_TREE_SUCCESS 0

struct avl_tree_node_t {
    struct avl_tree_node_t *parent; 
    struct avl_tree_node_t *left; 
    struct avl_tree_node_t *right;
    int value;
};

struct avl_tree_t {
    struct avl_tree_node_t *root;
};

/* rotation type */
enum {
    LL_ROTATE,
    LR_ROTATE,
    RL_ROTATE,
    RR_ROTATE 
};

static struct avl_tree_node_t* successor(struct avl_tree_node_t *node);
static int is_avl_empty(const struct avl_tree_t *t);
static void free_node(struct avl_tree_node_t *node);
static struct avl_tree_node_t* search_node(const struct avl_tree_t *t, int v);
static struct avl_tree_node_t* insert_node(struct avl_tree_t *t, int v);
static void balance(struct avl_tree_t *t, struct avl_tree_node_t *node, int trace_ancestor);
static void left_rotate(struct avl_tree_t *t, struct avl_tree_node_t *subtree_root);
static void right_rotate(struct avl_tree_t *t, struct avl_tree_node_t *subtree_root);
static void left_right_rotate(struct avl_tree_t *t, struct avl_tree_node_t *subtree_root);
static void right_left_rotate(struct avl_tree_t *t, struct avl_tree_node_t *subtree_root);
static void inorder_traverse_node(const struct avl_tree_node_t *nd);
static int node_height(const struct avl_tree_node_t *node);
static __inline__ int balance_factor(const struct avl_tree_node_t *node);
static struct avl_tree_node_t* is_balanced(struct avl_tree_node_t *node);
void remove_node(struct avl_tree_t *t, struct avl_tree_node_t *node);
static int recognize_rotate_type(struct avl_tree_node_t *least_unbalanced_subtree_root);


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
    struct avl_tree_node_t *node = insert_node(t, v);

    if (node == NULL) return -1;
    balance(t, node, 0);
    return AVL_TREE_SUCCESS;
}

int 
avl_tree_remove_node(struct avl_tree_t *t, int v)
{
    struct avl_tree_node_t *node;
    struct avl_tree_node_t *parent;

    node = search_node(t, v);
    if (node == NULL)
        return -1; /* not found the value */

    parent = node->parent;
    remove_node(t, node);
    balance(t, parent, 1);

    return 0;
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
    if (tn == NULL) return;

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

static __inline__ int 
balance_factor(const struct avl_tree_node_t *node)
{
    return node_height(node->left) - node_height(node->right);
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
    node->value = v;

    /* case 1: empty tree, new node is treated as root node */
    if (is_avl_empty(t)) {
        t->root = node;
        return node;
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
                return node;
            }
        } else if (v > tn->value) {
            if (tn->right) {
                tn = tn->right;
            } else {
                tn->right = node;
                node->parent = tn;
                return node;
            }
        } else { /* equal */
            free(node);
            return NULL; /* exist */
        }
    }
}

/*

   P is new insert node, right_rotate(A) is like this:
   A is the least_unbalanced_tree's root

            A                   B
           / \                 / \
          /   \               /   \
         B     C   =>        D     A
        / \                 /     / \
       /   \               /     /   \
      D    E              P      E   C
     /
    /
    P

            A                    B
           / \                 /  \
          /   \               /    \
         B     C   =>        D       A
        / \                   \     / \
       /   \                   \   /   \
      D    E                   P   E   C
       \
        \
        P
*/

/* 
 * subtree_root is the root node of the least unbalanced subtree
 */
static void 
right_rotate(struct avl_tree_t* t, struct avl_tree_node_t *subtree_root)
{
    struct avl_tree_node_t *old_subtree_root = subtree_root;
    struct avl_tree_node_t *new_subtree_root = subtree_root->left;

    if (old_subtree_root->parent != NULL) {
        if (old_subtree_root->value > old_subtree_root->parent->value) {
            old_subtree_root->parent->right = new_subtree_root;
        } else {
            old_subtree_root->parent->left = new_subtree_root;
        }
        new_subtree_root->parent = old_subtree_root->parent;
    } else {
        /* subtree_root is root of this avl tree */
        t->root = new_subtree_root;
        new_subtree_root->parent = NULL;
    }

    old_subtree_root->left = new_subtree_root->right;
    if (new_subtree_root->right != NULL) 
        new_subtree_root->right->parent = old_subtree_root;
    new_subtree_root->right = old_subtree_root;
    old_subtree_root->parent = new_subtree_root;
}

/*

   P is new insert node, left_rotate(A) is like this:
   A is the least_unbalanced_tree's root

            A                           C 
           / \                         / \
          /   \                       /   \
         B     C        =>           A     E
              / \                   / \     \
             /   \                 /   \     \
            D     E               B     D     P
                   \
                    \ 
                     P

            A                           C 
           / \                         / \
          /   \                       /   \
         B     C        =>           A     E
              / \                   / \    / 
             /   \                 /   \  /
            D     E               B     D P
                 /
                / 
               P
*/
/* 
 * subtree_root is the root node of the least unbalanced subtree
 */
static void 
left_rotate(struct avl_tree_t *t, struct avl_tree_node_t *subtree_root)
{
    struct avl_tree_node_t *old_subtree_root = subtree_root;
    struct avl_tree_node_t *new_subtree_root = subtree_root->right;

    if (old_subtree_root->parent != NULL) {
        if (old_subtree_root->value > old_subtree_root->parent->value) {
            old_subtree_root->parent->right = new_subtree_root;
        } else {
            old_subtree_root->parent->left = new_subtree_root;
        }
        new_subtree_root->parent = old_subtree_root->parent;
    } else {
        /* subtree_root is root of this avl tree */
        t->root = new_subtree_root;
        new_subtree_root->parent = NULL;
    }

    old_subtree_root->right = new_subtree_root->left;
    if (new_subtree_root->left != NULL) 
        new_subtree_root->left->parent = old_subtree_root;
    new_subtree_root->left = old_subtree_root;
    old_subtree_root->parent = new_subtree_root;
}

/*

   P is new insert node, right_left_rotate(A) is like this:
   A is the least_unbalanced_tree's root

     A                           A                              D
    / \                         / \                            / \
   /   \   right_rotate(C)     /   \    left_rotate(A)        /   \
  B     C        =>           B     D     =>                 A     C
       / \                           \                      /     / \
      /   \                           \                    /     /   \
     D     E                           C                  B     P    E
     \                                / \
      \                              /   \
       P                             P   E

     A                           A                              D
    / \                         / \                            / \
   /   \   right_rotate(C)     /   \    left_rotate(A)        /   \
  B     C        =>           B     D     =>                 A     C
       / \                         / \                      / \     \
      /   \                       /   \                    /   \     \
     D     E                     P     C                  B     P     E
    /                                   \
   /                                     \
  P                                       E

*/  
static void 
right_left_rotate(struct avl_tree_t *t, struct avl_tree_node_t *subtree_root) 
{
    right_rotate(t, subtree_root->right);
    left_rotate(t, subtree_root);
}

/*
   P is new insert node, left_right_rotate(A) is like this:
   A is the least_unbalanced_tree's root

         A                          A                           E 
        / \                        / \                         / \
       /   \   left_rotate(B)     /   \ right_rotate(A)       /   \
      B     C        =>          E     C     =>              B     A     
     / \                        /                           / \     \
    /   \                      /                           /   \     \
    D    E                    B                           D    P      C
        /                    / \       
       /                    /   \
      P                    D    P

         A                          A                           E 
        / \                        / \                         / \
       /   \   left_rotate(B)     /   \ right_rotate(A)       /   \
      B     C        =>          E     C     =>              B     A     
     / \                        / \                          /    / \
    /   \                      /   \                        /    /   \
    D    E                    B     P                      D     P    C
          \                  /        
           \                /   
            P              D   
*/

static void 
left_right_rotate(struct avl_tree_t *t, struct avl_tree_node_t *subtree_root)
{
    left_rotate(t, subtree_root->left);
    right_rotate(t, subtree_root);
}


/*
 * if unbalanced , return the root node of least unbalanced tree
 * else return NULL;
 */
static struct avl_tree_node_t*
is_balanced(struct avl_tree_node_t *node)
{
    if (!node)
        return NULL;

    struct avl_tree_node_t *p = node;
    int factor;

    /* 
     * check whether it is balanced 
     */
    while (p) {
        factor = balance_factor(p);
        if (abs(factor) > 1) {
            return p;
        }
        p = p->parent;
    }

    return NULL;
}

static int
recognize_rotate_type(struct avl_tree_node_t *least_unbalanced_subtree_root)
{
    struct avl_tree_node_t *node = least_unbalanced_subtree_root;
    int factor = balance_factor(node);

    if (factor < -1) {
        if (node->right != NULL && balance_factor(node->right) < 0) {
            return LL_ROTATE;
        } else {
            return RL_ROTATE;
        }
    } else if (factor > 1) {
        if (node->left != NULL && balance_factor(node->left) > 0) {
            return RR_ROTATE;
        } else {
            return LR_ROTATE;
        }
    }
}

/*
 *   
 * calculate the balance factor of ancestors of the insert node
 * if unbalance, rebalance the avl tree
 *
 * node - new insert node
 * trace_ancestor: if 1, we go on tracing ancestors' balance factor
 *                 otherwise, we just balance the first least unbalanced
 *                 subtree
 */
static void 
balance(struct avl_tree_t *t, struct avl_tree_node_t *node, int trace_ancestor)
{
    struct avl_tree_node_t *least_unbalanced_subtree_root;
    struct avl_tree_node_t *start_node = node;
    struct avl_tree_node_t *parent;

    do {
        least_unbalanced_subtree_root = is_balanced(start_node);
        if (!least_unbalanced_subtree_root) {
            /* no balance occurs */
            return;
        }
        parent = least_unbalanced_subtree_root->parent;

        int type = recognize_rotate_type(least_unbalanced_subtree_root);
        switch (type) {
            case LL_ROTATE:
                left_rotate(t, least_unbalanced_subtree_root);
                //printf("left rotate %d\n", least_unbalanced_subtree_root->value);
                break;

            case LR_ROTATE:
                left_right_rotate(t, least_unbalanced_subtree_root);
                //printf("left right rotate %d\n", least_unbalanced_subtree_root->value);
                break;

            case RL_ROTATE:
                right_left_rotate(t, least_unbalanced_subtree_root);
                //printf("right left rotate %d\n", least_unbalanced_subtree_root->value);
                break;

            case RR_ROTATE:
                right_rotate(t, least_unbalanced_subtree_root);
                //printf("right rotate %d\n", least_unbalanced_subtree_root->value);
                break;
            default:
                break;
        }
        start_node = parent;
    } while (trace_ancestor);

    return;
}

static struct avl_tree_node_t*
successor(struct avl_tree_node_t *node)
{
    if (node->right == NULL) return NULL;

    node = node->right;
    while (node->left != NULL) {
        node = node->left;
    }
    return node;
}

void 
remove_node(struct avl_tree_t *t, struct avl_tree_node_t *node)
{
    struct avl_tree_node_t *parent = NULL;
    int likely_unbalanced;

    /*
     * leaf node case
     */
    if ((node->left == NULL) && (node->right == NULL)) {
        if (node->parent == NULL) { /* root node */
            free(t->root);
            t->root = NULL;
        } else {
            parent = node->parent;
            if (node->value < parent->value) {
                parent->left = NULL;
            } else {
                parent->right = NULL;
            }
            free(node);
        }
        return;
    }

    /*
     * both left and right subtree are not empty
     */
    if ((node->left != NULL) && (node->right != NULL)) {
        struct avl_tree_node_t *sn = successor(node);
        parent = node->parent;
        if (sn->value < sn->parent->value) {
            sn->parent->left = NULL; /* remove the successor node */
        } else {
            sn->parent->right = NULL; /* remove the successor node */
        }
        node->value = sn->value;
        free(sn);
        return;
    }

    /*
     * either left subtree or right subtree
     */
    struct avl_tree_node_t *sub_n = NULL;
    if (node->left != NULL) sub_n = node->left;
    if (node->right != NULL) sub_n = node->right;
    parent = node->parent;

    if (node->value < node->parent->value) { /* node is left subtree of parent node */
        node->parent->left = sub_n;
    } else {
        node->parent->right = sub_n;
    }
    sub_n->parent = node->parent;

    free(node);
    return;
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
    printf("create avl tree ok\n");

    int arr[] = {15, 10, 4, 20, 30, 16, 13, 11, 12}; /* rr->ll->rl->lr */
    int retv, i = 0;
    for (i = 0; i < sizeof(arr)/sizeof(arr[0]); i++) {
        if ((retv = avl_tree_insert_node(t, arr[i])) != 0) {
            printf("err insert %d, err = %d\n", arr[i], retv);
            return -1;
        }
        printf("insert %d ok\n", arr[i]);
    }
    avl_tree_levelorder_traverse(t);

    printf("search 15 = %d\n", avl_tree_search_node(t, 15));
    printf("search 12 = %d\n", avl_tree_search_node(t, 12));
    printf("search 30 = %d\n", avl_tree_search_node(t, 30));
    printf("search 23 = %d\n", avl_tree_search_node(t, 23));

    printf("remove 4 = %d\n", avl_tree_remove_node(t, 4)); /* rl */
    avl_tree_levelorder_traverse(t);
    printf("remove 15 = %d\n", avl_tree_remove_node(t, 15)); /* remove root */
    avl_tree_levelorder_traverse(t); /* */
    printf("remove 30 = %d\n", avl_tree_remove_node(t, 30));  /* lr */
    avl_tree_levelorder_traverse(t);
    printf("remove 16 = %d\n", avl_tree_remove_node(t, 16));
    avl_tree_levelorder_traverse(t);
    printf("remove 20 = %d\n", avl_tree_remove_node(t, 20));
    avl_tree_levelorder_traverse(t);
    printf("remove 13 = %d\n", avl_tree_remove_node(t, 13));
    avl_tree_levelorder_traverse(t);
    printf("remove 11 = %d\n", avl_tree_remove_node(t, 11)); /* remove root */
    avl_tree_levelorder_traverse(t);
    printf("remove 10 = %d\n", avl_tree_remove_node(t, 10)); 
    avl_tree_levelorder_traverse(t);
    printf("remove 12 = %d\n", avl_tree_remove_node(t, 12));
    avl_tree_levelorder_traverse(t);

    for (i = 0; i < 1000000; i++) {
        avl_tree_insert_node(t, i);
        //printf("insert %d ok\n", i);
    }

    avl_tree_free(&t);
    return 0;
}
