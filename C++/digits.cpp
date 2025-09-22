#include <iostream>
int main()
{
    int num,digit;
    std::cout << "Enter the number: \n";
    std::cin >> num;
    while (num > 0)
    {
        
        digit = num % 10;
        num = num / 10;
        
        std::cout << "The digits are: " << digit << "\n";
    }

}