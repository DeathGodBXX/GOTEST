#include <stdio.h>
#include <time.h>

#define N 3

void test(); //预先声明，为了让main函数调用的时候可以找到
void main()
{
    // int start = time();
    clock_t start, end;
    double elapsed;
    start = clock();
    for(int i=0;i<N;i++){
        test();
    }
    end = clock();
    elapsed = (double)(end - start)/CLOCKS_PER_SEC;
    printf("耗时是:%f\n",elapsed/N);
}

void test()
{
    for (int i = 1; i < 10; i++)
    {
        for (int j = 1; j <= i; j++)
        {
            printf("%dx%d=%d  ", i, j, i * j);
        }
        printf("\n");
    }
}