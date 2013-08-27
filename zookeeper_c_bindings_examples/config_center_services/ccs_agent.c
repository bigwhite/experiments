#include <stdio.h>
#include <stdlib.h>
#include <string.h>
#include <unistd.h>
#include <stdbool.h>
#include <pthread.h>
#include "zookeeper.h"

bool i_am_leader = false;

struct watch_func_para_t {
    zhandle_t *zkhandle;
    char node[64];
};

static int 
is_leader(zhandle_t* zkhandle, char *myid);

static void 
get_node_name(const char *buf, char *node);

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

    int ret = 0;

    struct watch_func_para_t* para= (struct watch_func_para_t*)watcherCtx;

    struct String_vector strings;
    struct Stat stat;
    ret = zoo_wget_children2(para->zkhandle, "/election", ccs_children_watcher, watcherCtx, &strings, &stat);
    if (ret) {
        fprintf(stderr, "child: zoo_wget_children2 error [%d]\n", ret);
        exit(EXIT_FAILURE);
    }

    if (is_leader(para->zkhandle, para->node)) {
        i_am_leader = true;
        printf("This is [%s], i am a leader\n", para->node);
    } else {
        i_am_leader = false;
        printf("This is [%s], i am a follower\n", para->node);
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


static 
int join_the_election(zhandle_t *zkhandle, char node[]) 
{
    char buf[512] = {0};
    int ret;

    ret = zoo_create(zkhandle,
                        "/election/ccs-member", 
                        "hello", 
                        5,
                        &ZOO_OPEN_ACL_UNSAFE,  /* a completely open ACL */
                        ZOO_SEQUENCE|ZOO_EPHEMERAL,
                        buf, 
                        sizeof(buf)-1);
    if (ret) {
        fprintf(stderr, "zoo_create error [%d]\n", ret);
        return ret;
    }

    get_node_name(buf, node);

    struct Stat stat;
    memset(&stat, 0, sizeof(stat));
    struct String_vector strings;
    struct watch_func_para_t para;
    memset(&para, 0, sizeof(para));
    para.zkhandle = zkhandle;
    strcpy(para.node, node);
    ret = zoo_wget_children2(zkhandle, "/election", ccs_children_watcher, &para, &strings, &stat);
    if (ret) {
        fprintf(stderr, "zoo_wget_children2 error [%d]\n", ret);
        return ret;
    }

    return ret;
}


static void*
trigger_listen_thread(void *arg)
{

}

static void*
item_expire_thread(void *arg)
{

}

int 
main(int argc, const char *argv[])
{
    const char* host = "127.0.0.1:2181";
    zhandle_t* zkhandle;
    int timeout = 5000;
    char node[512] = {0};
    
    zoo_set_debug_level(ZOO_LOG_LEVEL_WARN);
    zkhandle = zookeeper_init(host, def_ccs_watcher, timeout, 
                              0, "Zookeeper examples: config center services", 0);
    if (zkhandle == NULL) {
        fprintf(stderr, "Connecting to zookeeper servers error...\n");
        exit(EXIT_FAILURE);
    }

    /* join the election group */
    int ret = join_the_election(zkhandle, node);
    if (zkhandle == NULL) {
        fprintf(stderr, "join the election error...\n");
        exit(EXIT_FAILURE);
    }

    /* leader election */
    if (is_leader(zkhandle, node)) {
        i_am_leader = true;
        printf("This is [%s], i am a leader\n", node);
    } else {
        i_am_leader = false;
        printf("This is [%s], i am a follower\n", node);
    }

    /* start the trigger listen thread */
    pthread_attr_t attr1;
    ret = pthread_attr_init(&attr1);
    if (ret != 0) {
        fprintf(stderr, "pthread_attr_init error...\n");
        exit(EXIT_FAILURE);
    }

    pthread_t trigger_listen_thr;
    ret = pthread_create(&trigger_listen_thr, &attr1, trigger_listen_thread, zkhandle);
    if (ret != 0) {
        fprintf(stderr, "pthread_create error...\n");
        exit(EXIT_FAILURE);
    }

    ret = pthread_attr_destroy(&attr1);
    if (ret != 0) {
        fprintf(stderr, "pthread_attr_destroy error...\n");
        exit(EXIT_FAILURE);
    }
    
    /* start the item expire thread */
    pthread_attr_t attr2;
    ret = pthread_attr_init(&attr2);
    if (ret != 0) {
        fprintf(stderr, "pthread_attr_init error...\n");
        exit(EXIT_FAILURE);
    }

    pthread_t expire_thr;
    ret = pthread_create(&expire_thr, &attr2, item_expire_thread, zkhandle);
    if (ret != 0) {
        fprintf(stderr, "pthread_create error...\n");
        exit(EXIT_FAILURE);
    }

    ret = pthread_attr_destroy(&attr2);
    if (ret != 0) {
        fprintf(stderr, "pthread_attr_destroy error...\n");
        exit(EXIT_FAILURE);
    }

    void *res;
    ret = pthread_join(trigger_listen_thr, (void**)&res);
    if (ret != 0) {
        fprintf(stderr, "pthread_join error...\n");
        exit(EXIT_FAILURE);
    }

    ret = pthread_join(expire_thr, (void**)&res);
    if (ret != 0) {
        fprintf(stderr, "pthread_join error...\n");
        exit(EXIT_FAILURE);
    }

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
