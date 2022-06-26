#define _GNU_SOURCE
#include <errno.h>
#include <grp.h>
#include <signal.h>
#include <spawn.h>
#include <stdio.h>
#include <stdlib.h>
#include <sys/mman.h>
#include <sys/resource.h>
#include <sys/wait.h>
#include <sys/sendfile.h>
#include <unistd.h>

#define MAX_BINARY_SIZE 64 * 1024

void execHelper(char** arguments) {
    pid_t pid;
    int r = posix_spawnp(&pid, arguments[0], NULL, NULL, &arguments[0], NULL);
    if (r != 0) {
        perror("posix_spawnp");
        exit(1);
    }

    int wstatus; 
    if (waitpid(pid, &wstatus, 0) != pid) {
        perror("waitpid");
        exit(1);
    }
    
    if (!WIFEXITED(wstatus) || WEXITSTATUS(wstatus) != 0) {
        fprintf(stderr, "Failed %s\n", arguments[0]);
        exit(1);
    }
}

void catch_alarm (int sig) {
    kill(0, SIGKILL);
}

int main(int argc, char** argv) {
    if (argc != 2) {
        fprintf(stderr, "usage: %s <path_to_stdlib.c>\n", argv[0]);
        return 1;
    }
    
    // kill after 3 sec
    signal(SIGALRM, catch_alarm);
    alarm(3);
    
    char* stdlibFile = argv[1];

    // limit output file size using prlimit
    struct rlimit limits = {.rlim_cur = MAX_BINARY_SIZE,
                            .rlim_max = MAX_BINARY_SIZE};
    setrlimit(RLIMIT_FSIZE, &limits);

    if (getuid() == 0) {
        // drop privileges
        setgroups(0, NULL);
        setgid(9999);
        setuid(9999);
    }

    int asmOutFd = memfd_create("asm_out", 0);
    char asmOutFile[100];
    snprintf(asmOutFile, 100, "/proc/self/fd/%d", asmOutFd);

    char* asArgs[] = {"as", "-o", asmOutFile, "-", NULL};
    execHelper(asArgs);

    int binaryOutFd = memfd_create("binary_out", 0);
    char binaryOutFile[100];
    snprintf(binaryOutFile, 100, "/proc/self/fd/%d", binaryOutFd);

    char* ldArgs[] = {
        "ld",
        "-pie",
        "-x",
        "--no-dynamic-linker",
        "-z", "norelro",
        "-z", "noseparate-code",
        "-z", "noexecstack",
        "--hash-style=sysv",
        "-o", binaryOutFile,
        asmOutFile,
        stdlibFile,
        NULL};
    execHelper(ldArgs);

    close(asmOutFd);

    //  copy linked elf to stdout, the linker can't directly write to stdout cause it needs seek 
    off_t offset = 0;
    sendfile(STDOUT_FILENO, binaryOutFd, &offset, MAX_BINARY_SIZE);

    close(binaryOutFd);

    return 0;
}
