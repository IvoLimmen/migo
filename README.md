# mi

Running `mi` will create a list of artifacts to build to shorten the build time.

## Build

    go build

Or use:

    ./build.sh

## Installation

Download and put into a 'bin' directory that is on your path.
Depends on the GIT command line utility.

## Usage

Run 'mi' in a directory of a maven project that contains changes.
Will return show the command (and executes) to build only what needed.
If a maven `settings.xml` is found in the parent directory it is added to the command to use that as settings file.

## Contributing

1. Fork it (<https://github.com/ivolimmen/migo/fork>)
2. Create your feature branch (`git checkout -b my-new-feature`)
3. Commit your changes (`git commit -am 'Add some feature'`)
4. Push to the branch (`git push origin my-new-feature`)
5. Create a new Pull Request

## Contributors

- [Ivo Limmen](https://github.com/ivolimmen) - creator and maintainer

## Credits
- [etree](https://github.com/beevik/etree) - for parsing XML
