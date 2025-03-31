// SPDX-License-Identifier: MIT
pragma solidity ^0.8.20; 

library SafeMath {
    function add(uint a, uint b) external   pure returns (uint){
        uint c = a + b;
        require(c >= a, "SafeMath: addition overflow");
        return c;
    }
}