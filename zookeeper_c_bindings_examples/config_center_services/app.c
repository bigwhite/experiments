#include <stdio.h>
#include <stdlib.h>
#include <string.h>
#include <unistd.h>
#include <stdbool.h>
#include <time.h>
#include "zookeeper.h"

time_t startup_time_stamp = 0;

static int
add_children_watch_on(zhandle_t *zh, const char *path, watcher_fn watcher, void *watcherCtx) 
{
    int ret = 0;

    struct String_vector strings;
    struct Stat stat;
    ret = zoo_wget_children2(zh, path, watcher, watcherCtx, &strings, &stat);
    if (ret) {
        fprintf(stderr, "zoo_wget_children2 error [%d]\n", ret);
    }
    return ret;
}

static int
cmpid(const void *p1, const void *p2)
{
    return strcmp(* (char * const *) p1, * (char * const *) p2);
}

void 
ccs_children_watcher(zhandle_t* zh, int type, int state,
                      const char* path, void* watcherCtx)
{
    
    printf("child event happened: type[%d]\n", type);

    /*
    struct Stat {
        int64_t czxid;
        int64_t mzxid;
        int64_t ctime; //use this
        int64_t mtime;
        int32_t version;
        int32_t cversion;
        int32_t aversion;
        int64_t ephemeralOwner;
        int32_t dataLength;
        int32_t numChildren;
        int64_t pzxid;
    };
    */


    int ret = 0;
    char *cur_id = (char*)watcherCtx;

    struct String_vector strings;
    struct Stat stat;
    ret = zoo_wget_children2(zh, "/ccs/employee_info_tab", ccs_children_watcher, watcherCtx, &strings, &stat);
    if (ret) {
        fprintf(stderr, "child: zoo_wget_children2 error [%d]\n", ret);
        return;
    }

    if (strings.count == 0) return;

    /* only care child event */
    if (type != ZOO_CHILD_EVENT) return;

    /* routine for item creating */
    char* *p = NULL;
    p = malloc(strings.count * sizeof(char*));
    if (p == NULL) {
        fprintf(stderr, "child: malloc error\n");
        return;
    }
    memset(p, 0, strings.count * sizeof(char*));

    for (int i = 0;  i < strings.count; i++) {
        p[i] = strings.data[i];
    }
    qsort(&p[0], strings.count, sizeof(p[0]), cmpid);

    for (int i = 0;  i < strings.count; i++) {
        puts(p[i]);
    }

    for (int i = 0;  i < strings.count; i++) {
        char path[128] = {0};
        char value[128] = {0};
        int value_len = sizeof(value);
        struct Stat item_stat;
        sprintf(path, "/ccs/employee_info_tab/%s", p[i]);
        
        if (strcmp(p[i], cur_id) <= 0) {
            printf("employee_info_tab: skip [%s]\n", p[i]);
            continue;
        };

        ret = zoo_get(zh, path, 0, value, &value_len, &item_stat); 
        if (ret != 0) {
            fprintf(stderr, "child: zoo_get error\n");
            continue;
        }

        if ((item_stat.ctime/1000) > startup_time_stamp) {
            printf("employee_info_tab: execute [%s]\n", value);
            strcpy(cur_id, p[i]);
        } else {
            printf("employee_info_tab: skip [%s]\n", value);
        }
    }

    free(p);

    return;
}

int 
main(int argc, const char *argv[])
{
    const char* host = "127.0.0.1:2181";
    zhandle_t* zkhandle;
    int timeout = 5000;
    char node[512] = {0};
    char cur_id[64] = "item-";

    (void)time(&startup_time_stamp);
    
    zoo_set_debug_level(ZOO_LOG_LEVEL_WARN);
    zkhandle = zookeeper_init(host, NULL, timeout, 
                              0, "Zookeeper examples: config center services", 0);
    if (zkhandle == NULL) {
        fprintf(stderr, "Connecting to zookeeper servers error...\n");
        exit(EXIT_FAILURE);
    }

    int ret = add_children_watch_on(zkhandle, "/ccs/employee_info_tab", ccs_children_watcher, cur_id);
    if (ret) {
        fprintf(stderr, "zoo_wget_children2 error [%d]\n", ret);
        exit(EXIT_FAILURE);
    }

    sleep(50000); // only for experiments

    zookeeper_close(zkhandle);
}

