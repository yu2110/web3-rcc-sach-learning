pragma solidity ^0.8.20;

contract TyperConvert{


    int8 y = -3; 
    // uint x = uint(y); 
    //如果转换成更小的类型，变量的值会丢失高位。 
    uint32 a = 0x12345678; 
    uint16 b = uint16(a); // b = 0x5678 
    //转换成更大的类型，将向左侧添加填充位 
    uint16 c = 0x1234;
    uint32 d = uint32(c); // b = 0x00001234 
    //转换到更小的字节类型，会丢失后面数据。 
    bytes2 e = 0x1234; 
    bytes1 f = bytes1(e); // b = 0x12 
    //转换为更大的字节类型时，向右添加填充位。 
    bytes2 g = 0x1234; 
    bytes4 h = bytes4(g); // b = 0x12340000 
    //把整数赋值给整型时，不能超出范围，发生截断，否则会报错。 
    uint8 i = 12; // 正确 
    uint32 l = 1234; // 正确 
    // uint16 m= 0x123456; // 错误,
}

