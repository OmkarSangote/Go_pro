#include <iostream>
using namespace std;

class A
{
    public:

    friend class B;

    private:
    int x = 10;
    void pr()
   {
     cout << x << endl;
   }

    
};

class B
{
private:
int y = 5;
public:
void show(A &a)
{
    cout << y << endl;
    cout << a.x << endl;
}

};

int main()
{
    A a;
    B b;
    b.show(a);
    
    return 0;
}