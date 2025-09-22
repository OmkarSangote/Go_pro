#include <iostream>

/* Prints Contents of Memory Blocks */
static void print_bytes(const void *object, size_t size){
    #ifdef __cplusplus
    const unsigned char * const bytes = static_cast<const unsigned char *>(object);
    #else // __cplusplus
    const unsigned char * const bytes = object;
    #endif // __cplusplus

    size_t i;

    printf("[-");
    for(i = 0; i < size; i++)
    {
        //printf(bytes[i]);
        int binary[8];
        for(int n = 0; n < 8; n++){
            binary[7-n] = (bytes[size -1 - i] >> n) & 1;
        }
        /* print result */
        for(int n = 0; n < 8; n++){
            printf("%d", binary[n]);
        }
        printf("%c", '-');
    }
    printf("]\n\n");
}


int main () {

    std::cout << "\nStoring a Char in Memory";
    std::cout << "\n------------------------\n\n";

    char firstLetter = 'abc';

    std::cout << "Address is "<< static_cast<void *>(&firstLetter) << "\n\n";
    std::cout << "Size is "<<  sizeof(firstLetter) << " bytes\n\n";
    std::cout << "Value is " <<  firstLetter << "\n\n";

    std::cout << "Memory Blocks : \n";
    print_bytes(&firstLetter, sizeof(firstLetter));

    return 0;
}