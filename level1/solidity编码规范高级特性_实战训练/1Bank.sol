// SPDX-License-Identifier: GPL-3.0

pragma solidity >=0.7.0 <0.9.0;

contract Bank {

    address public immutable owner;

    event Deposit(address _sender, uint256 amount);
    event Withdraw(address _sender, uint256 amount);

    constructor() payable{
        owner = msg.sender;
    }

    function withdraw() external{
        require(msg.sender == owner, "Not Owner");
        emit Withdraw(msg.sender, address(this).balance);
        selfdestruct(payable(msg.sender));
    }

    function getBalance() external view returns (uint256){
        return address(this).balance;
    }

    receive() external payable{
        emit Deposit(msg.sender, msg.value);
    }


}