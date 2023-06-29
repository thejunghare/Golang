## Methods
Methods are similar to the function in Golang with one difference that is methods contain a receiver argument in it, with help of this receiver arguments methods can access properties of receiver. *Receiver can be struct or non-stuct type*


## Syntax to declare **Methods**

```
func (reciver_name type ) method_name (parameter_list)  {}
```
         
## Methods with Pointer Receiver

```
func (p *type) method_name(...Type) Type {}
```
Note: *Go methods can accpet both value and pointer.*