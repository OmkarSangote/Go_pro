#include <iostream>
using namespace std;

class a
{
    public:
    int x = 0;
    int y = 0;
    
    a(int x, int y)
    {
      this->x =x;
      this->y =y;
    }

    a &ex(int a)
    {
        x = a;
        return *this;
    }

    a &ex1(int b)
    {
        y = b;
        return *this;
    }

    void show()
    {
        cout << x << y << endl;
    }
    
};

int main()
{
     a a1(5,10);
     a1.ex(11);
     a1.ex1(12);
     a1.show();
    return 0;
}