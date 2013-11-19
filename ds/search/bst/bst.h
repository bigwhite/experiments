/*
 * bst.h
 *
 * 
 */

#ifndef _BST_H_
#define _BST_H_

#ifdef __cplusplus
extern "C" {
#endif

struct bst_tree_t;

struct bst_tree_t* bst_tree_new();

void bst_tree_free(struct bst_tree_t **t);

int bst_tree_search_node(const struct bst_tree_t *t, int v);

int bst_tree_insert_node(struct bst_tree_t *t, int v);

int bst_tree_delete_node(struct bst_tree_t *t, int v);

int is_bst_tree_empty(const struct bst_tree_t *t);

void bst_tree_levelorder_traverse(const struct bst_tree_t *t);

void bst_tree_inorder_traverse(const struct bst_tree_t *t);

int bst_tree_height(const struct bst_tree_t *t);


#ifdef __cplusplus
}
#endif

#endif /* _BST_H_ */
