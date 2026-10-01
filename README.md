# CRDPA

ContainedReality's Data Protection Algorithm.

## Design

It's an ARX-based algorithm that uses 256 bit keys, has a block consisting of four 32 bit words. 56 rounds.

This is my first cipher that I've designed and released, so expect many changes that may completely change the specification.

## Notes

### Design

CRDPA is based off of/inspired by SPECK.

### Security

Don't use it. CRDPA Hasn't seen formal cryptanalysis, if it does I'll add it here. Otherwise I think it's secure, but I could totally be wrong, and it's not worth the risk when other formally reviewed and battletested ciphers exist. [Use AES instead](https://pkg.go.dev/crypto/aes@latest).

Only use it if you really know what you're doing and accept the risks of using a cipher that hasn't been analyzed or audited.

CRDPA was mostly designed as a learning exercise for ARX ciphers, and trying to make a reasonably secure cipher. So If it turns out I made an actually secure encryption algorithm, then I'm happy, that's the goal, but it just isn't a risk someone should take. It could have some wicked vulnerabilities that I don't see.
