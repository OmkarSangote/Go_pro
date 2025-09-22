#include <iostream>
int main()
{
    int rev = 0,num,digit,n;
    std::cout << "Enter the number: ";
    std::cin >> num;
    n = num;
    while (num>0)
    {
        digit = num % 10;
        num = num / 10;
        rev = rev * 10 + digit;
    }

    std::cout << "The original number is: " << n << "\n";
    std::cout << "The reverse number is: " << rev << "\n";

}