# Problem Set

## 5.1
Change the `findlinks` program to traverse the `n.FirstChild` linked list using recursive calls to `visit` instead of a loop.

## 5.2
Write a function to populate a mapping from element names - `p`, `div`, `span`, and so on - to the number of elements with that name is an HTML document tree.

## 5.3
Write a function to print the contents of all text nodes in an HTML document tree. Do not descend onto `<script>` or `<style>` elements, since their contents are not visible in a web browser.

## 5.4
Extend the `visit` function so that it extracts other kinds of links from the document, such as images, scripts, and style sheets. 

## 5.5
Implement `countWordsAndImages`. (See Exercise 4.9 for word-splitting.)

## 5.6
Modify the `corner` function in `gopl.io/ch3/surface` (3.2) to use named results and a bare return statement.

## 5.7
Develop `startElement` and `endElement` into a general HTML pretty-printer. Print comment nodes, text nodes, and the attributes of each element (`<a html='...'>`). Use short forms like `<img/>` instead of `<img></img>` when an elemnt has no children. Write a test to ensure that the output can be parsed successfully. (See Chapter 11).

## 5.8
Modify `forEachNode` so that the pre and post functions return a boolean result indicating whether to continue the traversal. Use it to write a function `ElementByID` with the following signature that finds the first HTML element with the specified id attribute. The function should stop the traversal as soon as a atch is found.

```
func ElementByID(doc *html.Node, id string) *html.Node
```

## 5.9
Write a function `expand(s string, f func(string) string) string` that replaces each substring "$foo" within `s` by the text returned by `f("foo")`.

## 5.10
Rewrite `topoSort` to use maps instead of slices and eliminate the initial sort. Verify the result, though nondeterminisic, are valid topological orderings.

## 5.11
The instructor of a linear algebra course decides that calculus is now a prerequisite. Extend `topoSort` function to report cycles.

## 5.12
The `startElement` and `endElement` functions in `gopl.io./ch5/outlines2` (5.5) share a global varialble, `depth`. Turn them into anonymous functions that share a variable local to the `outline` function.

## 5.13
Modify `crawl` to make local copies of the pages it finds, creating directories as necessay. Don't make copies of pages that come from a differnt domain. For example, if the original page comes from `golang.org` save all files from there, but exclude ones from `vimeo.org`.

## 5.14
Use the `breadthFirst` function to explore a different structure. For example, you could use the course dependencies from the `topoSort` example (a directed graph), the file system hierarchy on your computer (a tree), or a list of bus or subway routes downloaded from your city government's web site (an undirected graph).

## 5.15
Write variadic functions `max` and `min`, analogous to `sum`. What should these functions do when called with no arguments? Write variants that require at least one argument.

## 5.16
Write a variadic function to `strings.Join`.

## 5.17
Write a variadic function `ElementsByTagName` that, given an HTLM node tree and zero or more names, returns all the elements that match one of those names. Here are two example calls:

```
func ElementsByTagName(doc *html.Nodes, name ...string) []*html.Node

images := ElementsByTagName(doc, "img")
headings := ElementsByTagName(doc, "h1", "h2", "h3", "h4")
```

## 5.18
Without changing its behavior, rewrite the `fetch` functions to use `defer` to close the writable file.

## 5.19
Use `panic` and `recover` to write a function that contains no `return` statement yet returns a non-zero value. 