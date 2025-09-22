#include<iostream>
int main()
{
    char str[20],ch;
    int i = 0;
    std::cout << "Enter the string : \n";
    std::cin.get(ch);
    while (ch!='\n')
    {
        str[i] = ch;
        std::cin.get(ch);
        i++;
    }

    str[i]='\0';
    std::cout << "The string is : " << str << "\n"; 

}