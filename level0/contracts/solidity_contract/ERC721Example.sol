// SPDX-License-Identifier: MIT
pragma solidity ^0.8.20; 

//ERC721 实现示例：收藏品合约

// import "@openzeppelin/contracts/token/ERC721/ERC721.sol";
// import "@openzeppelin/contracts/token/ERC721/ERC721Metadata.sol";
// import "@openzeppelin/contracts/token/ERC721/ERC721Enumerable.sol";
// contract Collectible is ERC721, ERC721Metadata, ERC721Enumerable {
//     constructor() ERC721Metadata("My Collectible", "MCL") public {}
//     // 创建新的 NFT
//     function mintCollectible(address to, uint256 tokenId, string memory uri) public {
//         _mint(to, tokenId);
//         _setTokenURI(tokenId, uri);
//     }
// }

//扩展功能：ERC721 枚举

// import "@openzeppelin/contracts/token/ERC721/ERC721Enumerable.sol";
// contract CollectibleWithEnumeration is ERC721Enumerable {
//     constructor() public {}
//     // 返回所有NFT的总数量
//     function getTotalSupply() public view returns (uint256) {
//         return totalSupply();
//     }
//     // 获取特定索引位置的 NFT
//     function getTokenByIndex(uint256 index) public view returns (uint256) {
//         return tokenByIndex(index);
//     }
// }