# PrimitiveGif
#### Input:
![input image](https://github.com/Genji-MS/PrimitiveGif/blob/main/SampleImage/cat.png)
#### Output:
![output_image](https://github.com/Genji-MS/PrimitiveGif/blob/main/SampleImage/cat.gif)

> Uses the primitive library to automate creating gifs

Primitive takes in a single image (.png) and recreates the image using primitive shapes. The process is random and creating multiple images is a nice effect for a gif. This script automates the process of creating 8 primitive images, and then converts them into a gif.

## Installing / Getting started

This script requires primitive be installed.
(Primitive has no public class for import)

```shell
git clone https://github.com/Genji-MS/PrimitiveGif.git
go get -u github.com/fogleman/primitive
go test
```

The test script will ensure that primitive has been installed correctly
If the test fails, follow the instructions in the error message and re test

To run the script, place any number of .png files in the local directory and use

```shell
go build && ./main -dir=.
```

The completed gif will be written into the same directory as <filename>.gif

## Links

- Fogleman primitive: https://github.com/fogleman/primitive
