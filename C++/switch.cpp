#include <iostream>
int main()
{
    int i;
    std::cout << "Enter the Case ID: ";
    std::cin >> i;
    switch (i)
    {
        case 1: std::cout << "one\n";
                break;
        case 2: std::cout << "Two\n";
                break; 
        case 3: std::cout << "Three\n";
                break;
        default: std::cout << "Out of count\n";
    }
    return 0;
}