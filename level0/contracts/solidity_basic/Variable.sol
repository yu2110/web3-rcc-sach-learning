// SPDX-License-Identifier: MIT
pragma solidity 0.8.20;

//solidity静态类型语言，这意味这申明期间我们就需要指定变量类型，每个变量申明的时候，都有一个基于类型的默认值, 像java一样 undefined null
//状态变量:变量值永久存储到区块链的变量，合约存储空间
//局部变量:仅在函数执行中有效的变量，函数退出后，变量无效
//全局变量:存在全局命名空间的变量，可以获取区块链的相关相关信息，比如区块链高度


contract Variables {
   
   //状态变量
    string public text = "welcome";
    
    uint256 public  num = 456;

    function doWelcoming() public {
        uint256 i = 789; //局部变量
        uint256 timestamp = block.timestamp; //全局变量
        address sender = msg.sender;
    }

   
    // block.blockhash;
    // block.coinbase;
    // block.difficulty;
    // block.gaslimit;
    // block.number;
    // block.timestamp;
    // msg.data;
    // msg.sender;
    // msg.sig;
    // msg.value;
    // now;
    // tx.gas;
    // tx.origin;

   //变量的使用规范

   // 1，不应该不用关键词作为变量名，比如now,break,for,if,bool,block,blockhash;
   // 2,不应该以数字开通0-9,123test,无效变量名，_123test,局部变量。
   // 3，区分大小写，Name 和name，是两个不同的变量。

    
}