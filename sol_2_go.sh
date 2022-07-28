Names=(Event CrowdFund ERC20Interchange ETHERC20Interchange)
for Name in ${Names[@]} ; do
    solc contracts/$Name/contracts/$Name.sol --combined-json bin,bin-runtime,srcmap,srcmap-runtime,abi,userdoc,devdoc,metadata,hashes --optimize -o ./contracts/$Name/contracts/ --overwrite;
    abigen --pkg contracts --out contracts/$Name/contracts/$Name.go --combined-json ./contracts/$Name/contracts/combined.json;
done
