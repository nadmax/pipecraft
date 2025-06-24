use std::io::prelude::*;
use std::net::TcpStream;
use std::sync::Arc;
use crate::routes::users::UserRoutes;
use crate::middlewares::error_handler::ErrorHandler;

pub struct App {
    user_routes: UserRoutes,
    error_handler: ErrorHandler,
}

impl App {
    pub fn new() -> Self {
        let user_store = Arc::new(std::sync::Mutex::new(std::collections::HashMap::new()));
        let next_id = Arc::new(std::sync::Mutex::new(1u32));
        
        App {
            user_routes: UserRoutes::new(user_store, next_id),
            error_handler: ErrorHandler::new(),
        }
    }

    pub fn handle_request(&self, mut stream: TcpStream) {
        let mut buffer = [0; 1024];
        match stream.read(&mut buffer) {
            Ok(_) => {
                let request = String::from_utf8_lossy(&buffer[..]);
                let response = self.route_request(&request);
                let _ = stream.write_all(response.as_bytes());
                let _ = stream.flush();
            }
            Err(_) => {
                let response = self.error_handler.internal_server_error();
                let _ = stream.write_all(response.as_bytes());
            }
        }
    }

    pub fn route_request(&self, request: &str) -> String {
        let lines: Vec<&str> = request.lines().collect();
        if lines.is_empty() {
            return self.error_handler.bad_request();
        }

        let request_line = lines[0];
        let parts: Vec<&str> = request_line.split_whitespace().collect();

        if parts.len() < 3 {
            return self.error_handler.bad_request();
        }

        let method = parts[0];
        let path = parts[1];
        let version = parts[2];
        if version != "HTTP/1.1" {
            return self.error_handler.bad_request();
        }

        if path.starts_with("/users") {
            self.user_routes.handle(method, path, request)
        } else {
            self.error_handler.not_found()
        }
    }
}
