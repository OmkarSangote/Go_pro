#include <iostream>
using namespace std;

void Show(int &x)
{
    cout << x << endl;

}

int main ()
{
   int a= 10;
   Show (a);
    return 0;
}