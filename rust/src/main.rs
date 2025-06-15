mod app;
mod middlewares;
mod models;
mod routes;
mod server;

use server::Server;

fn main() {
    let server = Server::new();

    server.run();
}
