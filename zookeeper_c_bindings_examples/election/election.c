#include <stdio.h>
#include <stdlib.h>
#include <string.h>
#include <unistd.h>
#include "zookeeper.h"

struct watch_func_para_t {
    zhandle_t *zkhandle;
    char node[64];
};

static int 
is_leader(zhandle_t* zkhandle, char *myid);

static void 
get_node_name(const char *buf, char *node);

void 
election_parent_watcher(zhandle_t* zh, int type, int state,
                        const char* path, void* watcherCtx)
{
    
    printf("\n\n");
    printf("event happened =====>\n");
    printf("event: node [%s] \n", path);
    printf("event: state [%d]\n", state);
    printf("event: type[%d]\n", type);
    printf("event done  <=====\n");

    int ret = 0;

    zhandle_t *zkhandle = (zhandle_t*)watcherCtx;
    struct Stat stat;

    ret = zoo_wexists(zkhandle, "/election", election_parent_watcher, watcherCtx, &stat);
    if (ret) {
        fprintf(stderr, "Error %d for %s\n", ret, "wexists");
        exit(EXIT_FAILURE);
    }

}

void 
election_children_watcher(zhandle_t* zh, int type, int state,
                      const char* path, void* watcherCtx)
{
    
    printf("\n\n");
    printf("children event happened =====>\n");
    printf("children event: node [%s] \n", path);
    printf("children event: state [%d]\n", state);
    printf("children event: type[%d]\n", type);
    printf("children event done  <=====\n");

    int ret = 0;

    struct watch_func_para_t* para= (struct watch_func_para_t*)watcherCtx;

    struct String_vector strings;
    struct Stat stat;
    ret = zoo_wget_children2(para->zkhandle, "/election", election_children_watcher, watcherCtx, &strings, &stat);
    if (ret) {
        fprintf(stderr, "child: zoo_wget_children2 error [%d]\n", ret);
        exit(EXIT_FAILURE);
    }

    if (is_leader(para->zkhandle, para->node))
        printf("This is [%s], i am a leader\n", para->node);
    else
        printf("This is [%s], i am a follower\n", para->node);

    return;
}

void def_election_watcher(zhandle_t* zh, int type, int state,
        const char* path, void* watcherCtx)
{
    printf("Something happened.\n");
    printf("type: %d\n", type);
    printf("state: %d\n", state);
    printf("path: %s\n", path);
    printf("watcherCtx: %s\n", (char *)watcherCtx);
}


static int 
is_leader(zhandle_t* zkhandle, char *myid) 
{
    int ret = 0;
    int flag = 1;

    struct String_vector strings;
    ret = zoo_get_children(zkhandle, "/election", 0, &strings);
    if (ret) {
        fprintf(stderr, "Error %d for %s\n", ret, "get_children");
        exit(EXIT_FAILURE);
    }

    // printf("leader election ==>\n");
    for (int i = 0;  i < strings.count; i++) {
        // printf("[%s] vs [%s]\n", myid, strings.data[i]); 
        if (strcmp(myid, strings.data[i]) > 0) {
            flag = 0;
            break; 
        }
    }
    // printf("leader election end<==\n");

    return flag;
}

int 
main(int argc, const char *argv[])
{
    const char* host = "127.0.0.1:2181";
    zhandle_t* zkhandle;
    int timeout = 5000;
    char buf[512] = {0};
    char node[512] = {0};
    
    zoo_set_debug_level(ZOO_LOG_LEVEL_WARN);
    zkhandle = zookeeper_init(host, def_election_watcher, timeout, 
                              0, "Zookeeper examples: election", 0);
    if (zkhandle == NULL) {
        fprintf(stderr, "Connecting to zookeeper servers error...\n");
        exit(EXIT_FAILURE);
    }

    int ret = zoo_create(zkhandle,
                        "/election/member", 
                        "hello", 
                        5,
                        &ZOO_OPEN_ACL_UNSAFE,  /* a completely open ACL */
                        ZOO_SEQUENCE|ZOO_EPHEMERAL,
                        buf, 
                        sizeof(buf)-1);
    if (ret) {
        fprintf(stderr, "zoo_create error [%d]\n", ret);
        exit(EXIT_FAILURE);
    }

    get_node_name(buf, node);
    if (is_leader(zkhandle, node)) {
        printf("This is [%s], i am a leader\n", node);
    } else {
        printf("This is [%s], i am a follower\n", node);
    }

    struct Stat stat;
    memset(&stat, 0, sizeof(stat));
    ret = zoo_wexists(zkhandle, "/election", election_parent_watcher, zkhandle, &stat);
    if (ret) {
        fprintf(stderr, "zoo_wexists error [%d]\n", ret);
        exit(EXIT_FAILURE);
    }

    struct String_vector strings;
    struct watch_func_para_t para;
    memset(&para, 0, sizeof(para));
    para.zkhandle = zkhandle;
    strcpy(para.node, node);
    ret = zoo_wget_children2(zkhandle, "/election", election_children_watcher, &para, &strings, &stat);
    if (ret) {
        fprintf(stderr, "zoo_wget_children2 error [%d]\n", ret);
        exit(EXIT_FAILURE);
    }

    /* just wait for experiments*/
    sleep(10000);

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
