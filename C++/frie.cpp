#include <iostream>
using namespace std;

class a
{
    private:
    int x;

    friend void show(a);

};

void show(a a1)
{
    cout << a1.x ;
}

int main()
{
    a a1;
    show(a1);

}
