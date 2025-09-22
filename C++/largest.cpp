#include <iostream>
int main()
{
    int a,b,c,large,small,seclar;
    std::cout << "Enter the 3 numbers:\n";
    std::cin >> a >> b >> c;
    if (a>b)
    {
        large = a;
        small = b;
    }
    else 
    {
        small = a;
        large = b;
    }
    if (large > c)
    {
        std::cout << "The largest number is:" << large << "\n";
    }
    else
    {
        std::cout << "The largest number is:" << c << "\n";
        
    }
    return 0;
}