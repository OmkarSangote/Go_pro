#include <iostream>
using namespace std;

class Complex
{
    public:
    float real =0;
    float img =0;

    void input()
    {
        cout << "enter real and imaginary parts" << endl;
        cin >> real;
        cin >> img;
    }

    Complex operator + (const Complex& obj)
    {
        Complex temp;
        temp.real = real + obj.real;
        temp.img = img + obj.img;
        return temp;
    }

    void output()
    {
        if(img < 0)
        cout << real << img << "i" <<endl;
        else
        cout << real << "+" << img << "i" << endl;
    }


};

int main()
{
    Complex complex1, complex2, result;
    complex1.input();
    complex2.input();
    result = complex1 + complex2;
    result.output();

    return 0;
}



