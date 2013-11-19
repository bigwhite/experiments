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

struct bst_t;

struct bst_t* bst_new();

void bst_free(struct bst_t **t);

int bst_search_node(const struct bst_t *t, int v);

int bst_insert_node(struct bst_t *t, int v);

int bst_remove_node(struct bst_t *t, int v);

int is_bst_empty(const struct bst_t *t);

void bst_levelorder_traverse(const struct bst_t *t);

void bst_inorder_traverse(const struct bst_t *t);

int bst_height(const struct bst_t *t);


#ifdef __cplusplus
}
#endif

#endif /* _BST_H_ */
