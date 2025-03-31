// SPDX-License-Identifier: MIT
pragma solidity ^0.8.20;

contract ExternalContract{
    function getValue() public  pure returns (uint){
        return 42;
    }

    function willRevert() public  pure{
        revert("this function always fails");
    }

}

contract TryCatchExample{
    ExternalContract externalContract;

    constructor(){
        externalContract = new ExternalContract();
    }

    function tryCatchTest() public view returns (uint, string memory) {
         try externalContract.getValue() returns (uint value){
                return (value, "Sucess");
         }catch {
                return (0, "Failed");
         }

    }

    function tryCatchWithRevert() public view returns (string memory) {
        try externalContract.willRevert(){
              return "This will not execute";
        }catch  Error(string memory reason){
            return reason;
        }catch{
            return "unknow error";
        }
    }
}