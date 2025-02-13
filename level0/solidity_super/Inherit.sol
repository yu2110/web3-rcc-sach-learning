// SPDX-License-Identifier: MIT
pragma solidity ^0.8.20; 

//2.继承的基本用法

contract Owned {
       address payable   public    owner;  //public  private internal

       constructor(){
         owner = payable(msg.sender);
       }

       function setOwner(address _owner) public  virtual {
          owner = payable (_owner);
       }

}

contract Mortal is Owned {
   
     event SetOwner(address  indexed  owner);

     function kill() public {
        //  if(msg.sender == owner) selfdestruct(owner);
     }

     function setOwner(address _owner) public override{
         super.setOwner(_owner);//调用父合约的setOwner
         emit SetOwner(_owner);

     }
}
//第二节：多重继承
// contract Named is Owned, Mortal{

// }   

//多重继承的顺序
contract X {}
contract A is X {}
// contract C is A, X{} // 编译出错，X 出现在继承关系中两次
//方式2：在子合约构造函数中通过修饰符调用父合约

abstract contract A1 {
    uint public a;

    constructor(uint _a) {
        a = _a;
    }
}

contract B1 is A1 {
    uint public  b;

    constructor() A1(1){
        b = 2;
    }
}

//抽象合约 纯虚函数

abstract contract A2{
    function get() virtual  public ;//纯虚函数没有实现，用 virtual 关键字修饰

}

//函数重写
contract Base {
    function foo() virtual public {}
}
contract Middle is Base{

}
contract Inherited is Middle{
    function foo() public override  {}
}

//多重继承中的重写
contract Base1 {
    function foo() virtual public {}
}

contract Base2 {
    function foo() virtual public {}
}

contract Inherited1 is Base1, Base2{
    function foo() public  override(Base1,Base2) {}
}


