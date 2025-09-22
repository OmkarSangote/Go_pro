#include <iostream>
using namespace std;

class A
{
    public:
    int x= 6;
};

class B:virtual public A
{
    public:
    int z= 9;
};

class C:virtual public A
{
    public:
    int s= 10;
};

class D:public B, public C
{
    public:
    int y =5;

};

int main()
{
    
    D d;
    cout << d.x;
    cout << d.y;
    cout << d.z;
    cout << d.s;
}