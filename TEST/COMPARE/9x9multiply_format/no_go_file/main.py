import time

def test():
    for i in range(1,10):
        for j in range(1,i+1):
            print(f"{j}x{i}={j*i}  ",end="")
        print()

N=3

if __name__ == "__main__":
    # test()
    # timeit.repeat(lambda : test(),repeat=3)
    start = time.time()
    for i in range(N):
        test()
    end = time.time()
    print(f"耗时是:{(end-start)/N}")

