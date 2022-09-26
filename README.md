# WRecon

Simple tools for pinging a list of websites.

### How to use

> wrecon --config ./sites.json

This will print the result in your console.

### Config example

````json
{
  "sites": [
    {
      "name": "Example!",
      "address": "http://example.com"
    },
    {
      "address": "https://google.com"
    }
  ]
}
````

### Watch

Use flag ``interval``
> wrecon --config ./sites.json --interval 30s

This will ping your sites every 30 seconds until you quit the program.
