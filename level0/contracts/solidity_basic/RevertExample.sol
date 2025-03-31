// SPDX-License-Identifier: MIT
pragma solidity ^0.8.20;

contract RevertExample{

    function checkValue(uint value) public  pure {
        if(value > 10){
            revert("value cannot exceed 10");
        }
    }

}