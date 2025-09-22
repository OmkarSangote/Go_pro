#include <iostream>
int factorial(int n);
int main()
{
    int n;
    std::cout << "Enter the number foe which we need factorial: ";
    std::cin >> n;
    if (n>0) std::cout << "The factorial of " << n << " is = " << factorial(n) << "\n";
    else std::cout << "Invalid input (Enter POSITIVE NUMBERS only)";
}

int factorial(int n)
{
    if (n==0) return 1;
    else return (n*factorial(n-1));
}