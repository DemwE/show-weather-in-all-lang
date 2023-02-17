require 'dotenv/load'
require 'net/http'
require 'json'

Dotenv.load('../.env')
city = "Warsaw"
API_KEY = ENV['API_KEY']

url = "http://api.openweathermap.org/data/2.5/weather?q=#{city}&appid=#{ENV['API_KEY']}&units=metric"
uri = URI(url)
response = Net::HTTP.get(uri)
data = JSON.parse(response)

puts "Weather in #{city}: #{data['weather'][0]['description']}. Temperature: #{data['main']['temp']}°C"
