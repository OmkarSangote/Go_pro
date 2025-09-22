#include <iostream>
struct date
{
    int day;
    int month;
    int year;
};

struct student
{
    char name[20];
    int regno;
    struct date dob;
};

int main()
{
    struct student std1;
    std::cout << "Enter Student's Register number: ";
    std::cin >> std1.regno;
    std::cin.ignore(std::numeric_limits<std::streamsize>::max(), '\n');
    std::cout << "Enter Student's Name: ";
    std::cin.getline(std1.name,20);
    std::cout << "Enter Student's date of birth: ";
    std::cin >> std1.dob.day >> std1.dob.month >> std1.dob.year;
    std::cout << "\n";
    std::cout << " STUDENT DETAILS \n";
    std::cout << "Student's Register number: " << std1.regno << "\n";
    std::cout << "Student's Name: " << std1.name << "\n";
    std::cout << "Student's date of birth: " << std1.dob.day << "-" << std1.dob.month << "-" << std1.dob.year << "\n";

}