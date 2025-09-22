#include <iostream>
using namespace std;

enum R
{
   w,
   e,
   r
};
R a = w;
R b = e;
R c = r;

int main ()
{
    // union R R1;
    // R1.x = 5;
    // R1.c ='a';
    // R1.y = 7.5;
    cout << a<< b << c<< endl;
    // cout << R2.c << endl;
    // cout << R3.y ;
    return 0;
}