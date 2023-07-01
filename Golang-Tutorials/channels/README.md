## Channels

Channels can be thought of pipes using which goroutines communicate. Data can be spent from one end to and received from the other end using channels.

Zero value of a channel is nil, nil channels are not of any use and hence the channel has to be defined using make.

## Types of channel
1. Bidirectional - channels that can both sent and received on them.
2. Unidirectional - channels that only send or receive data.

*It is possible to convert a bidirectional channel to a send only or receive only channel but not the vice versa.*


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

## Closing a Channel

You can aslo close a channel with help of close() function.

```
close(Mychannel)
```

## Length of Channel

```
len(Mychannel)
```

## Capacity of Channel

```
cap(Mychannel)
```

## Deadlock

One important factor to consider while using channels is deadlock. If a Goroutine is sending data on a channel, then it is expected that some other Goroutine should be receiving the data. If this does not happen, then the program will panic at runtime with Deadlock.

## Closing channels and for range loops on channels

sender have the ability to close the channel to notify receivers that no more data will be sent on the channel.

Receivers can use an additional variable while receiving data from the channel to check wether the channel has been closed.

```
v, ok := <- ch
```

