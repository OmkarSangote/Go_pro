#include <iostream>
using namespace std;

int main()
{
    const int q = 10;
    int const* p = &q;
 
    cout << *p << endl;
    
    const int y =11;
    p = &y;

    cout <<*p<< endl;
    return 0;
}