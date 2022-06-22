# Simple DApp

This project is created for learning and trying some DApp stuff on Ethereum.

Truffle/Ganache is used for running test environment

## Starting Test Network

If you prefer Ganache, simply create new workspace by pointing to `testnet/truffle-config.js` file.

Otherwise, [install Truffle](https://trufflesuite.com/docs/truffle/getting-started/installation/) and,

```
$ cd testnet
$ truffle develop
```

Then you need to deploy the contracts from Truffle CLI,

```
$ (truffle) deploy
```

## Testing Go App

After starting the network, edit `goapp/.env` file according to your parameters, then run the application.
This document provides a docker way to run the app;

```
$ docker run -it --network host --name dappgo -v $(pwd)/goapp:/usr/local/go/src/simple-dapp -w /usr/local/go/src/simple-dapp golang:1.18
$ go run ./...
```

# TODO

[ ] Add web app GUI for the dApp