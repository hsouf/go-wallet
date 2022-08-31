// SPDX-License-Identifier: MIT
pragma solidity ^0.8.0;

import "./Wallet.sol";

contract WalletFactory {
    uint256 public constant mintCost = .003 ether;
    address private factoryGov;

    event WalletMinted(address wallet, address admin);

    function setAdmin(address _factoryGov) public {
        factoryGov = _factoryGov;
    }

    function mintWallet(address _admin) public payable returns (address) {
        require(msg.value >= mintCost, "Factory: insufficient value");
        Wallet newWallet = new Wallet(_admin);

        emit WalletMinted(address(newWallet), _admin);

        return address(newWallet);
    }

    function withdraw() public payable {
        require(msg.sender == factoryGov);
        (bool os, ) = payable(factoryGov).call{value: address(this).balance}(
            ""
        );
        require(os);
    }
}
