#include <iostream>
using namespace std;

class a
{
    public:
    int *age;
    a()
    {
       age = new int(10);
    }
    a(const a &a1)
    {
        age = a1.age;
    }

    ~a()
    {
        delete age;
    }

    int getage()
    {
        return *age;
    }

    int setage(int a)
    {
       *age = a;
    }

};

int main ()
{
    a a1;
    a *pointerobj = &a1;
   // cout << pointerobj->getage();
     cout << pointerobj->getage() << endl;
     pointerobj->setage(45);
    cout << pointerobj->getage();

}