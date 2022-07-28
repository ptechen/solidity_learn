// SPDX-License-Identifier: MIT
pragma solidity ^0.8.0;

contract SelfDesctruct {
    constructor() payable{}

    function kill() external {
        selfdesctruct(payable(msg.sender));
    }

    function testCall() external pure returns (uint){
        return 256;
    }
}
