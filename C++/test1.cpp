#include <iostream>
using namespace std;

class AC
{
public:
float temperature = 0.0;
protected:
 AC()
{
    cout << "AC is on" << endl;

}
 ~AC()
{
    cout << "AC is off" << endl;
}
void setTemp(float t)
{
    temperature = t;
    cout << "AC temp:" << temperature << endl;

}
};

class CAR:public AC
{
public:
char* model;
CAR()
{
    cout << "CAR is on" << endl;
}
 ~CAR()
{
    cout << "CAR is off" << endl;
}

void switchOnAc()
{
 cout <<"Enter the Temperature:" << endl;
 cin >> temperature;
 setTemp(temperature);
}
};



int main()
{
CAR c;
c.switchOnAc();
}