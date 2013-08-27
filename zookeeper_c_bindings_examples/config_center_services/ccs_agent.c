#include <stdio.h>
#include <stdlib.h>
#include <string.h>
#include <unistd.h>
#include <stdbool.h>
#include <pthread.h>
#include <sys/types.h>
#include <sys/socket.h>
#include <netinet/in.h>
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
    printf("children event happened: type[%d]\n", type);

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
int join_the_election(zhandle_t *zkhandle, char node[], struct watch_func_para_t *para) 
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
    memset(para, 0, sizeof(*para));
    para->zkhandle = zkhandle;
    strcpy(para->node, node);
    ret = zoo_wget_children2(zkhandle, "/election", ccs_children_watcher, para, &strings, &stat);
    if (ret) {
        fprintf(stderr, "zoo_wget_children2 error [%d]\n", ret);
        return ret;
    }

    return ret;
}

static void get_section();
static int 
parse_trigger_pkg(const char *buf, int buf_len, char table[], 
                  char oper_type[], char id[])
{
    /*
     * pkg format:
     *   ^table_name|oper_type|id$
     *
     * oper_type: 
     *   ADD/DEL/MOD/BAT
     */

    if (buf[0] != '^') return -1;
    if (buf[buf_len-1] != '$') return -1;
    int pos[2];

    for (int i = 1, j = 0; i < buf_len; i++) {
        if (buf[i] == '|') {
            pos[j] = i;
            j++;
        }
    }
    
    strncpy(table, &buf[1], pos[0] - 1);
    strncpy(oper_type, &buf[pos[0]+1], pos[1] - pos[0] - 1);
    strncpy(id, &buf[pos[1]+1], buf_len -1 - pos[1] -1);

    printf("table[%s], oper_type[%s], id[%s]\n", table, oper_type, id);

    return 0;
}

static void*
trigger_listen_thread(void *arg)
{
    fprintf(stderr, "trigger listen thread start up!\n");
    int ret;
    zhandle_t* zkhandle = (zhandle_t*)arg;
    int sock;
    int cli_sock;
    struct sockaddr_in sin ;
    struct sockaddr_in cin;
    socklen_t len = sizeof(struct sockaddr);

    sock = socket(AF_INET, SOCK_STREAM, 0);
    if (sock < 0) {
        fprintf(stderr, "socket error!\n");
        return;
    }

    memset(&sin, 0, sizeof(sin));
    sin.sin_family       = AF_INET;
    sin.sin_port         = htons(8877);
    sin.sin_addr.s_addr = INADDR_ANY;

    ret = bind(sock, (struct sockaddr*)&sin, len);
    if (ret < 0) {
        close(sock);
        fprintf(stderr, "socket bind error!\n");
        return;
    }

    ret = listen(sock, 5);
    if (ret < 0) {
        close(sock);
        fprintf(stderr, "socket listen error!\n");
        return;
    }

    while (1) {
        memset(&cin, 0, sizeof(cin));
        cli_sock = accept(sock, (struct sockaddr*)&cin, &len);
        if (cli_sock < 0) {
            fprintf(stderr, "accept client socket error!\n");
        }

        /* recv and parse the trigger package */
        char buf[512];
        int num = 0;
        char table[64];
        char oper_type[64];
        char id[64];
        memset(buf, 0, sizeof(buf));
        memset(table, 0, sizeof(table));
        memset(oper_type, 0, sizeof(oper_type));
        memset(id, 0, sizeof(id));

        num = recv(cli_sock, buf, sizeof(buf), 0);
        if (num < 0) {
            fprintf(stderr, "recv client socket error");
            close(cli_sock);
            continue;
        }

        ret = parse_trigger_pkg(buf, num, table, oper_type, id);
        if (ret != 0) {
            fprintf(stderr, "parse trigger pkg error");
            close(cli_sock);
            continue;
        }

        /* create sequenced item on ccs */
        char path[128] = {0};
        sprintf(path, "/ccs/%s/item-", table);
        char value[128] = {0};
        sprintf(value, "%s %s", oper_type, id);

        ret = zoo_create(zkhandle,
                path,
                value,
                strlen(value),
                &ZOO_OPEN_ACL_UNSAFE,  /* a completely open ACL */
                ZOO_SEQUENCE,
                buf,
                sizeof(buf)-1);
        if (ret) {
            fprintf(stderr, "zoo_create error [%d]\n", ret);
        }

        close(cli_sock);
    }
}


