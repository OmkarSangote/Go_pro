#include <iostream>
using namespace std;

class a
{
    public:
    int x;
    a()
    {}
    a(const a &a1)
    {
        x= a1.x;
        cout << " copy is called" << endl;
    }

    a& operator =(const a &a1)
    {
        cout << "Assignment is called" << endl;
        return *this;
    }

};

int main()
{
    a a1, a2;
    a a3 = a1;
    a2 = a1;
  

    return 0;

}