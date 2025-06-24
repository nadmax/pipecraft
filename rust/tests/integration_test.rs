use crud_api::app::App;

#[test]
fn test_get_all_users_empty() {
    let app = App::new();
    let request = "GET /users HTTP/1.1\r\n\r\n";
    let response = app.route_request(request);

    assert!(response.contains("200 OK"));
    assert!(response.contains("[]"));
}

#[test]
fn test_post_user_and_get_by_id() {
    let app = App::new();
    let post_body = r#"{"name":"Alice","email":"alice@example.com"}"#;
    let request = format!(
        "POST /users HTTP/1.1\r\nContent-Length: {}\r\n\r\n{}",
        post_body.len(),
        post_body
    );
    let response = app.route_request(&request);
    assert!(response.contains("201 Created"));
    assert!(response.contains("\"name\":\"Alice\""));
    assert!(response.contains("\"email\":\"alice@example.com\""));

    let request = "GET /users/1 HTTP/1.1\r\n\r\n";
    let response = app.route_request(request);
    assert!(response.contains("200 OK"));
    assert!(response.contains("\"id\":1"));
}

#[test]
fn test_put_user() {
    let app = App::new();
    let post_body = r#"{"name":"Bob","email":"bob@example.com"}"#;
    let post_request = format!(
        "POST /users HTTP/1.1\r\nContent-Length: {}\r\n\r\n{}",
        post_body.len(),
        post_body
    );
    app.route_request(&post_request);

    let put_body = r#"{"name":"Bobby","email":"bobby@example.com"}"#;
    let put_request = format!(
        "PUT /users/1 HTTP/1.1\r\nContent-Length: {}\r\n\r\n{}",
        put_body.len(),
        put_body
    );
    let response = app.route_request(&put_request);

    assert!(response.contains("200 OK"));
    assert!(response.contains("\"name\":\"Bobby\""));
    assert!(response.contains("\"email\":\"bobby@example.com\""));
}

#[test]
fn test_delete_user() {
    let app = App::new();
    let body = r#"{"name":"Charlie","email":"charlie@example.com"}"#;
    let request = format!(
        "POST /users HTTP/1.1\r\nContent-Length: {}\r\n\r\n{}",
        body.len(),
        body
    );
    app.route_request(&request);

    let response = app.route_request("DELETE /users/1 HTTP/1.1\r\n\r\n");
    assert!(response.contains("204 No Content"));

    let response = app.route_request("GET /users/1 HTTP/1.1\r\n\r\n");
    assert!(response.contains("404 Not Found"));
}

#[test]
fn test_bad_request() {
    let app = App::new();
    let response = app.route_request("BAD REQUEST");

    assert!(response.contains("400 Bad Request"));
}

#[test]
fn test_not_found() {
    let app = App::new();
    let response = app.route_request("GET /unknown HTTP/1.1\r\n\r\n");

    assert!(response.contains("404 Not Found"));
}

#[test]
fn test_method_not_allowed() {
    let app = App::new();
    let response = app.route_request("PATCH /users HTTP/1.1\r\n\r\n");

    assert!(response.contains("405 Method Not Allowed"));
}
