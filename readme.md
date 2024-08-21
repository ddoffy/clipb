# A tool help you to share clipboard between devices

## Features

- Share clipboard between devices using Redis
- Support Windows, MacOS, Linux
- Support text
-

## How to use

1. Install Redis on your server
2. Open Redis port on your server that you installed Redis to LAN (default port
   is 6379) for the other devices in the same network to connect to it
3. Install this tool on your devices
4. Run this tool on your devices

## How to install

1. Clone this repository
2. Run `make [operating system name]` to build the tool. .e.g `make linux`, `make
window`, `make macos`
3. Run the tool in the `bin` folder
