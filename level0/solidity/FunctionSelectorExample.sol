pragma solidity ^0.8.20;

contract FunctionSelectorExample{

   function square(uint x) public  pure returns (uint){
     return x * x;
   }

   function double(uint x) public  pure returns (uint){
     return 2 * x;
   }


   //返回函数选择器
   function getSelectorSquare() external  pure  returns(bytes4){
     return this.square.selector;
   }

   function getSelectorDouble() external  pure  returns(bytes4){
    return this.double.selector;

   }

   function getSelectorDouble1() external  pure  returns(bytes4){
     return bytes4(keccak256("double(uint256)"));

   }

   //函数： 根据传入的选择器动态调用 square 或者double函数
    function select(bytes4 selector, uint x) external returns (uint z) {
        (bool success, bytes memory data) = address(this).call(abi.encodeWithSelector(selector, x));
        require(success, "Function call failed");
        z = abi.decode(data, (uint));
    }

     bytes4 public  storedSelector;

    function storeSelector(bytes4 selector) public  {
       storedSelector = selector;
    }

    function executeStoredFunction(uint x) external  returns (uint z){
        bytes4 selector = storedSelector;
        require(bytes4(0) != selector, "sSelector not set");
        (bool success, bytes memory data) = address(this).call(abi.encodeWithSelector(selector, x));
        require(success, "Function call failed");
        z = abi.decode(data, (uint));
    }


}