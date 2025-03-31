// SPDX-License-Identifier: MIT
pragma solidity ^0.8.20;

contract CustomErrorExample{
    error Unauthorized(address caller);// self 定义 error
    address public owner;
    constructor(){
        owner = msg.sender;
    }

    function restrictedFunction() public  view{
        if(msg.sender != owner){
            revert Unauthorized(msg.sender);
        }
    }

}
