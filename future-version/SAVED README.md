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

The full index can be downloaded, but it will consume somewhere around 80GB of storage. However, it will make initial queries to new addresses much faster.

## Configuration

After installing this package, go to [configure.trueblocks.public.dappnode] to add or remove chains and change settings.
When you configure it for a first time, you do not need to restart the package.

## Exposing Publicly

It is highly recommended that you don't expose TrueBlocks publicly. It is intended for decentralized single user use. But, you do you.

This API has destructive options. You should block the following:
* HTTP DELETE method
  * Any HTTP DELETE call is destructive
* /scrape
  * This API can turn your scraper on and off, ideally you don't want someone to be able to do this.

## Contributing

We love contributors. Please see information about our [work flow](https://github.com/TrueBlocks/trueblocks-core/blob/develop/docs/BRANCHING.md) before proceeding.

1. Fork this repository into your own repo.
2. Create a branch: `git checkout -b <branch_name>`.
3. Make changes to your local branch and commit them to your forked repo: `git commit -m '<commit_message>'`
4. Push back to the original branch: `git push origin TrueBlocks/trueblocks-core`
5. Create the pull request.

## List of Contributors

Thanks to the following people who have contributed to this project:

* [@tjayrush](https://github.com/tjayrush)
* [@dszlachta](https://github.com/dszlachta)
* [@wildmolasses](https://github.com/wildmolasses)

## Contact

If you have questions, comments, or complaints, please join the discussion on our discord server which is [linked from our website](https://trueblocks.io).
