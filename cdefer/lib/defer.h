#ifndef _DEFER_H_
#define _DEFER_H_

typedef void (*defer_func)();

struct zero_params_func_ctx {
    defer_func df;
};

struct one_params_func_ctx {
    defer_func df;
    void *p1;
};

struct two_params_func_ctx {
    defer_func df;
    void *p1;
    void *p2;
};

struct defer_func_ctx {
    int params_count;
    union {
        struct zero_params_func_ctx zp;
        struct one_params_func_ctx op;
        struct two_params_func_ctx tp;
    } ctx;
};

void stack_push(struct defer_func_ctx *ctx);
struct defer_func_ctx* stack_pop();
int stack_top();

#endif
