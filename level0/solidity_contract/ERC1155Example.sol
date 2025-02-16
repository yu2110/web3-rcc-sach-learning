// SPDX-License-Identifier: MIT
pragma solidity ^0.8.20; 

// import "@openzeppelin/contracts/token/ERC1155/ERC1155.sol";
// contract MyGameAssets is ERC1155 {
//     uint256 public constant GOLD = 0;
//     uint256 public constant SILVER = 1;
//     uint256 public constant SWORD = 2;
//     uint256 public constant SHIELD = 3;
//     constructor() public ERC1155("https://mygame.io/api/token/{id}.json") {
//         _mint(msg.sender, GOLD, 1000, "");
//         _mint(msg.sender, SILVER, 5000, "");
//         _mint(msg.sender, SWORD, 10, "");
//         _mint(msg.sender, SHIELD, 5, "");
//     }
// }

/**
GOLD 和 SILVER 表示同质化代币，每种代币有不同的数量。
SWORD 和 SHIELD 表示非同质化代币，每个代币都是独特的且数量有限。
*/