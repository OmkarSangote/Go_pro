#include <iostream>
using namespace std;

// class A
// {
//     public:
//     static int x ;
//     A()
//     {

//     }


// };

// int A::x = 10;

// int main()
// {
//     A a1;
//     cout << a1.x ;
// }

class A
{
    public:
    int show(int x)
    {
      cout << "class A";
    }
      void show(int y, int z)
    {
        cout << "Derived";
    }

};


int main ()
{
    A a;
    a.show(10, 11);

}
