# Countdown App

This is a simple countdown app built with Go and Ebiten. The application allows you to set a countdown timer, which will be displayed in a floating window. Once the timer completes, the app will terminate. It's a lightweight utility that can be installed and run locally.

**Note**: 95% of this code was written with the assistance of ChatGPT.

## Features

![Example](screenshot.png)

- Set a countdown timer in minutes.
- Displays the remaining time in hours, minutes, and seconds.
- A transparent window with minimal UI, ideal for overlay use.
- The application is windowless (no borders, no title bar).
- Once the countdown is complete, the app will automatically exit.

### Install and Compile with Make

To build and install the app, simply run:

```bash
make
```

Install the app already build and present here:
```bash
make install
```

## Usage

put the desired minutes for the countdown as the first argument

```bash
countdown minutes
countdown 10
countdown 160
```
