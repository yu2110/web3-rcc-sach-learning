// SPDX-License-Identifier: UNLICENSED
pragma solidity ^0.8.20;

import "@openzeppelin/contracts/token/ERC20/ERC20.sol";

contract ERC20MinerReward is ERC20 {

      event LogNewAlert(string description, address indexed _from, uint256 _n);

      constructor () ERC20("MinerRewoard","MRW"){}

    function _reward() public { 
        _mint(block.coinbase, 20); 
        emit LogNewAlert('_rewarded', block.coinbase, block.number); 
    } 

}
//“代币”这个词只是⼀个隐喻。 它是指由计算机网络或区块链网络共同管理的资产或访问权限。

