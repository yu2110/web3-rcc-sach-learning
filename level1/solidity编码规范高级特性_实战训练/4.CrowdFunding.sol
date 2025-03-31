// SPDX-License-Identifier: GPL-3.0

pragma solidity >=0.7.0 <0.9.0;

contract CrowdFunding{

    address public immutable beneficPerson;
    uint256 public immutable fundingGoal;
    uint256 public fundingAmout;

    mapping (address => uint256) public funders;

    mapping (address => bool) private fundersInserted;

    address[] public fundersKey;

     bool public AVAIBABLE = true;


    constructor(address _beneficPersion, uint256 _fundingGoal){
        beneficPerson = _beneficPersion;
        fundingGoal = _fundingGoal;
    }

    function contribute() external payable {
          require(AVAIBABLE, "crowdFunding is closed");

          uint256 targetFundingAmout = fundingAmout + msg.value;

          uint256 backAmount = 0;


          if(targetFundingAmout > fundingGoal){
            backAmount = targetFundingAmout - fundingGoal;
            funders[msg.sender] += (msg.value - backAmount);
            fundingAmout += (msg.value - backAmount);
          }else{
            funders[msg.sender] += msg.value;
            fundingAmout += msg.value;
          }
          if(!fundersInserted[msg.sender]){
            fundersInserted[msg.sender] = true;
            fundersKey.push(msg.sender);
          }
          if (backAmount > 0){
            payable(msg.sender).transfer(backAmount);
          }
    }


    function close() external returns (bool){

       if(fundingAmout < fundingGoal){
         return false;
       }

       uint256  amount = fundingAmout;
       fundingAmout = 0;
       AVAIBABLE = false;

       payable(beneficPerson).transfer(amount);
       return true;
    }

    function fundersSize() public view returns (uint256){
        return fundersKey.length;
    }


}