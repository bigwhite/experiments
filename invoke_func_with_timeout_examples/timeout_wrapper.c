
#include <setjmp.h>
#include <stdarg.h>
#include <unistd.h>
#include <stdio.h>
#include <signal.h>
#include "timeout_wrapper.h"

volatile int invoke_count = 0;
sigjmp_buf invoke_env;

void 
timeout_signal_handler(int sig) 
{
    invoke_count++;
    siglongjmp(invoke_env, 1);
}

sigfunc *
my_signal(int signo, sigfunc *func)
{
    struct sigaction act, oact;

    act.sa_handler = func;
    sigemptyset(&act.sa_mask);
    act.sa_flags = 0;
    if (signo == SIGALRM) {
#ifdef SA_INTERRUPT
        act.sa_flags |= SA_INTERRUPT;
#endif
    } else {
#ifdef SA_RESTART
        act.sa_flags |= SA_RESTART;
#endif
    }
    if (sigaction(signo, &act, &oact) < 0)
        return SIG_ERR;
    return oact.sa_handler;
}

