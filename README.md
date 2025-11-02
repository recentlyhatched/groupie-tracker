# Groupie Tracker

## Usage

### `git clone <git-remote-url>`

### `cd groupie-tracker`

### `go run .`
Server will run on port 8080 e.g [http://localhost:8080](http://localhost:8080)\
or... for ChromeOS users, [http://penguin.linux.test:8080](http://penguin.linux.test:8080)

## About the project

### 🧭 About Groupie Tracker

In this project, I worked with a provided API and manipulated the data it contained to build a dynamic, user-friendly website that displays information about musical artists and bands.

The API is divided into four main parts:

artists — includes details such as the band or artist’s name, image, formation year, first album date, and members.

locations — lists their most recent or upcoming concert locations.

dates — contains the corresponding concert dates.

relation — links together all the other parts (artists, dates, and locations).

Using this data, I designed a website that visually presents each band’s information in an organized and engaging way. I used different types of data visualization elements such as cards, lists, and tables to make the browsing experience intuitive and clear.

### ⚙️ Event and Server Interaction

A key part of this project involved implementing an event/action system that demonstrates client–server interaction. I created a feature where a client request triggers an event on the server, which processes the request and sends back the appropriate response.

This allowed me to explore how web applications handle user-triggered events — where the system reacts dynamically to actions initiated by the user.