// SPDX-License-Identifier: MIT
pragma solidity ^0.8.20; 

/**
库仅由函数组成，不能定义状态变量。
库不能直接修改 msg.sender 或 this，因为它们没有自己独立的上下文。
库函数可以是内部函数，也可以是公共或外部函数，使用不同的函数类型决定库的使用方式
*/

// library SafeMath {
//     function add(uint a, uint b) internal  pure returns (uint){
//         uint c = a + b;
//         require(c >= a, "SafeMath: addition overflow");
//         return c;
//     }
// }


// 使用 SafeMath 库的合约
import "./LibrarySafeMath.sol";

contract AddTest{
    function add(uint x, uint y) public  pure returns (uint){
        return SafeMath.add(x, y);
    }
}
//使用 ** using for 绑定类型
contract TestLib {
    using SafeMath for uint;
    function add(uint x, uint y) public pure returns (uint){
        return x.add(y);
    }
}

library Search{
    function indexOf(uint[] storage self, uint value) public  view returns (uint) {
        for ( uint i = 0; i < self.length; i++){
            if(self[i] == value) return i;
        }
        return uint(0); // 如果找不到返回 -1
    }
}

contract C {
    using Search for uint[];
    uint[] data;
    function append(uint value) public {
        data.push(value);
    }

    function replace(uint _old, uint _new) public {
        uint index = data.indexOf(_old);
        if (index == uint(0)) {
            data.push(_new);
        } else {
            data[index] = _new;
        }
    }
}



