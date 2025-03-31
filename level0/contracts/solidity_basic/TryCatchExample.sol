// SPDX-License-Identifier: MIT
pragma solidity ^0.8.20;


contract TryCatchExample{
   
     function tryCatchDemo(address _contractAddress) public {
           // 尝试调用外部合约的函数
            try ExternalContract(_contractAddress).someFunction() {
            // 处理成功
            } catch {
            // 处理失败
            }
     }

}

contract ExternalContract{
    function someFunction() public returns(bool){
        //可能抛出异常的函数
    }
}