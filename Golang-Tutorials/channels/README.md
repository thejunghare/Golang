## Channels

Channels can be thought of pipes using which goroutines communicate. Data can be spent from one end to and received from the other end using channels.

Zero value of a channel is nil, nil channels are not of any use and hence the channel has to be defined using make.

## declaring channels

```
var Channel_name chan Type

<!-- using make keyword -->
channel_name := make(chan Type)
```

Channel work with two principal operations

1. sending
2. receving

This both operations collectively known as communication.
The direction of <- operator indicates whether the data is received or send.

## Send operation

The send operation is used to send data from one goroutine to another goroutine with help of channel.

Values like int, float64, and bool are safe and easy to send, whereas sending pointers or reference like a slice, map, etc aren't secure.

```
Mychannel <- element
```

The above statement indicates that the data(element) send to the channel with help of a <- operator.

## Receive operation

The receive operation is used to receive the data sent to by the send operator.

```
elemnt := <-Mychannel
```
