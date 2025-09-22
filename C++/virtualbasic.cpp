#include <iostream>
using namespace std;

class a
{
    public:
     void show()
    {
        cout << "base" << endl;

    }
};

class b: public a{
    public:
    void show()
    {
        cout << "derived" << endl;
    }
};

int main()
{
   b b1;
    b1.a::show();
}