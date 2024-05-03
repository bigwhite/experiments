pub mod bar;
pub mod foo;

pub fn func1() {
    println!("called {}::func1()", module_path!());
    foo::func1();
    bar::func1();
}
