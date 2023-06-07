#include <crow/app.h>
#include <iostream>
int main(){
    std::cout << "im alive"<< std::endl;
    crow::SimpleApp app;
    return EXIT_SUCCESS;
}