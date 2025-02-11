
pragma solidity ^0.8.20;

/**
作业：
任务 1：基于映射创建一个简单的用户余额管理系统，并实现余额的增加与查询功能。
任务 2：扩展系统，使其能够记录每个用户的交易历史，模拟 Java 的 Map 中键集合的概念。
任务 3：结合映射与数组，实现一个简单的排行榜系统，记录并排序用户的积分。

*/

contract MappingUserPointsSystem{
    mapping(address => uint) public  points;

    address[] public  users;


    function addUser(address _user) public {
        require(points[_user] == 0, "User already exists.");
        users.push(_user);
        points[_user] = 100;//init 100 
    }


    function transferPoints(address _to, uint _amount) public {
        require(points[msg.sender] >= _amount, "Insufficient points.");
        points[msg.sender] -= _amount;
        points[_to] +=_amount;
    }

    function getAllUsers() public view returns (address[] memory){
        return users;
    }


    

}

