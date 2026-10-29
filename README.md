# CRDPA

[CRDPA - containedreality](https://containedreality.github.io/2026/10/26/crdpa.html)

### Security

Don't use it. CRDPA Hasn't seen formal cryptanalysis, if it does I'll add it here. Otherwise I *think* it's "secure", but I could totally be wrong, and I probably am. It's not worth the risk when other formally reviewed and battletested ciphers exist. [Use AES instead](https://pkg.go.dev/crypto/aes@latest).

Only use it if you really know what you're doing and accept the risks of using a cipher that hasn't been analyzed or audited.

CRDPA was mostly designed as a learning exercise for ARX ciphers, and trying to make a reasonably secure *toy* cipher. So If it turns out I made an actually secure encryption algorithm, then I'm happy, that's the goal, but it just isn't a risk someone should take. It could have some wicked vulnerabilities that I don't see.
