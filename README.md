# Generaidor

A Static Site Generator, currently built with PHP. The name is a bit of wordplay, but it accidentally matches with the YugiOh card.


## Motivation & Spirit
Every project is built with some sort of motivation behind it, as it exists for a reason. For example, Go exists because Go's founders were frustrated with the languages they were using at Google. And Generaidor is no different. This project was built because I want to learn and I want to complete something on my own.

## Installation
Make sure you have PHP installed on your local machine, you can check it by checking the PHP version:
```bash
php --version
```

If you don't receive a messages such as below, then look for an online guide to install PHP on your machine.
```
PHP 8.3.6 (cli) (built: Jan 01 2025 10:08:38) (NTS)
Copyright (c) The PHP Group
Zend Engine v4.3.6, Copyright (c) Zend Technologies
with Zend OPcache v8.3.6, Copyright (c), by Zend Technologies
```

Pull the project to your local workspace, then make sure you set the bin file to be executable. After that, you can run it:
```bash
# Clone project
git clone https://github.com/cohekoma/generaidor.git

# Move into directory then set bin file to be executable
cd generaidor
sudo chmod +x generaidor

# Try to run it
./generaidor
```

And that's it!

## Folder Structure
Below is the explanation of what role each folder plays in this project. It will be updated.
```
generaidor
|   README.md           # Information about project
|   generaidor.php      # The entry file that can be executed.
|   src                 # Project source
|   dist                # The static outcomes
|   templates           # All layout templates
|   content             # Post content written in markdown
|   assets              # Static assets such as CSS, JS, images or font
```