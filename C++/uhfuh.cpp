#include <iostream>
using namespace std;

class unary
{
    public:
    
    int x =0;

    unary(int y)
    {
        x =y;
    }

    unary operator ++ ()
    {
        ++x;
    }

    unary operator ++(int)
    {
        x++;
    }

    void show()
    {
        cout << x << endl;
    }

};

int main()
{
    unary u(8);
    ++u;
    u.show();
    u++;
    u.show();

    return 0;
}