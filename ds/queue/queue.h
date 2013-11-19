/*
 * queue.h
 *
 * brief description of this file
 */

#ifndef _QUEUE_H_
#define _QUEUE_H_

#ifdef __cplusplus
extern "C" {
#endif

struct queue_t;

struct queue_t* queue_new();
int enqueue(struct queue_t *q, void *item);
void* dequeue(struct queue_t *q);
int is_queue_empty(const struct queue_t *q);
void queue_free(struct queue_t **q);

#ifdef __cplusplus
}
#endif

#endif /* _QUEUE_H_ */
