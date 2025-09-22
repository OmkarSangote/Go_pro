#include <iostream>
using namespace std;

union R
{
    int x;
    char c;
    float y;
};

int main ()
{
    union R R1;
    R1.x = 5;
    R1.c ='a';
    R1.y = 7.5;
    cout << R1.x << R1.c << R1.y<< endl;
    // cout << R2.c << endl;
    // cout << R3.y ;
    return 0;
}

// #include <iostream>
// using namespace std;
 
// union GFG {
//     int Geek1;
//     char Geek2;
//     float Geek3;
// };
 
// int main()
// {
//     union GFG G1, G2, G3;
 
//     G1.Geek1 = 34;
//     G2.Geek2 = 34;
//     G3.Geek3 = 34.34;
 
//     // Printing values
//     cout << "The first value at "
//          << "the allocated memory : "
//          << G1.Geek1 << endl;
 
//     cout << "The next value stored "
//          << "after removing the "
//          << "previous value : "
//          << G2.Geek2 << endl;
 
//     cout << "The Final value value "
//          << "at the same allocated "
//          << "memory space : "
//          << G3.Geek3 << endl;
//     return 0;
// }