#include <iostream>
using namespace std;

class a
{
    public:
    int x = 10;
     void show()
    {
        cout << x  << endl;
    }
    
};
class b : public a
{
    public:
    int x = 11;
    int y = 12;
     void show()
    {
        cout << x <<" " << y << endl;
    }
};


int main()
{
    b *b1 = new b;
    //a *a1 = b1;
  
    b1->show();
    return 0;
}