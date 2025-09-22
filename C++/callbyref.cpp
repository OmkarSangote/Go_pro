#include<iostream>
void func(int *x,int *y);
int main()
{
    int a=10,b=20;
    std::cout << "The values of a and b before calling func: " << a << " and " << b << "\n";
    func(&a,&b);
    std::cout << "The values of a and b after calling func: " << a << " and " << b << "\n";
    return 0;
}

void func(int *x,int *y)
{
    *x = 20;
    *y = 10;
}