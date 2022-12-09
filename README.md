# TrueBlocks DAppNode Package

<<<<<<< HEAD
This dAppNode package contains both [TrueBlocks-Core](https://github.com/TrueBlocks/trueblocks-core) and [TrueBlocks-Exporer](https://github.com/TrueBlocks/trueblocks-explorer).

These two packages work together to provide decentralized indexing and address exploring. See the underlying repos for more information.
=======
## Building

As long as https://github.com/dappnode/DAppNodeSDK/pull/258 is not merged, we have to use forked DAppNodeSDK to build the package:

```bash
git clone https://github.com/dszlachta/DAppNodeSDK.git
cd DAppNodeSDK
git checkout dszlachta/fix_wizard_target_service_schema
yarn
yarn build

# Go back to this directory
cd ../trueblocks-dappnode
../DAppNodeSDK/dist/dappnodesdk.js build
```

## Description

This DAppNode package contains both [TrueBlocks-Core](https://github.com/TrueBlocks/trueblocks-core), a tool providing decentralized indexing and address monitoring/exploring.
See official repo for more information. There are two primary requirements:
>>>>>>> 3b976ed189228b91222d82d39f7b5a7f2175172b

In order for this package to work, you must provide two things:

1. An Ethereum RPC endpoint enabling the `trace_` namespace. (We suggest Erigon or Nethermind. Geth does not work.)
2. An [Etherscan](https://etherscan.io/) API key. (This is optional, but enables certain features in the Explorer).

While you may configure the following, by default this package:

* Downloads the Unchained Index bloom filters from IPFS (about 3GB), and
* Starts the TrueBlocks indexing scraper.

The scraper "picks up" from where the bloom filters leave off, meaning your implementation will be producing its own index from that point forward. This will keep your index "fresh."

As you query against the history of the chain for particular addresses, the system will download only those portions of the Unchained Index needed for that particular address. This keeps the size of the index on your machine to a minimum. In effect, you get that portion of the index that you are interested in and no more.

The Bloom filters allow quick searching of the chain's history. The index portions themselves, once downloaded, provide exact locations for an addresses' appearances.

It is possible to download the entire index, and in this case, the overall system behaves more quickly, although this mode takes up nearly 110 GB on your machine.

, but it will consume somewhere around 80GB of storage. However, it will make initial queries to new addresses much faster.

## Configuration

This package was designed to be user friendly and simple, as such it only supports Ethereum Mainnet out of the box and there are config options only
for Ethereum Mainnet exposed under the configs section of the DAppNode package UI.

After installing this package, go to [configure.trueblocks.public.dappnode] to add or remove chains and change settings. When you configure it for a
first time, you do not need to restart the package.

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

## Contact

If you have questions, comments, or complaints, please join the discussion on our discord server which is [linked from our website](https://trueblocks.io).

## List of Contributors

Thanks to the following people who have contributed to this project:

- [@tjayrush](https://github.com/tjayrush)
- [@dszlachta](https://github.com/dszlachta)
- [@wildmolasses](https://github.com/wildmolasses)
