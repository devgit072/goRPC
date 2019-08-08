Tutorials how to implement RPC in several ways.

<h5>*First of all how RPC is different from normal function call?*</h5>
Normal function call happens when method and caller are present in same address space.  
But when method is not present is same address space, then we need to do RPC over the network.
To make RPC succesfull we need Client and Server: Client will establish connection
to Server and will tell Server that I need to call this method and 
then Server will serve the request.
Client can create TCP connection or create HTTP connection which is on top of TCP only.

<h5>How RPC are different from Web-Service?</h5>
Web-Service is higher level implementation of RPC. Here method is accessed via 
Hypertext Transfer Protocol.

<H5>RPC in Golang</H5>
RPC is Golang can be well implemented by gRPC. But even using net/rpc package in 
Golang you can create RPC server/client. Though in presence of gRPC, there is 
no need to implement RPC using native library(net/rpc), BUT just for learning 
purpose we are trying here to implement RPC using native library.

<h5>RPC based on HTTP vs TCP</h5>
So RPC needs transport mechanism: it can either HTTP or TCP. The advantage of using HTTP 
is that it can leverage HTTP support library.
We will implement both HTTP and TCP based RPC.

<h5>Repository Details</h5>
1) simpleRPC : It contains basic implementation of RPC based on both TCP vs HTTP. 
