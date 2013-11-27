/*
 * bpt.h
 *
 * an implementation of b+ tree
 */

#ifndef _BPTREE_H_
#define _BPTREE_H_

/* headers included */

#ifdef __cplusplus
extern "C" {
#endif

struct bpt_t* bpt_new(int order);
void bpt_free(struct bpt_t **t);
int bpt_search_node(const struct bpt_t *t, int key);
int bpt_insert_node(struct bpt_t *t, int v);
int bpt_remove_node(struct bpt_t *t, int v);

#ifdef __cplusplus
}
#endif

#endif /* _BPT_H_ */
