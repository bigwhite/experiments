extern crate rand;
use rand::Rng;

mod sub1;
mod sub2;

mod sub3 {
    pub fn func1() {
        println!("called {}::func1()", module_path!());
    }
    pub fn func2() {
        self::func1();
        println!("called {}::func2()", module_path!());
        super::func1();
    }
}

fn func1() {
    println!("called {}::func1()", module_path!());
}

fn main() {
    println!("current module: {}", module_path!());
    let mut rng = rand::thread_rng();
    let num: u32 = rng.gen();
    println!("Random number: {}", num);

    sub1::func1();
    sub2::func1();
    sub3::func2();
}
