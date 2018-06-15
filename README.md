# Assignment  
* Upon execution without any command-line parameters, display the top 20 stories from Hacker
News (https://news.ycombinator.com) to stdout in a readable format.  

* When executed with a -csv parameter, save the above results to a CSV file rather than sending
it to stdout.  

# Solving
* Thank you for assignment and your time first.  

* I did this assignment with simple pipeline  
    * get the 20 top stories from HA
    * make goroutine with channals for communicating.
    * write stdout or create CSV through streaming.

# Useage  
    * git clone https://github.com/YoungsoonLee/bpi.git
    * cd bpi
    * go build .
    * run) ./bpi or ./bpi -csv
    * test) go test -run=nope ./...

# Answer
* Does the resulting code build successfully?  
    > Yes  

* What is the test coverage, and do tests pass? What about go test -race ?  
    > Yes I tested funtions.  
    >    Especailly, tested pipeline funtion(./getnews/getnews.go/process function) with race condition.  
    >    It can check with run go test -run nope ./... -race  
    >    all passed

* Is the code readable? Are confusing bits commented sufficiently?  
    > Yes, I add comments sufficiently.  

* What do go fmt and go lint say about the code?  
    > Go fmt, Go lint  are a tool that automatically formats Go source code. 
    > It make code to easier to write, easier to read, easier to maintain.  

* Given the requirements, what is the approach to the application design?  
    > I always worry, design, test, and write programs for good performance.  
    > In this requirements, I wrote the pipeline code using goroutine and channel for fast stdout.  
* Is the code easy to extend? Is the code reusable? If asked to instead display the top 20 stories  
  from r/golang would it be a 20 minute tweak or a three hour rewrite?  
  > I Think It is easy to extend. If you want to extend, just add some const value in ./utils/consts.go like fixed url  
  > and then, I wrote to the code with switch in core function, So write to function for extending simply.  
  > and then can call start function with the category for extended function.  

* Does it scrape the page, or rely on the HN API?  
  > I used HA API. Because I love open source or opened API. I think this is part of reusability.  

* Is it actually written in Perl, or mining bitcoin?
  > No, But If need it, I can do it~!  Thank you ~!  
