#include <iostream>
using namespace std;

class a
{
    public:
    int x;


    void accept()
    {
        cout << "Enter x" << endl;
        cin >> x;
    }
    a operator +(a &a2)
    {   
        a temp;
        temp.x =x+a2.x;
        return temp;
        
    }
    void print()
    {
        cout << x << endl;

    }


};

int main()
{
 a a1,a2, a3;
 a1.accept();
 a2.accept();
 a3=a1+a2;
 a3.print();
}
