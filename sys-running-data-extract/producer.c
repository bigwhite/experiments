

#include <sys/types.h>
#include <sys/mman.h>

#include <stdio.h>
#include <errno.h>
#include <unistd.h>

#define STAT_FILE "./perf/xxstat"

int 
main() 
{
    FILE *fp = NULL;

    errno = 0;
    fp = fopen(STAT_FILE, "w+");
    if (fp == NULL) {
        printf("can not create stat file , err = %d\n", errno);
        return -1;
    }

    errno = 0;
    long size = sysconf(_SC_PAGESIZE);
    if (ftruncate(fileno(fp), size) != 0) {
        printf("can not set stat file size, err = %d\n", errno);
        fclose(fp);
        return -1;
    }

    errno = 0;
    char *p = NULL;
    p = mmap(NULL, size, PROT_WRITE|PROT_READ, 
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
    q[0] = 1;
    q[1] = 1000;
    q[2] = 10000;
    q[3] = 100000;

    while(1) {
        q[0] += 1;
        q[1] += 10;
        q[2] += 100;
        q[3] += 1000;
        usleep(200);
    }
    
    return 0;
}
