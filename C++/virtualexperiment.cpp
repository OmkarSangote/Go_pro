#include <iostream>
using namespace std;

class base
{
    public:
    int x =10;

    virtual void show()
    {
        cout << x << endl;
    }
};

class derived:public base
{
    public:
    int y =11;

     virtual void show()
    {
        cout << y << x << endl;
    }
};

int main ()
{
  derived *d;
  base *b;
  b =d;

    b->show();
    return 0;
}