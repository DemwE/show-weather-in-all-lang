use std::env;

#[derive(serde::Deserialize)]
struct Main {
    temp: f32,
}

#[derive(serde::Deserialize)]
struct WeatherData {
    main: Main,
}

#[tokio::main]
async fn main() -> Result<(), Box<dyn std::error::Error>> {
    dotenv::dotenv().ok();
    let api_key = env::var("API_KEY").expect("API key is not set");

    let city = "Warsaw";

    let url = format!(
        "https://api.openweathermap.org/data/2.5/weather?q={}&appid={}&units=metric",
        city, api_key
    );

    let response = reqwest::get(&url).await?;

    let response_text = response.text().await?;

    let weather_data: WeatherData = serde_json::from_str(&response_text)?;

    println!("Temperature in {} is {}°C", city, weather_data.main.temp);

    Ok(())
}
