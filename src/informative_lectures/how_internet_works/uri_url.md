# URI / URL

## How Internet Works

Internet is a global network of interconnected computers that communicate using standardized protocols.

**Key Components**
- Clients and Servers
- Protocols
- IP Addresses
- Domain Name Systems(DNS)

## A Web Request's Journey 

- Step-1: Entering a URL
- Step-2: DNS Lookup, DNS Server Interaction
- Step-3: Establicshing a TCP connection
    - browser sends a TCP SYN (synchronize) packet to the server.
    - server responds with a SYN-SCK (synchronize-acknowledgement) packet.
    - browser sends an ACK (acknowledgement) packet, completing the three-way handshake.
- Step-4: Sending an HTTP Request
- Step-5: Server Processing and Response
- Step-6: Rendering the Webpage


## URI & URL

### URI (Uniform Resource Locator)

Components:
- URL (Uniform Resource Locator)
- URN (Uniform Resource Name)

**Components of a URL**
- Scheme
- Host
- Port
- Path
- Query
- Fragment