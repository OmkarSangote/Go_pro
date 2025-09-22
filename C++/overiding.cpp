#include <iostream>
using namespace std;

class A
{
    public:
   virtual int print()
    {
        cout << "Base print fun called";
    }

     int show()
    {
      cout << "Base show fun called";
    }

};

class B: public A
{
    public:
    int print()
    {
        cout << "Derived print fun called";
    }

     int show()
    {
      cout << "Derived show fun called";
    }

};

int main()
{
 A *a;
 B b;
 a = &b;
a->print();
a->show();
return 0;
}