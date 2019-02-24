#include <cstdio>

const int MAX_N = 50;

int main(int argc, char **argv) 
{
    int n, m, k[MAX_N];
    bool flag = false;

    scanf("%d %d", &n, &m);
    for (int i=0; i< n ; i++ ){
        scanf("%d", &k[i]);
    }
    
    printf("n = %d, m = %d\n", n, m);
    for (int i = 0; i < n; i++) {
        printf("%d ", k[i]);
    }
    puts("");
    
    int a, b, c, d;

    for (a =0 ; a< n ;a++ ){
        for (b=0; b<n; b++){
            for (c=0; c<n ;c++){
                for (d=0;d<n;d++) {
                    if (k[a] + k[b] + k[c] + k[d] == m) {
                        flag = true;
                        printf("The series is %d %d %d %d", k[a], k[b], k[c], k[d]);
                        goto result;
                    }

                }
            }

        }
    }

    
result:
    if(flag) {
        puts("Yes");
    }else {
        puts("NO");
    }

    return 0;
}
