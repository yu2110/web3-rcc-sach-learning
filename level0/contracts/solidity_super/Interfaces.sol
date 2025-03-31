// SPDX-License-Identifier: MIT
pragma solidity ^0.8.20; 

/**
Solidity 中的接口具有以下限制：

无法继承其他合约或接口：接口不能扩展其他合约或接口。
无法定义构造函数：接口不允许定义任何构造函数，因为它不能有内部状态。
无法定义状态变量：接口不能有状态变量，因为它不存储数据。
无法定义结构体或枚举：接口不能包含结构体或枚举。
*/

//接口的使用
// 定义一个接口 IToken

interface  IToken {
   function transfer(address recipient, uint256 amount) external ;
} 

// 实现接口的合约
contract SimpleToken is IToken{
    mapping (address => uint256) public  balances;
    constructor(){
        balances[msg.sender] = 1000;//初始化代币余额
    }

    // 实现接口中的 transfer 函数
    function transfer(address recipient, uint256 amount) public  override {
       require(balances[msg.sender] >= amount, "Insufficient balance");
       balances[msg.sender] -= amount;
       balances[recipient] += amount;
    }

}
//合约之间通过接口进行通信

contract Award {
    IToken immutable token;

    constructor(IToken _token){
        token = _token;
    }

    // 调用 SimpleToken 合约的 transfer 函数来发送奖励
    function sendBonus(address user) public {
        token.transfer(user, 100);//向用户发送100个代币作为奖励
    }
}

//ERC20 标准的实现

// interface IERC20 {
//     function totalSupply() external view returns (uint256);
//     function balanceOf(address account) external view returns (uint256);
//     function transfer(address recipient, uint256 amount) external returns (bool);
//     function allowance(address owner, address spender) external view returns (uint256);
//     function approve(address spender, uint256 amount) external returns (bool);
//     function transferFrom(address sender, address recipient, uint256 amount) external returns (bool);
// }