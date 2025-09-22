#include<iostream>
int main()
{
    int i,count=0;
    char str[100],ch;
    std::cout << "Enter the string elements : ";
    std::cin.getline(str,100,'\n');
    std::cout << "Enter the character whose occurance is counted : ";
    std::cin.get(ch);
    for(i=0;str[i]!='\0';i++)
    {
    if(str[i] == ch ) count++;
    }
    std::cout << ch << " occurs " << count << " times \n";
}