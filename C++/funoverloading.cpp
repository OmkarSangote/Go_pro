#include <iostream>
using namespace std;

class over
{
    public:
    int ride(int x)
    {
        cout << x <<endl;
    }

     int ride(int x, int y)
    {
        cout << x << "&" << y <<endl;
    }

      int ride(int x, double y)
    {
        cout << x  << "&" << y << endl;
    }

};
int main()
{
 over o;
 o.ride(5);
 o.ride(6,8);
 o.ride(7,9.8);
return 0;

}