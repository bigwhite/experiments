

#include <sys/types.h>
#include <sys/mman.h>

#include <stdio.h>
#include <errno.h>
#include <unistd.h>
#include <time.h>

#define STAT_FILE "./perf/xxstat"

int 
main() 
{
    FILE *fp = NULL;

    errno = 0;
    fp = fopen(STAT_FILE, "r");
    if (fp == NULL) {
        printf("can not open stat file , err = %d\n", errno);
        return -1;
    }

    long size = sysconf(_SC_PAGESIZE);
    errno = 0;
    char *p = NULL;
    p = mmap(NULL, size, PROT_READ, 
             MAP_SHARED, fileno(fp), 0);
    if (p == MAP_FAILED) {
        printf("can not mmap file, error = %d\n", errno);
        fclose(fp);
        return -1; 
    }

    errno = 0;
    if (fclose(fp) != 0) {
        printf("can not close file, error = %d\n", errno);
        return -1;
    }

    /* round up to a multiple of 8 */
    while((int)p % 8 != 0) {
        p++;
    }
    
    long long *q = (long long*)p;

    while(1) {
        printf("%lld\t\t%lld\t\t%lld\t\t%lld\n", q[0], q[1], q[2], q[3]);
        sleep(1);
    }
    
    return 0;
}
