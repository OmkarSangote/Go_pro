#include<iostream>
int reverse(int n);
int main()
{
    int num,rev;
    std::cout << "Enter the number: ";
    std::cin >> num;
    rev = reverse(num);
    std::cout << "Reverse of the number: " << rev << "\n";
    if (num == rev)
    {
        std::cout << "The given number is pallindrome \n";
    }
    else {std::cout << "The given number is not pallindrom \n";}
}
    int reverse(int n)
    {
        int r,rev = 0;
        while (n!=0)
        {
            r = n % 10;
            n = n / 10;
            rev = rev * 10 + r ;

        }

        return rev;
    }
