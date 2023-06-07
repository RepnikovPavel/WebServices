#include <crow.h>
#include <iostream>
#include <thread>

int main(){
    std::uint16_t port{8080};

    crow::SimpleApp app;
    app.port(port);
    CROW_ROUTE(app, "/route1")(
        [](){return "OK";}
    );
    CROW_ROUTE(app, "/route2").\
    methods(crow::HTTPMethod::POST)([&](const crow::request &req){
        return crow::response(200, "POST METHOD");
    });
    std::thread t([&](){
        std::cout << "listening on port 8080"<< std::endl;
        app.run();
    });
    t.join();
    return EXIT_SUCCESS;
}