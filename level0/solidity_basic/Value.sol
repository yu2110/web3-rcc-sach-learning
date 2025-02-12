//值类型 没有浮点型，溢出
/*
布尔类型
整型
定长浮点型
定长字节数组
有理数和整型常量
字符串常量
十六进制常量
枚举
函数类型
地址类型
地址常量*/
//引用类型
// SPDX-License-Identifier: MIT
pragma solidity 0.8.20;

contract ValueType{
    bool public isOpen = true;

    function toggleOpen() public {
        isOpen = !isOpen;
    }

  /**
  整型支持以下几种运算符：
    比较运算符： <=（小于等于）、<（小于） 、==（等于）、!=（不等于）、>=（大于等于）、>（大于）
    位操作符： &（和）、|（或）、^（异或）、~（位取反）
    算术操作符：+（加号）、-（减）、-（负号）、* （乘法）、/ （除法）, %（取余数）, **（幂）
    移位： <<（左移位）、 >>（右移位）
   */  
    int8 signedInt = 32; // -2^8 - 1 -------- 2^8 - 1  (2*2*2*2*2*2*2*2=256)
    uint8 uInt = 32;   // 0 ------- 2^8 -1 
    

    function add(uint x, uint y) public  pure  returns (uint z){
        z = x + y;
    }
    
    function divide(uint x, uint y) public pure returns (uint z){
        z = x / y;
    }

    function shift(int x, uint y) public pure returns (int z){
        z = x >> y;
    }


    int256 fixedPoint = 1e18; //模拟定长浮点型
    //定长字节数据
    bytes1 singleByte = 0x12;
    // bytes256 byte256 = 0x0000000000000;

    //有理数和整型常量
    uint256 integer = 2345;
    uint256 rationalLitecal = 1.5 * 1e18; //表示1.5以太

    //字符串常量
    string greeting = "welcome";//不可变的文本常量

    //枚举
    enum Status {Valid, InValid}

    Status public  currentStatus;

    //函数类型 存储和传递函数的引用
    function setStatus(Status _status) public {
        currentStatus = _status;
    }
    //地址类型，存储以太坊地址，特殊字符串类型，20字符
    address owner = msg.sender;
    uint256 balance = owner.balance;

     address payable recipient;
    //  recipient.transfer(1 ether);
    //  bool success = recipient.send(1 ether);
    //  (bool success,) = recipient.call{value: 1 ether}("");


    //地址常量
    // address constant fixedAddress = 0x5B38Da6a701c568545dCfcB03FcB875f56beddC4;

    // 定义一个 address 类型的变量
    address public myAddress1 = 0x1234567890123456789012345678901234567890;
    address myAddress2 = msg.sender;  // 当前合约调用者的地址

    // if (myAddress1 == myAddress2) {
    // // 两个地址相等时的操作
    // } else {
    // // 两个地址不等时的操作
    // }

    address payable payableAddress = payable(myAddress1);  // 将 address 转换为 address payable

}