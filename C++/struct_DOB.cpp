#include <iostream>
int main()
{
    struct date
    {
        int day;
        int month;
        int year;
    };

    struct date DOB;
    std::cout << "Enter your date of birth : ";
    std::cin >> DOB.day >> DOB.month >> DOB.year;
    std::cout << "The date of birth is: " << DOB.day << "-" << DOB.month << "-" << DOB.year << "\n";
}