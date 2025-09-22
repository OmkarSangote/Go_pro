#include <iostream>
int length (char ch[]);
void reverse (char ch[]);
int main()
{
    char str[100];
    std::cout << "Enter the elements of the string: ";
    std::cin.getline(str,100);
    std::cout << "Original string is: " << str << "\n";
    reverse(str);
    std::cout << "Reversed string is: " << str << "\n";
}

int length(char ch[])
{
    int i;
    for (i=0;ch[i]!='\0';i++);
    return i;
}

void reverse(char ch[])
{
    int n = length(ch);
    char temp;
    for(int i=0,j=n-1;i<n/2;i++,j--)
    {
        temp = ch[i];
        ch[i] = ch[j];
        ch[j] = temp;
    }

}
