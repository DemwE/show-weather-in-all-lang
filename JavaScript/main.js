require('dotenv').config({ path: '../.env' });
const axios = require('axios');

const apiKey = process.env.API_KEY;
const city = "Warsaw";

axios.get(`http://api.openweathermap.org/data/2.5/weather?q=${city}&appid=${apiKey}`)
    .then(response => {
        const tempCelsius = response.data.main.temp - 273.15;
        console.log(`Current temperature in ${city}: ${tempCelsius.toFixed(2)} degrees Celsius`);
    })
    .catch(error => {
        console.error(`Error getting weather data: ${error}`);
    });