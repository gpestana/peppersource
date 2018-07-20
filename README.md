# :fire: Pepper source :fire:

Pepper source is a secure and decentralized software distribution 
utility built on top of [IPFS](https://ipfs.io).

It allows software providers to quickly and securely share software with their
users without having to worry about infrastructure, security, notification
schemes and maintenance. In the future, it will also allow providers
to offer incentives to clients who allocate disk space and bandwidth by 
storing and distributing software bundles.

Pepper source exposes a single primitive which provides **storage**, **replication**,  
**security**, **notification mechanisms** and **incentives** for software
releases by leveraging the IPFS protocol stack and distributed file system.

### Security 

Pepper source provides a transparent mechanism for verifying content integrity 
of the software release and identity of the provider:

- The software binaries are signed by the software provider before uploading it
  to IPFS. The signature is performed using asymmetric crypto algorithms so that
the client can verify the provenience of the software bundle. This process is
all transparent to the provider.

- Binaries are stored based on their cryptographic hashes, making it easy for
clients to verify the integrity of the software, even when coming from unknown
sources (e.g. untrusted peers which are not software providers).

### Notification

The provider can notify its clients that new software bundles (releases,
patches, updates, ...) are available through a p2p
[publish/subscription](https://en.wikipedia.org/wiki/Publish%E2%80%93subscribe_pattern) 
mechanism. The notification contains the hash (content address) of the newly 
released bundle and metadata about the software bundle being published. With 
this notification mechanism in place, those interested in the software 
(subscribers) know *when* freshly baked software is ready for download and 
*where from*.

### Replication

Anyone can store and serve copies of the bundled software.  This is done through
the
[pinning](https://ipfs.io/ipfs/QmTkzDwWqPbnAh5YiV5VwcTLnGdwSNsNTn2aDxdXBFca7D/example#/ipfs/QmQwAP9vFjbCtKvD8RkJdCvPHqLQjZfW7Mqbbqx18zd8j7/pinning/readme.md) IPFS primitives and opens up interesting 
opportunities to build incentives so that the provider can rely on its users to
securely store and share replicas of software bundles.

### How to use it (CLI)

```go
// TODO
```

### More cool stuff coming up:

- **Client side interfaces**: SDK for software providers to embed to their
  client side applications which automatically fetches, verifies and apply 
bundles of software released by the provider.

- **Binary encryption**: allow only certain users to access the binaries;

- **Replication incentices**: by providing monetary (or other types) incentives,
  the provider can make sure its software bundles are 


:fire: :fire: :fire:
