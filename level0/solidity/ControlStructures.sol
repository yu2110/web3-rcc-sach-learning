// SPDX-License-Identifier: MIT
pragma solidity ^0.8.20;

contract ControlTest{

    function testWhile() public  pure returns (uint){
         uint i = 0;
         uint sumOfOdd = 0;
         while (true){
             i++;
             if(i % 2 == 0){
                continue ;//跳过偶数
             }
             if(i > 10 ){
                break ;
             }
             sumOfOdd +=i;//1+3+5+7+9=25
         }
         return sumOfOdd;
    }

    function testFor() public  pure returns(uint, uint){
        uint sumOfOdd =0;
        uint sumofEven = 0;
        for (uint i=0; i < 10; i++){
            if(i % 2 == 0){
                sumofEven += i;//累加偶数
            } else {
                sumOfOdd +=i;
            }
        }
        return (sumOfOdd, sumofEven);
    }

}