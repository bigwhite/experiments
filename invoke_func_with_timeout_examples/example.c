
#include "timeout_wrapper.h"


int 
sim_oci_call(int a, int b) {
    int ret;
    while(1)
        ret = a + b;

    return a+b;
}

int 
test1()
{
    int ret = 0;
    int try_times = 3;
    int interval = 1000;
    add_timeout_to_func(sim_oci_call, try_times, interval, ret, 1, 2);
    if (ret == E_CALL_TIMEOUT) {
        printf("invoke sim_oci_call timeouts for 3 times\n");
        return -1;
    } else if (ret == 0) {
        printf("invoke sim_oci_call ok\n");
        return 0;
    } else {
        printf("timeout_func error = %d\n", ret);
    }
}

int 
test2()
{
    #define MAXLINE 1024
    char line[MAXLINE];

    int ret = 0;
    int try_times = 3;
    int interval = 1000;
    add_timeout_to_func(read, try_times, interval, ret, STDIN_FILENO, line, MAXLINE);
    if (ret == E_CALL_TIMEOUT) {
        printf("invoke read timeouts for 3 times\n");
        return -1;
    } else if (ret == 0) {
        printf("invoke read ok\n");
        return 0;
    } else {
        printf("timeout_func error = %d\n", ret);
    }
}

int 
main() 
{
    test1();
    test2();
}
