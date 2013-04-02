#include <stdio.h>
#include <stdlib.h>
#include <unistd.h>
#include <string.h>
#include <limits.h> /* for PATH_MAX */

static char path[PATH_MAX];

__attribute__((constructor))
static void executable_path_init() {
    char    buf[PATH_MAX];

    memset(buf, 0, sizeof(buf));
    memset(path, 0, sizeof(path));

#ifdef _SOLARIS_TRACE
    getcwd(buf, PATH_MAX);
    sprintf(path, "%s/%s", buf, getexecname());
#elif _LINUX_TRACE
    readlink("/proc/self/exe", path, PATH_MAX);
#else
    #error "The OS has not been supported!"
#endif
}

__attribute__((no_instrument_function))
void __cyg_profile_func_enter(void *this_fn, void *call_site) {
    char buf[PATH_MAX];
    char cmd[PATH_MAX];

    memset(buf, 0, sizeof(buf));
    memset(cmd, 0, sizeof(cmd));

    sprintf(cmd, "addr2line %p -e %s -f", this_fn, path);

    FILE *ptr = NULL;
    memset(buf, 0, sizeof(buf));

    if ((ptr = popen(cmd, "r")) != NULL) {
        fgets(buf, PATH_MAX, ptr);
        buf[strlen(buf) - 1] = '\0';
        printf("enter func => %p:%s:", this_fn, buf);
        fgets(buf, PATH_MAX, ptr); /* filename and lineno */
        printf("%s\n", buf);
    }

    (void) pclose(ptr);
}

__attribute__((no_instrument_function))
void __cyg_profile_func_exit(void *this_fn, void *call_site) {
    char buf[PATH_MAX];
    char cmd[PATH_MAX];

    memset(buf, 0, sizeof(buf));
    memset(cmd, 0, sizeof(cmd));

    sprintf(cmd, "addr2line %p -e %s -f", this_fn, path);

    FILE *ptr = NULL;
    memset(buf, 0, sizeof(buf));

    if ((ptr = popen(cmd, "r")) != NULL) {
        fgets(buf, PATH_MAX, ptr);
        buf[strlen(buf) - 1] = '\0';
        printf("exit func <= %p:%s:", this_fn, buf);
        fgets(buf, PATH_MAX, ptr); /* filename and lineno */
        printf("%s\n", buf);
    }

    (void) pclose(ptr);
}
