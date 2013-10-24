
#include "timeout_wrapper.h"


int 
sim_oci_call(int a, int b) {
    int ret;
    while(1)
        ret = a + b;

    return a+b;
}

int main() {
    int ret = 0;
    int try_times = 3;
    int interval = 2;
    timeout_func(sim_oci_call, try_times, interval, ret, 1, 2);
    if (ret == E_CALL_TIMEOUT) {
        printf("invoke sim_oci_call timeouts for 3 times\n");
        return -1;
    } else {
        printf("invoke sim_oci_call ok\n");
        return 0;
    }
}
