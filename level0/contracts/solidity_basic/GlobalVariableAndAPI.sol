// SPDX-License-Identifier: MIT
pragma solidity ^0.8.20;

contract GlobalVariableAndAPI{
  //1. 区块属性
  //返回指定区块的哈希值
//   bytes32 hash = block.blockhash(100);
//   返回挖出当前区块的矿工地址
//   address miner= block.coinbase;
//返回当前区块的难度
     uint difficulty = block.difficulty;
//返回当前区块的 Gas 上限
     uint gasLimit = block.gaslimit; 
//返回当前区块号
     uint blockNumber = block.number;
//返回当前区块的时间戳（单位：秒）
     uint timestamp = block.timestamp;
//2. 交易属性
 //返回当前合约执行剩余的 Gas 数量
     uint remainingGas = gasleft();    
//返回当前调用的完整 calldata
     bytes  data = msg.data;
//返回当前调用的发送者地址。
     address sender = msg.sender;
 //返回当前调用的函数选择器
     bytes4 functionSelector = msg.sig;
 //返回此次调用发送的以太币数量（单位：wei）
     uint sentValue = msg.value;
 //返回当前交易的 Gas 价格
     uint gasPrice = tx.gasprice;
 //返回交易的最初发起者地址
     address origin = tx.origin;
 //二节：ABI 编码及解码函数 API
 //对输入的参数进行 ABI 编码，返回字节数组
     bytes  encodedData = abi.encode(uint(1),address(0x123));
//将多个参数进行紧密打包编码，不填充到 32 字节。适用于哈希计算
     bytes  packedData = abi.encodePacked(uint(1),address(0x123));
//将参数编码，并在前面加上函数选择器（用于外部调用）
      bytes4 selector = bytes4(keccak256("transfer(address,uint256)"));
      bytes  encodedWithSelector = abi.encodeWithSelector(selector, address(0x123),100);

 //通过函数签名生成函数选择器，并将参数编码。
 bytes  encodedWithSignature = abi.encodeWithSignature("transfer(address,uint256)", address(0x123), 100);     
//对编码的数据进行解码，返回解码后的参数。
    //   (uint a, address b) = abi.decode(encodedData, (uint, address));

//第三节：数学和密码学函数 API
//在任意精度下执行加法再取模，支持大数运算。
    uint result = addmod(10, 20, 7);//结果为 2
    //先进行乘法再取模
     uint result2 = mulmod(10, 20, 7);//结果为 6
//2. 密码学哈希函数.
//使用 Keccak-256 算法计算哈希值（以太坊的主要哈希算法）
     bytes32 hash = keccak256(abi.encodePacked("Hello, World"));
//计算 SHA-256 哈希值。
     bytes32 hash1 = sha256(abi.encodePacked("Hello, World!"));
//计算 RIPEMD-160 哈希值，生成较短的 20 字节哈希值。
     bytes20 hash2 = ripemd160(abi.encodePacked("Hello, World!"));

 //通过椭圆曲线签名恢复公钥对应的地址，常用于验证签名。
// address signer = ecrecover(hash, v, r, s);    
//第四节：时间单位及其应用
/**
1 seconds = 1 秒。
1 minutes = 60 秒。
1 hours = 3600 秒。
1 days = 86400 秒。
1 weeks = 604800 秒。
1 years = 31536000 秒（非精确，忽略了闰年）。
*/
}

//2. 示例：使用时间单位进行锁定
//下面的代码示例演示了如何使用时间单位创建一个定时锁定合约，只有在锁定时间之后才能解锁代币： 
contract TimeLock{
    uint public  unlockTime;
    address public owner;
    constructor(uint _lockTime){
        owner = msg.sender;
        unlockTime = block.timestamp + _lockTime* 1 days;//锁定制定天数

    }
    
    function withdraw() public {
        require(block.timestamp >= unlockTime, "Funds are locked");
        require(msg.sender == owner, "Only owner can withdraw");
        
    }



}