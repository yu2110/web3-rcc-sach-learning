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

pragma solidity 0.8.20;

contract ValueType{
    bool public isOpen = true;

    function toggleOpen() public {
        isOpen = !isOpen;
    }

    
    int8 signedInt = 32; // -2^8 - 1 -------- 2^8 - 1  (2*2*2*2*2*2*2*2=256)
    uint8 uInt = 32;   // 0 ------- 2^8 -1 
    
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
    //地址常量
    address constant fixedAddress = 0x5B38Da6a701c568545dCfcB03FcB875f56beddC4;

}