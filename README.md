# find-twitch-streamer
Golang backend with basic html / css frontend. You type in the twitch streamer's username. If they are live you are redirected to their stream. If they are not live, you are forwarded back to the home page.

Setup:

1) Clone the repo in am empty folder  --> change go.mod file to '{folder name}/find-twitch-streamer'
2) Create a .env file 
3) Read through the twitch API docs  --> create your twitch account --> navigate to 'twitch developer console' {https://dev.twitch.tv/console/apps} and registrer this app. {Use whatever name you want}
  3a) find your client id
    - can be found in your twitch developer dashboard
  3b) add http://localhost:500/home as an authorized app link
   - load the home page {localhost:500/home} --> {go run main.go}
   - paste your version of this url in the url bar: 
   https://id.twitch.tv/oauth2/authorize?response_type=token&client_id={add clinet id here}&redirect_uri=http://localhost:5500/home&scope=channel%3Amanage%3Apolls+channel%3Aread%3Apolls&state=c3ab8aa609ea11e793ae92361f002671
   
  3c) your access token credential should be returned in the url after you authenticatt with twitch
  - http://localhost:5500/home#access_token={access token is here}&scope=channel%3Amanage%3Apolls+channel%3Aread%3Apolls&state=c3ab8aa609ea11e793ae92361f002671&token_type=bearer

4) add your {Client ID} & {Access Token or Bearer Token} to your .env file 

5) all good to go




