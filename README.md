# NID Native Registry

NID Native Registry(NNR) is a set of registry services specifically for localized registries, which can be used to quickly build NNS secondary resolution services.

## NNS resolution mechanism

- The user enters the NNS domain name and queries the root server to obtain the registrar service address;
- Initiate a request to the registrar's service address and carry the NNS domain name;
- NNR parses the NNS domain name, queries the records in the registry, obtains the NFT information and returns it.

## Features

NNR adopts the principle of simplification and uses the database self-incrementing ID as the NNS second-level domain name allocation rule. Distributed deployment is not currently supported. In the future, distributed IDs will be considered as the NNS second-level domain name allocation rules.

NFT metadata is synchronously stored in IPFS, which is convenient for data synchronization and backup.

The domain name record is verified by `accessToken`, which is generated when it is created, and the `accessToken` must be carried in the update for permission verification. User must provide HTTP Authorization header with format `Bearer {accessToken}`. The `accessToken` is regenerated for each update, and the newly generated `accessToken` must be carried in the next update.
