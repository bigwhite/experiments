#include <stdio.h>
#include "defer.h"

extern struct defer_func_ctx ctx_stack[10];

__attribute__((no_instrument_function))
void
__cyg_profile_func_enter(void *this_fn, void *call_site) 
{
    printf("enter func => %p\n", this_fn);
}

__attribute__((no_instrument_function))
void 
__cyg_profile_func_exit(void *this_fn, void *call_site) 
{
    printf("exit func <= %p\n", this_fn);

    struct defer_func_ctx *ctx = NULL;
    printf("in funcexit: stack top is [%d]\n", stack_top());

    while ((ctx = stack_pop()) != NULL) {
        printf("params count is [%d]\n", ctx->params_count);
        if (ctx->params_count == 0) {
            printf("defer pop function: [%p]\n", ctx->ctx.zp.df);
            ctx->ctx.zp.df();
        } else if (ctx->params_count == 1) {
            printf("defer pop function: [%p]\n", ctx->ctx.op.df);
            ctx->ctx.op.df(ctx->ctx.op.p1);
        } else if (ctx->params_count == 2) {
            printf("in exit: df = %p\n", (ctx->ctx).tp.df);
            printf("in exit: p1 = %p\n", (ctx->ctx).tp.p1);
            printf("in exit: p2 = %p\n", (ctx->ctx).tp.p2);
            printf("defer pop function: [%p]\n", ctx->ctx.tp.df);
            ctx->ctx.tp.df(ctx->ctx.tp.p1, ctx->ctx.tp.p2);
        }
    }
}
