/*
 * avl.h
 *
 */

#ifndef _AVL_H_
#define _AVL_H_

#ifdef __cplusplus
extern "C" {
#endif

struct avl_tree_t;

struct avl_tree_t* avl_tree_new();

void avl_tree_free(struct avl_tree_t **t);

int avl_tree_search_node(const struct avl_tree_t *t, int v);

int avl_tree_insert_node(struct avl_tree_t *t, int v);

int avl_tree_remove_node(struct avl_tree_t *t, int v);

void avl_tree_levelorder_traverse(const struct avl_tree_t *t);

void avl_tree_inorder_traverse(const struct avl_tree_t *t);

#ifdef __cplusplus
}
#endif

#endif /* _AVL_H_ */
