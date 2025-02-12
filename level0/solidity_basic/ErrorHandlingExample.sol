// SPDX-License-Identifier: MIT
pragma solidity ^0.8.20;


contract ErrorHandlingExample{
       uint public balance;

       function sendHalf(address payable  addr) public  payable {
          require(msg.value % 2 ==0, "Even value required");

          uint balanceBeforeTransfer = address(this).balance;

          addr.transfer(msg.value / 2);

          assert(address(this).balance == balanceBeforeTransfer - msg.value /2);
       }

}