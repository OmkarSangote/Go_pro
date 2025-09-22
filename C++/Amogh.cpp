#include <iostream>

int main ()
{
    int rupee,dollar;
    std::cout << "Enter the rupee value:\n";
    std::cin >> rupee;
    dollar = rupee / 80;
    std::cout<<"Dollar value is:" << dollar;
    return 0;
}