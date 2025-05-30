// SPDX-License-Identifier: MIT
pragma solidity ^0.8.20;

/**
Data Locations
Data Locations - Storage, Memory and Calldata
Variables are declared as either storage, memory or calldata to explicitly specify the location of the data.

storage variable is a state variable (store on blockchain)
memory variable is in memory and it exists while a function is being called
calldata special data location that contains function arguments
*/

contract DataLocations{
    uint256[] public arr;
    mapping(uint256 => address) map;


    struct MyStruct {
        uint256 foo;
    }

    mapping(uint256 => MyStruct) myStructs;


    function f() public {
        // call _f with state variables
        _f(arr, map, myStructs[1]);

        // get a struct from a mapping
        MyStruct storage myStruct = myStructs[1];
        // create a struct in memory
        MyStruct memory myMemStruct = MyStruct(0);
    }

    function _f(
        uint256[] storage _arr,
        mapping(uint256 => address) storage _map,
        MyStruct storage _myStruct
    ) internal {
        // do something with storage variables
    }


    // You can return memory variables
    function g(uint256[] memory _arr) public returns (uint256[] memory) {
        // do something with memory array
    }

    function h(uint256[] calldata _arr) external {
        // do something with calldata array
    }
}

//Transient Storage

// Make sure EVM version and VM set to Cancun

// Storage - data is stored on the blockchain
// Memory - data is cleared out after a function call
// Transient storage - data is cleared out after a transaction

interface ITest {
    function val() external view returns (uint256);
    function test() external;
}

contract Callback {
    uint256 public val;

    fallback() external {
        val = ITest(msg.sender).val();
    }

    function test(address target) external {
        ITest(target).test();
    }
}

contract TestStorage {
    uint256 public val;

    function test() public {
        val = 123;
        bytes memory b = "";
        msg.sender.call(b);
    }
}

contract TestTransientStorage {
    bytes32 constant SLOT = 0;

    function test() public {
        assembly {
            // tstore(SLOT, 321)；
        }
        bytes memory b = "";
        // msg.sender.call(b);
    }

    function val() public view returns (uint256 v) {
        assembly {
            // v := tload(SLOT)
        }
    }
}

contract ReentrancyGuard {
    bool private locked;

    modifier lock() {
        require(!locked);
        locked = true;
        _;
        locked = false;
    }

    // 35313 gas
    function test() public lock {
        // Ignore call error
        bytes memory b = "";
        msg.sender.call(b);
    }
}

contract ReentrancyGuardTransient {
    bytes32 constant SLOT = 0;

    modifier lock() {
        assembly {
            // if tload(SLOT) { revert(0, 0) }
            // tstore(SLOT, 1)
        }
        _;
        assembly {
            // tstore(SLOT, 0)
        }
    }

    // 21887 gas
    function test() external lock {
        // Ignore call error
        bytes memory b = "";
        msg.sender.call(b);
    }
}


