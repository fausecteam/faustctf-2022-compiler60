#include "syscalls.c"
#include "strings.c"
#include "sha1.c"

// algol top level function
extern void $top();
// elf entry
void _start() {
    // gcc wants to align stack, but on entry it is already aligned
    asm volatile (
            "addq $8, %%rsp" : : :
        );
    $top();
    exit(0);
}

void writechar(int fd, long c) {
    write(fd, (char*) &c, 1);
}

void outchar(long c) {
    writechar(1, c);
}

void outinteger(long x) {
    char buf[20];
    char* start = integer2string_impl(&buf[0], x);
    write(1, start, &buf[20] - start);
}

void writestring(int fd, char* s) {
    write(fd, s, strlen(s));
}

void outstring(char* s) {
    writestring(1, s);
}

long readchar(int fd) {
    unsigned char c;
    int r = read(fd, &c, 1);
    if (r != 1) {
        return -1;
    }
    return c;
}

void readstring(char* dest, long fd) {
    char* end = dest + 127;
    for (;dest < end; dest++) {
        long c = readchar(fd);
        if (c <= 0) {
            break;
        }
        dest[0] = (char) c;
    }
    dest[0] = 0;
}

#define O_RDONLY	00000000
#define O_WRONLY	00000001
#define O_RDWR		00000002
#define O_CREAT	    00000100

int openRW(char* passwd) {
    sha1HashResult filename = sha1Hash(passwd);
    return open((char*)&filename, O_CREAT | O_RDWR, 0644); // mode rw-r--r-- 
}

int openRO(char* passwd) {
    sha1HashResult filename = sha1Hash(passwd);
    return open((char*)&filename, O_RDONLY, 0);
}

int openWOConfidential(char* filename) {
    return open(filename, O_CREAT | O_WRONLY, 0644); // mode rw-r--r-- 
}
