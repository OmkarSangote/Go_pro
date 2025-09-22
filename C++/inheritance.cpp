#include <iostream>
using namespace std;
class A
{
public:
int a=1;
};

class B: public A
{
    public:
    int b=2;

};

class C: public A
{
    public:
    int c=3; 
};

class D: public B, public C
{
    public:
    int d=4; 
};

int main ()
{
 D d1;
  cout << d1.B::a << endl;
 cout << d1.b << endl;
 cout << d1.c << endl;
 cout << d1.d << endl;
    return 0;
}
