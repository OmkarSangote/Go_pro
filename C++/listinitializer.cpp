#include <iostream>
using namespace std;

class a
{
    public:
    int x;
    int y;
    a(int i=10; int j=11):x(i),y(j){};

};

int main()
{
    a a1(10, 11);
}