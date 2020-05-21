# crypto-go

Code examples and slides for the talk: "Crypto Fails and How to Tackle Them in Go" at go get -u community conference in May 2020.

## Try out the two code examples

```
$ go run main.go aes 
$ go run main.go random
```

You can find the template usage in the corresponding go-file in the cmd package, e.g., `cmd/aes.go` to see the example to encrypt and decrypt a message.

## Further Resources

- Check out the [cryptopasta repo by George Tankersley](https://github.com/gtank/cryptopasta)
  - [Slide deck](https://speakerdeck.com/gtank/crypto-for-go-developers)
  - [Talk at Youtube (GopherCon 2016)](https://www.youtube.com/watch?v=2r_KMzXB74w)
  - [Code examples](https://github.com/gtank/cryptopasta)
- The Gopher slack has a *crypto* channel which provides further information
- [Gosec](https://github.com/securego/gosec) can check for some crypto misuses like a usage of MD5. 

- Egele, M., Brumley, D., Fratantonio, Y., & Kruegel, C. (2013, November). An empirical study of cryptographic misuse in android applications. In Proceedings of the 2013 ACM SIGSAC conference on Computer & communications security (pp. 73-84).
[Paper](https://dl.acm.org/doi/pdf/10.1145/2508859.2516693)
- Krüger, S., Späth, J., Ali, K., Bodden, E., & Mezini, M. (2019). Crysl: An extensible approach to validating the correct usage of cryptographic apis. IEEE Transactions on Software Engineering.
[Paper](https://ieeexplore.ieee.org/document/8880510)
- Veracode report: "The State of Software Security Today". [Wepage](https://www.veracode.com/state-of-software-security-report)
