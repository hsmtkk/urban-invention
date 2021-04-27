mod beersong;

fn main() {
    let mut n = 100;
    loop {
        let phrase = beersong::get_phrase(n);
        println!("{}", phrase);
        n -= 1;
        if n == 0 {
            break;
        }
    }
}
