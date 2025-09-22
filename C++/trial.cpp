#include <iostream>
using namespace std;

// class A
// {
//     public:

//     void Show(int x)
//     {
//         cout << x << endl;

//     }

//     void Show(int a, int b)
//     {
//         cout << a << endl;
//         cout << b << endl;

//     }
// };

// int main ()
// {
//     A a;
//     a.Show(10);
//     a.Show(20,52);
//     return 0;

// }

class A
{
    public:
    int x;

    A()
    {
       x = 5;
    }

    A operator ++()
    {
        
     ++x;
    }

    A operator ++(int)
    {
     
        x++;
    }

    void Show()
{
    cout << x << endl;

}

};

int main ()
{
    A a;
    ++a;
   a.Show();
   a++;
    a.Show();
    return 0;
}