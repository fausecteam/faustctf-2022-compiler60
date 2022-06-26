static inline int syscall(int sys_num, long rdi, long rsi, long rdx) {
    int ret;
    asm volatile (
        "syscall"
        : "=a" (ret)
        : "0"(sys_num), "D"(rdi), "S"(rsi), "d"(rdx)
        : "rcx", "r11", "memory"
    );
    return ret;
}

static int open(char* name, int flags, int mode) {
    return syscall(0x02, (long)name, flags, mode);
}

static int read(unsigned int fd, char* buf, unsigned int count) {
    return syscall(0x00, fd, (long)buf, count);
}

static int write(unsigned int fd, char* buf, unsigned long count) {
    return syscall(0x01, fd, (long)buf, count);
}

void exit(int code) {
    long rax = 0x3c;
    asm volatile (
        "syscall"
        :
        : "a"(rax), "D"(code)
        : "rcx", "r11", "memory"
    );
}

// not a syscall but inline asm to prevent gcc from implemention the whole big vectorized memcpy
static void _memcpy(char* dest, char* src, long len) {
    asm volatile (
            "rep movsb"
            : "+D"(dest), "+S"(src), "+c"(len)
            :
            : "memory"
        );
}
