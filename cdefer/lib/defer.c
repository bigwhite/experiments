#include <stdio.h>
#include <stdarg.h>
#include <stdlib.h>
#include <string.h>
#include "defer.h"

struct defer_func_ctx ctx_stack[10];
int top_of_stack = 0; /* stack top from 1 to 10 */

void
stack_push(struct defer_func_ctx *ctx) 
{
    if (top_of_stack >= 10) {
        return;
    }

    ctx_stack[top_of_stack] = *ctx;
    top_of_stack++;
}

struct defer_func_ctx* 
stack_pop() 
{
    if (top_of_stack == 0) {
        return NULL;
    }

    top_of_stack--;
    return &ctx_stack[top_of_stack];
}

int 
stack_top() 
{
    return top_of_stack;
}

void 
defer(defer_func fp, int arg_count, ...) 
{
    va_list ap;
    va_start(ap, arg_count);

    struct defer_func_ctx ctx;
    memset(&ctx, 0, sizeof(ctx));
    ctx.params_count = arg_count;
    printf("in defer: params count is [%d]\n", ctx.params_count);

    if (arg_count == 0) {
        ctx.ctx.zp.df = fp;

    } else if (arg_count == 1) {
        ctx.ctx.op.df = fp;
        ctx.ctx.op.p1 = va_arg(ap, void*);

    } else if (arg_count == 2) {
        ctx.ctx.tp.df = fp;
        ctx.ctx.tp.p1 = va_arg(ap, void*);
        ctx.ctx.tp.p2 = va_arg(ap, void*);
        ctx.ctx.tp.df(ctx.ctx.tp.p1, ctx.ctx.tp.p2);
    }

    va_end(ap);
    stack_push(&ctx);
    printf("defer push function: [%p]\n", fp);
    printf("in defer: stack top is: [%d]\n", stack_top());
}

