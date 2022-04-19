# lookup
simple cli tool to check domain

```shell
lookup ycombinator.com

 --- [ whois whois.verisign-grs.com:43 ] ---
   Domain Name: YCOMBINATOR.COM
   Registry Domain ID: 147225527_DOMAIN_COM-VRSN
   Registrar WHOIS Server: whois.gandi.net
   ... <removed content to make example easier to read> ...

 --- [ certs TLS 1.3 ] ---
SANs: ycombinator.com, *.ycombinator.com
Expiry: 2023-05-12 23:59:59 +0000 UTC

 --- [ lookup ] ---
ycombinator.com has address
  54.192.137.12 domain name pointer server-54-192-137-12.lhr62.r.cloudfront.net.
  54.192.137.71 domain name pointer server-54-192-137-71.lhr62.r.cloudfront.net.
  54.192.137.3 domain name pointer server-54-192-137-3.lhr62.r.cloudfront.net.
  54.192.137.20 domain name pointer server-54-192-137-20.lhr62.r.cloudfront.net.
CNAME ycombinator.com.
NS records
  ns-1411.awsdns-48.org.
  ns-1914.awsdns-47.co.uk.
  ns-225.awsdns-28.com.
  ns-556.awsdns-05.net.
MX records
  aspmx.l.google.com.
  alt1.aspmx.l.google.com.
  alt2.aspmx.l.google.com.
  aspmx4.googlemail.com.
TXT records
  ZOOM_verify_2ndw8KZxSRa8PT8NmdyXvw
  apple-domain-verification=WG0sP5Alm7N6h1Te
  google-site-verification=GJKdQskycEclAGPua3yXB9m_nVhxbrsVps_y-t9SXV0
  google-site-verification=KsI69Y_jEVkp4eXqSQ9R9gwxjIpZznvuvrus6UolB9Y
  google-site-verification=rivq8jKu6AADGtbbEzJhmOpcqq08B7QxIzXxYV8DtyU
  v=spf1 include:_spf.google.com include:mailgun.org a:rsweb1-36.investorflow.com include:_spf.createsend.com include:servers.mcsv.net -all
```

## release

Releases are published when the new tag is created e.g.
`git tag -m "add super cool feature" v1.0.0 && git push --follow-tags`
