# Problem Set

## 1.1
Modify the ```echo``` program to also print ```os.Args[0]```, the name of the command that invoked it.

## 1.2
Modify the ```echo``` program to print the index and value of each of its arguments, one per line.

## 1.3
Experiment to measure the difference in running time between our potentially inefficient versions and the one that uses ```strings.Join```. (Section 1.6 illustrates part of the time package, and Section 11.4 shows how to write benchmark tests for systematic performance evaluation.)

## 1.4
Modify ```dup2``` to print the names of all files in which each duplicated line occurs.

## 1.5
Change the Lissajous program’s color palette to green on black, for added authenticity. To create the web color ```#RRGGBB```, use ```color.RGBA{0xRR, 0xGG, 0xBB, 0xff}```, where each pair of hexadecimal digits represents the intensity of the red, green, or blue component of the pixel.

## 1.6
Modify the Lissajous program to produce images in multiple colors by adding more values to palette and then displaying them by changing the third argument of ```SetColorIndex``` in some interesting way.

## 1.7
The function call ```io.Copy(dst, src)``` reads from ```src``` and writes to ```dst```. Use it instead of ```ioutil.ReadAll``` to copy the response body to ```os.Stdout``` without requiring a buffer large enough to hold the entire stream. Be sure to check the error result of ```io.Copy```.

## 1.8
Modify ```fetch``` to add the prefix ```http://``` to each argument URL if it is missing. You might want to use ```strings.HasPrefix```.

## 1.9
Modify ```fetch``` to also print the HTTP status code, found in ```resp.Status```.

## 1.10
Find a web site that produces a large amount of data. Investigate caching by running ```fetchall``` twice in succession to see whether the reported time changes much. Do you get the same content each time? Modify ```fetchall``` to print its output to a file so it can be examined.

## 1.11
Try ```fetchall``` with longer argument lists, such as samples from the top million web sites available at ```alexa.com```. How does the program behave if a web site just doesn’t respond? (Section 8.9 describes mechanisms for coping in such cases.)

## 1.12
Modify the Lissajous server to read parameter values from the URL. For example, you might arrange it so that a URL like ```http://localhost:8000/?cycles=20``` sets the number of cycles to 20 instead of the default 5. Use the ```strconv.Atoi``` function to convert the string parameter into an integer. You can see its documentation with go doc ```strconv.Atoi```.
