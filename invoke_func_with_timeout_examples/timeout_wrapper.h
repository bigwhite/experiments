
#include <setjmp.h>
#include <stdarg.h>
#include <unistd.h>
#include <stdio.h>
#include <signal.h>
#include <string.h>
#include <errno.h>
#include <sys/time.h>

extern int invoke_count;
extern sigjmp_buf invoke_env;

void timeout_signal_handler(int sig);
typedef void (*sighandler_t)(int);
#define E_CALL_TIMEOUT (-9)

/* interval: microseconds */
#define timeout_func(func, n, interval, ret, ...) \
    { \
        invoke_count = 0; \
        sighandler_t h = signal(SIGALRM, timeout_signal_handler); \
        if (h == SIG_ERR) { \
            ret = errno; \
            goto end; \
        }  \
\
        if (sigsetjmp(invoke_env, SIGALRM) != 0) { \
            if (invoke_count >= n) { \
                ret = E_CALL_TIMEOUT; \
                goto err; \
            } \
        } \
\
        struct itimerval tick;  \
        struct itimerval oldtick;  \
        tick.it_value.tv_sec = interval/1000; \
        tick.it_value.tv_usec = (interval%1000) * 1000; \
        tick.it_interval.tv_sec = interval/1000; \
        tick.it_interval.tv_usec = (interval%1000) * 1000; \
\
        if (setitimer(ITIMER_REAL, &tick, &oldtick) < 0) { \
            ret = errno; \
            goto err; \
        } \
\
        ret = func(__VA_ARGS__);\
        setitimer(ITIMER_REAL, &oldtick, NULL); \
err:\
        signal(SIGALRM, h); \
end:\
        ;\
    }
