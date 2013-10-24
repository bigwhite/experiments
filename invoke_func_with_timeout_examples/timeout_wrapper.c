
#include <setjmp.h>
#include <stdarg.h>
#include <unistd.h>
#include <stdio.h>
#include <signal.h>
#include "timeout_wrapper.h"

int invoke_count = 0;
sigjmp_buf invoke_env;

void 
timeout_signal_handler(int sig) 
{
    invoke_count++;
    siglongjmp(invoke_env, 1);
}
