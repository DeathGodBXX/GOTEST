#include <iostream>
#include <string>
#include "test.h"
// #include <tuple>

int main()
{
    Tuple<int, float, std::string> t(41, 6.3, std::string("hello"));
    std::cout << t.head() << std::endl;
    std::cout << "t的首地址是:" << &t << std::endl;
    std::cout << t.tail() << std::endl; //this指针移动了，输出的结果this指向的inherited
    std::cout << t.tail().head() << std::endl;
    // &(t.tail());
    return 0;
}
