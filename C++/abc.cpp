#include <iostream>
using namespace std;

class complex
{
    public:
    float real =0;
    float img=0;

    void input()
    {
        cin >> real;
        cin >> img;
    }

    complex operator + (const complex& obj)
    {
        complex temp;
        temp.real = real + obj.real;
        temp.img =img + obj.img;
        return temp;
    }

    void output()
    {
        if(img < 0)
        cout << real << img << "i" << endl;
        else
        cout << real << "+" << img << endl;
    }
};

int main()
{
    complex c1, c2, res;
    c1.input();
    c2.input();

    res = c1 + c2;
    res.output();
        return 0;
}