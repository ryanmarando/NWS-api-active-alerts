# Ryan Marando Meteorologist Portfolio
Besides showcasing Ryan Marando's broadcasting skills, this project provides a streamlined interface to access active weather alerts from the National Weather Service (NWS) API. It enables users to receive personalized and localized alerts, as well as wet bulb globe temperature forecasts.

## Features
- **Personalized Alerts**: Receive weather alerts tailored to your specific location.
- **Localized Data**: Access detailed weather information pertinent to your area.
- **Wet Bulb Globe Temperature Forecasts**: Obtain forecasts that consider temperature, humidity, wind speed, and solar radiation.

## Prerequisites
Before setting up the project, ensure you have the following installed:
- Node.js (version 14.x or higher)
- [Node.js](https://nodejs.org/en) (version 14.x or higher)
- [npm](https://www.npmjs.com/) (version 6.x or higher)

## Installation
### Clone the Repository:
```
git clone https://github.com/ryanmarando/NWS-api-active-alerts.git
cd NWS-api-active-alerts
```
### Install Dependencies:
Navigate to the project directories and install the necessary packages:
```
# For the API
cd weather-alert-system-API
npm install

# For the Application
cd ../weather-alert-system-app
npm install
```
## Configuration
### API Configuration:
1. Create a  `.env` file in the `weather-alert-system-API` directory.
2. Update the environment variables as needed, such as setting the port number or API keys.

### Application Configuration:
1. Create a  `.env` file in the `weather-alert-system-app` directory.
2. Configure the environment variables, including the API endpoint and any other necessary settings.

## Usage
### Start the API Server:
```
cd weather-alert-system-API
npm start
```
The API server should now be running, ready to handle requests.

### Start the Application:
```
cd ../weather-alert-system-app
npm start
```
The application will launch, allowing you to interact with the weather alert system.

## National Weather Service API
This project utilizes the National Weather Service (NWS) API to fetch active weather alerts. The NWS API provides access to critical forecasts, alerts, and observations, along with other weather data. For more information, refer to the [NWS API Documentation](https://www.weather.gov/documentation/services-web-api).

## Acknowledgements
[National Weather Service](https://www.weather.gov/) for providing the weather data API.
Ryan Marando for developing this project.
For more information, visit the [project repository](https://github.com/ryanmarando/NWS-api-active-alerts).
