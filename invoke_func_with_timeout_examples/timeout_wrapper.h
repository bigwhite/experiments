
#include <setjmp.h>
#include <stdarg.h>
#include <unistd.h>
#include <stdio.h>
#include <signal.h>
#include <string.h>
#include <errno.h>
#include <sys/time.h>

extern volatile int invoke_count;
extern sigjmp_buf invoke_env;

void timeout_signal_handler(int sig);
typedef void sigfunc(int sig);
sigfunc *my_signal(int signo, sigfunc* func);
#define E_CALL_TIMEOUT (-9)

/* interval: microseconds */
#define add_timeout_to_func(func, n, interval, ret, ...) \
    { \
        invoke_count = 0; \
        sigfunc *sf = my_signal(SIGALRM, timeout_signal_handler); \
        if (sf == SIG_ERR) { \
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
        my_signal(SIGALRM, sf); \
end:\
        ;\
    }
