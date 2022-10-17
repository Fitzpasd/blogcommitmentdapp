// SPDX-License-Identifier: MIT
pragma solidity >=0.4.22 <0.9.0;

contract AccessControl { 
    address private _owner;

    mapping (address => bool) private _admins;

    error OnlyOwner();
    error OnlyAdmin();

    constructor() {
        _owner = msg.sender;
        _admins[_owner] = true;
    }

    modifier onlyOwner() {
        if (msg.sender != _owner)
            revert OnlyOwner();
        _;
    }

    modifier onlyAdmin() {
        if (!_admins[msg.sender])
            revert OnlyAdmin();
        _;
    }

    function setAdmin(address other, bool val) external onlyOwner {
        require(other != _owner);
        _admins[other] = val;
    }
}