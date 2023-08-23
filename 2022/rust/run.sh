rm ./target/release/adv-of-code
cargo build --release &&
	clear &&
	./target/release/adv-of-code $@

