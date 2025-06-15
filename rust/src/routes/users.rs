use std::sync::{Arc, Mutex};
use crate::models::user::{User, UserStore};
use crate::middlewares::error_handler::ErrorHandler;

pub struct UserRoutes {
    users: UserStore,
    next_id: Arc<Mutex<u32>>,
    error_handler: ErrorHandler,
}

impl UserRoutes {
    pub fn new(users: UserStore, next_id: Arc<Mutex<u32>>) -> Self {
        UserRoutes {
            users,
            next_id,
            error_handler: ErrorHandler::new(),
        }
    }

    pub fn handle(&self, method: &str, path: &str, request: &str) -> String {
        match method {
            "GET" => self.handle_get(path),
            "POST" => self.handle_post(path, request),
            "PUT" => self.handle_put(path, request),
            "DELETE" => self.handle_delete(path),
            _ => self.error_handler.method_not_allowed(),
        }
    }

    fn handle_get(&self, path: &str) -> String {
        if path == "/users" {
            self.get_all_users()
        } else if path.starts_with("/users/") {
            let id_str = &path[7..];
            match id_str.parse::<u32>() {
                Ok(id) => self.get_user_by_id(id),
                Err(_) => self.error_handler.bad_request(),
            }
        } else {
            self.error_handler.not_found()
        }
    }

    fn handle_post(&self, path: &str, request: &str) -> String {
        if path == "/users" {
            let body = self.extract_body(request);
            match self.parse_user_json(&body) {
                Some((name, email)) => self.create_user(name, email),
                None => self.error_handler.bad_request(),
            }
        } else {
            self.error_handler.not_found()
        }
    }

    fn handle_put(&self, path: &str, request: &str) -> String {
        if path.starts_with("/users/") {
            let id_str = &path[7..];
            match id_str.parse::<u32>() {
                Ok(id) => {
                    let body = self.extract_body(request);
                    match self.parse_user_json(&body) {
                        Some((name, email)) => self.update_user(id, name, email),
                        None => self.error_handler.bad_request(),
                    }
                }
                Err(_) => self.error_handler.bad_request(),
            }
        } else {
            self.error_handler.not_found()
        }
    }

    fn handle_delete(&self, path: &str) -> String {
        if path.starts_with("/users/") {
            let id_str = &path[7..];
            match id_str.parse::<u32>() {
                Ok(id) => self.delete_user(id),
                Err(_) => self.error_handler.bad_request(),
            }
        } else {
            self.error_handler.not_found()
        }
    }

    fn get_all_users(&self) -> String {
        let users = self.users.lock().unwrap();
        let users_json: Vec<String> = users
            .values()
            .map(|u| u.to_json())
            .collect();
        let body = format!("[{}]", users_json.join(","));
        self.ok_response(&body)
    }

    fn get_user_by_id(&self, id: u32) -> String {
        let users = self.users.lock().unwrap();
        match users.get(&id) {
            Some(user) => {
                let body = user.to_json();
                self.ok_response(&body)
            }
            None => self.error_handler.not_found(),
        }
    }

    fn create_user(&self, name: String, email: String) -> String {
        let mut next_id = self.next_id.lock().unwrap();
        let id = *next_id;
        *next_id += 1;
        drop(next_id);

        let user = User::new(id, name, email);
        let mut users = self.users.lock().unwrap();
        users.insert(id, user.clone());
        drop(users);

        let body = user.to_json();
        self.created_response(&body)
    }

    fn update_user(&self, id: u32, name: String, email: String) -> String {
        let mut users = self.users.lock().unwrap();
        match users.get_mut(&id) {
            Some(user) => {
                user.name = name;
                user.email = email;
                let body = user.to_json();
                self.ok_response(&body)
            }
            None => self.error_handler.not_found(),
        }
    }

    fn delete_user(&self, id: u32) -> String {
        let mut users = self.users.lock().unwrap();
        match users.remove(&id) {
            Some(_) => self.no_content(),
            None => self.error_handler.not_found(),
        }
    }

    fn extract_body(&self, request: &str) -> String {
        let lines: Vec<&str> = request.lines().collect();
        let mut body_start = false;
        let mut body_lines = Vec::new();

        for line in lines {
            if body_start {
                body_lines.push(line);
            } else if line.is_empty() {
                body_start = true;
            }
        }

        body_lines.join("\n").trim_matches('\0').to_string()
    }

    fn parse_user_json(&self, json: &str) -> Option<(String, String)> {
        let json = json.trim();
        if !json.starts_with('{') || !json.ends_with('}') {
            return None;
        }

        let content = &json[1..json.len()-1];
        let mut name: Option<String> = None;
        let mut email: Option<String> = None;

        for part in content.split(',') {
            let part = part.trim();
            if let Some(colon_pos) = part.find(':') {
                let key = part[..colon_pos].trim().trim_matches('"');
                let value = part[colon_pos+1..].trim().trim_matches('"');

                match key {
                    "name" => name = Some(value.to_string()),
                    "email" => email = Some(value.to_string()),
                    _ => {}
                }
            }
        }

        match (name, email) {
            (Some(n), Some(e)) => Some((n, e)),
            _ => None,
        }
    }

    fn ok_response(&self, body: &str) -> String {
        format!(
            "HTTP/1.1 200 OK\r\nContent-Type: application/json\r\nContent-Length: {}\r\n\r\n{}",
            body.len(),
            body
        )
    }

    fn created_response(&self, body: &str) -> String {
        format!(
            "HTTP/1.1 201 Created\r\nContent-Type: application/json\r\nContent-Length: {}\r\n\r\n{}",
            body.len(),
            body
        )
    }

    fn no_content(&self) -> String {
        "HTTP/1.1 204 No Content\r\n\r\n".to_string()
    }
}
