// SPDX-License-Identifier: UNLICENSED
pragma solidity ^0.8.20;

/**
 合约⽤于跟踪从在线市场购买的商品的状态。
  创建该合约时，装运状态设置为 Pending 。 
  装运商品后，则将装运状态设置为 Shipped 并会发出事件。 
  交货后，则将商品的装 运状态设置为 Delivered ，并发出另⼀个事件。 
*/

contract Shipping {

    // Our predefined values for shipping listed as enums 
    enum ShippingStatus {Pending, Shipped, Delivered}
    // Save enum ShippingStatus in variable status 
    ShippingStatus private status;
     
    // Event to launch when package has arrived 
    event LogNewAlert(string descriptiong);

   
    // This initializes our contract state (sets enum to Pending once the program starts) 
     constructor(){status = ShippingStatus.Pending;}

    // Function to change to Shipped 
     
    function Shipped() public{
       status = ShippingStatus.Shipped;
       emit LogNewAlert("your package has shipped");
    }

    // Function to change to Delivered 
    function Delivered() public { 
        status = ShippingStatus.Delivered; 
        emit LogNewAlert("Your package has arrived"); 
    }
    // Function to get the status of the shipping 
    
    function getStatus(ShippingStatus _status) internal pure returns (string memory statusText){
        // Check the current status and return the correct name
        if (ShippingStatus.Pending == _status) return "Pending"; 
        if (ShippingStatus.Shipped == _status) return "Shipped"; 
        if (ShippingStatus.Delivered == _status) return "Delivered";
    }

    // Get status of your shipped item
    function Status() public view returns (string memory) { 
        ShippingStatus _status = status; 
        return getStatus(_status); 
    }


}