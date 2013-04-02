#include <stdio.h>
#include <stdlib.h>
#include "defer.h"

int
bar(int a, char *s) 
{
    printf("a = [%d], s = [%s]\n", a, s);

}

int
main() 
{
    FILE *fp = NULL;
    fp = fopen("main.c", "r");
    if (!fp) return;
    defer(fclose, 1, fp);

    int *p = malloc(sizeof(*p));
    if (!p) return;
    defer(free, 1, p);
    
    defer(bar, 2, 13, "hello");
    return 0;
}

