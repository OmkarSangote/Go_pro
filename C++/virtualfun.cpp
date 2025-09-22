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
      cout << "Base show fun called" << endl;
    }

    virtual int can(int y)
    {
        cout << y;
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

    int can(int y, int r)
    {
        cout << y;
        cout << r;
    }


};

int main ()
{
    A *a;
    B b;
    a = &b;
    a->can(5);
    a->print();
    a->show();
    return 0;
}