## **Concurrency**

Concurrency capability to deal with lots of thigs at once.

Concurrency is just an illusion of parallelism.

## **Parallelism**

Parallelism is doing lots of things at the same time.

## **Technical example**

We are programming a web browser that has two comppnents

1. Downloading files from internet
2. rendering pages user requested

Lets assume that we have structured our browser's code in such a way that each of these components can be executed independtly. In Golang this is achieved using goroutines.

When this browser is run the processor will context switch between two components of the browser. It might download a file for some time and then might switch to render the html user requested for sometime. *This is called Concurrency.*

Same browser is running on multicore processor, then this 2 components might run simultaneously in different cores *Knows as parallelism.*

*Concurrency is handled in golang using Goroutines and Channels.*

## **Goroutines**
Goroutines are functions and methods that run simultaneously with other functions or method.

*Cost of creating a goroutine is tiny when compared to a thread.*

