#include <stdio.h>

static void foo2() {

}

void foo1() {
    foo2();
}

void foo(){
    chdir("/home/tonybai");
    foo1();
}

int main(int argc, const char *argv[])
{
    foo();    
    return 0;
}
