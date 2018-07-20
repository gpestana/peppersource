## :fire: Pepper source :fire:

:fire: Pepper source :fire: is a secure and decentralized software distribution 
utility built on top of [IPFS](https://ipfs.io).

It allows software providers to quickly and securely share software with their
users without having to worry about infrastructure, security and maintenance. 

### Security 

:fire: Pepper source :fire: provides a transparent mechanism for verifying
content integrity of the software and identity of the provider.

- The release binaries are signed by the software provider before uploading it
  to IPFS. The signature is performed using asymmetric crypto algorithms so that
the client can verify the provenience of the software released. This is all done
transparently and automatically.

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
(subscribers) know *when* an *where from* freshly baked software is ready for 
download.

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
