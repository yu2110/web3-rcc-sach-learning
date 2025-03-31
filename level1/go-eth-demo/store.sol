pragma solidity ^0.8.26;

contract Store {
  event ItemSet(bytes32 key, bytes32 value);
  event ItemGet(bytes32 key);

  string public version;
  mapping (bytes32 => bytes32) public items;

  constructor(string memory _version) {
    version = _version;
  }

  function setItem(bytes32 key, bytes32 value) external {
    items[key] = value;
    emit ItemSet(key, value);
  }

  function getItem(bytes32 key) external returns (bytes32) {
    bytes32 value = items[key];
    emit ItemGet(key);
    return value;
  }
}