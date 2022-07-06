# TrueBlocks DAppNode Package

This DAppNode package contains both [TrueBlocks-Core](https://github.com/TrueBlocks/trueblocks-core) AND [TrueBlocks-Exporer](https://github.com/TrueBlocks/trueblocks-explorer)

These two packages work together to provide decentralized indexing and address monitoring/exploring. See their official repos for more information. There are two primary requirements:

1. An [Etherscan](https://etherscan.io/) API Key
    * This key is used to download contract ABI data
2. An Ethereum endpoint capable of `trace_` calls (Erigon, OpenEthereum, Nethermind?)

For everything else it is generally recommended to just run with these defaults:

* Enable scraper
* Bootstrap bloom filters
* Do not download entire index

The scraper will fill in the missing gaps from the bloom filters and continually update the chain data from your endpoint.

The Bloom Filters are what allow quick searching of the appearance data, it's downloaded from IPFS and consumes around 3GB of storage.

The full index can be download, but it will consume somewhere around 80GB of storage. However, it will make initial queries to new addresses much faster.

## Configuration

This package was designed to be user friendly and simple, as such it only supports Ethereum Mainnet out of the box and there are config options only for Ethereum Mainnet exposed under the configs section of the DAppNode package UI.

However, you can download, modify, and re-upload the `/root/.local/share/trueblocks/trueBlocks.toml` file and add your own chain configs from inside the DAppNode UI.

## Exposing Publicly

It is highly recommended that you don't expose TrueBlocks publicly. It is intended for decentralized single user use. But, you do you.

This API has destructive options. You should block the following:
* HTTP DELETE method
  * Any HTTP DELETE call is destructive
* /scrape
  * This API can turn your scraper on and off, ideally you don't want someone to be able to do this.
