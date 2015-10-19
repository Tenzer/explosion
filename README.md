# Explosion

Command line image viewer which supports 24-bit (16 million) colors, much based upon the post by [maato](http://softwarebakery.com/maato/image_in_terminal.html). The implementation in that post is however only with 256 colors while a [bunch of terminal emulators](https://gist.github.com/XVilka/8346728) has support for True Colors, so I had an attempt at reimplementing the same but with more colors.

That's also where the name comes from, as this is more of an explosion of colors, compared to maato's implementation.


## Screenshots

Taken in iTerm 2.9 on OS X:

![Lenna](screenshots/lenna.png)


## Installation

If you have [Go](https://golang.org/) installed, then you should simply run:

```
go get github.com/Tenzer/explosion
```

and you will get an `explosion` executable. There is currently no flags for the program, it just takes a list of images and prints them out in the same order.


## Change log

### 1.0.0 - 2015-10-19
* Initial release


## To do

* Detect the terminal window size and use that as the base of the image output size.
* Allow override of the output image size and resize interpolation function through flags.
* Attempt to implement maato's solution with the higher image resolution, by making use of the extra characters available.
