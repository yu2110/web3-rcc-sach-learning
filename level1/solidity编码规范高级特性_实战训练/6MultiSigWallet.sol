// SPDX-License-Identifier: GPL-3.0

pragma solidity >=0.7.0 <0.9.0;

contract MultiSigWallet {
  
  address[] public ownerArray;

  mapping(address => bool) public ownerHasMap;

  uint256 public agreeCount;

  struct Transaction{
     address to;
     uint256 amount;
     bytes data;
     bool invoked;
  }  

  Transaction[] public transactions;

  mapping (uint256 => mapping (address => bool)) txOwnerApprovedMap;

  event Deposit(address indexed from, uint256 amount);

  event Submit(uint256 indexed txId);

  event Approve(address indexed owner, uint256 indexed txId);

  event Revoke(address indexed owner, uint256 indexed txId);

  event Invoke(uint256 indexed txId);


   receive() external payable{
      emit Deposit(msg.sender, msg.value);
   }

   modifier isOwner(){
      require(ownerHasMap[msg.sender], "not owner");
      _;
   }

   modifier txExists(uint256 _txId){
      require(_txId < transactions.length, "tx does not exist");
      _;
   }

   modifier notApproved(uint256 _txId){
       require(!txOwnerApprovedMap[_txId][msg.sender], "tx already approved");
       _;
   }

   modifier notInvoked(uint256 _txId){
      require(!transactions[_txId].invoked, "tx is invoked");
      _;
   }


   constructor(address[] memory _ownerArray, uint256 _agreeCount){

       require(_ownerArray.length > 0, "required owner");
       require(_agreeCount>0 && _agreeCount<=_ownerArray.length,
       "invalid required number of owners");

       for(uint256 i=0; i< _ownerArray.length; i++){
          address owner = _ownerArray[i];
          require(owner != address(0), "invalid owner");
          require(!ownerHasMap[owner], "owner is not unique");

          ownerHasMap[owner] = true;

          ownerArray.push(owner);
       }

       agreeCount = _agreeCount;
   }

   function  getBalance() external view returns (uint256) {
        return address(this).balance;
   }

   function submit(address _to, uint256 _amount, bytes calldata _data) external isOwner returns (uint256){
       transactions.push(
        Transaction({to:_to, amount:_amount, data:_data, invoked: false})
       );

      emit Submit(transactions.length-1);
      return transactions.length -1;
   }

   function approve(uint256 _txId) external isOwner txExists(_txId) notApproved(_txId) notInvoked(_txId){
        txOwnerApprovedMap[_txId][msg.sender] = true;
        emit Approve(msg.sender, _txId);
   }

   function invoke(uint256 _txId) external isOwner txExists(_txId)  notInvoked(_txId){
       require(getAgreeCount(_txId) >= agreeCount, "agrees < agreeCount");

       Transaction storage tr = transactions[_txId];
       tr.invoked = true;
       (bool success, ) = tr.to.call{value: tr.amount}(tr.data);
        require(success, "tx failed");
        emit Invoke(_txId);  
   }

   function getAgreeCount(uint256 _txId) public view returns (uint256 count){
      for(uint256 i=0; i< ownerArray.length; i++){
        if(txOwnerApprovedMap[_txId][ownerArray[i]]){
            count +=1;
        }
      }
   }


   function revoke(uint256 _txId) external isOwner txExists(_txId)  notInvoked(_txId){
       require(txOwnerApprovedMap[_txId][msg.sender], "tx not approved");
       txOwnerApprovedMap[_txId][msg.sender] = false;
       emit Revoke(msg.sender, _txId);

   }
   
}