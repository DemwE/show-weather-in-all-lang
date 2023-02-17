import os
import requests

from dotenv import load_dotenv
dotenv_path = os.path.join(os.path.dirname(__file__), '../.env')
load_dotenv(dotenv_path)

API_KEY = os.getenv('API_KEY')
city = "Warsaw"

url = f'http://api.openweathermap.org/data/2.5/weather?q={city}&appid={API_KEY}&units=metric'
response = requests.get(url)

data = response.json()
print(f'The temperature in {city} is {data["main"]["temp"]}°C.')
