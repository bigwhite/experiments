pub fn hello_from_b() {
    println!("Hello from B begin");
    C::hello_from_c();
    println!("Hello from B end");
}
