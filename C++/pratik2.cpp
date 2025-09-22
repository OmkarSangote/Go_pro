#include <iostream>
using namespace std;
#include <typeinfo>

class a
{
    public:
    void show()
    {
        cout << "Base" << endl;

    }
    
};

class b:public a
{
    public:
     void show()
    {
        cout << "Derived" << endl;
        
    }
    
};

class c: public b
{
      public:
    virtual void show()
    {
        cout << "Derived C" << endl;
        
    }

};

int main()
{
    c c1;
   // a * a1= &b1;
   b *b1 = &c1;

    b1->show();


//      int a = 2.0;
//     // char c= 'A';



//     //void *ptr;

//     //a a1, a2;
//     //const type_info& p1 = typeid(a);
//     cout << typeid(a).hash_code() << endl;
//     //cout << typeid(&a).hash_code() << endl;
//     cout << &a;
//     //cout << typeid(a2).name() << endl;

//     // cout << a*c;
//     // cout << sizeof(ptr);
//


 }