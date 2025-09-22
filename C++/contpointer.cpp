#include <iostream>
using namespace std;

int main()
{
   int q = 10;
    int *const p = &q;

    *p = 8;
    cout << *p ;

    // int y = 11;
    // p = &y;
    return 0;
}