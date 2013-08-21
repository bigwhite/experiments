#include <stdio.h>
#include <stdlib.h>
#include <string.h>
#include <unistd.h>
#include "zookeeper.h"

int 
isLeader( zhandle_t* zkhandle, char *myid);

struct Stat stat;
zhandle_t* zkhandle;
char buf[512] = {'\0'};
char node[512] = {'\0'};

void 
zktest_parent_watcher(zhandle_t* zh, int type, int state,
                      const char* path, void* watcherCtx)
{
    
    printf("\n\n");
    printf("event happened =====>\n");
    printf("event: node [%s] \n", path);
    printf("event: state [%d]\n", state);
    printf("event: type[%d]\n", type);
    printf("event done  <=====\n");

    int ret = 0;

    zkhandle = (zhandle_t*)watcherCtx;

    ret = zoo_wexists(zkhandle, "/election", zktest_parent_watcher, zkhandle, &stat);
    if (ret) {
        fprintf(stderr, "Error %d for %s\n", ret, "wexists");
        exit(EXIT_FAILURE);
    }

}

void 
zktest_children_watcher(zhandle_t* zh, int type, int state,
                      const char* path, void* watcherCtx)
{
    
    printf("\n\n");
    printf("children event happened =====>\n");
    printf("children event: node [%s] \n", path);
    printf("children event: state [%d]\n", state);
    printf("children event: type[%d]\n", type);
    printf("children event done  <=====\n");

    int ret = 0;

    zkhandle = (zhandle_t*)watcherCtx;

    struct String_vector strings;
    ret = zoo_wget_children2(zkhandle, "/election", zktest_children_watcher, zkhandle, &strings, &stat);
    if (ret) {
        fprintf(stderr, "Error %d for %s\n", ret, "wexists");
        exit(EXIT_FAILURE);
    }

    if (isLeader(zkhandle, node))
        printf("[%s] is leader!\n", node);
    else
        printf("[%s] is follower!\n", node);

    /*
    printf("children nodes:\n");
    for (int i = 0;  i < strings.count; i++) {
        printf("%s\n", strings.data[i]);
    }

    printf("children nodes end:\n");
    */
}

void zktest_watcher_g(zhandle_t* zh, int type, int state,
        const char* path, void* watcherCtx)
{
    printf("Something happened.\n");
    printf("type: %d\n", type);
    printf("state: %d\n", state);
    printf("path: %s\n", path);
    printf("watcherCtx: %s\n", (char *)watcherCtx);
}


int 
isLeader( zhandle_t* zkhandle, char *myid) 
{
    int ret = 0;
    int flag = 1;

    struct String_vector strings;
    ret = zoo_get_children(zkhandle, "/election", 0, &strings);
    if (ret) {
        fprintf(stderr, "Error %d for %s\n", ret, "get_children");
        exit(EXIT_FAILURE);
    }

    printf("leader election ==>\n");
    for (int i = 0;  i < strings.count; i++) {
        printf("[%s] vs [%s]\n", myid, strings.data[i]);
        if (strcmp(myid, strings.data[i]) > 0) {
            flag = 0;
            break; 
        }
    }
    printf("leader election end<==\n");

    return flag;
}

int 
main(int argc, const char *argv[])
{
    const char* host = "127.0.0.1:2181";
    int timeout = 5000;
    
    zoo_set_debug_level(ZOO_LOG_LEVEL_WARN);
    zkhandle = zookeeper_init(host,
            zktest_watcher_g, timeout, 0, "hello zookeeper.", 0);
    if (zkhandle == NULL) {
        fprintf(stderr, "Error when connecting to zookeeper servers...\n");
        exit(EXIT_FAILURE);
    }

    int ret = zoo_create(zkhandle,
            "/election/member", 
            "hello", 
            5,
           &ZOO_OPEN_ACL_UNSAFE,  /* a completely open ACL */
           ZOO_SEQUENCE|ZOO_EPHEMERAL, /* ZOO_SEQUENCE */
           buf, 
           sizeof(buf)-1);
    if (ret) {
        fprintf(stderr, "Error %d for %s\n", ret, "zoo_create");
        exit(EXIT_FAILURE);
    }

    printf("I am [%s]\n", buf);

    char *p = buf;
    int i;
    for (i = strlen(buf) - 1; i >= 0; i--) {
        if (*(p + i) == '/') {
            break;
        }
    }

    strcpy(node, p + i + 1);

    if (isLeader(zkhandle, node))
        printf("[%s] is leader!\n", node);
    else
        printf("[%s] is follower!\n", node);

    ret = 0;

    memset(&stat, 0, sizeof(stat));
    ret = zoo_wexists(zkhandle, "/election", zktest_parent_watcher, zkhandle, &stat);
    if (ret) {
        fprintf(stderr, "Error %d for %s\n", ret, "wexists");
        exit(EXIT_FAILURE);
    }
    ret = 0;

    struct String_vector strings;
    ret = zoo_wget_children2(zkhandle, "/election", zktest_children_watcher, zkhandle, &strings, &stat);
    if (ret) {
        fprintf(stderr, "Error %d for %s\n", ret, "wget_children2");
        exit(EXIT_FAILURE);
    }

    /* just wait for experiments*/
    sleep(10000);

    zookeeper_close(zkhandle);
}
