
#include <stdio.h>
#include <string.h>
#include <sys/types.h>
#include <unistd.h>
#include <libmemcached/memcached.h>

static const char *product_id = "nexus5";
static const int component_in_total = 5;
static const int component_order[] = {2, 3, 1, 2, 5, 4};

//code from <Algorithms.for.Programmers.Ideas.and.Source.Code>
static inline unsigned long long 
bit_count(unsigned long long x)
{
    x = (0x5555555555555555UL & x) + (0x5555555555555555UL & (x >> 1));
    x = (0x3333333333333333UL & x) + (0x3333333333333333UL & (x >> 2));
    x = (0x0f0f0f0f0f0f0f0fUL & x) + (0x0f0f0f0f0f0f0f0fUL & (x >> 4));
    x = (0x00ff00ff00ff00ffUL & x) + (0x00ff00ff00ff00ffUL & (x >> 8));
    x = (0x0000ffff0000ffffUL & x) + (0x0000ffff0000ffffUL & (x >> 16));
    x = (0x00000000ffffffffUL & x) + (0x00000000ffffffffUL & (x >> 32));
    return x;
}

int 
main(int argc, char *argv[])
{
    memcached_st *memc;
    memcached_return_t rc = MEMCACHED_SUCCESS;
    memcached_server_st *server = NULL;

    memc = memcached_create(NULL);
    if (NULL == memc) {
        printf("memcached_create error\n");
        return -1;
    }

    server = memcached_server_list_append(server, "127.0.0.1", 11211, &rc);
    if (rc != MEMCACHED_SUCCESS) {
        printf("memcached_server_list_append error: %s\n", memcached_strerror(memc, rc));
        return -1;
    }

    rc = memcached_server_push(memc, server);
    if (rc != MEMCACHED_SUCCESS) {
        printf("memcached_server_push error: %s\n", memcached_strerror(memc, rc));
        return rc;
    }

    memcached_server_list_free(server);

    rc = memcached_behavior_set(memc, MEMCACHED_BEHAVIOR_SUPPORT_CAS, 1);
    if (rc != MEMCACHED_SUCCESS) {
        printf("memcached_behavior_set support cas error: %s\n", memcached_strerror(memc, rc));
        return -1;
    }

    /* pack the component one by one */
    int ret = 0;
    int i = 0;
    for (i = 0; i < sizeof(component_order)/sizeof(component_order[0]); i++) {
        ret = pack_component(memc, component_order[i]);
        if (ret == 0) {
            printf("pack component [%d] ok\n", component_order[i]);
        } else if (ret == 1) {
            printf("pack component [%d] exists\n", component_order[i]);
        } else {
            printf("other error occurs\n");
            return -1;
        }
        getchar();
    }

    return 0;
}

int 
pack_component(memcached_st *memc, int i) 
{
    memcached_return_t rc = MEMCACHED_SUCCESS;

    uint32_t mask = 1 << (i - 1);
    uint32_t value_added = 1 << (i - 1);
    char value_added_str[11] = {0};
    uint32_t value = 0;
    char *pvalue = NULL;
    size_t value_len = 0;
    uint32_t flags = 0;

    while(1) {
        pvalue = memcached_get(memc, product_id, strlen(product_id), &value_len, &flags, &rc);
        if (!pvalue) {
            if (rc == MEMCACHED_NOTFOUND) {
                printf("componet [%d] - memcached_get not found product key: [%s]\n", i, product_id);
                memset(value_added_str, 0, sizeof(value_added_str));
                sprintf(value_added_str, "%u", value_added);
                rc = memcached_add(memc, product_id, strlen(product_id), value_added_str, 
                                   strlen(value_added_str), 1000, 0);
                if (rc == MEMCACHED_DATA_EXISTS) {
                    printf("componet [%d] - memcached_add key[%s] exist\n", i, product_id);
                    pvalue = memcached_get(memc, product_id, strlen(product_id), &value_len, &flags, &rc);
                    if (!pvalue) return -1;
                } else if (rc != MEMCACHED_SUCCESS) {
                    printf("componet [%d] - memcached_add error: %s, [%d]\n", i, memcached_strerror(memc, rc), rc);
                    return -1;
                } else {
                    printf("assign a package for product[%s]\n", product_id);
                    printf("componet [%d] - memcached_add key[%s] successfully, its value = %u, cas = %llu\n", i,product_id,
                            value_added, (memc->result).item_cas);
                    return 0;
                }
            } else {
                printf("componet [%d] - memcached_get error: %s, %d\n", i, memcached_strerror(memc, rc), rc);
                return -1;
            }
        } 
        
        value = atoi(pvalue);
        printf("componet [%d] - memcached_get value = %u, cas = %llu\n", i, value, (memc->result).item_cas);

        if (value & mask) {
            free(pvalue);
            return 1;
        } else {
            uint64_t cas_value = 0;
            cas_value = (memc->result).item_cas;
            memset(value_added_str, 0, sizeof(value_added_str));
            sprintf(value_added_str, "%d", value_added + value);

            rc = memcached_cas(memc, product_id, strlen(product_id), value_added_str, strlen(value_added_str), 1000, 0, cas_value);
            if (rc != MEMCACHED_SUCCESS) {
                printf("componet [%d] -  memcached_cas error = %d,  %s\n", i, rc, memcached_strerror(memc, rc));
                free(pvalue);
            } else {
                printf("componet [%d] -  memcached_cas ok\n", i);
                free(pvalue);
                if (bit_count(value_added + value) == component_in_total) {
                    rc = memcached_delete(memc, product_id, strlen(product_id), 0);
                    if (rc != MEMCACHED_SUCCESS) {
                        printf("memcached_delete error: %s\n", memcached_strerror(memc, rc));
                        return -1;
                    } else {
                        printf("memcached_delete key: %s ok\n", product_id);
                    }
                }
                return 0;
            }
        }
        getchar();
    }

    return 0;
}
