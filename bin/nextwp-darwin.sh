/Users/nickramsay/Documents/Projects/go-auto-wallpaper/bin/getwp /Users/nickramsay/Documents/Projects/go-auto-wallpaper/bin/next.png
echo Setting wallpaper 
sqlite3 /Users/nickramsay/Library/Application\ Support/Dock/desktoppicture.db "UPDATE data SET value = \"/Users/nickramsay/Documents/Projects/go-auto-wallpaper/bin/next.png\""
echo Resetting dock
killall Dock