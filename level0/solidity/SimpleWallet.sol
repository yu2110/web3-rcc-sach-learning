// SPDX-License-Identifier: MIT
pragma solidity 0.8.20;

/**
练习任务:

编写一个智能合约，允许用户存款并设置一个特定的地址为白名单地址，只有该地址能够提取合约中的资金。
实现一个功能，允许用户查询自己在合约中的余额，并且测试存款、提取和余额查询功能的正确性。
测试案例:

说明: 通过测试用例验证学生实现的功能，确保合约代码的正确性和安全性。
测试用例示例:
测试 deposit() 函数，确保余额正确更新。
测试 withdraw() 函数，检查余额减少和以太币成功转移。
测试白名单功能，确保只有授权的地址可以提取资金。

*/

contract SimpleWallet{
    

    function deposit (uint256 amount) public  returns (bool) {

    }

    function withdraw () public  returns (uint256 amount) {

    }

    function balance () public  returns (uint256 amount) {

    }


}



