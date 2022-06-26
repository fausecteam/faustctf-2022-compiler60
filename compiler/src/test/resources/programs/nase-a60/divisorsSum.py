
def abs(i):
    return i if i >= 0 else -i

B = [0] * 32
S = [0] * 202

B[1] = 1

K = 2
while K + 200 - B[K - 1] - K >= 0:
#while K <= 30:
    B[K] = B[K - 1] + K // 2
    B[K+1] = B[K] + K + 1
    K += 2

for N in range(1, 201):
    S[0] = N
    U = -1
    T = 0
    L = N - 1

    K = 2
    while K + L - K >= 0:
        T = S[L] + T*U
        L = N - B[K]
        U = -U
        K += 1
    S[N] = abs(T)

    #print(N, S[N])

#for N in range(1, 201):
for N in range(200, 201):
    S[0] = N
    S[N] = 0
    H = -1
    U = -1

    K = N-1
    while K >= 0:
    #for K in range(N-1, -1, -1):
        S[N] = S[K] + S[N]*U
        
        H = ((3+U) * H - 2) // (3 - U)

        U = -U
        K += H
        print(K, H)
    S[N] = abs(S[N])
    print(N, S[N])