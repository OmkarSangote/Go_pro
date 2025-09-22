#include <iostream>
int main()
{
    int num,i;
    std::cout << "Enter the number: ";
    std::cin >> num;
    for (i=0;i<=num;i++)
    {
        if (num%i==0)
        std::cout << "The factors are: "<< i << "\n";
    }
}

