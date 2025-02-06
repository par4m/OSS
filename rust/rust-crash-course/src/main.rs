// use futures::executor::block_on;

use std::process::Output;

use futures::Future;
use tokio::time::{sleep, Duration};

fn call_api_one() -> impl Future<Output = String> {
    async {
        sleep(Duration::from_secs(1)).await;
        "One".to_string()
    }
}
fn call_api_two() -> impl Future<Output = String> {
    async {
        sleep(Duration::from_secs(1)).await;
        "Two".to_string()
    }
}

#[tokio::main]
async fn main() {
    let one = call_api_one().await;
    let two = call_api_two().await;
    println!("{:?}", one);
    println!("{:?}", two);
}
