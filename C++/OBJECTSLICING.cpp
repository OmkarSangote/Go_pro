#include <iostream>
using namespace std;

class Base
{
    public:
    int i;
    Base(int a)
    {
        i =a;
    }

    virtual void Display()
    {
        cout << i <<" Base class" << endl;
    }
};

class Derived: public Base
{
    public:
    int j;
    Derived(int a, int b): Base(a)
    {
        j= b;
    }

   virtual void Display()
    {
        cout << i << " "<< j << " Devided class " << endl;
    }
};


  void xyz(Base *obj)
   {
       obj->Display();
   }

int main()
{
    Base *b = new Base(10);
    Derived *d = new Derived(11,12);
    xyz(b);
    xyz(d);
    
    return 0;

}

