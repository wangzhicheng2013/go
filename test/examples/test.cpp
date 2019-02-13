#include <stdint.h>
#include <iostream>
using namespace std;
uint64_t fun(uint64_t n)
{
    if (n > 0)
    {
        return n * fun(n - 1);
    }
    return 1;
}
int main()
{
    cout << fun(20) << endl;
    return 0;
}

