#include <iostream>
using namespace std;
class A
{
protected:
int a=1;
};

class B: public A
{
    public:
    void show()
    {
        cout << a ;
    }
};

int main ()
{
    B b;
    b.show();
    return 0;
}
