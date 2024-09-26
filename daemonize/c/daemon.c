
#include <stdio.h>
#include <stdlib.h>
#include <unistd.h>
#include <fcntl.h>
#include <sys/stat.h>
#include <sys/types.h>
#include <syslog.h>
#include <signal.h>

void daemonize()
{
    pid_t pid;

    // 1. Fork off the parent process
    pid = fork();
    if (pid < 0) {
        exit(EXIT_FAILURE);
    }
    // If we got a good PID, then we can exit the parent process.
    if (pid > 0) {
        exit(EXIT_SUCCESS);
    }

    // 2. Create a new session.
    if (setsid() < 0) {
        exit(EXIT_FAILURE);
    }

    // 3. Fork again to ensure the process is not a session leader.
    pid = fork();
    if (pid < 0) {
        exit(EXIT_FAILURE);
    }
    if (pid > 0) {
        exit(EXIT_SUCCESS);
    }

    // 4. Change the current working directory to root.
    if (chdir("/") < 0) {
        exit(EXIT_FAILURE);
    }

    // 5. Set the file mode creation mask to 0.
    umask(0);

    // 6. Close all open file descriptors.
    for (int x = sysconf(_SC_OPEN_MAX); x>=0; x--) {
        close(x);
    }

    // 7. Reopen stdin, stdout, stderr to /dev/null
    open("/dev/null", O_RDWR); // stdin
    dup(0);                    // stdout
    dup(0);                    // stderr

    // Optional: Log the daemon starting
    openlog("daemonized_process", LOG_PID, LOG_DAEMON);
    syslog(LOG_NOTICE, "Daemon started.");
    closelog();
}

int main() {
    daemonize();

    // Daemon process main loop
    while (1) {
        // Perform some background task...
        sleep(30); // Sleep for 30 seconds.
    }

    return EXIT_SUCCESS;
}

