#include <iostream>
int main()
{
    char str[100];
    int ualpha=0,lalpha=0,digitcount=0,splchar=0;
    std::cout << "Enter the string: ";
    std::cin.getline(str,100);
    for (int i=0;str[i]!='\0';i++)
    {
        if(isupper(str[i])) ualpha++;
        else if (islower(str[i])) lalpha++;
        else if (isdigit(str[i])) digitcount++;
        else splchar++;
    }

    std::cout << "Number of upper case: " << ualpha << "\n";
    std::cout << "Number of lower case: " << lalpha << "\n";
    std::cout << "Number of digits: " << digitcount << "\n";
    std::cout << "Number of spl case: " << splchar << "\n";
}