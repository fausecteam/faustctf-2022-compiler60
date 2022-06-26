#include <alloca.h>

static inline char* integer2string_impl(char* dest, long x) {
    char* p = &dest[20];
    if (x >= 0) {
        do {
            *--p = '0' + (x % 10);
            x /= 10;
        } while (x > 0);
    } else {
        do {
            *--p = '0' - (x % 10);
            x /= 10;
        } while (x < 0);
        *--p = '-';
    }
    return p;
}

static inline long strlen(char* s) {
    long len = 0;
    while (*(s++)) {
        len++;
    }
    return len;
}

void integer2string(char* dest, long x) {
    char buf[20];
    char* start = integer2string_impl(&buf[0], x);
    int size = &buf[20] - start;
    // manually use rep movdb here to prevent gcc implementing the whole big vectorized memcpy()
    asm volatile (
            "rep movsb"
            : "+D"(dest), "+S"(start), "+c"(size)
            :
            : "memory"
        );
    *dest = '\0';
}

void $strcpy(char* dest, char* src) {
    while(*(dest++) = *(src++)) {}
}

void $strcat(char* dest, char* left, char* right) {
    if (dest == right) {
        long len = strlen(right) + 1;
        char* rBuffer = alloca(len);
        _memcpy(rBuffer, right, len);
        right = rBuffer;
    }
    while(*left) {
        *(dest++) = *(left++);
    }
    $strcpy(dest, right);
}

int $strget(char* s, unsigned long index) {
    while (index > 0) {
        if (*s == 0) {
            return -1; // illegal string index
        }
        index--;
        s++;
    }
    return (unsigned char)s[0];
}

int $strput(char* s, unsigned long index, char c) {
    while (1) {
        if (*s == 0) {
            return -1; // illegal string index
        }
        if (index == 0) {
            s[0] = c;
            return 0;
        }
        index--;
        s++;
    }
}
