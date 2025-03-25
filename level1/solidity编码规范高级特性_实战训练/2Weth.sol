// SPDX-License-Identifier: GPL-3.0

pragma solidity >=0.7.0 <0.9.0;

contract WETH{
    string public name = "wrapped ether";
    string public symbol = "weth";
    uint8 public decimails = 18;

    event Approval(address indexed from, address indexed  to, uint256 amount);

    event Transfer(address indexed from, address indexed to, uint256 amount);

    event  Deposit(address indexed to, uint256 amount);

    event Withdraw(address indexed from, uint256 amount);

    mapping (address => uint256) public balances;

    mapping (address => mapping (address => uint256) )  public allows;


    function deposit() public payable{
        balances[msg.sender] += msg.value;
        emit Deposit(msg.sender, msg.value);
    }

    function withdraw(uint256 amount) public{
        require(balances[msg.sender] >= amount);
        balances[msg.sender] -= amount;
        payable(msg.sender).transfer(amount);
        emit Withdraw(msg.sender, amount);
    }

    function totalSupply() public view returns (uint256){
        return address(this).balance;
    }

    function approve(address to, uint256 amount) public returns (bool){
        allows[msg.sender][to] = amount;
        emit Approval(msg.sender, to, amount);
    }

    function transfer(address to, uint256 amount) public returns (bool){
        return transferFrom(msg.sender, to, amount);
    }

    function transferFrom(address from, address to, uint256 amount) public returns (bool) {
          require(balances[from] >= amount);
          if( from != msg.sender){
            require(allows[from][msg.sender]>=amount);
            allows[from][msg.sender] -= amount;
          }

          balances[from] -= amount;
          balances[to] += amount;
          emit Transfer(from, to, amount);
          return true;
    }


    fallback() external payable{
        deposit();
    }

    receive() external payable{
        deposit();
    }


}