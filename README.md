# Generaidor

A Static Site Generator, currently built with Go. The name is a bit of wordplay, but it accidentally matches with the YugiOh card.


## Motivation & Spirit
Every project is built with some sort of motivation behind it, as it exists for a reason. For example, Go exists because Go's founders were frustrated with the languages they were using at Google. And Generaidor is no different. This project was built because I want to learn and I want to complete something on my own.

## Installation

- Ensure you have Go installed (tested with go1.25.1).
- Pull the project to your local machine.
- `cd` into it.
- Run `go build .`
- Run the executable binary file `./generaidor`.


And that's it!

## Folder Structure
Below is the explanation of what role each folder plays in this project. It will be updated.
```
generaidor
|   README.md           # Information about project
|   public              # The outcome of the generated static content to serve to the public
|   src                 # Project source code
|   templates           # All layout templates
|   content             # Post content written in markdown
|   assets              # Static assets such as CSS, JS, images or font
```