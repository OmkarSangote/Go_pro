#include <iostream>

int a, b, sum;

int main()
{
    std::cout << "Enter two numbers:  \n";
    std::cin >> a >> b;
    sum = a + b;
    std::cout << "Sum: " << sum;

    return 0; // Optional, but good practice to explicitly return 0 from main
}

