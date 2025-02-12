// SPDX-License-Identifier: MIT
pragma solidity ^0.8.20;

//第一节：合约的基本结构
contract BasicContract {
    uint public data;

    function setData(uint _data) public {
        data = _data;
    }
}

//第二节：可见性修饰符（Visibility Modifiers）
/**
external：只能从合约外部调用，不能从合约内部调用（除非使用 this 关键字）。
public：支持内部和外部调用，对于 public 状态变量，编译器会自动生成访问器函数。
internal：只能在当前合约或继承的合约中访问。
private：只能在当前合约内部访问，派生合约中无法访问。
*/

contract VisibilityExample {
    uint public data;
    uint internal iData = 10;
    function setData(uint a) internal {
        data = a;
    }
    function f(uint a) private pure returns (uint) {
        return a + 1;
    }
}
//第三节：构造函数（Constructor）
contract ConstructorExample {
    uint public x;
    constructor(uint _x) public {
        x = _x;
    }
}

//第四节：常量状态变量与不可变量
contract ConstantsExample {
    uint constant x = 42;
}

contract ImmutableExample {
    uint immutable maxBalance;

    constructor(uint _maxBalance) public {
        maxBalance = _maxBalance;
    }
}
//第五节：视图函数与纯函数

contract ViewFunctionExample {

    uint public data;

    // function getData() public view returns (uint) {
    //     return data;
    // }
}

contract PureFunctionExample {
    function add(uint a, uint b) public pure returns (uint) {
        return a + b;
    }
}
//第六节：事件（Events）
contract EventExample {
    event DataChanged(uint newValue);
    uint public data;
    function setData(uint _data) public {
        data = _data;
        emit DataChanged(_data);  // 触发事件
    }
}
//第七节：接收函数与回退函数
contract ReceiveExample {
    event Received(address sender, uint amount);
    // 仅用于接收以太币
    receive() external payable {
        emit Received(msg.sender, msg.value);
    }
}

contract FallbackExample {
    event FallbackCalled(address sender, uint amount);
    // 当调用不存在的函数时触发
    fallback() external payable {
        emit FallbackCalled(msg.sender, msg.value);
    }
}
//第八节：函数修改器（Modifiers）
//函数修改器用于改变函数的行为，可以用于验证条件、修改参数等。
contract ModifierExample {
    address public owner;
    constructor() public {
        owner = msg.sender;
    }
    modifier onlyOwner() {
        require(msg.sender == owner, "Only owner can call this function.");
        _;
    }
    function changeOwner(address newOwner) public onlyOwner {
        owner = newOwner;
    }
}

//第九节：函数重载（Function Overloading）
contract OverloadingExample {
    function f(uint _in) public pure returns (uint) {
        return _in;
    }
    function f(uint _in, bool _flag) public pure returns (uint) {
        return _flag ? _in : 0;
    }
}













