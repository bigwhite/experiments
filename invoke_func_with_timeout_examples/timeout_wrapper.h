
#include <setjmp.h>
#include <stdarg.h>
#include <unistd.h>
#include <stdio.h>
#include <signal.h>
#include <string.h>

extern int invoke_count;
extern sigjmp_buf invoke_env;

void timeout_signal_handler(int sig);
typedef void (*sighandler_t)(int);

#define timeout_func(func, n, interval, ret, ...) \
    { \
        invoke_count = 0; \
        sighandler_t h = signal(SIGALRM, timeout_signal_handler); \
        if (h == SIG_ERR) { \
            perror("install sigal error"); \
            ret = -1 ; \
            goto end; \
        }  \
\
        if (sigsetjmp(invoke_env, SIGALRM) != 0) { \
            printf("invoke_count is %d\n", invoke_count); \
            if (invoke_count >= n) { \
                ret = -1; \
                goto end; \
            } \
        } \
\
        alarm(interval);\
        ret = func(__VA_ARGS__);\
        alarm(0); \
        printf("call func ok\n");\
end:\
        signal(SIGALRM, h);\
    }