static void*
item_expire_thread(void *arg)
{
    zhandle_t* zkhandle = (zhandle_t*)arg;
    int timeout = 30;
    struct String_vector strings;
    int ret;
    time_t now;

    fprintf(stderr, "item expire thread start up!\n");

    while(1) {
        if (i_am_leader) {
            /* clear the expired items */
            ret = zoo_get_children(zkhandle, "/ccs/employee_info_tab", 0,
                                   &strings);
            if (ret != 0) {
                fprintf(stderr, "zoo_get_children error [%d]\n", ret);
            }
            time(&now);

            for (int i = 0; i < strings.count; i++) {
                char path[128];
                memset(path, 0, sizeof(path));
                char value[128] = {0};
                int value_len = sizeof(value);
                struct Stat item_stat;

                sprintf(path, "/ccs/employee_info_tab/%s", strings.data[i]);
                ret = zoo_get(zkhandle, path, 0, value, &value_len, &item_stat);
                if (ret != 0) {
                    fprintf(stderr, "zoo_get error [%d]\n", ret);
                    continue;
                }

                if ((item_stat.ctime/1000) < now - 30) {
                    ret = zoo_delete(zkhandle, path, -1);
                    if (ret != 0) {
                        fprintf(stderr, "zoo_delete error [%d]\n", ret);
                        continue;
                    }
                    printf("[expire]: employee_info_tab: expire [%s]\n", strings.data[i]);
                } else {
                    printf("[expire]: employee_info_tab: skip [%s]\n", strings.data[i]);
                }

            }
        } 
        sleep(timeout);
    }

}

static int
create_tables(zhandle_t *zkhandle, char p[][64], int count)
{
    int ret = 0;
    char path[128];
    char buf[128] = {0};

    ret = zoo_create(zkhandle,
                        "/ccs",
                        "hello",
                        5,
                        &ZOO_OPEN_ACL_UNSAFE,  /* a completely open ACL */
                        0,
                        buf, 
                        sizeof(buf)-1);
    if (ret == ZNODEEXISTS) {
        ret = 0;
        fprintf(stderr, "/ccs node exists\n");
    } else if (ret != 0) {
        fprintf(stderr, "zoo_create error [%d]\n", ret);
        return ret;
    } else {
        fprintf(stderr, "create [%s] ok\n", "/ccs");
    } 

    for (int i = 0; i < count; i++) {
        memset(path, 0, sizeof(path));
        sprintf(path, "/ccs/%s", p[i]);
        ret = zoo_create(zkhandle,
                            path,
                            "hello", 
                            5,
                            &ZOO_OPEN_ACL_UNSAFE,  /* a completely open ACL */
                            0,
                            buf, 
                            sizeof(buf)-1);
        if (ret == ZNODEEXISTS) {
            ret = 0;
            fprintf(stderr, "%s node exists\n", path);
        } else if (ret != 0) {
            fprintf(stderr, "zoo_create error [%d]\n", ret);
            return ret;
        } else {
            fprintf(stderr, "create [%s] ok\n", path);
        } 
    }

    return ret;
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

int 
main(int argc, const char *argv[])
{
    const char* host = "127.0.0.1:2181";
    zhandle_t* zkhandle;
    int timeout = 5000;
    char node[512] = {0};
    char tables[][64] = {
        {"employee_info_tab"},
        {"boss_info_tab"}
    };
    
    zoo_set_debug_level(ZOO_LOG_LEVEL_WARN);
    zkhandle = zookeeper_init(host, NULL, timeout, 
                              0, "Zookeeper examples: config center services", 0);
    if (zkhandle == NULL) {
        fprintf(stderr, "Connecting to zookeeper servers error...\n");
        exit(EXIT_FAILURE);
    }

    /* join the election group */
    struct watch_func_para_t para;
    int ret = join_the_election(zkhandle, node, &para);
    if (zkhandle == NULL) {
        fprintf(stderr, "join the election error...\n");
        exit(EXIT_FAILURE);
    }

    /* leader election */
    if (is_leader(zkhandle, node)) {
        i_am_leader = true;
        printf("This is [%s], i am a leader\n", node);
        (void)create_tables(zkhandle, tables, sizeof(tables)/sizeof(tables[0]));
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

