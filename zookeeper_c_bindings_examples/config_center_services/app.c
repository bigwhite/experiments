#include <stdio.h>
#include <stdlib.h>
#include <string.h>
#include <unistd.h>
#include <stdbool.h>
#include "zookeeper.h"

bool i_am_leader = false;
int64_t startup_time_stamp = 0;

struct watch_func_para_t {
    zhandle_t *zkhandle;
    char node[64];
};

static void zktest_dump_stat(const struct Stat *stat)
{
    char tctimes[40];
    char tmtimes[40];
    time_t tctime;
    time_t tmtime;

    if (!stat) {
        fprintf(stderr,"null\n");
        return;
    }
    tctime = stat->ctime/1000;
    tmtime = stat->mtime/1000;

    ctime_r(&tmtime, tmtimes);
    ctime_r(&tctime, tctimes);

    fprintf(stderr, "ctime = [%ld]\n", tctime);
    fprintf(stderr, "\tctime = %s\tczxid=%llx\n"
            "\tmtime=%s\tmzxid=%llx\n"
            "\tversion=%x\taversion=%x\n"
            "\tephemeralOwner = %llx\n",
            tctimes, stat->czxid,
            tmtimes, stat->mzxid,
            (unsigned int)stat->version, (unsigned int)stat->aversion,
            stat->ephemeralOwner);
}

static int 
is_leader(zhandle_t* zkhandle, char *myid);

static void 
get_node_name(const char *buf, char *node);

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
    zktest_dump_stat(&stat);
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
    
    printf("\n\n");
    printf("children event happened =====>\n");
    printf("children event: node [%s] \n", path);
    printf("children event: state [%d]\n", state);
    printf("children event: type[%d]\n", type);
    printf("children event done  <=====\n");

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
    zktest_dump_stat(&stat);

    if (strings.count == 0) return;

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




    return;
}

void def_ccs_watcher(zhandle_t* zh, int type, int state,
        const char* path, void* watcherCtx)
{
    printf("Something happened.\n");
    printf("type: %d\n", type);
    printf("state: %d\n", state);
    printf("path: %s\n", path);
    printf("watcherCtx: %s\n", (char *)watcherCtx);
}


int 
main(int argc, const char *argv[])
{
    const char* host = "127.0.0.1:2181";
    zhandle_t* zkhandle;
    int timeout = 5000;
    char node[512] = {0};
    char cur_id[64] = "item-";
    
    zoo_set_debug_level(ZOO_LOG_LEVEL_WARN);
    zkhandle = zookeeper_init(host, def_ccs_watcher, timeout, 
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

static void 
get_node_name(const char *buf, char *node) 
{
    const char *p = buf;
    int i;
    for (i = strlen(buf) - 1; i >= 0; i--) {
        if (*(p + i) == '/') {
            break;
        }
    }

    strcpy(node, p + i + 1);
    return;
}
