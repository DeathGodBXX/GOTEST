#include <iostream>
#include <time.h>
#define N 3

using namespace std;
void test(); //预先声明，为了让main函数调用的时候可以找到

int main(){
    clock_t start,end;
    double elapsed;
    start = clock();
    for(int i=0;i<N;i++){
       test(); 
    }
    end = clock();
    elapsed = (double)(end-start)/CLOCKS_PER_SEC;
    cout<<"耗时是:"<<(elapsed/N)<<endl;
}

void test(){
    for(int i=1;i<10;i++){
        for(int j=1;j<=i;j++){
            cout << i << "x" << j << "=" << i*j <<"\t";
        }
        cout<<"\n";
    }
}