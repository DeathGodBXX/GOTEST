import time


def decorator(func):
    def wrapper(*args, **kwargs):
        print('=======开始执行========')
        print(func(*args, **kwargs))
        print('=======执行结束========')

    return wrapper


@decorator
def add(x, y):
    print(f'{x}+{y}={x + y}')


start = time.time()
for i in range(0, 100000):
    add(1, i)
collapsed = time.time() - start
print(f"耗时是：{collapsed}")

        