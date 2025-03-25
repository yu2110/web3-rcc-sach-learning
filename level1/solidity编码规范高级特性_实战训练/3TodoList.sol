// SPDX-License-Identifier: GPL-3.0

pragma solidity >=0.7.0 <0.9.0;


contract TodoList{
  
   struct Todo{
     string name;
     bool isFinished;
   }

   Todo[] public list;

   function create(string memory _name) external{
     list.push(
        Todo({name:_name,
              isFinished:false
        })
     );
   }


   function rename1(uint256 _i,string  memory _name) external{
         list[_i].name = _name;
   }

    function rename2(uint256 _i,string  memory _name) external{
         list[_i].name = _name;
         Todo storage temp = list[_i];
         temp.name = _name;
   }
 
    function setStatus1(uint256 _i, bool _status) external{
        list[_i].isFinished = _status;
    }

     
    function setStatus2(uint256 _i) external{
        list[_i].isFinished = ! list[_i].isFinished;
    }


    function get1(uint256 _i) external view returns (string memory _name, bool _status){

        Todo memory temp = list[_i];
        return (temp.name, temp.isFinished);
    }


    function get2(uint256 _i) external view returns (string memory _name, bool _status){

        Todo storage temp = list[_i];
        return (temp.name, temp.isFinished);
    }


}
