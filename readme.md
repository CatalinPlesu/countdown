# Countdown App

This is a simple countdown app built with Go and Ebiten. The application allows you to set a countdown timer, which will be displayed in a floating window. Once the timer completes, the app will terminate. It’s a lightweight utility that can be installed and run locally.

**Note**: 95% of this code was written with the assistance of ChatGPT.

## Features

![Example](screenshot.png)

- Set a countdown timer in minutes.
- Displays the remaining time in hours, minutes, and seconds.
- A transparent window with minimal UI, ideal for overlay use.
- The application is windowless (no borders, no title bar).
- Once the countdown is complete, the app will automatically exit.

## Prerequisites

- **Go** (1.18 or newer)
- **Ebiten** (for window management and rendering)
- A font like **DejaVu Sans** (or another font of your choice)

### Install Go

You can download and install Go from the [official website](https://golang.org/dl/).

### Install Dependencies

Make sure you have the required dependencies by running:

```bash
go mod tidy
