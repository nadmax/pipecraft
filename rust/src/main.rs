use crud_api::server::Server;

fn main() {
    let server = Server::new();

    server.run();
}
