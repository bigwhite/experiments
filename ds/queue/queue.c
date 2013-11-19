/*
 * queue.c
 *
 * brief description of this file
 */

#include "queue.h"
#include <stdlib.h>
#include <string.h>


struct queue_node_t {
    struct queue_node_t *next;
    struct queue_node_t *prev;
    void *data;
};

struct queue_t {
    struct queue_node_t *head;
    struct queue_node_t *tail;
};

struct queue_t* 
queue_new() 
{
    struct queue_t *q = NULL;
    q = malloc(sizeof(*q));
    if (!q) return NULL;
    q->head = q->tail = NULL;
    return q;
}

void 
queue_free(struct queue_t **q)
{
    if ((*q) != NULL)
        free(*q);

    (*q) = NULL;
}

int 
enqueue(struct queue_t *q, void *item)
{
    struct queue_node_t *qn = NULL;
    qn = malloc(sizeof(*qn));
    if (!qn) return -1;
    qn->next = NULL;
    qn->prev = NULL;
    qn->data = item;

    if (q->head == NULL) {
        q->head = q->tail = qn;
    } else {
        qn->prev = q->tail;
        q->tail->next = qn;
        q->tail = qn;
    }
    
    return 0;
}

void* 
dequeue(struct queue_t *q)
{
    if (q->head == NULL) return NULL;

    struct queue_node_t *qn = q->head;

    if (q->head == q->tail) {//only one node
        q->head = q->tail = NULL;
    } else {
        q->head = qn->next;
        q->head->prev = NULL;
    }
    void *p = qn->data;
    free(qn);
    return p;
}

int
is_queue_empty(const struct queue_t *q) 
{
    return q->head == NULL;
    
}
