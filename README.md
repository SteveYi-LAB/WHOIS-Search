# WHOIS Search

A simple WHOIS search tool, which was used on whois.steveyi.net.

## Install

Download the latest release version, port will listen at 30010.

## Usage

If you are trying to specify the IRR DB, please refer to the below list.

```
"afrnic"    whois.afrinic.net
"apnic"     whois.apnic.net
"arin"      whois.arin.net
"lacnic"    whois.lacnic.net
"ripe"      whois.ripe.net
"RADB"      whois.radb.net
```

HTTP GET:

```
curl http://localhost:30010/whois/{target}
curl http://localhost:30010/whois/{target}/{IRR}
```

