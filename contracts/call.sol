// SPDX-License-Identifier: GPL-3.0

pragma solidity ^0.8.0 ;


contract IterableMapping {

   mapping(address => uint) private  data;

   address[] private keys;

   mapping(address => bool) private keyExists;

   function add(address _key, uint _value) public {
     
     if(!keyExists[_key]){
        keys.push(_key);
        keyExists[_key] = true;
     }
     data[_key] = _value;
   }

   function remove(address _key) public {
      if(keyExists[_key]){
        delete data[_key];
        keyExists[_key] = false;
      }

      for(uint i =0; i < keys.length; i++){
        if(keys[i] == _key){
            keys[i] = keys[keys.length -1];
            keys.pop();
            break;
        }
      }
   }

   function get(address _key) public view returns (uint){
       return data[_key];
    }

    function getAllKeys() public  view returns (address[] memory){
        return keys;
    }

    function getAll() public  view returns (address[] memory, uint[] memory){
        uint[] memory values = new uint[](keys.length);
        for (uint i =0;  i < keys.length ; i++){
            values[i] = data[keys[i]];
        }
         return  (keys,values);
        
    }

}