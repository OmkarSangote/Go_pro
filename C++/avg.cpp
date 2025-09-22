#include <iostream>
int main()
{
    int a,b,c;
    float avg;
    std::cout << "Enter the numbers:\n";
    std::cin >> a >> b >> c;
    avg = (a+b+c)/3.0;
    std::cout << "The average of 3 numbers is: " << avg << "\n";
    return 0;
}