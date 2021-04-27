pub fn get_phrase(n: i64) -> String {
    if n == 0 {
        return "No more bottles of beer on the wall, no more bottles of beer.
Go to the store and buy some more, 99 bottles of beer on the wall."
            .to_string();
    }
    if n == 1 {
        return "1 bottle of beer on the wall, 1 bottle of beer.
Take it down and pass it around, no more bottles of beer on the wall."
            .to_string();
    }
    if n == 2 {
        return "2 bottles of beer on the wall, 2 bottles of beer.
Take one down and pass it around, 1 bottle of beer on the wall."
            .to_string();
    }
    return format!(
        "{} bottles of beer on the wall, {} bottles of beer.
Take one down and pass it around, {} bottles of beer on the wall.",
        n,
        n,
        n - 1
    );
}

#[test]
fn test0() {
    let want = "No more bottles of beer on the wall, no more bottles of beer.
Go to the store and buy some more, 99 bottles of beer on the wall.";
    let got = get_phrase(0);
    assert_eq!(want, got);
}

#[test]
fn test3() {
    let want = "3 bottles of beer on the wall, 3 bottles of beer.
Take one down and pass it around, 2 bottles of beer on the wall.";
    let got = get_phrase(3);
    assert_eq!(want, got);
}
