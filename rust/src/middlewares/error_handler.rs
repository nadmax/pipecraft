pub struct ErrorHandler;

impl ErrorHandler {
    pub fn new() -> Self {
        ErrorHandler
    }

    pub fn bad_request(&self) -> String {
        let body = r#"{"error":"Bad Request"}"#;
        format!(
            "HTTP/1.1 400 Bad Request\r\nContent-Type: application/json\r\nContent-Length: {}\r\n\r\n{}",
            body.len(),
            body
        )
    }

    pub fn not_found(&self) -> String {
        let body = r#"{"error":"Not Found"}"#;
        format!(
            "HTTP/1.1 404 Not Found\r\nContent-Type: application/json\r\nContent-Length: {}\r\n\r\n{}",
            body.len(),
            body
        )
    }

    pub fn method_not_allowed(&self) -> String {
        let body = r#"{"error":"Method Not Allowed"}"#;
        format!(
            "HTTP/1.1 405 Method Not Allowed\r\nContent-Type: application/json\r\nContent-Length: {}\r\n\r\n{}",
            body.len(),
            body
        )
    }

    pub fn internal_server_error(&self) -> String {
        let body = r#"{"error":"Internal Server Error"}"#;
        format!(
            "HTTP/1.1 500 Internal Server Error\r\nContent-Type: application/json\r\nContent-Length: {}\r\n\r\n{}",
            body.len(),
            body
        )
    }
}
