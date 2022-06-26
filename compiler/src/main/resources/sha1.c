#include <inttypes.h>

typedef struct {
    char hash[41];
} sha1HashResult;

typedef struct {
    uint32_t h0, h1, h2, h3, h4;
} sha1HashState;

// do not put this into .rodata to prevent tampering
__attribute__((section(".text")))
static char hexDigits[16] = "0123456789abcdef";

static sha1HashState initalState __attribute__((section(".text"))) = {
        .h0 = 0x67452301,
        .h1 = 0xEFCDAB89,
        .h2 = 0x98BADCFE,
        .h3 = 0x10325476,
        .h4 = 0xC3D2E1F0,
};

# define BLOCKLEN (512 / 8)

static uint32_t rotl(uint32_t x, uint32_t n) {
    return (x << n) | (x >> (32-n));
}

static void hashBlock(sha1HashState* state, char* block) {
    uint32_t w[80];
    for (int i = 0; i < 16; i++) {
        w[i] = __builtin_bswap32(((uint32_t*)block)[i]);
    }
    for (int i = 16; i < 80; i++) {
        w[i] = rotl(w[i-3] ^ w[i-8] ^ w[i-14] ^ w[i-16], 1);
    }
    uint32_t a=state->h0, b=state->h1, c=state->h2, 
      d=state->h3, e=state->h4;
    uint32_t f, k, temp;
    
    for (int i = 0; i < 80; i++) {
        if (i < 20) {
            f = (b & c) | ((~ b) & d);
            k = 0x5A827999;
        } else if (i < 40) {
            f = b ^ c ^ d;
            k = 0x6ED9EBA1;
        } else if (i < 60) {
            f = (b & c) | (b & d) | (c & d);
            k = 0x8F1BBCDC;
        } else {
            f = b ^ c ^ d;
            k = 0xCA62C1D6;
        }
        temp = rotl(a, 5) + f + e + k + w[i];
        e = d;
        d = c;
        c = rotl(b, 30);
        b = a;
        a = temp;
    }
    state->h0 += a;
    state->h1 += b;
    state->h2 += c;
    state->h3 += d;
    state->h4 += e;
}

static sha1HashResult sha1Hash(char * msg) {
    sha1HashState state = initalState;
    
    long msgLen = strlen(msg);
    long remLen = msgLen;
    while (remLen >= BLOCKLEN) {
        hashBlock(&state, msg);
        msg += BLOCKLEN;
        remLen -= BLOCKLEN;
    }
    
    char lastBlock[BLOCKLEN];
    _memcpy(lastBlock, msg, remLen);
    lastBlock[remLen] = 0x80;
    remLen++;
    // check if 8 bytes for length do not fit
    if (remLen > BLOCKLEN - 8)  {
        // need two padding blocks
        // fill second to last block with padding
        for (int i = remLen; i < BLOCKLEN; i++) { lastBlock[i] = 0; }
        hashBlock(&state, &lastBlock[0]);
        remLen = 0;
    }
    // add padding to last block
    for (int i = remLen; i < BLOCKLEN - 8; i++) { lastBlock[i] = 0; }
    // add length in bits to last block in big endian
    msgLen *= 8;
    for (int i = 8; i > 0; i--) {
        lastBlock[BLOCKLEN - i] = msgLen >> (8 * i - 8);
    }
    hashBlock(&state, &lastBlock[0]);
    
    // swap byte order for output
    state.h0 = __builtin_bswap32(state.h0);
    state.h1 = __builtin_bswap32(state.h1);
    state.h2 = __builtin_bswap32(state.h2);
    state.h3 = __builtin_bswap32(state.h3);
    state.h4 = __builtin_bswap32(state.h4);
    sha1HashResult result;
    for (int i = 0; i < 20; i++) {
        char c = ((char*)&state)[i];
        result.hash[i*2] = hexDigits[(c >> 4) & 0x0F];
        result.hash[i*2 + 1] = hexDigits[c & 0x0F];
    }
    // add null termination
    result.hash[40] = 0;
    return result;
}
