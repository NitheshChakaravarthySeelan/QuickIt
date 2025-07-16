#[tokio::test]
async fn test_health() {
    let client = reqwest::Client::new();

    let res = client
        .get("http://localhost:3000/health")
        .send()
        .await
        .unwrap()
        .json::<serde_json::Value>()
        .await
        .unwrap();

    assert_eq!(res["status"], "ok");
}
