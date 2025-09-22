#include <iostream>
using namespace std;

class A
{
    private:
    int x=0;
    public:
    A()
    {
        cout << "sss";
    }
    A(A& a3)
    {
      cout << a3.x << endl;
    }
   int get(int r)
   {
       x=r;
       return x;
   }


};

int main()
{
    A a;
    A b = a;
    cout << a.get(6);
    return 0;
}