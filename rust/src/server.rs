use std::net::TcpListener;
use std::sync::Arc;
use std::thread;
use crate::app::App;

pub struct Server {
    app: Arc<App>,
}

impl Server {
    pub fn new() -> Self {
        Server {
            app: Arc::new(App::new()),
        }
    }

    pub fn run(&self) {
        let listener = TcpListener::bind("127.0.0.1:8003").unwrap();
        println!("Server running on http://localhost:8003");

        for stream in listener.incoming() {
            match stream {
                Ok(stream) => {
                    let app_clone = Arc::clone(&self.app);
                    thread::spawn(move || {
                        app_clone.handle_request(stream);
                    });
                }
                Err(_) => {}
            }
        }
    }
}
