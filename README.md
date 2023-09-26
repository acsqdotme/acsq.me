# personal website of [ángel castañeda](https://www.angel-castaneda.com)

I wrote my site in go to escape the comforts of nginx and actually learn how a
webserver works.

## features

* templating html

* real multilingual content/url localization for
  [spanish](https://es.angel-castaneda.com) and
  [german](https://de.angel-castaneda.com)

* [atom feed](https://www.angel-castaneda.com/atom.xml)

* [fancy error pages](https://www.angel-castaneda.com/page/doesnt/exist)

* md -> html for articles

* blog organized w/ tagging system

* gzipping

* sqlite server

## License

This site uses my [dblog repo](https://git.acsq.me/dblog) as an interface
between the sqlite database and my code. That package is licensed under the
LGPLv3. Check it's [`LICENSE`](https://git.acsq.me/dblog/tree/LICENSE/)
