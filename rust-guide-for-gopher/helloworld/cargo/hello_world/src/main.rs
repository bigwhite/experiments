use std::io;
use std::io::Write;

fn main() {
    let mut output = io::stdout();
    output.write(b"Hello, World!").unwrap();
    output.flush().unwrap();
}
