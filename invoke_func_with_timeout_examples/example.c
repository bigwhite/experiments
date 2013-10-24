
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
    timeout_func(sim_oci_call, 3, 2, ret, 1, 2)
    printf("ret = %d\n", ret);
}
