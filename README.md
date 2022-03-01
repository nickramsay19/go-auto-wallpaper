# Get Nature Wallpaper | go-auto-wallpaper
> Developed by Nicholas Ramsay thanks to Unsplash.com image API.

A very simple program that retreives a nature image of the correct orientation for my desktop wallpaper.

## Compilation
1. You are required to get an API key from [api.unsplash.com](http://api.unsplash.com). Once done, place it in a file "secret.json" in the format:
```json
{
     "api_key": "YOUR_ACCESS_KEY_HERE"
}
```
2. Simply run `sh compile.sh` (assuming you have Go version 1.17 or later installed).


## Project Reflection
* Despite my best efforts to avoid using JSON (instead of YAML) for storing the secret api key I eventually conceded in large part due to an aversion to using a third-party library for parsing other formats. JSON notably has parsing tools in the Golang standard library. 
* Initially, I was to programatically change the user's desktop wallpaper to the downloaded image. I decided against this as interfacing with the operating system to change the wallpaper is different on each system. The program without this feature works on all systems. Instead I will write my own local shell script to execute the program then set my wallpaper to the image downloaded.
