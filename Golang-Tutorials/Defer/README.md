## Defer

The Defer keyword is used to enusre that the function call is performed later in a program's execution, ususally for purpose of cleanup.

## Example 

Assuming that you are engaged in a task that requires the completion of certain tasks, such as opening a file, reading it, performing required operations, and closing it. Using the defer keyword in conjunction with the close function will ensure that the file is closed after all the operations have been completed.