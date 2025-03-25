// SPDX-License-Identifier: GPL-3.0

pragma solidity >=0.7.0 <0.9.0;

contract EtherWallet{

    address payable public immutable owner;
    event Log(string funcName, address from ,uint256 amount, bytes data);

    constructor(){
        owner = payable(msg.sender);
    }

    receive() external payable{
        emit Log("receive", msg.sender,msg.value,"");
    }

    modifier  isOwner {
        require(msg.sender == owner, "Not owner");
        _;
    }

    function withdraw1() external isOwner{
        payable(msg.sender).transfer(100);
    }

    function withdraw2() external isOwner{
        bool success = payable(msg.sender).send(100);
        require(success, "send failed");
    }

    function withdraw3() external isOwner{
        (bool success,) = msg.sender.call{value: address(this).balance}("");
         require(success, "call failed");
    }

    function getBalance() external view returns (uint256){
        return address(this).balance;
    }


}