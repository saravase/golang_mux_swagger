# golang_mux_swagger

<h2>Golang:</h2>
Go is a statically typed, compiled programming language designed at Google by Robert Griesemer, Rob Pike, and Ken Thompson. Go is syntactically similar to C, but with memory safety, garbage collection, structural typing, and CSP-style concurrency

<h2>REST API</h2>
REST is acronym for REpresentational State Transfer. It is architectural style for distributed hypermedia systems and was first presented by Roy Fielding in 2000 in his famous dissertation.

<h3>Guiding Principles of REST</h3>
<ol>
  <li>
Client–server – By separating the user interface concerns from the data storage concerns, we improve the portability of the user interface across multiple platforms and improve scalability by simplifying the server components.
  </li><li>
Stateless – Each request from client to server must contain all of the information necessary to understand the request, and cannot take advantage of any stored context on the server. Session state is therefore kept entirely on the client.
  </li><li>
Cacheable – Cache constraints require that the data within a response to a request be implicitly or explicitly labeled as cacheable or non-cacheable. If a response is cacheable, then a client cache is given the right to reuse that response data for later, equivalent requests.
  </li><li>
Uniform interface – By applying the software engineering principle of generality to the component interface, the overall system architecture is simplified and the visibility of interactions is improved. In order to obtain a uniform interface, multiple architectural constraints are needed to guide the behavior of components. REST is defined by four interface constraints: identification of resources; manipulation of resources through representations; self-descriptive messages; and, hypermedia as the engine of application state.
  </li><li>
Layered system – The layered system style allows an architecture to be composed of hierarchical layers by constraining component behavior such that each component cannot “see” beyond the immediate layer with which they are interacting.
  </li><li>
Code on demand (optional) – REST allows client functionality to be extended by downloading and executing code in the form of applets or scripts. This simplifies clients by reducing the number of features required to be pre-implemented.
  </li>
</ol>



<h2>Used Packages:</h2>
<ul>
  <li>mux</li>
  <li>validator</li>
  <li>middleware</li>
  <li>json</li>
  <li>context</li>
  <li>os/signal</li>
  <li>net/http</li>
  <li>os</li>
  <li>regexp</li>
  <li>strconv</li>
  <li>time</li>
  <li>os</li>
</ul>

<h2>Implemented REST Resource Method through Gorilla Mux</h2>
<ol>
  <li>GET</li>
  <li>POST</li>
  <li>PUT</li>
  <li>DELETE</li>
</ol> 


<h2>Swagger Environment Setup:</h2>

<ol>
<li>Install Swagger :</br>  go get -u github.com/go-swagger/go-swagger/cmd/swagger </li>
<li>Initialize Swagger Meta Data :</br>  swagger:meta</li>
<li>Generate swagger.yaml:</br>  GO111MODULE=off swagger generate spec -o ./swagger.yaml --scan-models </li>
<li>Swagger Docs Endpoint:</br> http://localhost:9090/docs </li>
</ol> 

