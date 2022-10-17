// SPDX-License-Identifier: MIT
pragma solidity ^0.8.0;

import "@openzeppelin/contracts/token/ERC20/IERC20.sol";

contract Wallet {
    // To reduce deployment costs this variable is internal and needs to be retrieved via `getStorageAt`
    address private admin;

    event EtherDeposited(
        address indexed wallet,
        address depositor,
        uint256 amount
    );

    constructor(address _admin) public payable {
        admin = _admin;
    }

    function withdrawERC20(address _token, uint256 amount) public {
        require(msg.sender == admin);
        require(
            IERC20(_token).balanceOf(address(this)) >= amount,
            "insufficient funds"
        );
        IERC20(_token).transfer(msg.sender, amount);
    }

    function withdrawEther() public payable {
        require(msg.sender == admin);
        (bool os, ) = payable(admin).call{value: address(this).balance}("");
        require(os);
    }

    fallback() external payable {
        emit EtherDeposited(address(this), msg.sender, msg.value);
    }
}
